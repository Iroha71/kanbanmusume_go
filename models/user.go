package models

type User struct {
	Id       uint   `json:"id" binding:"requried" gorm:"AUTO_INCREMENT" gorm:"primaryKey"`
	Name     string `json:"name" binding:"required" gorm:"size:30"`
	Nickname string `json:"nickname" binding:"required" gorm:"size:10"`
	Coin     int    `json:"coin" sql:"DEFAULT:0"`
}
