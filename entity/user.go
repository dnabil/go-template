package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// uuid is used as 'subject' in user claims
	UUID        string `gorm:"type:CHAR(36);NOT NULL;UNIQUE;INDEX"`
	
	Email 		string	`gorm:"type:VARCHAR(200);UNIQUE"`
	Password	string	`gorm:"type:VARCHAR(255);NOT NULL"`
	FName		string	`gorm:"type:VARCHAR(50);NOT NULL"`
	LName 		string	`gorm:"type:VARCHAR(50);"`
}