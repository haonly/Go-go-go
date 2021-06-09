package home

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	l *log.Logger
}

func New(l *log.Logger) *Handlers {
	return &Handlers{
		l: l,
	}
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	h.l.Printf("context=%v", ctx)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello go rest"))
}

func (h *Handlers) SetupRoutes(m *mux.Router) {
	m.HandleFunc("/", h.ReqLatencyMiddleware(h.Home))
}

func (h Handlers) ReqLatencyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := time.Now()
		defer h.l.Printf("[mid] req processed for %v\n", time.Now().Sub(s))
		next(w, r)
	}
}
