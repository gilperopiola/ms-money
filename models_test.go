package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMoney(t *testing.T) {
	cfg.Setup("test")
	db.Setup(cfg)
	defer db.Close()

	money := GetMoney()
	assert.Equal(t, 0, money)
}

func TestGetTransactions(t *testing.T) {
	cfg.Setup("test")
	db.Setup(cfg)
	defer db.Close()

	transaction := &Transaction{}
	transactions, err := transaction.GetAll()
	assert.Zero(t, len(transactions))
	assert.NoError(t, err)

	transaction = &Transaction{Name: "1", Description: "1", Amount: 50, Date: time.Now()}
	transaction.Create()

	transactions, err = transaction.GetAll()
	assert.Equal(t, 1, len(transactions))
	assert.Equal(t, "1", transactions[0].Name)
	assert.Equal(t, 50, transactions[0].Amount)
	assert.NoError(t, err)
}

func TestCreateTransaction(t *testing.T) {
	cfg.Setup("test")
	db.Setup(cfg)
	defer db.Close()

	transaction := &Transaction{Name: "1", Description: "1", Amount: 50, Date: time.Now()}
	transaction, err := transaction.Create()
	assert.Equal(t, "1", transaction.Name)
	assert.Equal(t, "1", transaction.Description)
	assert.Equal(t, 50, transaction.Amount)
	assert.NotZero(t, transaction.Date)
	assert.NoError(t, err)

	updatedMoney := GetMoney()

	assert.Equal(t, 50, updatedMoney)
}
