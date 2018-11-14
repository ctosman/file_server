package main 

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

//curl -F "uploaded=@test" "127.0.0.1:8888/"

func savefile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(50 << 20)//50M
		file, handler, err := r.FormFile("filename")
		if err != nil {
			fmt.Println("get file name failed")
			return;
		}
		defer file.Close()

		f, err := os.OpenFile("/home/file/" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Println("create file failed")
			return;
		}
		defer f.Close()
		io.Copy(f, file)
		fmt.Fprintln(w, "200 success")
	}
}

func main() {
	http.HandleFunc("/", savefile)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("Server listen failed")
	}
} 
