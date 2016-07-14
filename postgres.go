package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/lib/pq"
)

const CREATE_TWEET_TABLE string = "create table tweet(id serial primary key, data varchar, timestamp varchar, author varchar, favoritecount integer, retweetcount integer)"

const CREATE_WORD_TABLE string = "create table word(id serial primary key, word varchar, value integer, count integer)"

type PostgresConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Dbname   string
}

type Word struct {
	Id    int64
	Word  string
	Value int64
	Count int64
}

func LoadPostgreSQLConfig(filename string) (PostgresConfig, error) {
	pgc := PostgresConfig{}

	file, err := os.Open(filename)
	if err != nil {
		return pgc, err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&pgc)
	if err != nil {
		return pgc, err
	}

	return pgc, nil
}

func InsertTweet(db *sql.DB, data []anaconda.Tweet) error {
	txn, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := txn.Prepare(pq.CopyIn("tweet", "data", "timestamp", "author", "favoritecount", "retweetcount"))
	if err != nil {
		return err
	}

	for _, d := range data {
		_, err = stmt.Exec(d.Text, d.CreatedAt, d.User.ScreenName, d.FavoriteCount, d.RetweetCount)
		if err != nil {
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	return nil
}

func InsertWord(db *sql.DB, data []Word) error {
	txn, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := txn.Prepare(pq.CopyIn("word", "word", "value", "count"))
	if err != nil {
		return err
	}

	for _, d := range data {
		_, err = stmt.Exec(d.Word, d.Value, d.Count)
		if err != nil {
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	return nil
}

func NewDatabase(config PostgresConfig) (*sql.DB, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", config.Username, config.Password, config.Host, config.Dbname)
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	db.Query(CREATE_TWEET_TABLE)
	db.Query(CREATE_WORD_TABLE)

	return db, nil
}
