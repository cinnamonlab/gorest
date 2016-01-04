package validator
import (
	"strconv"
)

type IsFloatValidator struct { // implement validator interface
}

func (obj IsFloatValidator) execute(param string) bool  {
	_,err := strconv.ParseFloat(param,64);
	if err != nil {
		return  false
	} else {
		return  true
	}
}

