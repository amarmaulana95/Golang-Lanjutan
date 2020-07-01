package main
import ("fmt"; "database/sql")
import _ "mysql-master"

type pelajar struct{
	id int
	name string
  password string
  token string
  created_at string
  updated_at string
	status string
	image string
}

func koneksi() (*sql.DB, error){
	db, err := sql.Open("mysql", "root:AmikBsi12119617@tcp(127.0.0.1:3306)/post") // namadb ny
	if err != nil {
		return nil,err
	}
	return db, nil
}

func sql_tampil(){
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
  var status_user = "admin"
	rows, err := db.Query("select * from user where status = ?", status_user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []pelajar

	for rows.Next() {
		var each = pelajar{}
		var err = rows.Scan(&each.id, &each.name, &each.password, &each.token, &each.created_at, &each.updated_at, &each.status, &each.image)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result{
		fmt.Println(each.name)
	}
}

func sql_tambah() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into user values (?,?,?,?,?,?,?,?)", nil, "Nida shofa", "nidashoga_", "token1234", "12-2-2222", "12111111", "admin", "nida.png")

	if err != nil {
		fmt.Println(err.Error())
    return
	}

	fmt.Println("tambah data berhasil")
}

func ubah_data(){
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("update user set password = ? where id = ?", "NidaShofa_", 17)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ubah berhasil!")
}

func hapus_data(){
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("delete from user where id = ?", 18)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("del berhasil!")
}

func main() {
  hapus_data()
  // ubah_data()
	sql_tampil()
  // sql_tambah()
}
