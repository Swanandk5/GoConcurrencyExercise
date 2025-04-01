package main

import (
	"GoConcurrencyExercise/bank"
	"GoConcurrencyExercise/images"
)

func main() {
	///exercise 1
	images.Start()

	///exercise 2
	bank.ProcessTransactions()
}
