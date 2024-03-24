package middleware

import (
	"fmt"
	"net/http"
	"p1-backend/api/pkg/config"
)

func AppKeyAuthorization(next http.Handler, config *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appKey := r.Header.Get("AppKey")

		if len(appKey) == 0 || !contains(config.AppKeys, appKey) {
			fmt.Println(r.Host + r.RequestURI + " AppKey invalido")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("AppKey invalido"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
