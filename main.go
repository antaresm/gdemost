package main

import (
	"gdemost/adminConf"
	_ "gdemost/bindatafs"
	"gdemost/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	println("Hello bridges!")

	r := mux.NewRouter()
	//r.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir("./media/"))))

	r.Handle("/", handlers.BasicAuth(handlers.MainHandler)).Methods("GET")

	r.Handle("/bridges/{id:[0-9]+}", handlers.BridgeHandler).Methods("GET")
	r.Handle("/bridges", handlers.BridgeHandler).Methods("GET")

	m := adminConf.InitAdmin()
	r.PathPrefix("/admin").Handler(m)

	http.ListenAndServe(":1235", r)
}
