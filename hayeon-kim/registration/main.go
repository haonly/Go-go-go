package main

import (
	"Go-go-go/hayeon-kim/registration/app"
	"Go-go-go/hayeon-kim/registration/app/database"
	"log"
	"net/http"
	"os"
)

func main() {
	//TODO:
	// DB 캡슐화
	// dependancy injection --> 지금 코드는 의존성이 있으니 명시적으로 넣어줘라
	// app.New() 할 때 db를 인자로 받아서(소문자로 안에서만 사용하게 함) db 변경 안되도록 수정
	app := app.New()
	app.DB = &database.DB{}
	err := app.DB.Open()
	check(err)

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("App running...")

	err = http.ListenAndServe(":9000", nil)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
