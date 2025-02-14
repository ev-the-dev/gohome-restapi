package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type cmdresult struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func main() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/api/v1/getdate", getdate)
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		os.Exit(1)
	}
}

func getdate(write http.ResponseWriter, _ *http.Request) {
	result := cmdresult{}

	out, err := exec.Command("date").Output()
	if err == nil {
		result.Message = "The date is " + string(out)
		result.Success = true
	}

	json.NewEncoder(write).Encode(result)
}

func homepage(write http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(write, "GoHome Simple REST API Server")
}
