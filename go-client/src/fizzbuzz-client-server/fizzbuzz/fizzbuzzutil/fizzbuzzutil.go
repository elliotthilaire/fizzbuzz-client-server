package fizzbuzzutil

import (
  "fmt"
  "net/http"
  "log"
  "bytes"
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
func (api *Api) request(action string, apiPath string, requestBody string) (*http.Response)  {

  url := api.url(apiPath)
  requestBodyBytes := []byte(requestBody)

  request, err := http.NewRequest(action, url, bytes.NewBuffer(requestBodyBytes))
  if err != nil {
    log.Fatal("NewRequest: ", err)
    //return
  }

  request.Header.Add("Content-Type", "application/json")
  request.Header.Add("Accept", "application/json")

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

  defer response.Body.Close()

  return response
}


