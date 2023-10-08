package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string `json:"username" gorm:"unique;not null;default:null"`
	Email   string `json:"email" gorm:"unique;not null;default:null"`
	Picture string `json:"picture" gorm:"not null;default:null"`
	ID      string `json:"id" gorm:"unique;not null;default:null"`
}

type Token struct {
	gorm.Model
	UserID       string `json:"user_id" gorm:"unique;not null;default:null"`
	AccessToken  string `json:"access_token" gorm:"unique;not null;default:null"`
	RefreshToken string `json:"refresh_token" gorm:"unique;not null;default:null"`
}
