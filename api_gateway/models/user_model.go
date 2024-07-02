package models

type User struct {
	UserID   uint    `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Username string  `json:"username" gorm:"type:varchar(255);not null"`
	Email    string  `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password string  `json:"password" gorm:"type:varchar(255);not null"`
	Role     string  `json:"role" gorm:"type:varchar(50);not null"`
	Deposit  float64 `json:"deposit" gorm:"type:decimal(10,2);check:deposit>=0"`
}

type UserRegister struct {
	Username string `json:"username" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserID   uint    `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Username string  `json:"username" gorm:"type:varchar(255);not null"`
	Email    string  `json:"email" gorm:"type:varchar(255);unique;not null"`
	Deposit  float64 `json:"deposit" gorm:"type:decimal(10,2);check:deposit>=0"`
}

type UserLoginResponse struct {
	UserID   uint    `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Username string  `json:"username" gorm:"type:varchar(255);not null"`
	Email    string  `json:"email" gorm:"type:varchar(255);unique;not null"`
	Deposit  float64 `json:"deposit" gorm:"type:decimal(10,2);check:deposit>=0"`
}

type TopUpReq struct {
	Deposit float64 `json:"deposit"`
}
