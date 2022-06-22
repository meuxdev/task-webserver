package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello root! 🚀")
}

func HandleHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is the API endpoint")
}

func UserPostRequest(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var userModel User
	// send decoder as reference
	err := decoder.Decode(&userModel)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	fmt.Println("Name:", userModel.Name)
	// %v so that the message is formated
	fmt.Fprintf(w, "Payload: %v", userModel)
}
