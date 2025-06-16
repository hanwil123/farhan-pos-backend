package dto

type UserModel struct {
	Id       uint64 `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
