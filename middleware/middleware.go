package middleware

import (
	"github.com/julienschmidt/httprouter"
	"go-belajar-restfull/helper"
	"net/http"
)

type ApiMiddleware struct {
	Handler http.Handler
}

func NewApiMiddleware(handler *httprouter.Router) *ApiMiddleware {
	return &ApiMiddleware{
		Handler: handler,
	}
}

func (m *ApiMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-Api-Key")
	w.Header().Set("Content-Type", "application/json")
	if apiKey == "123-xyz-ok" {
		m.Handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		helper.WriteResponse(w, helper.CreateResponse(http.StatusUnauthorized, "Unauthorized", nil))
	}
}
