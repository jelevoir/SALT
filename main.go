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
    v.Add("geocode","42.6725469,-83.2170434,5mi")
    v.Add("lang","en")
	searchResult, _ := api.GetSearch("OU", v)
//    for _, []tweet := range searchResult.Statuses {

//    }
	err = InsertTweet(db, searchResult.Statuses)
	if err != nil {
		fmt.Println(err)
	}

	var words []Word
	words = append(words, Word{0, "the", 0, 1})
	InsertWord(db, words)
}
