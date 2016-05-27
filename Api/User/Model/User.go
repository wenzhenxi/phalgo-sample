package Model
import "github.com/wenzhenxi/phalgo"


type User struct {
	Id     int       `gorm:"column:aId"`
	Name   string    `gorm:"column:name"`
	Passwd string    `gorm:"column:passwd"`
	Email  string    `gorm:"column:email"`
}

func (User) TableName() string {
	return "user"
}

func (this *User)GetInfoById(id int) (User, error) {

	User := User{}
	err := phalgo.GetORM().Where("id = ?", id).First(&User).Error
	return User, err
}