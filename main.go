package main

import (
	"fmt"
)

func main() {
	config, _ := LoadConfig("twitterconfig.json")
	api := NewTwitter(config)

	searchResult, _ := api.GetSearch("golang", nil)
	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}
}
