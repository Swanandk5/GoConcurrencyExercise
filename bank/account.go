package bank

import (
	"sync"
	"time"
)

type Account struct {
	balance      int
	transactions []Transaction
	mu           sync.RWMutex
}

type Transaction struct {
	amount    int
	timestamp time.Time
}
