package internal

import (
	"lemonadee/types"
	"net/http"
)

var db map[string][]any

const (
	UserTableName        = "users"
	TransactionTableName = "transactions"
)

func SaveToDB(tableName string, data any) (string, *Error) {
	db[tableName] = append(db[tableName], data)
	return retrieveID(tableName, data)
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
		return data.(types.User).ID, nil
	case TransactionTableName:
		return data.(types.Transactions).ID, nil
	}

	return "", NewError("db error occurred", http.StatusInternalServerError)
}
