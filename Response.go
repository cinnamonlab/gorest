package gorest
import (
	"net/http"
	"encoding/json"
)

type Response struct {
	response http.ResponseWriter
}

func NewResponse(response http.ResponseWriter) *Response {

	rsp := &Response{response : response}

	return  rsp
}

func (rsp * Response) Json(content interface{}, code int)  {
	rsp.response.WriteHeader(code)
	rsp.response.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(content)
	if err != nil {
		http.Error(rsp.response, err.Error(), http.StatusInternalServerError)
		return
	}

	rsp.response.Write(js)
}
func (rsp * Response) Text(content string, code int)  {
	rsp.response.WriteHeader(code)
	rsp.response.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rsp.response.Write([]byte(content))
}

func (rsp *Response) Redirect(redirectTo string)  {
	rsp.response.Header().Add("Location:",redirectTo)
}
