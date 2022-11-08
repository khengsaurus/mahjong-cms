package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/khengsaurus/mahjong-cms/controllers"
	"github.com/khengsaurus/mahjong-cms/middlewares"
)

const defaultPort = "8080"

var (
	route_test = "/test"
	route_mj   = "/mj"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(middlewares.EnableCors)

	router.HandleFunc(route_test, test)
	router.Route(route_mj, controllers.MahjongContent)
	router.HandleFunc("/*", others)

	fmt.Printf("Server listening at port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		panic(err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))
}

func others(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
