package model

import "time"

type User struct {
	ID            uint32  `gorm:"primary_key"`
	AccountNumber string  `gorm:"type:varchar(100);unique"`
	Name          string  `gorm:"type:varchar(100)"`
	Surname       string  `gorm:"type:varchar(100)"`
	Email         string  `gorm:"type:varchar(100);unique"`
	Password      string  `gorm:"type:varchar(100)"`
	Balance       float64 `gorm:"type:decimal(10,2)"`
}

type Permission struct {
	ID         uint32 `gorm:"primary_key"`
	UserID     uint32
	Permission string
}

type Transaction struct {
	ID        uint32    `gorm:"primary_key"`
	Sender    uint32    `gorm:"index"`
	Receiver  uint32    `gorm:"index"`
	Amount    float64   `gorm:"type:decimal(10,2)"`
	Fee       float64   `gorm:"type:decimal(10,2)"`
	Type      string    `gorm:"type:enum('deposit', 'withdraw', 'transfer')"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
