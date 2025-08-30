package helper

import (
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/log"
)

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

type Result struct {
	Value     any    `json:"value,omitempty"`
	IsSuccess bool   `json:"is_success"`
	Error     *string `json:"error,omitempty"`
}

//TODO:  add func to return a good log by result
func(r *Result) Log(logger *log.Logger){
	var value any
	if r.Value != nil{
		value = r.Value
	}else{
		value = nil
	}
	logger.Error("Return Result:", "Value" , value , "IsSuccess" , r.IsSuccess  ,"Error" , *r.Error)
} 
