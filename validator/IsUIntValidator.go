package validator
import (
	"strconv"
)

type IsUIntValidator struct { // implement validator interface
}

func (obj IsUIntValidator) execute(param string) bool  {
	_,err := strconv.ParseUint(param,0,64);
	if err != nil {
		return  false
	} else {
		return  true
	}
}

