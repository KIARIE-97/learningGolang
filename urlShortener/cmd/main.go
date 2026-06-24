package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "URL shortener api is running")
}

func main () {
	http.HandleFunc("/", homeHandler)

	fmt.Println("server is running on port 8080")

	err:=http.ListenAndServe(":8080", nil) 
	if err != nil {
		fmt.Println("error starting server:", err)
	}
}