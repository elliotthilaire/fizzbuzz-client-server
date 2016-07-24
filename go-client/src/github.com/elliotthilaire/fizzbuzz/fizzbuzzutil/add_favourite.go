package fizzbuzzutil

import (
  "fmt"
  "log"
  "net/http"
  "net/url"
)

func (api *Api) AddFavourite(number string) {
  safeNumber := url.QueryEscape(number)

  apiPath := fmt.Sprintf("favourites?number=%s", safeNumber)
  url := api.url(apiPath)

  request, err := http.NewRequest("POST", url, nil)
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

  fmt.Println(response.Status)
}
