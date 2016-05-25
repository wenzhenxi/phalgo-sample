package Api
import (
	"phalgo-sample/Api/appmarket/1.0/Domain"
)

type Base struct {
	App App_Api
}

type App_Api struct {
	Domain.Base
}
