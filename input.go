package gorest
import (
	"net/http"
)

type RequestInput struct {
	params map[string]string
}

func (input RequestInput) parserFromRequest(request *http.Request)  {

}

func (input RequestInput) has(key string) bool  {
	if _, ok := input.params[key]; ok {
		return  true
	} else  {
		return  false
	}
}
func (input RequestInput) setParam(key string, value string)  {
	if(input.params==nil) {
		input.params = make(map[string]string)
	}
	input.params[key]=value;
}
func (input RequestInput) get(key string, defaultVal interface{}  ) interface{} {
	if(input.has(key)) {
		value := input.params[key]

		return  value;

	} else {
		return defaultVal
	}
}

