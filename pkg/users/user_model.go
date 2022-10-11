package users

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model       `json:"-"`
	AccountNumber    string  `gorm:"unique;not null" json:"account"`
	Balance          float64 `gorm:"not null" json:"balance"`
	Beneficiaries    []User  `json:"beneficiaries"`
	SentTransactions []User  `json:"sent"`
}

// `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"

type User struct {
	gorm.Model `json:"-"`
	FullName   string   `gorm:"not null" json:"first"`
	Email      string   `gorm:"unique;not null" json:"email"`
	Phone      int      `gorm:"unique;not null" json:"phone"`
	Account    *Account `gorm:"unique;not null" json:"account"`
}

type Transaction struct {
	Amount        float64   `gorm:"not null" json:"amount"`
	Time          time.Time `gorm:"default:CURRENT_TIMESTAMP; not null" json:"time"`
	Sender        *User     `gorm:"not null" json:"sender"`
	Receiver      *User     `gorm:"not null" json:"receiver"`
	TransactionID int       `gorm:"unique;not null" json:"txn_id"`
	Location      string    `json:"location"`
}
