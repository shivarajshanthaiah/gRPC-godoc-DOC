package models

type Doctor struct {
	ID          uint64 `json:"doctor_id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null;unique" validate:"required"`
	Age         uint64 `json:"age" gorm:"not null" validate:"required"`
	Speciality  string `json:"speciality" gorm:"not null" validate:"required,len=4"`
	Gender      string `json:"gender" gorm:"not null" validate:"required"`
	Email       string `json:"email" gorm:"not null" validate:"required"`
}