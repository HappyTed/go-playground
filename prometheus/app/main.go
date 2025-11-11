package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"runtime/metrics"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const PORT = ":1234"

// своя метрика типа counter
var counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "my_metrics_group",
		Name:      "my_counter",
		Help:      "This is my counter",
	})

// своя метрика типа gauge
var gauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "my_metrics_group",
		Name:      "my_gauge",
		Help:      "This is my gauge",
	})

// своя метрика с типом гистограмы
var histogram = prometheus.NewHistogram(prometheus.HistogramOpts{
	Namespace: "my_metrics_group",
	Name:      "my_histogram",
	Help:      "This is my histogram",
})

// своя метрика с типом summury
var summury = prometheus.NewSummary(prometheus.SummaryOpts{
	Namespace: "my_metrics_group",
	Name:      "my_summary",
	Help:      "This is my summary",
})

func metreicsMiddleware(target http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("→ %s %s", r.Method, r.URL.Path)

		target.ServeHTTP(w, r) // вызываем основной обработчик

		log.Printf("← %s %s (%v)", r.Method, r.URL.Path, time.Since(start))
	})
}

// ---
// Получаем метрики по памяти и горутинам
var n_goroutines = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "base_metrics",
		Name:      "n_goroutines",
		Help:      "Number of goroutines",
	})

var n_memory = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "base_metrics",
		Name:      "n_memory",
		Help:      "Memory usage",
	})

func main() {
	rand.Seed(time.Now().Unix()) // deprecated, TODO: заменить на новый метод

	// регестрируем переменные метрик
	// custom:
	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summury)
	// base:
	prometheus.MustRegister(n_goroutines)
	prometheus.MustRegister(n_memory)

	// генерим рандомные данные для кажой метрики, пока работает веб-сервер
	go func() {
		for {
			counter.Add(rand.Float64() * 5)
			gauge.Add(rand.Float64()*15 - 5)
			histogram.Observe(rand.Float64() * 10)
			summury.Observe(rand.Float64() * 10)
			time.Sleep(2 * time.Second)
		}
	}()

	// собираем базовые метрики
	const nGo = "/sched/goroutines:goroutines"
	const nMem = "/memory/classes/heap/free:bytes"
	getBaseMetrics := make([]metrics.Sample, 2)
	getBaseMetrics[0].Name = nGo
	getBaseMetrics[1].Name = nMem

	go func() {
		runtime.GC()
		metrics.Read(getBaseMetrics)
		goVal := getBaseMetrics[0].Value.Uint64()
		memVal := getBaseMetrics[1].Value.Uint64()
		time.Sleep(time.Duration(5 * time.Second))

		n_goroutines.Set(float64(goVal))
		n_memory.Set(float64(memVal))
	}()

	mux := http.NewServeMux()
	mux.Handle("/metrics", metreicsMiddleware(promhttp.Handler()))

	fmt.Println("Listening to port", PORT)
	fmt.Println(http.ListenAndServe(PORT, mux))
}
