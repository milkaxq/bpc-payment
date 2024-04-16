package utils

import (
	"math/rand"
	"time"
)

func GenerateOrderNumber(minLen, maxLen int) string {
	rand.NewSource(time.Now().UnixNano())
	length := rand.Intn(maxLen-minLen+1) + minLen
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var orderNumber string
	for i := 0; i < length; i++ {
		orderNumber += string(chars[rand.Intn(len(chars))])
	}
	return orderNumber
}
