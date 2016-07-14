package main

import (
	"fmt"
)

func main() {
	twitterconfig, _ := LoadTwitterConfig("twitterconfig.json")
	api := NewTwitter(twitterconfig)

	postgresqlconfig, _ := LoadPostgreSQLConfig("postgresql.json")
	db, err := NewDatabase(postgresqlconfig)
	if err != nil {
		fmt.Println(err)
	}

	searchResult, _ := api.GetSearch("golang", nil)
	err = InsertTweet(db, searchResult.Statuses)
	if err != nil {
		fmt.Println(err)
	}

	var words []Word
	words = append(words, Word{0, "the", 0, 1})
	InsertWord(db, words)
}
