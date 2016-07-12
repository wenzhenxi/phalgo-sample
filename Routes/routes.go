package routes
import (

	"github.com/wenzhenxi/phalgo"
	"phalgo-sample/Api/User/Api"
)


func GetRoutes() {

	var Api Api.Base

	phalgo.Echo.Any("/GetUserInfo", Api.User.GetUserInfo())
	phalgo.Echo.Any("/hello", Api.User.Hello())

}