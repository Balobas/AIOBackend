package main

import (
	"./app/server"
	"bytes"
	"encoding/json"

	"./data"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	user := data.User{
		UID:      "1574dbe2-b119-4ebb-8062-03000f17166d",
		Name:     "Set",
		FullName: "Petyar Pupkin",
		Login:    "ioiasasa",
		Phone:    "+79654731509",
		Email:    "pupan@gmail.com",
		Location: "Minsk",
	}
	//db := database.CouchDB{}
	//err := db.Init("balobas", "balobas", "http://localhost:5984")
	//if err != nil {
	//	fmt.Println("db init error ")
	//}
	//
	////fmt.Println(users.CreateUser(&db, user))
	//

	//_, err := uuid.FromString("0D693249-7c8a-48eb-a8a8-74ea77660b85")
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println("good")
	//}
	go server.Run()

	//r, err := http.Get("http://localhost:9505/Api/GetUser?uid=1574dbe2-b119-4ebb-8062-03000f17166d")

	b, err := json.Marshal(user)
	r, err := http.Post("http://localhost:9505/Api/CreateUser", "application/json", bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
	fmt.Println("end")
}