package fizzbuzzutil

import (
  "fmt"
  "encoding/json"
  "log"
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
  response := api.request("GET", apiPath, nil)

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


