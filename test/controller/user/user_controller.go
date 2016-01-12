package user
import (
	"github.com/cinnamonlab/gorest"
)

func (u *UserController) GetMySelf(w *gorest.Response, r *gorest.Request, vars map[string]string) {
	//sss := r.GetStringValue("sss")
}
func (u *UserController) GetMyByID(w *gorest.Response, r *gorest.Request, vars map[string]string) {
	w.Text("OK111",200)
}
