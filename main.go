package main 

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/{num}", Multiplicar)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Por favor, ingrese un número.")
}

func Multiplicar(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    num, err := strconv.Atoi(vars["num"])
    if err != nil {
    	fmt.Fprintln(w, "Por favor, ingrese un número.")
    } else {
    	fmt.Fprintln(w, num*2)
    }
}