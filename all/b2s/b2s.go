package b2s

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func StartService(config Config) {
	log.Println("start b2s service ...")

	http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/static/", staticHandler)
	s := http.Server{Addr: fmt.Sprintf(":%d", config.Port)}
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	s.Serve(ln)
	// http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) > 1 {
		http.ServeFile(w, r, fmt.Sprintf("static/views/%s.html",r.URL.Path))
	} else {
		http.ServeFile(w, r, "static/views/index.html")
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/views/users.html")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/login/ - path:", r.URL.Path)
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/views/login.html")
	} else {
		name := r.FormValue("name")
		room := r.FormValue("room")
		log.Printf("name:%s, room:%s", name, room)
		http.Redirect(w, r, "/?name="+name+"&room="+room, http.StatusFound)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/static/ - path:", r.URL.Path)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	http.FileServer(http.Dir(wd)).ServeHTTP(w, r)
}
