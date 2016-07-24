package fizzbuzzutil

import (
  "fmt"
  "encoding/json"
  "log"
  "net/http"
  "net/url"
  "strconv"
)

func (api *Api) List(page int, per_page int) {

  pageAsString := strconv.Itoa(page)
  perPageAsString := strconv.Itoa(per_page)

  safePage := url.QueryEscape(pageAsString)
  safePerPage := url.QueryEscape(perPageAsString)

  type Fizzbuzz struct {
    Number int `json:"number"`
    Output string `json:"output"`
    Favourite bool `json:"favourite"`
  }

  apiPath := fmt.Sprintf("fizzbuzzes?page=%s&per_page=%s", safePage, safePerPage)
  url := api.url(apiPath)

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


