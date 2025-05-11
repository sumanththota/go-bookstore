package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

//it helps to parse the bodyin the create book function, json => unmarshall
//-> reads the json body
//-> Unmarshall(converts it into the provided interface structure)
//-> 
func ParseBody(r *http.Request, x interface{}){
	if body, err := io.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return
		}

	}
}