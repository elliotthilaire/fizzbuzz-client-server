package fizzbuzzutil

import (
  "fmt"
)

type Api struct {
  Host string
  Port string
}


func (api *Api) url(apiPath string) string {
  baseUrl := fmt.Sprintf("http://%s:%s/api/v1", api.Host, api.Port)
  url := fmt.Sprintf("%s/%s", baseUrl, apiPath)
  return url
}
