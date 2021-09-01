package app

import (
	"Go-go-go/hayeon-kim/registration/app/models"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Registration API")
	}
}

func (a *App) CreateRegistrationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.RegistrationRequest{}
		err := parse(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse registration body, err=%v, \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest) // http status를 같이 보냄
			return
		}

		//TODO:
		// MasterUserID -> (jwt 토큰) access 토큰을 활용 / userGetProfileAPI 사용하여 email 로
		// 클라이언트 정책에 따라 달라지겠는데, 그리고 토큰 유효성도 확인, jwt secret은 어떻게 가져오냐 -> 승윤매니저님이 알려줌(깃에 올라가있때)
		// 인증인가는 Ingress를 쓸 수 있다. --> 지효매니저님이 세미나를 해주신대!!!!!
		// Create registration
		reg := &models.Registration{
			"0",
			"2",
			req.MasterUserID,
			req.Company,
			req.Model,
			req.Type,
			req.Year,
			req.Capacity,
			time.Now(),
			time.Now(),
		}

		// save in db
		err = a.DB.CreateRegistration(reg)
		if err != nil {
			log.Printf("Cannot save registration in DB. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := mapRegistrationToJSON(reg)
		sendResponse(w, r, resp, http.StatusOK)
	}
}
