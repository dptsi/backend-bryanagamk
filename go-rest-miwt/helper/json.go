package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(req *http.Request, res interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(res)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, res interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(res)
	PanicIfError(err)
}
