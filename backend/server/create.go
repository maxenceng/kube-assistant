package server

import (
	"github.com/gorilla/mux"
	"kube-assistant/server/client"
	"kube-assistant/server/kubeconfig"
	"log"
	"net/http"
)

func CreateServer() {
	r := mux.NewRouter()
	addRoutes(r)
	r.Use(loggingMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func addRoutes(r *mux.Router) {
	r.HandleFunc("/kubeconfig", kubeconfig.ListKubeconfigLocations).Methods("GET")
	r.HandleFunc("/kubeconfig/file", kubeconfig.AddFileConfig).Methods("POST")
	r.HandleFunc("/kubeconfig/raw", kubeconfig.AddRawConfig).Methods("POST")
	r.HandleFunc("/client/connect", client.Connect).Methods("GET")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
