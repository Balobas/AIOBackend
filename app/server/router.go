package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	//"log"
	"net/http"
)
/*
Маршрутизация запросов
 */

func Route(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	fmt.Println(url, "////")
	var params interface{}
	if r.Method == http.MethodGet {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		params = r.Form
	}
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(body))
		params = body
	}
	//token := r.Form.Get("access_token")
	//if token == "" {
	//	fmt.Println("empty access token")
	//	return
	//}
	action, ok := RoutesMap[url]
	if !ok {
		fmt.Println("404 - method not found")
	}
	result := reflect.ValueOf(action).Call([]reflect.Value{reflect.ValueOf(params)})[0]
	_, err := fmt.Fprint(w, result)
	if err != nil {
		fmt.Println(err.Error())
	}
}