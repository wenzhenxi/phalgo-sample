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

func (this *User_Api)GetUserInfo() echo.HandlerFunc {

	return func(c echo.Context) error {
		Request := phalgo.Requser{Context:c}
		Response := phalgo.Response{Context:c}
		defer Request.ErrorLogRecover()

		id := Request.GetParam("id").GetInt()
		user, err := this.Domain.User.GetUserInfo(id)
		if err != nil {
			return Response.RetError(err, 400)
		}

		return Response.RetSuccess(user)
	}
}