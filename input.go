package gorest

type RequestInput struct {
	params map[string]string
}

func (input *RequestInput) has(key string) bool  {
	if value, ok := input.params[key]; ok {
		return  true
	} else  {
		return  false
	}
}
func (input *RequestInput) setParam(key string, value string)  {
	if(input.params==nil) {
		input.params = make(map[string]string)
	}
	input.params[key]=value;
}
func (input *RequestInput) get(key string) string  {
	if(input.has(key)) {
		return input.params[key]
	} else {
		return nil
	}
}

