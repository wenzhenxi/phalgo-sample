package Api

import (
	"github.com/labstack/echo"
	"github.com/wenzhenxi/phalgo"
)


func (this *User_Api)Hello() echo.HandlerFunc {

	return func(c echo.Context) error {
		Request := phalgo.Requser{Context:c}
		Response := phalgo.Response{Context:c}
		defer Request.ErrorLogRecover()
		return Response.RetSuccess("helloworld")
	}
}

func (this *User_Api)Login() echo.HandlerFunc {

	return func(c echo.Context) error {
		Request := phalgo.Requser{Context:c}
		Response := phalgo.Response{Context:c}
		defer Request.ErrorLogRecover()
		return Response.RetSuccess("helloworld")
	}
}