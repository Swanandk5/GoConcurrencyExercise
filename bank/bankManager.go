package bank

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ProcessTransactions() {
	var bankAccount = Account{
		balance:      100,
		transactions: []Transaction{},
	}

	transactions := []int{
		-200,
		100,
		500,
		-3200,
		1000,
		450,
	}

	wg := sync.WaitGroup{}

	for _, amount := range transactions {
		transaction := Transaction{
			amount:    amount,
			timestamp: time.Now(),
		}
		wg.Add(1)
		go processTransaction(transaction, &bankAccount, &wg)
		randomDelay := time.Duration(10+rand.Intn(990)) * time.Millisecond // 10–100ms
		time.Sleep(randomDelay)
	}
	wg.Wait()
	printTransactions(&bankAccount)
}

func printTransactions(account *Account) {
	for i, transaction := range account.transactions {
		fmt.Println("transaction", i, "amount", transaction.amount)
	}
	fmt.Println("account balance is ", account.balance)
}

func processTransaction(transaction Transaction, bankAccount *Account, wg *sync.WaitGroup) {
	bankAccount.mu.RLock()
	defer bankAccount.mu.RUnlock()
	bankAccount.balance += transaction.amount
	bankAccount.transactions = append(bankAccount.transactions, transaction)
	wg.Done()
}
