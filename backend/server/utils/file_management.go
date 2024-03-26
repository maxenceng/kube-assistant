package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

func homeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func LocalKubeconfigFile() string {
	return filepath.Join(homeDir(), ".kube", "config")
}

func AssistantDir() string {
	return filepath.Join(homeDir(), ".kube-assistant")
}
