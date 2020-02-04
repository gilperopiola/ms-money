package main

var createMoneyTable = `
CREATE TABLE IF NOT EXISTS money (
	amount int NOT NULL DEFAULT '0'
)
`

var createTransactionsQuery = `
CREATE TABLE IF NOT EXISTS transactions (
	id int UNIQUE NOT NULL AUTO_INCREMENT,
	name varchar(255) NOT NULL,
	description varchar(255) NOT NULL DEFAULT '',
	amount int NOT NULL,
	date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)
`
