package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/hidelbreq/goblueprints/chapter1/trace"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.templ =
		template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8081", "アプリケーションのアドレス")
	flag.Parse()
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	go r.run()
	log.Println("Web サーバーを起動します。ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
