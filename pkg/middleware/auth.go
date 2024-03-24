package middleware

import (
	"fmt"
	"net"
	"net/http"
	service "p1-backend/api/pkg/services"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			fmt.Println(ip + ":" + r.RequestURI + " Header de autorización faltante")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Header de autorización faltante"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := service.ValidateToken(tokenString)
		if err != nil {
			fmt.Println(ip + ":" + r.RequestURI + " Error al verificar el token JWT: " + err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error al verificar el token JWT: " + err.Error()))
			return
		}
		expirationTime := time.Unix(int64(claims.(jwt.MapClaims)["expiration"].(float64)), 0)
		if time.Now().After(expirationTime) {
			fmt.Println(ip + ":" + r.RequestURI + " Token expirado")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token expirado"))
			return
		}
		name := claims.(jwt.MapClaims)["name"].(string)
		ID := claims.(jwt.MapClaims)["id"].(string)
		email := claims.(jwt.MapClaims)["email"].(string)

		r.Header.Set("name", name)
		r.Header.Set("userID", ID)
		r.Header.Set("email", email)
		fmt.Println(ip + ":" + r.RequestURI + " Token aceptado")
		next.ServeHTTP(w, r)
	})
}
