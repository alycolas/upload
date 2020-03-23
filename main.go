package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
	"strconv"
	"html/template"
	"crypto/md5"
	"net/http"
//	"strings"
	"log"
)

const (
	noteTMPL =  "note.html"
	uploadTMPL = "upload.gtpl"
	noteFile = "1.db"
)

// func sayhelloName(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	fmt.Println(r.Form)
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Hello World!") // 这个写入到 w 的是输出到客户端的
// }

// 处理 /note 
func note(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
//		crutime := time.Now().Unix()
//		h := md5.New()
//		io.WriteString(h, strconv.FormatInt(crutime, 10))
//		token := fmt.Sprintf("%x", h.Sum(nil))
		note, _ := ioutil.ReadFile(noteFile)
		t, _ := template.ParseFiles("noteTMPL")
		t.Execute(w, note)
	} else {
		newNote = r.FormValue("note")
		err := ioutil.WriteFile(noteFile, []byte(newNote), 0666)
		if err != nil {
			fmt.Println("can not update note")
			return
		}
	}
}

// 处理 /upload
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("uploadTMPL")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("/var/www/file/upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
//	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/note", note)
	err := http.ListenAndServe("127.0.0.1:9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
