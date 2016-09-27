package Api

import (
	"github.com/labstack/echo"
	"github.com/wenzhenxi/phalgo"
	"fmt"
)


func (this *User_Api)Hello() echo.HandlerFunc {

	return func(c echo.Context) error {
		Request := phalgo.NewRequest(c)
		Response := phalgo.NewResponse(c)
		defer Request.ErrorLogRecover()

		name := Request.Param("name").Max(30).SetDefault("喵咪").GetString()

		//参数过滤error处理
		if err := Request.GetError(); err != nil {
			return Response.RetError(err, -1)
		}
		
		return Response.RetSuccess("hello " + name + " Welcome to PhalGo world")
	}
}

func (this *User_Api)GetUserInfo() echo.HandlerFunc {

	return func(c echo.Context) error {

		Request := phalgo.NewRequest(c)
		Response := phalgo.NewResponse(c)
		defer Request.ErrorLogRecover()

		id := Request.GetParam("id").GetInt()

		//参数过滤error处理
		if err := Request.GetError(); err != nil {
			return Response.RetError(err, -1)
		}

		user, err := this.Domain.User.GetUserInfo(id)
		if err != nil {
			return Response.RetError(err, 400)
		}

		return Response.RetSuccess(user)
	}
}

func (this *User_Api)GetUserList() echo.HandlerFunc {

	return func(c echo.Context) error {

		Request := phalgo.NewRequest(c)
		Response := phalgo.NewResponse(c)
		defer Request.ErrorLogRecover()

		Request.SetJson(`{"handle" : "fa0c7f9e6136ee5a42a384f069c4827f","memberidx" : 12,"leagueidx" : 22,"handleid" : "15062212900"}`)
		fmt.Println(Request.JsonParam("memberidx").GetInt())

		//参数过滤error处理
		if err := Request.GetError(); err != nil {
			return Response.RetError(err, -1)
		}

		user, err := this.Domain.User.GetUserList()
		if err != nil {
			return Response.RetError(err, 400)
		}

		return Response.RetSuccess(user)
	}
}

