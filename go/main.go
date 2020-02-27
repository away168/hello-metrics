package main

import (
	"flag"
	"log"
	// "math"
	// "fmt"
	"math/rand"
	"net/http"
	"time"
	"os"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var Port int
	var Mean, StdDev float64
	var Version, Environment string

	flag.IntVar(&Port, "port", LookupEnvOrInt("PROMETHEUS_PORT", 8080), "Listen Port")
	flag.Float64Var(&Mean, "mean", float64(LookupEnvOrInt("DUMMY_MEAN", 100)), "Mean Value")
	flag.Float64Var(&StdDev, "sd", float64(LookupEnvOrInt("DUMMY_SD", 20)), "Standard Deviation")
	flag.StringVar(&Version, "version", LookupEnvOrString("DUMMY_VERSION", "NA"), "App Version")
	flag.StringVar(&Environment, "environment", LookupEnvOrString("DUMMY_ENVIRONMENT", "NA"), "Environment")

	flag.Parse()

	var gauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "custom",
			Name:      "dummy_latency",
			Help:      "(Fake) Latency",
		},
		[]string{"version", "environment"},
	)

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.Handler())

	prometheus.MustRegister(gauge)

	go func() {
		for {
			v := (rand.NormFloat64() * StdDev) + Mean
			gauge.WithLabelValues(Version, Environment).Set(v)
			// gauge.Add(10)

			time.Sleep(time.Second)
		}
	}()

	listener := ":" + strconv.Itoa(Port)
	// fmt.Print(Port)
	// fmt.Print(listener)
	log.Fatal(http.ListenAndServe(listener, nil))
}


func LookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("LookupEnvOrInt[%s]: %v", key, err)
		}
		return v
	}
	return defaultVal
}

func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}