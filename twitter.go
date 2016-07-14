package main

import (
	"encoding/json"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

type TwitterConfig struct {
	ConsumerKey    string
	ConsumerSecret string
}

func LoadConfig(filename string) (TwitterConfig, error) {
	tc := TwitterConfig{}

	file, err := os.Open(filename)
	if err != nil {
		return tc, err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tc)
	if err != nil {
		return tc, err
	}

	return tc, nil
}

func NewTwitter(config TwitterConfig) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	api := anaconda.NewTwitterApi("", "")

	return api
}
