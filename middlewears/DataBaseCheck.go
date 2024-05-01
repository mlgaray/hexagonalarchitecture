package middlewears

import (
	"net/http"
)

func DataBaseCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*if bd.ChequeoConnection() == 0 {
			http.Error(w,"connection has been lost",500)
			return
		}*/
		next.ServeHTTP(w, r)
	}
}
