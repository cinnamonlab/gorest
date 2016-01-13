package gorest
import (
	"net/http"
	"encoding/json"
)

type Response struct {
	http.ResponseWriter
}

func NewResponse(response http.ResponseWriter) Response {

	rsp := Response{response}

	return  rsp
}

func (rsp * Response) Json(content interface{}, code int)  {
	rsp.WriteHeader(code)
	rsp.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(content)
	if err != nil {
		http.Error(rsp, err.Error(), http.StatusInternalServerError)
		return
	}

	rsp.Write(js)
}

func (rsp * Response) Text(content string, code int)  {
	rsp.WriteHeader(code)
	rsp.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rsp.Write([]byte(content))
}

func (rsp *Response) Redirect(redirectTo string)  {
	rsp.Header().Add("Location:",redirectTo)
}
