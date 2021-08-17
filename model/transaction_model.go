package model

import (
	"fmt"
	"errors"

	"sales-backend/utility"
)

type Transaction struct {
	TransactionId	string 	`json:"transactionId"`
	CustomerId	string 	`json:"customerId"`
	EmployeeId	string 	`json:"employeeId"`
	Date		string 	`json:"date"`
	Audit		Audit 	`json:"audit"`
}

func (t *Transaction) IsTransactionExistsById(transactionId int) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	check := 0
	err = db.QueryRow("SELECT COUNT(transaction_id) FROM transaction WHERE transaction_id = ?", transactionId).Scan(&check)
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}

	if check != 1 {
		return false, nil
	}

	return true, nil
}

func (t *Transaction) SaveTransaction() (*Transaction, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &Transaction{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	number := 0
	err = db.QueryRow("SELECT COUNT(transaction_id) FROM transaction").Scan(&number)
	if err != nil {
		return &Transaction{}, errors.New("Somethings wrong!")
	}

	count := utility.DigitsCount(number)
	formatTransactionId := "TRC00000000"
	transactionId := "TRC"
	for i := 0; i < len(formatTransactionId)-count-5; i++ {
		transactionId += "0"
	}
	number += 1
	transactionId += fmt.Sprintf("%d", number)

	t.TransactionId = transactionId
	return t, nil
}

func (t *Transaction) FindTransactionById(transactionId int) (*Transaction, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &Transaction{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	err = db.QueryRow("SELECT transaction_id, customer_id, employee_id, date, created_at, updated_at FROM transaction WHERE transaction_id = ?", transactionId).
		Scan(&t.TransactionId, &t.CustomerId, &t.EmployeeId, &t.Date, &t.Audit.CreatedAt, &t.Audit.UpdatedAt)
	if err != nil {
		return &Transaction{}, errors.New("Somethings wrong!")
	}

	if t == (&Transaction{}) {
		return &Transaction{}, errors.New("Somethings wrong!")
	}

	return t, nil
}

func (t *Transaction) UpdateTransactionById(transactionId int) (*Transaction, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &Transaction{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("UPDATE transaction SET customer_id = ?, employee_id = ?, date = ?, created_at = ?, updated_at = ? WHERE transaction_id = ?",
		t.CustomerId,
		t.EmployeeId,
		t.Date,
		t.Audit.CreatedAt,
		t.Audit.UpdatedAt,
		transactionId)
	if err != nil {
		return &Transaction{}, errors.New("Somethings wrong!")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &Transaction{}, errors.New("Somethings wrong!")
	}

	if rowsAffected != 1 {
		return &Transaction{}, errors.New("Somethings wrong!")
	}

	return t, nil
}
