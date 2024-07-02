package models

type UserDetail struct {
	UserDetailID      uint   `json:"user_detail_id" gorm:"primaryKey;autoIncrement"`
	UserID            uint   `json:"user_id" gorm:"unique;not null"`
	Fname             string `json:"fname" gorm:"type:varchar(255)"`
	Lname             string `json:"lname" gorm:"type:varchar(255)"`
	Address           string `json:"address" gorm:"type:text"`
	Age               int    `json:"age" gorm:"check:age>0"`
	PhoneNumber       string `json:"phone_number" gorm:"type:varchar(20)"`
	ProfilePictureUrl string `json:"profile_picture_url" gorm:"varchar(255)"`
}

type UserDetailResponse struct {
	UserID            uint    `json:"user_id" gorm:"unique;not null"`
	Username          string  `json:"username" gorm:"type:varchar(255);not null"`
	Email             string  `json:"email" gorm:"type:varchar(255);unique;not null"`
	Deposit           float64 `json:"deposit" gorm:"type:decimal(10,2);check:deposit>=0"`
	Fname             string  `json:"fname" gorm:"type:varchar(255)"`
	Lname             string  `json:"lname" gorm:"type:varchar(255)"`
	Address           string  `json:"address" gorm:"type:text"`
	Age               int     `json:"age" gorm:"check:age>0"`
	PhoneNumber       string  `json:"phone_number" gorm:"type:varchar(20)"`
	ProfilePictureUrl string  `json:"profile_picture_url" gorm:"varchar(255)"`
}

type UserUpdateRequest struct {
	Username          string `json:"username"`
	Fname             string `json:"fname"`
	Lname             string `json:"lname" `
	Address           string `json:"address"`
	Age               int    `json:"age"`
	PhoneNumber       string `json:"phone_number"`
	ProfilePictureUrl string `json:"profile_picture_url"`
}
