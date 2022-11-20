package main

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/configcat/go-sdk/v7"
)

var client = configcat.NewClient("ScDaCD8ETUuG7wYo3BdP2A/5s96HBVckk-RzI-iVf-zRA")

type Job struct {
	Title string `json:"title"`
	Experience string `json:"experience"`
	Salary string `json:"salary"`
}

type Jobs []Job

func allJobs(w http.ResponseWriter, r *http.Request) {
	jobs := Jobs{
		Job{Title: "Software Engineer II", Experience: "6 Years", Salary: "$75,000"},
		Job{Title: "Graphic Designer", Experience: "2 Years", Salary: "$54,000"},
	}

	json.NewEncoder(w).Encode(jobs)
}

func handleRequests() {
	isJobsEndpointEnabled := client.GetBoolValue("jobsendpoint", false, nil)

	if isJobsEndpointEnabled {
		http.HandleFunc("/jobs", allJobs)
	}

	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main() {
	handleRequests()
}



