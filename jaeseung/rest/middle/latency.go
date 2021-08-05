package middle

import (
	"log"
	"net/http"
	"time"
)

func LatencyPrint(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := time.Now()
		defer log.Printf("[mid] req processed for %v\n", time.Now().Sub(s))
		next(w, r)
	}
}
