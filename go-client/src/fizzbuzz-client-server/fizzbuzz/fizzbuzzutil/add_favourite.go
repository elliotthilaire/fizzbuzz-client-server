package fizzbuzzutil

import (
  "fmt"
  "net/url"
)

func (api *Api) AddFavourite(number string) {
  safeNumber := url.QueryEscape(number)

  apiPath := fmt.Sprintf("favourites")
  requestBody := fmt.Sprintf(`{"number":%s}`, safeNumber)

  response := api.request("POST", apiPath, requestBody)

  fmt.Println(response.Status)
}
