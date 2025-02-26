package handlers

import (
	"fmt"
	"net/http"
	"os" 
	"path/filepath"
); 

func DeleteFile(res http.ResponseWriter, req *http.Request) {
	fileName := req.URL.Query().Get("file"); 
	if fileName == "" {
		http.Error(res, "File name is empty", http.StatusBadRequest); 
		return;
	}; 

	fileName = filepath.Base(fileName);

	path := fmt.Sprintf("storage/files/%s", fileName); 
	if err := os.Remove(path); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError); 
		return;
	};
}