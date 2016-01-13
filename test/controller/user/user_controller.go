package user
import (
	"github.com/cinnamonlab/gorest"
)

func (u *UserController) GetMySelf(w gorest.Response, r *gorest.Request) {
	//sss := r.GetStringValue("sss")
}
func (u *UserController) GetMyByID(w gorest.Response, r *gorest.Request) {
	w.Text("OK111",200)
}
