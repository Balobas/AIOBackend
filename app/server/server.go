package server

/*
основная часть сервера
 */

import (
	"fmt"
	"github.com/gorilla/schema"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/", Route)
	err := http.ListenAndServe(":9505", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	type u struct {
		Uid string `json:"uid"`
	}
	var U u

	schema.NewDecoder().Decode(&U, r.Form)

	fmt.Println(U)
}

