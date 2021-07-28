package home

import (
	"github.com/gorilla/mux"
	"github.com/haonly/Go-go-go/jaeseung/rest/middle"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handlers struct {
	l *logrus.Entry
}

func New() *Handlers {
	return &Handlers{
		l: logrus.WithFields(logrus.Fields{"tag": "home"}),
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
	m.HandleFunc("/", middle.LatencyPrint(h.Home))
}
