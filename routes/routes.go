package routes
import (

	"github.com/wenzhenxi/phalgo"
	"phalgo-sample/Api/user/Api"
)


func GetRoutes() {

	var Api Api.Base

	phalgo.Echo.Any("/Login", Api.User.Login())
	phalgo.Echo.Any("/hello", Api.User.Hello())
}