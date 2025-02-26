package handlers

import (
	"fmt"
	"net/http"
	"os" 
	"path/filepath"
); 

func DownloadFile(res http.ResponseWriter, req *http.Request) {
	fileName := req.URL.Query().Get("file"); 
	if fileName == ""	 {
		http.Error(res, "File is empty", http.StatusBadRequest);
		return; 
	}; 

	fileName = filepath.Base(fileName);

	path := fmt.Sprintf("storage/files/%s", fileName);
	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.Error(res, err.Error(), http.StatusNotFound); 
		return;
	}; 

	res.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", fileName))
	res.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(res, req, path);
}; 