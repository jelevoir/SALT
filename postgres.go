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
const CREATE_SENTIMENT_TABLE string = "create table sentiment(id serial primary key, sentiment real, magnitude real, author varchar)"

type PostgresConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Dbname   string
}

type Sentiments struct {
    Id           int64
    SentimentVal float64
    MagnitudeVal float64
    Author       string
}

type NegativeSentiments struct {
    Author    string
    Sentiment float64
    Data      string
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

func InsertSentiment(db *sql.DB, data []Sentiments) error {
    txn, err := db.Begin()
    if err != nil {
        return err
    }

    stmt, err := txn.Prepare(pq.CopyIn("sentiment", "sentiment", "magnitude", "author"))
    if err != nil {
        return err
    }

    for _, d := range data {
        _, err = stmt.Exec(d.SentimentVal, d.MagnitudeVal, d.Author)
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

func GetNegativeSentiments(db *sql.DB) ([]NegativeSentiments, error) { 
    rows, err := db.Query("select tweet.author, sentiment.sentiment, tweet.data from tweet left join sentiment on tweet.author = sentiment.author where sentiment.sentiment<=0")
    if err != nil {
        return nil, err
    }

    var negsents []NegativeSentiments

    defer rows.Close()
    for rows.Next() {
        var author    string
        var data      string
        var sentiment float64
        err := rows.Scan(&author, &sentiment, &data)
        if err != nil {
            return nil, err
        }
        negsents = append(negsents, NegativeSentiments{author, sentiment, data})
    }
    err = rows.Err()
    if err != nil {
        return nil, err
    }
    return negsents, nil
}


func NewDatabase(config PostgresConfig) (*sql.DB, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", config.Username, config.Password, config.Host, config.Dbname)
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	db.Query(CREATE_TWEET_TABLE)
    db.Query(CREATE_SENTIMENT_TABLE)

	return db, nil
}
