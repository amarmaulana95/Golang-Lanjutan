package user

// import "encoding/json"
// import "net/http"
import "fmt"
import k "belajargolang/koneksi"

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

func Ambil_data() ([]pelajar, error){
  db, err := k.Connect()

  if err != nil {
    fmt.Println(err.Error())
    return nil, err
  }
  defer db.Close()

  rows, err := db.Query("select * from user")
  if err != nil {
    fmt.Println(err.Error())
    return nil, err
  }
  defer rows.Close()

  for rows.Next(){
    var each = pelajar{}
    var err = rows.Scan(&each.Id, &each.Name, &each.Password, &each.Token, &each.Created_at, &each.Updated_at, &each.Status, &each.Image)

    if err != nil {
      fmt.Println(err.Error())
      return nil, err
    }
    data_pelajar = append(data_pelajar, each)
  }

  if err = rows.Err(); err != nil {
    fmt.Println(err.Error())
    return nil, err
  }
  /*
  for _, each := range data_pelajar{
		fmt.Println(each.name)
	}
  */
  return data_pelajar, nil
}
