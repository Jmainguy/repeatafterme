package main

import (
    "github.com/dghubble/go-twitter/twitter"
    _ "github.com/mattn/go-sqlite3"
    "github.com/dghubble/oauth1"
    "github.com/ghodss/yaml"
    "io/ioutil"
    "time"
    "fmt"
    "log"
)

type Config struct {
    Consumer_key string `json:"consumer_key"`
    Consumer_secret string `json:"consumer_secret"`
    Access_token string `json:"access_token"`
    Access_secret string `json:"access_secret"`
    Follow []string
}

func config()  (consumer_key,consumer_secret,access_token,access_secret string, follow []string ){
    var v Config
    config_file, err := ioutil.ReadFile("/etc/repeatafterme/config.yaml")
    if err != nil {
        log.Fatal(err)
    }
    yaml.Unmarshal(config_file, &v)
    consumer_key = v.Consumer_key
    consumer_secret = v.Consumer_secret
    access_token = v.Access_token
    access_secret = v.Access_secret
    follow = v.Follow
    return
}

func check(e error) {
    if e != nil {
        fmt.Println(e)
    }
}

func likeTweet(client *twitter.Client, tweetID int64){
    _, _, err := client.Favorites.Create(&twitter.FavoriteCreateParams{
        ID: tweetID,
    })
    check(err)
}

func reTweet(client *twitter.Client, tweetID int64){
    _, _, err := client.Statuses.Retweet(tweetID, &twitter.StatusRetweetParams{
        ID: tweetID,
        TrimUser: twitter.Bool(true), 
    })
    check(err)
}

func main() {
    // For loop to continue into infinity, this will keep daemon up and running forever
    for {
        consumer_key,consumer_secret,access_token,access_secret,follow := config()
        configure := oauth1.NewConfig(consumer_key, consumer_secret)
        token := oauth1.NewToken(access_token, access_secret)
        httpClient := configure.Client(oauth1.NoContext, token)

        // Twitter client
        client := twitter.NewClient(httpClient)
        // open database, and create table if needed
        db := InitDB("/opt/repeatafterme/repeatafterme.db")
        CreateTable(db)

        // Loop over the usernames in the config file, and look up their tweets since last time.
        for username := range follow {
            user, _, _ := client.Users.Show(&twitter.UserShowParams{
                ScreenName: follow[username],
            })
            // What was the last tweet we have a record of?
            lastTweet := ReadItem(db, follow[username])
            // What has that user tweeted since last time?
            tweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
                UserID: user.ID, 
                TrimUser: twitter.Bool(true), 
                IncludeRetweets: twitter.Bool(false),
                ExcludeReplies: twitter.Bool(true),
                SinceID: lastTweet,
            })
            // If we dont have a record of their last tweet, only like / retweet their latest tweet.
            if lastTweet == 0 {
                // Like it
                likeTweet(client, tweets[0].ID)
                // Retweet it
                reTweet(client, tweets[0].ID)
                // Print when you do stuff, so we can follow along
                fmt.Printf("@%v said:\n", follow[username])
                fmt.Println(tweets[0].Text)
            } else {
                // Loop over their tweets and like / retweet each one
                for tweet := range tweets {
                    // Like it
                    likeTweet(client, tweets[tweet].ID)
                    // Retweet it
                    reTweet(client, tweets[tweet].ID)
                    // Print when you do stuff, so we can follow along
                    fmt.Printf("@%v said:\n", follow[username])
                    fmt.Println(tweets[tweet].Text)
                }
            }
            // If there are any tweets, store the latest tweetid in sqlite so we dont retweet / like it next time.
            if len(tweets) > 0 {
               items := []TestItem{
                    TestItem{follow[username], tweets[0].ID},
               }
               StoreItem(db, items)
            }
        }
        db.Close()
        fmt.Println("Loop complete")
        // Sleep for a minute, so that the bot doesnt go crazy and eat your cpu
        time.Sleep(60 * time.Second)
    }
}
