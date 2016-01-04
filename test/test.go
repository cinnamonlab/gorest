package main
import (
	"net/http"
	"github.com/cinnamonlab/gorest"
	"github.com/cinnamonlab/gorest/test/controller"
	"github.com/cinnamonlab/gorest/test/controller/user"
)

func main()  {

	route := gorest.GetRouteInstance()

	user.UserController{}

	http.ListenAndServe(":3000",route)
}
func api(rsp http.ResponseWriter, req *http.Request, vars map[string]string) error  {
//	rsp.Write("Hello");
	a :=0
	a++;

	return nil
}
