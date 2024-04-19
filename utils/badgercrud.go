package utils

import (
	"encoding/json"

	"github.com/dgraph-io/badger/v4"
	"github.com/milkaxq/bpcpayment/config"
)

func CreateNewEntry(orderID string, data []byte) error {
	err := config.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(orderID), data)
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func GetEntry(orderID string) ([]byte, error) {
	var valCopy []byte
	err := config.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(orderID))
		if err != nil {
			return err
		}
		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return valCopy, nil
}

func GetBankModel(orderID string) (config.BankModel, error) {
	var bankModel config.BankModel

	data, err := GetEntry(orderID)
	if err != nil {
		return config.BankModel{}, err
	}
	json.Unmarshal(data, &bankModel)

	return bankModel, nil
}
