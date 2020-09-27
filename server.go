package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests",
	}, []string{"code", "method"})
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// create response binary data
	data := []byte("Hello Hepsiburada from @BurakBandak") // slice of bytes

	// write `data` to response
	res.Write(data)
}

func main() {

	r := prometheus.NewRegistry()
	r.MustRegister(httpRequestsTotal)

	// create a new handler
	handler := HttpHandler{}

	// listen and serve
	http.ListenAndServe(":11130", handler)

	http.Handle("/", promhttp.InstrumentHandlerCounter(httpRequestsTotal, handler))
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
}
