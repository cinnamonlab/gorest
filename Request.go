package gorest
import (
	"net/http"
	"strconv"
	"io/ioutil"
	"errors"
)

type Request struct {
	*http.Request
	params map[string]string
}

func NewRequest(request *http.Request) *Request {
	input := &Request{request,nil}
	input.params = make(map[string]string)

	return  input
}

func (input *Request) GetJsonBody() []byte  {

	body, err := ioutil.ReadAll(input.Body)
	if err != nil {
		errors.New("can not read request body")
	}
	return body
}

// check input param is exiting or not
func (input *Request) GetStringValue(key string) string  {
	if val, ok := input.params[key]; ok {
		return  val
	} else {
		val := input.FormValue(key);
		if (val!="") {
			return val
		} else {
			val := input.PostFormValue(key);
			return val
		}
	}
}
// set param to resource
func (input *Request) SetParam(key string, value string)  {
	if(input.params==nil) {
		input.params = make(map[string]string)
	}
	input.params[key]=value;
}

// BoolValue transforms a form value in different formats into a boolean type.
func (input *Request) GetBoolValue( k string) bool {
	val := input.GetStringValue(k)
	value, err := strconv.ParseBool(val)

	if err != nil {
		return false
	} else {
		return value
	}
}
// BoolValue transforms a form value in different formats into a boolean type.
func (input *Request) GetIntValue( k string) int64 {
	val := input.GetStringValue(k)
	value, err := strconv.ParseInt(val, 0, 64)

	if err != nil {
		return -1
	} else {
		return value
	}
}
// BoolValue transforms a form value in different formats into a boolean type.
func (input *Request) GetUintValue( k string) uint64 {
	val := input.GetStringValue(k)
	value, err := strconv.ParseUint(val, 0, 64)

	if err != nil {
		return 0
	} else {
		return value
	}
}
// BoolValue transforms a form value in different formats into a boolean type.
func (input *Request) GetFloatValue( k string) float64 {
	val := input.GetStringValue(k)

	value, err := strconv.ParseFloat(val, 64)

	if err != nil {
		return -1
	} else {
		return value
	}
}
