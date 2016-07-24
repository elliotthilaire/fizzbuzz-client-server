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

  switch response.StatusCode {
    case 201:
        fmt.Println(safeNumber, "favourited")
    case 422:
        fmt.Println(safeNumber, "already favourited")
    case 400:
        fmt.Println("you can't favourite", safeNumber)
    }
}
