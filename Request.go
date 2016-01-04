package gorest
import (
	"net/http"
	"strconv"
)

type Request struct {
	req *http.Request
	params map[string]string
}

func NewRequest(request * http.Request) *Request {
	input := &Request{req:request}
	input.params = make(map[string]string)

	return  input
}

// check input param is exiting or not
func (input *Request) Has(key string) ( exist bool,val string)  {
	if val, ok := input.params[key]; ok {
		return  true,val
	} else if val, ok := input.req.Form[key]; ok {
		return false,""
	} else {
		return  true,val[0]
	}
}
// set param to resource
func (input *Request) SetParam(key string, value string)  {
	if(input.params==nil) {
		input.params = make(map[string]string)
	}
	input.params[key]=value;
}
// get param value
func (input *Request) Get(key string, defaultVal interface{}  ) interface{} {
	exist,val := input.Has(key)
	if(exist) {
		return  val;
	} else {
		return defaultVal
	}
}

// BoolValue transforms a form value in different formats into a boolean type.
func (input *Request) BoolValue( k string) bool {
	val :=input.req.FormValue(k)
	value, err := strconv.ParseBool(val)

	if err!=nil {
		return  false
	} else {
		return value
	}
}

// BoolValueOrDefault returns the default bool passed if the query param is
// missing, otherwise it's just a proxy to boolValue above
func (input *Request) BoolValueOrDefault( k string, d bool) bool {
	if _, ok := input.req.Form[k]; !ok {
		return d
	}
	return input.BoolValue(k)
}

// IntValueOrZero parses a form value into an int64 type.
// It returns 0 if the parsing fails.
func (input *Request) IntValueOrZero(k string) int64 {
	val, err := input.IntValueOrDefault(k, 0)
	if err != nil {
		return 0
	}
	return val
}

// Int64ValueOrDefault parses a form value into an int64 type. If there is an
// error, returns the error. If there is no value returns the default value.
func (input *Request) IntValueOrDefault(field string, def int64) (int64, error) {
	if input.req.Form.Get(field) != "" {
		value, err := strconv.ParseInt(input.req.Form.Get(field), 10, 64)
		if err != nil {
			return value, err
		}
		return value, nil
	}
	return def, nil
}

