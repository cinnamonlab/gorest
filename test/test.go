package main
import (
	"net/http"
	"github.com/cinnamonlab/gorest"
)

func main()  {
	route := gorest.GetRouteInstance()
	http.ListenAndServe(":3000",route)
}
