package main

import (
	"net/http"
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

	// create a new handler
	handler := HttpHandler{}

	// listen and serve
	http.ListenAndServe(":11130", handler)
}
