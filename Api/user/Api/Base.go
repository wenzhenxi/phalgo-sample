package Api
import (
	"phalgo-sample/Api/user/Domain"
)

type Base struct {
	User User_Api
}

type User_Api struct {
	Domain.Base
}
