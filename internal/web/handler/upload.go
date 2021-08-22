package handler

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/marki-eriker/kim-worker-go/internal/web/middleware"
	"io/ioutil"
	"net/http"
	"os"
)

type fileInfo struct {
	name        string
	size        int64
	contentType string
}

type response struct {
	Hash string `json:"hash"`
}

func uploadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_, err := middleware.GetCurrentUserFromCTX(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %v", err), http.StatusUnauthorized)
			return
		}

		err = r.ParseMultipartForm(10 * 1024 * 1024)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}
		defer file.Close()

		currentFile := fileInfo{
			name:        handler.Filename,
			size:        handler.Size,
			contentType: handler.Header.Get("Content-Type"),
		}

		if currentFile.contentType != "application/pdf" {
			http.Error(w, fmt.Sprintf("error: contentType must be application/pdf"), http.StatusBadRequest)
			return
		}

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		hash := sha1.New()
		hash.Write(fileBytes)
		sha1Hash := hex.EncodeToString(hash.Sum(nil))
		filePath := fmt.Sprintf("%s/%v/%v", os.Getenv("UPLOAD_LINK"), string(sha1Hash[0]), string(sha1Hash[1]))
		fileName := fmt.Sprintf("%s.pdf", sha1Hash)

		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %v", err), http.StatusInternalServerError)
			return
		}

		fileOnDisk, err := os.Create(fmt.Sprintf("%s/%s", filePath, fileName))
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %v", err), http.StatusInternalServerError)
			return
		}
		defer fileOnDisk.Close()

		_, err = fileOnDisk.Write(fileBytes)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %v", err), http.StatusInternalServerError)
			return
		}

		response := response{Hash: sha1Hash}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
