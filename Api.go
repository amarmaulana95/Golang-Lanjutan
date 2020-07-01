package main

import "encoding/json"
import "net/http"
import "fmt"
import "regexp"
import "reflect"
/*
import "database/sql"
import _ "mysql-master"
*/
import k "belajargolang/koneksi"
import u "belajargolang/model"

type user struct{
  ID int
  Name string
  Status string
}

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
var data_pelajar []pelajar
/*
var data_user = []user{
  user{1, "ppppppppppp", "admin"},
  user{2, "nnnn nnnnn nnnnnnn", "admin"},
}
*/
var data_user []user
/*
func koneksi() (*sql.DB, error){
  db, err := sql.Open("mysql", "root:passw@tcp(127.0.0.1:3306)/post")
  if err != nil {
    return nil, err
  }
  return db, nil
}
*/
func ambil_data2(){
  data, err := u.Ambil_data()

  if err != nil {
    fmt.Println(err.Error())
    return
  }
  var reflectdata = reflect.ValueOf(data)

  fmt.Println("type variabel: ", reflectdata.Type())
}

func ambil_data(){
  db, err := k.Connect()

  if err != nil {
    fmt.Println(err.Error())
    return
  }
  defer db.Close()

  rows, err := db.Query("select * from user")
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  defer rows.Close()

  for rows.Next(){
    var each = pelajar{}
    var err = rows.Scan(&each.Id, &each.Name, &each.Password, &each.Token, &each.Created_at, &each.Updated_at, &each.Status, &each.Image)

    if err != nil {
      fmt.Println(err.Error())
      return
    }
    data_pelajar = append(data_pelajar, each)
  }
  if err = rows.Err(); err != nil {
    fmt.Println(err.Error())
    return
  }
  /*
  for _, each := range data_pelajar{
		fmt.Println(each.name)
	}
  */
}

func search_user(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  if r.Method == "POST" {
    var namauser = r.FormValue("Name")
    var result []byte
    var err error
    var regex, _ = regexp.Compile(`(?i:`+namauser+`)`)
    var tampilkan []pelajar

    for _, each := range data_pelajar {

      if regex.MatchString(each.Name) { //each.Name == namauser
        tampilkan = append(tampilkan, each)
        /*
        result, err = json.Marshal(each)

        if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }
        fmt.Println("match: ",each)
        w.Write(result)
        return
        */
      }
    }
    result, err = json.Marshal(tampilkan)

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    if len(tampilkan) > 0 {
      w.Write(result)
      return
    }

    http.Error(w, "User tidak ditemukan", http.StatusBadRequest)
    return
  }
  http.Error(w, "", http.StatusBadRequest)
}

func get_users(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  if r.Method == "POST"{
    var result, err = json.Marshal(data_pelajar)

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Write(result)
    return
  }
  http.Error(w, "", http.StatusBadRequest)
}

func main(){
  ambil_data()
  http.HandleFunc("/api/users", get_users)
  http.HandleFunc("/api/search/user", search_user)

  fmt.Println("Menjalankan Web Server pada localhost: 8080")
  http.ListenAndServe(":8080", nil)
}
