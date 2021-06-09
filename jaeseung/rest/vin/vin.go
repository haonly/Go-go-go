package vin

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Handler struct {
	log  *log.Logger
	repo *MemRepository
}

func New(l *log.Logger, repo *MemRepository) *Handler {
	return &Handler{
		log:  l,
		repo: repo,
	}
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
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var cReq CreateRequest
	if err := json.Unmarshal(reqBody, &cReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Process
	newVID := cReq.VIN + cReq.VIN
	if _, err := h.repo.Save(cReq.VIN, newVID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Output response
	res := CreateResponse{VID: newVID}
	resJson, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.Write(resJson)
	h.log.Printf("response=%v", string(resJson))
}

type GetResponse struct {
	VID string `json:"vid"`
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	// Input request
	h.log.Printf("request=%v", r)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	vin := u.Path[strings.Index(u.Path, "vin/")+len("vin/"):]

	// Process
	vid, err := h.repo.Get(vin)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Output response
	res := GetResponse{VID: vid}
	resJson, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.Write(resJson)
	h.log.Printf("response=%v", string(resJson))
}

func (h *Handler) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/vin", h.Create).Methods(http.MethodPost)
	mux.HandleFunc("/vin/{id}", h.Get).Methods(http.MethodGet)
	mux.HandleFunc("/vin/{id}", h.Delete).Methods(http.MethodDelete)
}

type DeleteResponse struct {
	VID string
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	// Input request
	h.log.Printf("request=%v", r)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	vin := u.Path[strings.Index(u.Path, "vin/")+len("vin/"):]

	// Process
	vid, err := h.repo.Delete(vin)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Output response
	res := DeleteResponse{VID: vid}
	resJson, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.Write(resJson)
	h.log.Printf("response=%v", string(resJson))
}
