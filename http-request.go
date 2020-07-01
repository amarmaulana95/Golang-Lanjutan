package main
import (
  "encoding/json"
  "net/http"
  "fmt"
  "bytes"
  "net/url"
)

var baseURL = "http://localhost:8080"

type pelajar struct{
	Id int
	Name string
  Password string
  Token string
  Created_at string
  Updated_at string
	Status string
	Image string
}

func ambil_api(user string)([]pelajar, error){
  var err error
  var client = &http.Client{}
  var data []pelajar

  var param = url.Values{}
  param.Set("Name", user)
  var payload = bytes.NewBufferString(param.Encode())
  request, err := http.NewRequest("POST", baseURL + "/api/search/user", payload)

  if err != nil {
    return data, err
  }
  request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  response, err := client.Do(request)
  if err != nil {
    return nil, err
  }
  defer response.Body.Close()

  err = json.NewDecoder(response.Body).Decode(&data)
  if err != nil {
    return nil, err
  }

  return data, nil

}

func main(){
  var user, err = ambil_api("Shofa")
  if err != nil {
    fmt.Println("error!", err.Error())
    return
  }

  for _, each := range user {
    fmt.Println(each)
  }
}
