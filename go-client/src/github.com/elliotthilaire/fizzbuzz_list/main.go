package main

import (
  "fmt"
  "encoding/json"
  "strings"
  "log"
)

func main() {

  type Fizzbuzz struct {
    Number int `json:"number"`
    Output string `json:"output"`
    Favourite bool `json:"favourite"`
  }

  response := `[
    {
      "number": 1,
      "output": "1",
      "favourite": true
    },
    {
      "number": 2,
      "output": "2",
      "favourite": false
    },
    {
      "number": 2,
      "output": "Fizz",
      "favourite": true
    }]`

    var fizzbuzzes []Fizzbuzz

    if err := json.NewDecoder(strings.NewReader(response)).Decode(&fizzbuzzes); err != nil {
      log.Println(err)
    }

    for _,element := range fizzbuzzes {
      fmt.Println(element.Number, ": ", element.Output)
    }
}
