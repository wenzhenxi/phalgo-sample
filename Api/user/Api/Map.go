package Api

import (
	"github.com/labstack/echo"
	"github.com/wenzhenxi/phalgo"
)


func (this *App_Api)GetAreaByName() echo.HandlerFunc {

	return func(c echo.Context) error {
		Request := phalgo.Requser{Context:c}
		Response := phalgo.Response{Context:c}
		defer Request.ErrorLogRecover()
		return Response.RetSuccess("helloworld")
	}
}