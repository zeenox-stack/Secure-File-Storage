package handlers 

import (
	"os" 
	"net/http" 
	"encoding/json"
); 

func GetFiles(res http.ResponseWriter, req *http.Request) {
	files, err := os.ReadDir("storage/files"); 
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError); 
		return;
	}; 

	var fileNames []string; 
	for _, file := range files {
		fileNames = append(fileNames, file.Name());
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(fileNames)
};