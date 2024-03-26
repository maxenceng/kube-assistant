package kubeconfig

import (
	"encoding/json"
	"kube-assistant/server/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type KubeconfigLocation struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type RawConfig struct {
	Content string `json:"content"`
}

type FileConfig struct {
	Location string `json:"location"`
}

func ListKubeconfigLocations(w http.ResponseWriter, r *http.Request) {
	var kubeconfigLocations []KubeconfigLocation
	localKubeconfigLocation, err := localKubeconfigLocation()
	if err != nil {
		kubeconfigLocations = append(kubeconfigLocations, *localKubeconfigLocation)
	}
	assistantKubeconfigLocations, err := assistantKubeconfigLocations()
	if err != nil {
		kubeconfigLocations = append(kubeconfigLocations, *assistantKubeconfigLocations...)
	}
	jsonResponse, err := json.Marshal(kubeconfigLocations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func localKubeconfigLocation() (*KubeconfigLocation, error) {
	localFile := utils.LocalKubeconfigFile()
	return getKubeconfigLocation(localFile)
}

func assistantKubeconfigLocations() (*[]KubeconfigLocation, error) {
	dir := utils.AssistantDir()
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("Could not read directory %s", dir)
		return nil, err
	}
	var kubeconfigLocations *[]KubeconfigLocation
	for _, file := range files {
		location := filepath.Join(dir, file.Name())
		kubeconfigLocation, err := getKubeconfigLocation(location)
		if err != nil {
			*kubeconfigLocations = append(*kubeconfigLocations, *kubeconfigLocation)
		}
	}
	return kubeconfigLocations, nil
}

func getKubeconfigLocation(file string) (*KubeconfigLocation, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Printf("Could not read file %s", file)
		return nil, err
	}
	kubeconfig, err := utils.UnmarshalKubeconfig(content)
	if err != nil {
		log.Printf("Could not unmarshal %s", content)
		return nil, err
	}
	return &KubeconfigLocation{
		Name:     kubeconfig.Clusters[0].Name,
		Location: file,
	}, nil
}

func AddRawConfig(w http.ResponseWriter, r *http.Request) {
	var rawConfig RawConfig
	err := json.NewDecoder(r.Body).Decode(&rawConfig)
	if err != nil {
		log.Println("Malformed request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dest := filepath.Join(utils.AssistantDir(), "config"+utils.RandomString())
	err = os.WriteFile(dest, []byte(rawConfig.Content), 0644)
	if err != nil {
		log.Printf("Could not write file to %s", dest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(201)
}

func AddFileConfig(w http.ResponseWriter, r *http.Request) {
	var fileConfig FileConfig
	err := json.NewDecoder(r.Body).Decode(&fileConfig)
	if err != nil {
		log.Println("Malformed request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	content, err := os.ReadFile(fileConfig.Location)
	if err != nil {
		log.Printf("Could not read file %s", fileConfig.Location)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dest := filepath.Join(utils.AssistantDir(), "config"+utils.RandomString())
	err = os.WriteFile(dest, content, 0644)
	if err != nil {
		log.Printf("Could not write file to %s", dest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(201)
}
