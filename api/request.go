package api

import (
  "io/ioutil"
  "log"
  "net/http"
)

func GetStoriesFrom(category string) {
  resp, err := http.Get("https://hacker-news.firebaseio.com/v0/" + category + ".json")
  if err != nil {
    log.Fatalln(err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }

  //Convert the body to type string
  log.Printf("%T \n",body )
}