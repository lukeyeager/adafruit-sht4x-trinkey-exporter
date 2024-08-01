package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.bug.st/serial"
)

const (
	BaudRate int = 115200
)

var (
	tempGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "adafruit",
		Name:      "temp_celsius",
		Help:      "Temperature reading, in degrees C, from an Adafruit SHT4X trinkey",
	})
	rhGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "adafruit",
		Name:      "relative_humidity",
		Help:      "Relative humidity reading, in %%, from an Adafruit SHT4X trinkey",
	})
)

func readFromSerial(portName string) {
	// open serial connection
	mode := &serial.Mode{BaudRate: BaudRate}
	log.Printf("Opening serial connection to %v with BaudRate %v ...", portName, BaudRate)
	port, err := serial.Open(portName, mode)
	if err != nil {
		log.Fatalf("open %v: %v", portName, err)
	}

	reader := csv.NewReader(port)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// parse CSV row
		// header is: 'Serial number, Temperature in *C, Relative Humidity %, Touch'
		temp, err := strconv.ParseFloat(strings.TrimSpace(record[1]), 64)
		if err != nil {
			log.Fatal(err)
		}
		rh, err := strconv.ParseFloat(strings.TrimSpace(record[2]), 64)
		if err != nil {
			log.Fatal(err)
		}

		// update gauges
		tempGauge.Set(temp)
		rhGauge.Set(rh)
	}
}

func main() {
	// CLI args
	portName := flag.String("port-name", "", "path to the serial console port device (e.g. /dev/ttyABCD)")
	addr := flag.String("addr", ":8080", "address on which to listen for HTTP requests")
	flag.Parse()
	if *portName == "" {
		log.Fatal("-port-name is required")
	}

	// start thread to read data
	go readFromSerial(*portName)

	// serve prometheus metrics over http
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Listening for HTTP traffic on %v ...", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
