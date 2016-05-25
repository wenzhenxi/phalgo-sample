package Domain
import (
	"phalgo-sample/Api/user/Model"
)


type Base struct {
	Map Domain_Map
}

type Domain_Map struct {
	Model Model.Base
}
