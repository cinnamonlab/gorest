package validator

type Validator interface {
	execute(string) bool
}
