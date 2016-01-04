package user
import (
	"github.com/cinnamonlab/gorest"
)

func (u *UserController) GetMySelf(w *gorest.Response, r *gorest.Request, vars map[string]string) {
	mapavl := map[string]string {"sss":"ssssss"}
	w.Json(mapavl,200)
}
func (u *UserController) GetMyByID(w *gorest.Response, r *gorest.Request, vars map[string]string) {
	w.Text("OK111",200)
}
