package bank

import (
	"sync"
	"time"
)

type Account struct {
	balance      int
	transactions []Transaction
	sync.RWMutex
}

type Transaction struct {
	amount    int
	timestamp time.Time
}
