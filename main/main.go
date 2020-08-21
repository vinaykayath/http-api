package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Creating a server")

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Failed to start Server")
	}

}

func hello(w http.ResponseWriter, req *http.Request) {
	//fmt.Println("Hellow World")
	fmt.Println("Request Recieved")
	body, _ := ioutil.ReadAll(req.Body)

	responseString := "Hello, " + string(body)
	w.Write([]byte(responseString))
	//w.Write([]byte("Hello, ", + string(type)))
	//w.Write([]byte("This is my First API"))
}
