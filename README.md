# S.A.L.T.
Sentiment Analysis on Location-based Tweets

**Description**  
This program connects to both the Twitter API and Google Cloud's Natural Language API.

First it connects to the Twitter API and acquires 15 tweets containing a keyword and within a specified geo-location.  
Then it sends the tweet to Google Cloud's Natural Language API and returns the sentiment analysis.  
Finally it prints the username, sentiment score, and the tweet for any of the 15 tweets with a sentiment score of 0 or less.

## Prerequisites
1. `go get` these packages:  
* github.com/ChimeraCoder/anaconda - This package is a "Go client library for the Twitter 1.1 API"
* github.com/lib/pq - This package is a "Pure Go Postgres driver for database/sql"

2. Register an app at [dev.twitter](https://dev.twitter.com/apps) to get your ConsumerKey and ConsumerSecret.

3. Create an account with the Google Cloud Platform to get a Natural Language API Key.

## Set-Up
### 1 
Create an environment variable:

`NATLANG_API_KEY`

Set that variable to the API key generated for you by Google.

Generate one here:
[Google Cloud Platform](https://console.cloud.google.com/apis/credentials/wizard?api=language.googleapis.com)

### 2 
Create postgresql.json  
Use the template at postgresql.json-sample.

### 3 
Create twitterconfig.json  
Use the template at twitterconfig.json-sample.

## Using SALT
You have two options. You can either run or build.

go run/build main.go postgres.go twitter.go natlang.go
