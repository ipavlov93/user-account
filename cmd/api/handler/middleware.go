package handler

import (
	"event-calendar/internal/service/auth"
	"log"
	"net/http"
)

const AuthorizationHeader = "Authorization"

func requireValidIDToken(r *http.Request, rw http.ResponseWriter) {
	idToken := r.Header.Get(AuthorizationHeader)
	_, err := auth.Single().VerifyIDToken(idToken)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		_, err = rw.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}
}
