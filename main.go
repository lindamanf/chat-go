package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templは1つのテンプレートを指す
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// HTTPリクエストの処理
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() // フラグを解釈する

	r := newRoom()

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// チャットルームの開始
	go r.run()
	// Webサーバーを開始
	log.Println("Webサーバを開始します。ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
