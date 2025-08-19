package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var fileMap = map[string]string{
	"5kb":  "5kb.txt",
	"10kb": "10kb.txt",
	"25kb": "25kb.txt",
	"50kb": "50kb.txt",
	"100kb": "100kb.txt",
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/stream/{file}", streamHandler)

	fmt.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("ok"))
}

func streamHandler(w http.ResponseWriter, req *http.Request) {
	fileKey := req.PathValue("file")

	if fileKey == "" {
		http.Error(w, "Missing 'file' parameter", http.StatusBadRequest)
		return
	}

	filePath, ok := fileMap[fileKey]
	if !ok {
		http.Error(w, "Invalid file requested. Valid options are: 1mb, 5mb, 10mb, 25mb, 50mb", http.StatusBadRequest)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found.", http.StatusNotFound)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		http.Error(w, "Could not obtain file info.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "inline; filename=\""+filepath.Base(filePath)+"\"")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))

	_, err = io.Copy(w, file)
	if err != nil {
		fmt.Println("Error streaming file:", err)
	}
}
