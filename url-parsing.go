package main
import "fmt"
import "net/url"

func main() {
			var url_text = "http://google.com:8080/hal/?cari=buku&author=ilham"
			var u, err = url.Parse(url_text)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("url :", url_text)
			fmt.Println("Host :", u.Host) //google.com:8080
			fmt.Println("Path :", u.Path) // /hal/

			var cari = u.Query()["cari"][0]

			fmt.Println("cari :", cari) //buku
}
