package main

import (
	"fmt"
	"net/url"
)

func main() {
	twitterconfig, _ := LoadTwitterConfig("twitterconfig.json")
	api := NewTwitter(twitterconfig)

	postgresqlconfig, _ := LoadPostgreSQLConfig("postgresql.json")
	db, err := NewDatabase(postgresqlconfig)
	if err != nil {
		fmt.Println(err)
	}

	// Implements location specific tweet
	v := url.Values{}
	//Lookup geocodes at geocoder.us
	v.Add("geocode", "42.6725469,-83.2170434,20mi")
	v.Add("lang", "en")
	searchResult, _ := api.GetSearch("#StrangerThings", v)

	// Get Google Natural Language API Key
	key := GetAPIKey()

	var sentiments []Sentiments

	for _, tweet := range searchResult.Statuses {
		sentiment := GetSentiment(key, tweet.Text)
		sentiments = append(sentiments, Sentiments{0, sentiment.Sentiment.Polarity, sentiment.Sentiment.Magnitude, tweet.User.ScreenName})
	}

	// Insert Sentiments into database
	err = InsertSentiment(db, sentiments)
	if err != nil {
		fmt.Println(err)
	}

	// Insert Tweets into database
	err = InsertTweet(db, searchResult.Statuses)
	if err != nil {
		fmt.Println(err)
	}

	// Print Negative Sentiments
	negs, err := GetNegativeSentiments(db)
	if err != nil {
		fmt.Println(err)
	}
	for i, neg := range negs {
		fmt.Println("Tweet Number: ", i)
		fmt.Printf("Author: %v\nSentiment: %v \nData: %v\n\n", neg.Author, neg.Sentiment, neg.Data)
	}

}
