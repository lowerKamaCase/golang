package res

import (
	"encoding/json"
	"net/http"
)

func Json(rw http.ResponseWriter, res any, statusCode int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(res)
}
