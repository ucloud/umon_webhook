package utils

import (
	"log"
	"net/http"
	"time"
)

func Logger(h http.Handler, name string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			h.ServeHTTP(w, r)

			log.Printf("%s\t%s\t%s\t%s\n",
				r.Method, r.RequestURI, name, time.Since(start))
		})
}
