package main

import (
	"time"

	"github.com/gilperopiola/frutils"
)

type Transaction struct {
	ID          int
	Name        string
	Description string
	Amount      int
	HadBefore   int
	Date        time.Time
}

func GetMoney() int {
	money := 0

	err := db.DB.QueryRow("SELECT amount FROM money LIMIT 1").Scan(&money)
	if err != nil {
		db.DB.Exec("INSERT INTO money (amount) VALUES (0)")
	}

	return money
}

func (transaction *Transaction) GetAll() ([]*Transaction, error) {
	money := GetMoney()

	rows, err := db.DB.Query(`SELECT id, name, description, amount, date FROM transactions ORDER BY id DESC`)
	defer rows.Close()
	if err != nil {
		return []*Transaction{}, err
	}

	transactions := []*Transaction{}
	for rows.Next() {
		tempTransaction := &Transaction{}

		err = rows.Scan(&tempTransaction.ID, &tempTransaction.Name, &tempTransaction.Description, &tempTransaction.Amount, &tempTransaction.Date)
		if err != nil {
			return []*Transaction{}, err
		}

		tempTransaction.HadBefore = money - tempTransaction.Amount
		money -= tempTransaction.Amount

		transactions = append(transactions, tempTransaction)
	}

	return transactions, nil
}

func (transaction *Transaction) GetAllSince() ([]*Transaction, error) {
	rows, err := db.DB.Query(`SELECT id, name, description, amount, date FROM transactions WHERE date > ? ORDER BY id ASC`, transaction.Date)
	defer rows.Close()
	if err != nil {
		return []*Transaction{}, err
	}

	transactions := []*Transaction{}
	for rows.Next() {
		tempTransaction := &Transaction{}

		err = rows.Scan(&tempTransaction.ID, &tempTransaction.Name, &tempTransaction.Description, &tempTransaction.Amount, &tempTransaction.Date)
		if err != nil {
			return []*Transaction{}, err
		}

		transactions = append(transactions, tempTransaction)
	}

	return transactions, nil
}

func (transaction *Transaction) Create() (*Transaction, error) {
	result, err := db.DB.Exec("INSERT INTO transactions (name, description, amount, date) VALUES (?, ?, ?, ?)",
		transaction.Name, transaction.Description, transaction.Amount, transaction.Date)

	if err != nil {
		return &Transaction{}, err
	}

	transaction.ID = frutils.GetID(result)

	transaction.UpdateMoney()

	return transaction, nil
}

func (transaction *Transaction) UpdateMoney() {
	currentMoney := GetMoney()
	updatedMoney := currentMoney + transaction.Amount

	db.DB.Exec("UPDATE money SET amount = ?", updatedMoney)
}
