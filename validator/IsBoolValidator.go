package validator
import (
	"strconv"
)

type IsBoolValidator struct { // implement validator interface
}

func (obj IsBoolValidator) execute(param string) bool  {
	_,err := strconv.ParseBool(param);
	if err !=  nil {
		return  false
	} else {
		return  true
	}
}

