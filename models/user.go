package models

type User struct {
	Id       uint   `json:"id" binding:"requried" gorm:"AUTO_INCREMENT"`
	Name     string `json:"name" binding:"required" gorm:"size:30"`
	Nickname string `json:"nickname" binding:"required" gorm:"size:10"`
}

// func GetAll() ([]User, error) {

// }
