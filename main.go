package main

import (
	"fmt"

	"github.com/gorilla/mux"

	"net/http"
)

func main() {

	fmt.Println("Server running on port 3000")
	r := mux.NewRouter()

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
