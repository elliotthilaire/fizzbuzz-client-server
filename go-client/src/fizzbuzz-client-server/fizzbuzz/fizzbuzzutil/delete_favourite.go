package fizzbuzzutil

import (
  "fmt"
  "net/url"
)

func (api *Api) DeleteFavourite(number string) {
  safeNumber := url.QueryEscape(number)

  apiPath := fmt.Sprintf("favourites/%s", safeNumber)
  response := api.request("DELETE", apiPath, nil)

  fmt.Println(response.Status)
}

