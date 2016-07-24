package fizzbuzzutil

import (
  "fmt"
  "net/url"
)

func (api *Api) DeleteFavourite(number string) {
  safeNumber := url.QueryEscape(number)

  apiPath := fmt.Sprintf("favourites/%s", safeNumber)
  response := api.request("DELETE", apiPath, "")

  switch response.StatusCode {
    case 204:
        fmt.Println(safeNumber,"unfavourited")
    case 404:
        fmt.Println(safeNumber, "already unfavourited")
    case 400:
        fmt.Println("you can't unfavourite", safeNumber)
    }
}

