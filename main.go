package main

import (
	"net/http"

	configcat "github.com/configcat/go-sdk/v9"
	"github.com/gin-gonic/gin"
)

var client = configcat.NewClient("YOUR-CONFIGCAT-SDK-KEY")

type job struct {
	Title      string `json:"title"`
	Experience string `json:"experience"`
	Salary     string `json:"salary"`
}

var jobs = []job{
	{Title: "Software Engineer II", Experience: "6 Years", Salary: "$75,000"},
	{Title: "Graphic Designer", Experience: "2 Years", Salary: "$54,000"},
}

func getJobs(c *gin.Context) {
	isJobsEndpointEnabled := client.GetBoolValue("jobsEndpoint", false, nil)
	if isJobsEndpointEnabled {
		c.IndentedJSON(http.StatusOK, jobs)
	} else {
		c.IndentedJSON(http.StatusNotFound, "API endpoint disabled")
	}
}

func main() {
	router := gin.Default()
	router.GET("/jobs", getJobs)
	router.SetTrustedProxies(nil)
	router.Run("localhost:8080")
}
