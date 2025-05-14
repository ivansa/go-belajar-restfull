package exception

import (
	"fmt"
	"go-belajar-restfull/helper"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err any) {
	switch err.(type) {
	case NotFoundError:
		NotFoundErrotHandler(w)
	case BadRequestError:
		BadRequestErrotHandler(w)
	default:
		InternalErrorHandler(w, err)
	}
}

func BadRequestErrotHandler(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	helper.WriteResponse(w, helper.CreateResponse(http.StatusBadRequest, "Bad Request", nil))
}

func NotFoundErrotHandler(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	helper.WriteResponse(w, helper.CreateResponse(http.StatusNotFound, "Resource not found", nil))
}

func InternalErrorHandler(w http.ResponseWriter, err any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	helper.WriteResponse(w, helper.CreateResponse(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil))
}
