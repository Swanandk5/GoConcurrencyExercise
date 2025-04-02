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
		30,
		-200,
		100,
		500,
		-3200,
		1000,
		450,
		900,
		30,
		0,
		123,
		44,
		35,
		90,
		123,
		256,
		345,
		667,
		134,
		889,
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
	bankAccount.Lock()
	defer bankAccount.Unlock()
	defer wg.Done()
	bankAccount.balance += transaction.amount
	bankAccount.transactions = append(bankAccount.transactions, transaction)
}
