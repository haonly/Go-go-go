package vin

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/haonly/Go-go-go/jaeseung/rest/middle"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmlogrus"
	"net/http"
	"net/url"
	"strings"
)

type Handler struct {
	log  *logrus.Entry
	repo VinRepository
}

func New(repo VinRepository) *Handler {
	return &Handler{
		log:  logrus.WithFields(logrus.Fields{"tag": "vin"}),
		repo: repo,
	}
}

func (h *Handler) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/vin", middle.LatencyPrint(h.Create)).Methods(http.MethodPost)
	mux.HandleFunc("/vin/{id}", middle.LatencyPrint(h.Get)).Methods(http.MethodGet)
	mux.HandleFunc("/vin", middle.LatencyPrint(h.Update)).Methods(http.MethodPut)
	mux.HandleFunc("/vin/{id}", middle.LatencyPrint(h.Delete)).Methods(http.MethodDelete)
}

type ResponseError struct {
	Error   string   `json:"error"`
	Details []string `json:"details"`
}

type CreateRequest struct {
	VIN string `json:"vin"`
}

type CreateResponse struct {
	VID string `json:"vid"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	defer func(log *logrus.Entry) {
		h.log = log
	}(h.log)
	h.log = h.log.WithFields(apmlogrus.TraceContext(r.Context()))
	//traceContextFields := apmlogrus.TraceContext(r.Context())
	// Input request
	//h.log.WithFields(traceContextFields).Infof("111request=%v", r)
	msg := fmt.Sprintf("request=%v", r)
	h.log.WithFields(logrus.Fields{"message": msg}).Infof(msg)
	req := CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	// Process
	newVID := req.VIN + req.VIN
	if _, err := h.repo.Save(r.Context(), req.VIN, newVID); err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}

	// Output response
	res := CreateResponse{VID: newVID}
	h.writeJSONResponse(w, http.StatusOK, res)
}

type GetResponse struct {
	VID string `json:"vid"`
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	// Input request
	//traceContextFields := apmlogrus.TraceContext(r.Context())
	h.log = h.log.WithFields(apmlogrus.TraceContext(r.Context()))
	msg := fmt.Sprintf("request=%v", r)
	h.log.WithFields(logrus.Fields{"message": msg}).Infof(msg)
	//h.log.WithFields(traceContextFields).WithFields(logrus.Fields{"message": msg}).Infof(msg)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}
	vin := u.Path[strings.Index(u.Path, "vin/")+len("vin/"):]

	// Process
	vid, err := h.repo.Get(r.Context(), vin)
	if err != nil {
		h.handleError(w, http.StatusNotFound, err)
		return
	}

	// Output response
	res := GetResponse{VID: vid}
	h.writeJSONResponse(w, http.StatusOK, res)
}

type UpdateRequest struct {
	VIN string `json:"vin"`
	VID string `json:"vid"`
}

type UpdateResponse struct {
	VIN string `json:"vin"`
	VID string `json:"vid"`
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	// Input request
	h.log = h.log.WithFields(apmlogrus.TraceContext(r.Context()))
	msg := fmt.Sprintf("request=%v", r)
	h.log.WithFields(logrus.Fields{"message": msg}).Infof(msg)
	req := UpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	// Process
	if err := h.repo.Update(r.Context(), req.VIN, req.VID); err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}

	// Output response
	res := UpdateResponse{req.VIN, req.VID}
	h.writeJSONResponse(w, http.StatusOK, res)
}

type DeleteResponse struct {
	VID string
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	// Input request
	logger := logrus.New()
	logger.AddHook(&apmlogrus.Hook{})
	h.log = h.log.WithFields(apmlogrus.TraceContext(r.Context()))
	msg := fmt.Sprintf("request=%v", r)
	logger.WithFields(apmlogrus.TraceContext(r.Context())).WithField("message", msg).Info("ABCCCCC")
	logger.WithFields(apmlogrus.TraceContext(r.Context())).WithField("message", msg).Println("Aaaaaaaa")
	h.log.WithFields(logrus.Fields{"message": msg}).Infof(msg)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}
	vin := u.Path[strings.Index(u.Path, "vin/")+len("vin/"):]

	// Process
	vid, err := h.repo.Delete(r.Context(), vin)
	if err != nil {
		h.handleError(w, http.StatusNotFound, err)
		return
	}

	// Output response
	res := DeleteResponse{VID: vid}
	h.writeJSONResponse(w, http.StatusOK, res)
}

func (h *Handler) handleError(w http.ResponseWriter, statusCode int, err error) {
	errRes := ResponseError{http.StatusText(statusCode), []string{err.Error()}}
	h.log.Printf("err=%s", err.Error())
	h.writeJSONResponse(w, statusCode, errRes)
}

func (h *Handler) writeJSONResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	response, _ := json.Marshal(body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
	h.log.Printf("response=%v", response)
}
