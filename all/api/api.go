package api

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/lovebzhou/goprojs/all/api/model"
)

func StartService(config Config) {
	log.Println("start api service ...")

	users := model.Users{}
	users.Store = make(map[string]*model.User)
	users.Store["0"] = &model.User{ID: "0", Name: "张三", Age: 20}
	users.Store["1"] = &model.User{ID: "1", Name: "李四", Age: 30}
	users.Store["2"] = &model.User{ID: "2", Name: "王二麻子", Age: 40}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/users", users.GetAllUsers),
		rest.Post("/users", users.PostUser),
		rest.Get("/users/:id", users.GetUser),
		rest.Put("/users/:id", users.PutUser),
		rest.Delete("/users/:id", users.DeleteUser),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	// server := http.Server{Addr: fmt.Sprintf(":%d", config.Port),
	// 	Handler: api.MakeHandler(),
	// }
	// log.Fatal(server.ListenAndServe())

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	// http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
