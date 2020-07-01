package koneksi

import "database/sql"
import _ "mysql-master"

func Connect() (*sql.DB, error){
  db, err := sql.Open("mysql", "root:AmikBsi12119617@tcp(127.0.0.1:3306)/post")
  if err != nil {
    return nil, err
  }
  return db, nil
}
