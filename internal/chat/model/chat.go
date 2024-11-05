package model

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Messages []Messages
	Members  []ChatMembers
}
