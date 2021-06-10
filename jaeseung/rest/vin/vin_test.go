package vin

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestVin(t *testing.T) {
	suite.Run(t, new(VinTestSuite))
}

type VinTestSuite struct {
	suite.Suite
	cut *Handler
}

func (s *VinTestSuite) SetupTest() {
	logger := log.New(os.Stdout, "[motor] ", log.LstdFlags|log.Lshortfile)
	repo := NewMemRepository()
	s.cut = New(logger, repo)
}

func (s *VinTestSuite) TearDownTest() {

}

func (s *VinTestSuite) Test_Create_BadReq() {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("Content-Type", "text/json; charset=utf-8")
	w := httptest.NewRecorder()

	// Exercise
	s.cut.Create(w, req)
	actual := w.Result()

	// Verify
	assert.Equal(s.T(), http.StatusBadRequest, actual.StatusCode)
}

func (s *VinTestSuite) Test_Create() {
	// Setup
	createReq := CreateRequest{VIN: "1"}
	reqBody, err := json.Marshal(createReq)
	if err != nil {
		assert.Fail(s.T(), err.Error())
	}
	buf := bytes.NewBuffer(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/", buf)
	req.Header.Set("Content-Type", "text/json; charset=utf-8")
	w := httptest.NewRecorder()

	// Exercise
	s.cut.Create(w, req)
	actual := w.Result()
	body, _ := io.ReadAll(actual.Body)

	// Verify
	assert.Equal(s.T(), http.StatusOK, actual.StatusCode)
	assert.Equal(s.T(), "text/json; charset=utf-8", actual.Header.Get("Content-Type"))

	var expected CreateResponse
	if err := json.Unmarshal(body, &expected); err != nil {
		assert.Fail(s.T(), err.Error())
	}
	assert.Equal(s.T(), "11", expected.VID)
}

func (s *VinTestSuite) Test_Get_NotFound() {
	// Setup
	values := url.Values{}
	values.Add("id", "1")
	req := httptest.NewRequest(http.MethodGet, "/vin", strings.NewReader(values.Encode()))
	w := httptest.NewRecorder()

	// Exercise
	s.cut.Get(w, req)
	actual := w.Result()

	// Verify
	assert.Equal(s.T(), http.StatusNotFound, actual.StatusCode)
}

func (s *VinTestSuite) helperCreateVID() {
	createReq := CreateRequest{VIN: "1"}
	reqBody, err := json.Marshal(createReq)
	if err != nil {
		assert.Fail(s.T(), err.Error())
	}
	buf := bytes.NewBuffer(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/", buf)
	req.Header.Set("Content-Type", "text/json; charset=utf-8")
	w := httptest.NewRecorder()

	// Exercise
	s.cut.Create(w, req)
	actual := w.Result()

	// Verify
	assert.Equal(s.T(), http.StatusOK, actual.StatusCode)
	assert.Equal(s.T(), "text/json; charset=utf-8", actual.Header.Get("Content-Type"))
}

func (s *VinTestSuite) Test_Get() {
	// Setup
	s.helperCreateVID()
	req := httptest.NewRequest(http.MethodGet, "/vin/1", nil)
	w := httptest.NewRecorder()
	log.Println(req.URL.String())

	// Exercise
	s.cut.Get(w, req)
	actual := w.Result()
	body, _ := io.ReadAll(actual.Body)

	// Verify
	assert.Equal(s.T(), http.StatusOK, actual.StatusCode)
	var expected GetResponse
	if err := json.Unmarshal(body, &expected); err != nil {
		assert.Fail(s.T(), err.Error())
	}
	assert.Equal(s.T(), "11", expected.VID)
}

func (s *VinTestSuite) Test_Delete_NotFound() {
	// Setup
	req := httptest.NewRequest(http.MethodGet, "/vin/1", nil)
	w := httptest.NewRecorder()
	log.Println(req.URL.String())

	// Exercise
	s.cut.Delete(w, req)
	actual := w.Result()

	// Verify
	assert.Equal(s.T(), http.StatusNotFound, actual.StatusCode)
}

func (s *VinTestSuite) Test_Delete() {
	// Setup
	s.helperCreateVID()
	req := httptest.NewRequest(http.MethodGet, "/vin/1", nil)
	w := httptest.NewRecorder()
	log.Println(req.URL.String())

	// Exercise
	s.cut.Delete(w, req)
	actual := w.Result()
	body, _ := io.ReadAll(actual.Body)

	// Verify
	assert.Equal(s.T(), http.StatusOK, actual.StatusCode)
	var expected GetResponse
	if err := json.Unmarshal(body, &expected); err != nil {
		assert.Fail(s.T(), err.Error())
	}
	assert.Equal(s.T(), "11", expected.VID)
}
