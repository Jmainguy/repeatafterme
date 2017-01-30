# RepeatAfterMe
Retweet / Like anything a list of people post (exclude replies and retweets)
## Initial requirements
  1. bot that runs as a service
  2. interacts with twitter
  3. you supply yaml list of accounts to retweet
  4. it retweets / likes anything those accounts say
  5. write to a local sqlite3 db the tweetid of the tweet so it doesnt repeat retweets

## Purpose
The imagined use case is for people to follow political accounts and retweet / like everything they do to get the message out. Or just seem like a really obessesed friend that is stalking their buddy, I don't care, I am not your Dad.

### How to use.
#### Twitter Api AccessSetup twitter api access
  1. Follow [This](https://themepacific.com/how-to-generate-api-key-consumer-token-access-key-for-twitter-oauth/994/) to setup your api access with twitter.
  2. Grab the Consumer Key, Consumer Secret, Access Token, and Access Token Secret to be used by the service. For the Name / Description / Website put w/e you feel like. I used this github url for the website.

#### RPM
  1. Either compile yourself (read the Makefile to see odd setup I use, then run make rpm if you setup the same way, or come up with a more universal way and submit that as a PR)
  2. Or just yum install from this repo where I uploaded to.
  3. Edit /etc/repeatafterme/config.yaml and add your credentials from the API section above, and a list of users to follow.
  4. `systemctl start repeatafterme`

#### DEB
  1. I am super lazy, I don't use ubuntu / debian for much, if someone wants this packaged up in a DEB let me know and I can make that happen.
  2. Until then, follow the Golang section.


#### Golang
  1. run `go build` from inside the dir
  2. Copy the resulting binary somewhere like /usr/sbin/repeatafterme
  3. `mkdir -p /opt/repeatafterme`
  4. `mkdir /etc/repeatafterme`
  5. Edit /etc/repeatafterme/config.yaml and fill in all the details
  6. Run /usr/sbin/repeatafterme, if on a persistant server run it in screen / tmux to keep it going until the server dies.

#### Docker
  1. Run `make docker`
  2. go into the docker dir, and run `./build.sh`
  3. Edit the config.yaml file there (it gets mounted inside the container)
  4. then run with `./run.sh`
