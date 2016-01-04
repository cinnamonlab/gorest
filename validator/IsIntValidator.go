package validator
import (
	"strconv"
)

type IsIntValidator struct { // implement validator interface
}

func (obj IsIntValidator) execute(param string) bool  {
	_,err := strconv.ParseInt(param,0,64);
	if err != nil {
		return  false
	} else {
		return  true
	}
}

