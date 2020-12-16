package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	request1()
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to my Go Web API")
	fmt.Println("Endpoint Hit: Homepage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "A little bit about me")
	fmt.Println("Endpoint Hit: aboutMe")
}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutMe)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
