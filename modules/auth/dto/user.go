package dto

type User struct {
	Id       uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
	Role     string `json:"role"`
}
