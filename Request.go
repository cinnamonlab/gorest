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

const (
	defaultMaxMemory = 32 << 20 // 32 MB
)

func NewRequest(request *http.Request,params map[string]string) *Request {
	input := &Request{request,params}
	return  input
}

// set param to resource
func (input *Request) SetParam(key string, value string)  {
	if(input.params==nil) {
		input.params = make(map[string]string)
	}
	input.params[key]=value;
}

// Get query body
func (input *Request) GetJsonBody() []byte  {

	body, err := ioutil.ReadAll(input.Body)
	if err != nil {
		errors.New("can not read request body")
	}
	return body
}

// check param is exist or not
// return value if exist, error if not
func (input *Request) GetStringValue(key string) (string,error) {
	if val, ok := input.params[key]; ok {
		return val, nil
	} else {
		if input.Form == nil {
			input.ParseMultipartForm(defaultMaxMemory)
		}

		val, ok := input.Form[key]

		if ok && len(val) > 0 {
			return val[0], nil
		} else {
			return "", errors.New(key + " INVAID")
		}
	}
}

// check param is exist or not
// return value if exist, default value if not
func (input *Request) GetStringValueOrDefault(key string,defaultVal string) string {
	val,err := input.GetStringValue(key)

	if err!=nil {
		return defaultVal
	} else {
		return val
	}
}

// check param is exist or not
// return Bool value if exist, error if not
func (input *Request) GetBoolValue( k string) (bool,error) {
	val,err := input.GetStringValue(k)

	if(err!=nil) {
		return false,err
	} else {
		value, err := strconv.ParseBool(val)

		if err != nil {
			return false,err
		} else {
			return value,nil
		}
	}
}

// check param is exist or not
// return bool value if exist, default if not
func (input *Request) GetBoolValueOrDefault( k string,defaultVal bool) bool {
	val,err := input.GetBoolValue(k)

	if(err!=nil) {
		return defaultVal
	} else {
		return val
	}
}
// check param is exist or not
// return int value if exist, error if not
func (input *Request) GetIntValue( k string) (int64,error) {
	val,err := input.GetStringValue(k)

	if(err!=nil) {
		return -1,err
	} else {
		value, err := strconv.ParseInt(val, 0, 64)

		if err != nil {
			return -1,err
		} else {
			return value,nil
		}
	}
}

// check param is exist or not
// return int value if exist, default if not
func (input *Request) GetIntValueOrDefault( k string,defaultVal int64) int64 {
	val,err := input.GetIntValue(k)

	if(err!=nil) {
		return defaultVal
	} else {
		return val
	}
}
// check param is exist or not
// return unint value if exist, error if not
func (input *Request) GetUintValue( k string) (uint64,error) {
	val,err := input.GetStringValue(k)
	if(err!=nil) {
		return 0,err
	} else {
		value, err := strconv.ParseUint(val, 0, 64)

		if err != nil {
			return 0,err
		} else {
			return value,nil
		}
	}
}

// check param is exist or not
// return uint value if exist, default if not
func (input *Request) GetUintValueOrDefault( k string,defaultVal uint64) uint64 {
	val,err := input.GetUintValue(k)
	if(err!=nil) {
		return defaultVal
	} else {
		return val
	}
}
// check param is exist or not
// return float value if exist, error if not
func (input *Request) GetFloatValue( k string) (float64,error) {
	val,err := input.GetStringValue(k)

	if(err!=nil) {
		return 0,err
	} else {
		value, err := strconv.ParseFloat(val, 64)

		if err != nil {
			return -1,err
		} else {
			return value,nil
		}
	}
}
// check param is exist or not
// return float value if exist, default if not
func (input *Request) GetFloatValueOrDefault( k string,defaultVal float64) float64 {
	val, err := input.GetFloatValue(k)

	if (err != nil) {
		return defaultVal
	} else {
		return val
	}
}
