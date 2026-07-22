package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* struct- store the incoming url
struct to store the shortened url
*/

type shortenRequest struct {
	URL string `json:"url"`
}
type shortenResponse struct {
	SHORTURL string `json:"short_url"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "URL shortener api is running")
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	var req shortenRequest

	err:= json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	response := shortenResponse{
		SHORTURL: "http://localhost:8080/abc123",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main () {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/shorten", shortenHandler)

	fmt.Println("server is running on port 8080")

	err:=http.ListenAndServe(":8080", nil) 
	if err != nil {
		fmt.Println("error starting server:", err)
	}
}