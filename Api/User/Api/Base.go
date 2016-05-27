package Api
import (
	"phalgo-sample/Api/User/Domain"
)

type Base struct {
	User User_Api
}

type User_Api struct {
	Domain Domain.Base
}
