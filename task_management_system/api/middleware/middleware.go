// Package middleware contains the applications' HTTP endpoints and defines how they respond to client requests
package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"task_management_system/appcontext"
	"task_management_system/constant"
	"task_management_system/util"
)

// SetJSON sets the response header content-type to JSON
func SetJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

//Recovery
/*
adding recover middleware to make sure proper response returned to
upstream in panics
*/
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("[RECOVERED]: 500 ERROR\n%+v\nStackTrace: %s", err, string(debug.Stack()))
				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func ValidateAndGenerateHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, requestIDPresent := r.Header[constant.HeaderRequestID]
		if !requestIDPresent {
			log.Println("Important headers are missing")
			ctx := r.Context()
			requestID := util.GenerateUUIDPk()
			tenant := "task_management"
			reqContext := appcontext.NewRequestContext(tenant, requestID)
			ctx = appcontext.AddRequestContext(ctx, reqContext)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		next.ServeHTTP(w, r)
	})
}
