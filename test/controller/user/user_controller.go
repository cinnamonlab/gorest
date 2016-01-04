package user
import "net/http"

func (u *UserController) getMySelf(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	a := 1
	a++
	return nil
}
