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
	http.HandleFunc("/upload", uploadHandler)

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

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 100<<20) // 100 MB max
	fmt.Println("Incomming file to upload")
	if err := r.ParseMultipartForm(100 << 20); err != nil {
		http.Error(w, "File too large or invalid form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Upload failed: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save to disk in /tmp (each upload unique filename)
	dstPath := filepath.Join(os.TempDir(), handler.Filename)
	out, err := os.Create(dstPath)
	fmt.Println(os.TempDir())
	if err != nil {
		http.Error(w, "Could not create file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	written, err := io.Copy(out, file)
	if err != nil {
		http.Error(w, "Failed to save file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Saved %d bytes to %s\n", written, dstPath)
}
