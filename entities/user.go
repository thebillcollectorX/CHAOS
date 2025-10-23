package entities

import "time"

type User struct {
	DBModel
	Username    string    `form:"username" json:"username,omitempty" binding:"required"`
	Password    string    `form:"password" json:"password,omitempty" binding:"required"`
	Email       string    `json:"email" gorm:"unique;size:100"`
	FirstName   string    `json:"first_name" gorm:"size:50"`
	LastName    string    `json:"last_name" gorm:"size:50"`
	Avatar      string    `json:"avatar" gorm:"size:500"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	LastLoginAt *time.Time `json:"last_login_at"`
	
	// Relationships
	Tokens  []Token  `json:"tokens,omitempty" gorm:"foreignKey:UserID"`
	Wallets []Wallet `json:"wallets,omitempty" gorm:"foreignKey:UserID"`
}
