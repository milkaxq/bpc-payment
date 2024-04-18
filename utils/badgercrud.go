package utils

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/milkaxq/bpcpayment/config"
)

func CreateNewEntry(orderID string, data []byte, lifeTime int64) error {
	err := config.DB.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(orderID), data).WithTTL(time.Duration(lifeTime) * time.Minute)
		err := txn.SetEntry(e)
		return err
	})
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func GetEntry(orderID string) ([]byte, int64, error) {
	var valCopy []byte
	var expiresAt int64
	err := config.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(orderID))
		if err != nil {
			return err
		}
		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}
		expiresAt = int64(item.ExpiresAt())
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return valCopy, expiresAt, nil
}

func GetBankModel(orderID string) (config.BankModel, int64, error) {
	var bankModel config.BankModel

	data, expiresAt, err := GetEntry(orderID)
	if err != nil {
		return config.BankModel{}, 0, err
	}
	json.Unmarshal(data, &bankModel)

	return bankModel, expiresAt, nil
}
