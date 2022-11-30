package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// `Username` Generated value with disabled create/update permissions
	// [create with associations](https://gorm.io/docs/create.html#Create-With-Associations)
	// [field-level permissions](https://gorm.io/docs/models.html#field_permission)
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
	Username string `gorm:"->;GENERATED ALWAYS AS split_part(email, '@', 1)"  json:"username"`
	Role     string `gorm:"default:user" json:"role"`
}
