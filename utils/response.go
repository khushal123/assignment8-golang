package utils

import "net/http"

func Response(w http.ResponseWriter, status string, data []byte) {

	type response struct {
		Status string `json:"status"`
		Data   string `json:"data"`
	}
	r := []byte("{\"status\":" + status + ", \"data\":" + string(data) + "}")
	w.Header().Set("Content-Type", "application/json")
	w.Write(r)
}
