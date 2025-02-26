package middleware

import (
	"net/http"
	"os"
); 

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if key := req.Header.Get("X-Api-Key"); key != os.Getenv("KEY") {
			http.Error(res, "Forbidden", http.StatusForbidden); 
			return;
		}; 

		next.ServeHTTP(res, req);
	})
};
