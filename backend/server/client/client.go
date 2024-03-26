package client

import (
	"github.com/gorilla/mux"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"kube-assistant/cache"
	"log"
	"net/http"
)

func Connect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	kubeconfig, ok := vars["kubeconfig"]
	if !ok {
		log.Println("Invalid parameters")
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}
	c := GetCache()
	client, err := c.Get(kubeconfig)
	if client != nil {
		log.Println("Client already in cache")
		if !client.Selected {
			log.Println("Client set as selected")
			c.Set(kubeconfig, &cache.Value[*kubernetes.Clientset]{
				Value:    client.Value,
				Selected: true,
			})
		}
		w.WriteHeader(200)
		return
	}
	log.Println("Client not in cache, fetching...")
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Printf("Could not connect to client with config %s", kubeconfig)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create the clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("Could not create client for config %s", kubeconfig)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Set(kubeconfig, &cache.Value[*kubernetes.Clientset]{
		Value:    clientSet,
		Selected: true,
	})
	w.WriteHeader(200)
}
