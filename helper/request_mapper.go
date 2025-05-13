package helper

import (
	"encoding/json"
	"net/http"
)

func RequestBodyToStruct[T any](r *http.Request, requestBody *T) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)

	CheckError(err)
}
                                                