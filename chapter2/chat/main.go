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
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.templ =
		template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	data := map[string]interface{}{
		"Host": r.Host,
	}

	if authCo0kie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCo0kie.Value)
	}
	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8081", "アプリケーションのアドレス")
	flag.Parse()
	gomniauth.SetSecurityKey("")
	gomniauth.WithProviders(
		google.New(
			"",
			"",
			"",
		),
	)
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
