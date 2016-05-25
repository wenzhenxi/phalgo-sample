package Domain
import (
	"phalgo-sample/Api/appmarket/1.0/Model"
)


type Base struct {
	Map Domain_Map
}

type Domain_Map struct {
	Model Model.Base
}
