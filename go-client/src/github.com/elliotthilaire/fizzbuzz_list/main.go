
package main

import (
  "fmt"
  "encoding/json"
  "log"
  "net/http"
  "net/url"
)

func main() {
  page := "1"
  per_page := "100"

  fetch_fizzbuzzes(page, per_page)
}


func fetch_fizzbuzzes(page string, per_page string ) {
  safePage := url.QueryEscape(page)
  safePerPage := url.QueryEscape(per_page)

  type Fizzbuzz struct {
    Number int `json:"number"`
    Output string `json:"output"`
    Favourite bool `json:"favourite"`
  }

  url := fmt.Sprintf("http://localhost:3000/api/v1/fizzbuzzes.json?page=%s&per_page=%s", safePage, safePerPage)

  request, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal("NewRequest: ", err)
    return
  }

  // For control over HTTP client headers,
  // redirect policy, and other settings,
  // create a Client
  // A Client is an HTTP client
  client := &http.Client{}

  // Send the request via a client
  // Do sends an HTTP request and
  // returns an HTTP response
  response, err := client.Do(request)
  if err != nil {
    log.Fatal("Do: ", err)
    return
  }

  // Callers should close resp.Body
  // when done reading from it
  // Defer the closing of the body
  defer response.Body.Close()

  var fizzbuzzes []Fizzbuzz

  if err := json.NewDecoder(response.Body).Decode(&fizzbuzzes); err != nil {
    log.Println(err)
    return
  }

  for _,element := range fizzbuzzes {
    favourite := emojify(element.Favourite)
    fmt.Println(element.Number, ": ", element.Output, favourite)
  }

}

func emojify(value bool) (string) {
  emoji := ""
  if value { emoji = "\U0001f370" }
  return emoji
}


