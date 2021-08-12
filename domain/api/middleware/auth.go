package middleware

import (
	"context"
	"net/http"
	"transaction-api/client/userapi"
)

type Auth struct {
	userApiClient userapi.Client
	ctx           context.Context
}

func NewAuth(userApiClient userapi.Client) *Auth {
	return &Auth{userApiClient: userApiClient}
}

func (a *Auth) ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("Authorization")

		err := a.userApiClient.ValidateToken(a.ctx, token)
		if err == nil {
			next.ServeHTTP(writer, request)
		} else {
			http.Error(writer, "Forbidden", http.StatusForbidden)
			return
		}
	})
}
