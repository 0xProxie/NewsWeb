package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// news represents data about a new entry on a news report.
type news struct {
	ID          string    `json:"id"`
	Headline    string    `json:"title"`
	Autor       string    `json:"autor"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
}

// newsResponse represents the response format for a news entry.
type newsResponse struct {
	ID          string `json:"id"`
	Headline    string `json:"title"`
	Autor       string `json:"autor"`
	Description string `json:"description"`
	Time        string `json:"time"`
}

// news slice to seed record news entry.
var newsReport = []news{
	{ID: "1", Headline: "The Nigerian Economy", Autor: "Jude Akin", Description: "The Nigerian economy is becoming worse by the day and there is very little an average man can do about it.", Time: time.Now()},
	{ID: "2", Headline: "2024 bull run and what to expect", Autor: "John Terry", Description: "The 2024 bull run is fast approaching, better stack up your coin so as not to be left out.", Time: time.Now()},
	{ID: "3", Headline: "Crypto and the new evolution", Autor: "Emmanuel Mark", Description: "Seems like the world is now beginning to accept the new evolution and the fact that crypto is here to stay ", Time: time.Now()},
}

func main() {
	router := gin.Default()
	router.GET("/newsReport", getNewsReport)
	router.GET("/newsReport/:id", getNewsByID)
	router.POST("/newsReport", postNewsReport)

	router.Run("localhost:8080")
}

// getNewsReport fetches with the list of all news as JSON.
func getNewsReport(c *gin.Context) {
	var response []newsResponse
	for _, n := range newsReport {
		response = append(response, formatNewsResponse(n))
	}
	c.IndentedJSON(http.StatusOK, response)
}

// postNewsReport adds a new news from JSON received in the request body.
func postNewsReport(c *gin.Context) {
	var newNews news

	// Call BindJSON to bind the received JSON to newNews
	if err := c.BindJSON(&newNews); err != nil {
		return
	}

	// Add the new news to the slice.
	newNews.Time = time.Now()
	newsReport = append(newsReport, newNews)
	c.IndentedJSON(http.StatusCreated, formatNewsResponse(newNews))
}

// getNewsByID locates the news whose ID value matches the id
// parameter sent by the client, then returns that news as a response.
func getNewsByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of newsReport, looking for
	// a news whose ID value matches the parameter.
	for _, a := range newsReport {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, formatNewsResponse(a))
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "news not found"})
}

// formatNewsResponse formats a news entry for response, including formatting the time.
func formatNewsResponse(n news) newsResponse {
	return newsResponse{
		ID:          n.ID,
		Headline:    n.Headline,
		Autor:       n.Autor,
		Description: n.Description,
		Time:        n.Time.Format("2006-01-02 15:04:05"),
	}
}
