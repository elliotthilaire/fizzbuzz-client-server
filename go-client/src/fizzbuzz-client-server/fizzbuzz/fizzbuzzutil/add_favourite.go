package fizzbuzzutil

import (
  "fmt"
  "net/url"
)

func (api *Api) AddFavourite(number string) {
  safeNumber := url.QueryEscape(number)

  apiPath := fmt.Sprintf("favourites?number=%s", safeNumber)
  response := api.request("POST", apiPath, nil)

  fmt.Println(response.Status)
}
