package models

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model           `json:"-"`
	UserID               int     `gorm:"primaryKey;not null;unique" json:"userID"`
	AccountNumber        string  `gorm:"primaryKey;unique;not null" json:"account"`
	Balance              float64 `gorm:"not null" json:"balance"`
	Beneficiaries        []User  `gorm:"not null" json:"benefactors"`
	SentTransactions     []User  `gorm:"not null" json:"sent"`
	ReceivedTransactions []User  `gorm:"not null" json:"received"`
}

// `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"

type User struct {
	gorm.Model `json:"-"`
	UserID     int     `gorm:"primaryKey;not null;unique" json:"userID"`
	FullName   string  `gorm:"not null" json:"first"`
	Email      string  `gorm:"unique;not null" json:"email"`
	Phone      int     `gorm:"unique;not null" json:"phone"`
	Account    Account `gorm:"unique;not null" json:"account"`
}

type Transaction struct {
	Amount        float64   `gorm:"not null" json:"amount"`
	Time          time.Time `gorm:"default:CURRENT_TIMESTAMP; not null" json:"time"`
	Sender        Account   `gorm:"not null" json:"sender"`
	Receiver      Account   `gorm:"not null" json:"receiver"`
	TransactionID int       `gorm:"primaryKey;unique;not null" json:"txn_id"`
	Location      string    `json:"location"`
}
