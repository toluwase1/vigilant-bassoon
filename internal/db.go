package internal

import (
	"fmt"
	"lemonadee/types"
	"net/http"
)

var db = make(map[string][]any)

const (
	UserTableName        = "users"
	TransactionTableName = "transactions"
)

func SaveToDB(tableName string, data any) (string, *Error) {
	db[tableName] = append(db[tableName], data)
	return retrieveID(tableName, data)
}

func GetByID(tableName string, ID string) (any, *Error) {
	if val, ok := db[tableName]; ok {
		for _, d := range val {
			id, err := retrieveID(tableName, d)
			if err != nil {
				return nil, err
			}
			if id == ID {
				return d, nil
			}
		}
	}

	return nil, NewError("not found", http.StatusNotFound)
}

func UpdateDB(tableName string, item any) *Error {
	for i, data := range db[tableName] {
		id1, err := retrieveID(tableName, data)
		if err != nil {
			return err
		}
		id2, err := retrieveID(tableName, item)
		if err != nil {
			return err
		}
		if id1 == id2 {
			db[tableName][i] = item
			return nil
		}
	}

	return NewError("not found", http.StatusNotFound)
}

func retrieveID(tableName string, data any) (string, *Error) {
	switch tableName {
	case UserTableName:
		return types.UserFromDB(data).ID, nil
	case TransactionTableName:
		return types.TransactionFromDB(data).ID, nil
	}
	return "", NewError("db error occurred", http.StatusInternalServerError)
}

func UpdateDbTx(f func() []any, tableName string) *Error {
	// Begin transaction
	fmt.Println("Transaction started")
	// Save a snapshot of the current database state
	snapshot := make(map[string][]any)
	for key, value := range db {
		snapshot[key] = append([]any{}, value...)
	}
	// Execute the provided function to get the data to be updated
	data := f()
	// Update the database
	for _, v := range data {
		err := UpdateDB(tableName, v)
		if err != nil {
			fmt.Println("Error: ", err)
			// Rollback to the snapshot of the database state
			db = snapshot
			fmt.Println("Transaction rolled back")
			return err
		}
	}

	// Commit the transaction
	fmt.Println("Transaction committed")
	return nil
}

func GetAllFromDB(tableName string) ([]any, *Error) {
	err := &Error{
		Message:    "no user found",
		StatusCode: http.StatusNotFound,
	}
	if len(db[tableName]) == 0 {
		return nil, err
	}
	return db[tableName], nil
}

func EmptyDB() {
	for tableName := range db {
		db[tableName] = nil
	}
}
