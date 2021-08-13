package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/haonly/Go-go-go/jaeseung/carsharing/review/usecase"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmlogrus"
	"net/http"
)

type Handler struct {
	log *logrus.Entry
	uc  *usecase.ReviewService
}

func NewReview(uc *usecase.ReviewService) *Handler {
	return &Handler{
		log: logrus.WithFields(logrus.Fields{"tag": "carsharing"}),
		uc:  uc,
	}
}

func (h *Handler) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/carsharing", h.Create).Methods(http.MethodPost)
}

type ResponseError struct {
	Error   string   `json:"error"`
	Details []string `json:"details"`
}

type CreateRequestJSON struct {
	CreatorId      string  `json:"creator_id"`
	ReviewTargetId string  `json:"review_target_id"`
	Text           string  `json:"text,omitempty"`
	Rating         float32 `json:"rating,omitempty"`
}

type CreateResponse struct {
	VID string `json:"vid"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	// Input request
	h.log = h.log.WithFields(apmlogrus.TraceContext(r.Context()))
	msg := fmt.Sprintf("request=%v", r)
	h.log.WithFields(logrus.Fields{"message": msg}).Infof(msg)
	req := CreateRequestJSON{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.handleError(w, http.StatusBadRequest, err)
		return
	}
	reqDto := usecase.CreateReviewReqDTO{
		Ctx:            r.Context(),
		CreatorId:      req.CreatorId,
		ReviewTargetId: req.ReviewTargetId,
		Text:           req.Text,
		Rating:         req.Rating,
	}

	// Process
	res, err := h.uc.CreateReview(reqDto)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, err)
		return
	}

	// Output response
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
