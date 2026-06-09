package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(500 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "ошибка получения файла: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileType := r.URL.Query().Get("type")
	if fileType == "" {
		http.Error(w, "не указан тип файла", http.StatusBadRequest)
		return
	}

	uploadDir := fmt.Sprintf("/upload/%s", fileType)
	os.Mkdir(uploadDir, os.ModePerm)

	filename := filepath.Base(handler.Filename)
	filename = strings.ReplaceAll(filename, " ", "_")

	savePath := filepath.Join(uploadDir, filename)
	dst, err := os.Create(savePath)
	if err != nil {
		http.Error(w, "ошибка создания файла: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		http.Error(w, "ошибка сохранения файла: "+err.Error(), http.StatusInternalServerError)
		return
	}

	publicPath := fmt.Sprintf("/upload/%s/%s", fileType, filename)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"path":"%s"}`, publicPath)
}
