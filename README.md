# Prometheus exporter for an Adafruit SHT4X Trinkey

## Usage

```console
$ go build ./...
$ sudo ./adafruit-sht4x-trinkey-exporter -port-name /dev/ttyABCD
2024/08/01 10:33:08 Opening serial connection to /dev/ttyABCD with BaudRate 115200 ...
2024/08/01 10:33:08 Listening for HTTP traffic on :8080 ...
```
```console
$ curl -fsS localhost:8080/metrics | grep adafruit
# HELP adafruit_relative_humidity Relative humidity reading, in %%, from an Adafruit SHT4X trinkey
# TYPE adafruit_relative_humidity gauge
adafruit_relative_humidity 36.06
# HELP adafruit_temp_celsius Temperature reading, in degrees C, from an Adafruit SHT4X trinkey
# TYPE adafruit_temp_celsius gauge
adafruit_temp_celsius 30.77
```

## Product documentation

https://learn.adafruit.com/adafruit-sht4x-trinkey
