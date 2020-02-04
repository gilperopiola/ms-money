package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* GET */

func GetMoneyAmount(c *gin.Context) {
	money := GetMoney()

	c.JSON(http.StatusOK, money)
}

func GetTransactions(c *gin.Context) {
	transaction := &Transaction{}

	transactions, err := transaction.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, transactions)
}

/* POST */

func CreateTransaction(c *gin.Context) {
	var transaction *Transaction

	if bindError := c.BindJSON(&transaction); bindError != nil {
		c.JSON(http.StatusBadRequest, bindError.Error())
		return
	}

	if transaction.Name == "" || transaction.Amount == 0 {
		c.JSON(http.StatusBadRequest, "name and amount required")
		return
	}

	transaction, err := transaction.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, transaction)
}
