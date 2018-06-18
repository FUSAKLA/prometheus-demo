package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestDurationHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "HTTP request latency distribution",
	},
		[]string{"status", "endpoint"},
	)
)

func initializeMetrics() {
	prometheus.MustRegister(httpRequestDurationHistogram)
}

func setLogging() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func loadGopher(status int) (*os.File, error) {
	return os.Open(fmt.Sprintf("static/%d.png", status))
}

func renderGopher(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var (
		i   = rand.Intn(6) * 100
		img *os.File
		err error
	)

	if i >= 500 {
		img, err = loadGopher(i)
	} else {
		i = 200
		img, err = loadGopher(i)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(i)
	w.Header().Set("Content-Type", "image/png")
	io.Copy(w, img)
	duration := time.Since(start)
	httpRequestDurationHistogram.WithLabelValues(strconv.Itoa(i), "/demo").Observe(duration.Seconds())
	log.Debug("Request " + r.URL.RequestURI() + " returned status " + strconv.Itoa(i) + " in " + strconv.FormatFloat(duration.Seconds(), 'E', -1, 32) + "s")
	return
}

func main() {
	setLogging()
	log.Info("Starting up error-server...")
	initializeMetrics()
	http.HandleFunc("/demo", renderGopher)
	http.Handle("/demo/metrics", promhttp.Handler())
	log.Info("Server up and running ready to recieve requests!")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
