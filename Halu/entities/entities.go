package entities

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	// Automatically Includes ID, CreatedAt, UpdatedAt, DeletedAt Fields
	FirstName string `gorm:"type:varchar(50);not null" json:"first_name"`
	LastName  string `gorm:"type:varchar(50);not null" json:"last_name"`
	Age       int    `gorm:"type:int;not null;check:age >= 0" json:"age"`
}
