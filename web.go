package main
import "net/http"
import "fmt"

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "apa kaabr teman!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "hello")
	})
	http.HandleFunc("/index", index)
	fmt.Println("memulai web server pada localhost:8080")
	http.ListenAndServe(":8080", nil)
}
