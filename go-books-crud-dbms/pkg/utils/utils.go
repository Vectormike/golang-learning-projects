package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, v interface{}) {
	if body, err := ioutil.ReadAll(r.body); err == nil {
		if err := json.Unmarshal(body, v); err != nil {
			panic(err)
		}
	}
}
