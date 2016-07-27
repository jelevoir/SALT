package main

import (
	"encoding/json"
    "os"
    "net/http"
    "bytes"
)

// JSON object sent to the Google Natural Language API
type SentimentRequest struct {
    Document            `json:"document"`
}
type Document struct {
    Type string         `json:"type"`
    Language string     `json:"language"`
    Content string      `json:"content"`
}

// JSON object retrieved from the Google Natural Language API
type SentimentResponse struct {
	Sentiment           `json:"documentSentiment"`
	Language string     `json:"language"`
}
type Sentiment struct {
    Polarity float64    `json:"polarity"`
    Magnitude float64   `json:"magnitude"`
}

// Function to get the API Key from NATLANG_API_KEY environment variable
func GetAPIKey() string {
    key := os.Getenv("NATLANG_API_KEY")
    return key

}

//Function to construct the API request from the API Key and Tweet
func GetSentiment(key, content string) *SentimentResponse {
    // Google Cloud Natural Language API URI
    url := "https://language.googleapis.com/v1beta1/documents:analyzeSentiment?key=" + key

    content_sentiment := SentimentRequest{Document{"PLAIN_TEXT", "EN", content}}
    content_request, err := json.Marshal(content_sentiment)
    if err != nil {
        panic(err)
    }
    
    var natlang_request = []byte(content_request)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(natlang_request))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    s := new(SentimentResponse)

	err = json.NewDecoder(resp.Body).Decode(&s)
	if err != nil {
		panic(err)
	}
    
	return s
}

