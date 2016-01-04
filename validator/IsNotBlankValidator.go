package validator

type IsNotBlankValidator struct { // implement validator interface
}

func (obj IsNotBlankValidator) execute(param string) bool  {
	return len(param)>0
}

