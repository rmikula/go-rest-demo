package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type AppInfo struct {
	Version string    `json:"version"`
	Time    time.Time `json:"time"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	i := AppInfo{Version: "1.0", Time: time.Now()}

	json.NewEncoder(w).Encode(i)

	// fmt.Fprintf(w, p)
	fmt.Println("Endpoint Hit: homePage")
}

func personsHandler(w http.ResponseWriter, r *http.Request) {

	/*
		p := Person{
			Name: "Roman",
			Age:  30,
		}
	*/

	pp := []Person{
		{
			Name: "Roman",
			Age:  323,
		},
		{
			Name: "Monika",
			Age:  32,
		},
		{
			Name: "Tereza",
			Age:  19,
		},
	}

	json.NewEncoder(w).Encode(pp)

	// fmt.Fprintf(w, p)
	fmt.Println("Endpoint Hit: homepage")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	id := vars["id"]

	fmt.Fprintf(w, "Category: %v\n", id)

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homePage)

	r.HandleFunc("/persons", personsHandler)
	r.HandleFunc("/persons/{id}", personHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}
