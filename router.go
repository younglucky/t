package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func init() {
	fmt.Println("router.go init()...")

	router.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "hello world")
	})

	router.HandleFunc("/test/{weixinId}", func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		weixinId := vars["weixinId"]

		var u User
		if err := u.GetByWeixinId(weixinId); err != nil {
			r.HTML(rw, 200, "index", nil)
		} else {
			r.HTML(rw, 200, "index", u)
		}
	})

	router.HandleFunc("/create", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("create...")
		bytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("request body: ", string(bytes))
		var u User

		if err = json.Unmarshal(bytes, &u); err != nil {
			panic(err)
		}
		fmt.Println("u:", u)
		if err = u.Create(); err != nil {
			panic(err)
		} else {
			r.JSON(rw, 200, "success")
		}
	}).Methods("post")
}
