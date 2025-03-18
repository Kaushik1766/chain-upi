package models

import (
	"fmt"
	"time"
)

type Transaction struct {
	Amount          string    `json:"amount"`
	ReceiverAddress string    `json:"receiverAddress"`
	TransactionHash string    `json:"transactionHash"`
	TimeStamp       time.Time `json:"timeStamp"`
	SenderAddress   string    `json:"senderAddress"`
	Chain           string    `json:"chain"`
}

func (t *Transaction) ToString() string {
	s := fmt.Sprintf("Time=%s", t.TimeStamp.UTC())
	return s
}
