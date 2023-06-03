package model

import (
	"errors"
	"fmt"

	"sales-backend/utility"
)

type Transaction struct {
	TransactionId string `json:"transactionId"`
	CustomerId    string `json:"customerId"`
	EmployeeId    string `json:"employeeId"`
	Date          string `json:"date"`
	Audit         Audit  `json:"audit"`
}

func (t *Transaction) IsTransactionExistsById(transactionId string) (bool, error) {
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

	_, err = db.Exec("INSERT INTO transaction (transaction_id, customer_id, employee_id, date, created_at) VALUES (?,?,?,?,?)",
		t.TransactionId,
		t.CustomerId,
		t.EmployeeId,
		t.Date,
		t.Audit.CreatedAt)
	if err != nil {
		return &Transaction{}, errors.New("Somethings wrong!")
	}

	return t, nil
}

func (t *Transaction) FindTransactionById(transactionId string) (*Transaction, error) {
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

func (t *Transaction) UpdateTransactionById(transactionId string) (*Transaction, error) {
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

func (t *Transaction) DeleteTransactionByid(transactionId string) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return false, errors.New("Somethings wrong!")
	}

	check := 0
	err = tx.QueryRow("SELECT COUNT(transaction_detail_id) FROM transaction_detail WHERE transaction_id = ?", transactionId).Scan(&check)
	if err != nil {
		tx.Rollback()
		return false, errors.New("Somethings wrong!")
	}

	if check >= 1 {
		tx.Rollback()
		return false, errors.New("Can't delete the transaction because have details!")
	}

	result, err := tx.Exec("DELETE FROM transaction WHERE transaction_id = ?", transactionId)
	if err != nil {
		tx.Rollback()
		return false, errors.New("Somethings wrong!")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return false, errors.New("Somethings wrong!")
	}

	if rowsAffected != 1 {
		tx.Rollback()
		return false, errors.New("Somethings wrong!")
	}

	return true, nil
}

func (t *Transaction) FindAllTransaction(limit int, offset int) ([]*Transaction, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return []*Transaction{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	rows, err := db.Query("SELECT transaction_id, customer_id, employee_id, date, created_at, updated_at FROM transaction LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return []*Transaction{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result := []*Transaction{}
	for rows.Next() {
		each := &Transaction{}
		err = rows.Scan(&each.TransactionId, &each.CustomerId, &each.EmployeeId, &each.Date, &each.Audit.CreatedAt, &each.Audit.UpdatedAt)
		if err != nil {
			return []*Transaction{}, errors.New("Somethings wrong!")
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return []*Transaction{}, errors.New("Somethings wrong!")
	}

	return result, nil
}
