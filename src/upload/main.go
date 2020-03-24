package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
	//"strings"
	"flag"
	"log"
	"templ"
)

var (
	noteFile  string
	uploadDir string
	local     string
	port      string
)

func init() {
	flag.StringVar(&noteFile, "n", "/var/www/note.db", "指定 note 文件")
	flag.StringVar(&uploadDir, "u", "/var/www/file/upload/", "指定上传目录")
	flag.StringVar(&local, "l", "127.0.0.1", "指定监听地址")
	flag.StringVar(&port, "p", "9090", "指定端口")
}

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
		note, _ := ioutil.ReadFile(noteFile)
		t, _ := template.New("note").Parse(templ.NoteTemp)
		t.Execute(w, string(note))
	}
}
func save(w http.ResponseWriter, r *http.Request) {
	newNote := r.FormValue("note")
	err := ioutil.WriteFile(noteFile, []byte(newNote), 0666)
	if err != nil {
		fmt.Println("can not update note")
		return
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

		t, _ := template.New("upload").Parse(templ.UploadTemp)
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
		f, err := os.OpenFile(uploadDir+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	flag.Parse()
	//	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/", note)
	http.HandleFunc("/save", save)
	err := http.ListenAndServe(local+":"+port, nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
