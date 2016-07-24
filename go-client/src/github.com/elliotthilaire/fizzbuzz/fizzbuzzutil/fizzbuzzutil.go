package fizzbuzzutil

import (
  "fmt"
  "io"
  "net/http"
  "log"

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

// request wraps creating and running a http client
func (api *Api) request(action string, apiPath string, body io.Reader) (*http.Response)  {

  url := api.url(apiPath)

  request, err := http.NewRequest(action, url, nil)
  if err != nil {
    log.Fatal("NewRequest: ", err)
    //return
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
    //return
  }

  return response
}


