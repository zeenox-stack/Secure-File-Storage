package handlers 

import (
	"fmt" 
	"os" 
	"net/http" 
	"io"
); 

func UploadFile(res http.ResponseWriter, req *http.Request) {
	files, err := os.ReadDir("storage/files"); 
	if err != nil || len(files) >= 10 {
		http.Error(res, "Limit Reached", http.StatusBadRequest); 
		return;
	}; 

	req.ParseMultipartForm(10 << 20); 

	file, handler, err := req.FormFile("file"); 
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError); 
		return;
	}; 
	defer file.Close();
    
	dest, err := os.Create(fmt.Sprintf("storage/files/%s", handler.Filename)); 
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError); 
		return;
	}; 
	defer dest.Close();

	if _, err := io.Copy(dest, file); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(res, "File saved successfully");
}