package vin

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
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
	mux.HandleFunc("/vin", h.Create).Methods(http.MethodPost)
	mux.HandleFunc("/vin/{id}", h.Get).Methods(http.MethodGet)
	mux.HandleFunc("/vin", h.Update).Methods(http.MethodPut)
	mux.HandleFunc("/vin/{id}", h.Delete).Methods(http.MethodDelete)
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
	// Input request
	h.log.Printf("request=%v", r)
	req := CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	// Process
	newVID := req.VIN + req.VIN
	if _, err := h.repo.Save(req.VIN, newVID); err != nil {
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
	h.log.Printf("request=%v", r)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}
	vin := u.Path[strings.Index(u.Path, "vin/")+len("vin/"):]

	// Process
	vid, err := h.repo.Get(vin)
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
	h.log.Printf("request=%v", r)
	req := UpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}

	// Process
	if err := h.repo.Update(req.VIN, req.VID); err != nil {
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
	h.log.Printf("request=%v", r)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}
	vin := u.Path[strings.Index(u.Path, "vin/")+len("vin/"):]

	// Process
	vid, err := h.repo.Delete(vin)
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
	h.log.Printf("err=%s", err.Error)
	h.writeJSONResponse(w, statusCode, errRes)
}

func (h *Handler) writeJSONResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	response, _ := json.Marshal(body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
	h.log.Printf("response=%v", response)
}
