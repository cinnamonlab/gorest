package main
import (
	"net/http"
	"github.com/cinnamonlab/gorest"
	"regexp"
	"fmt"
)

func main()  {

	r,_ := regexp.Compile("^:(.*)$")

	str := r.FindStringSubmatch(":id");

	fmt.Println(str)



	route := gorest.GetRouteInstance()
	http.ListenAndServe(":3000",route)
}
