package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type (

	TXRepository struct {
		db *gorm.DB
	}
)

func NewTXRepository(db *gorm.DB) *TXRepository {
	return &TXRepository{db}
}

func (repo TXRepository) AddCartItem() error {
	fmt.Println("Transaction")
	return nil
}