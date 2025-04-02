package main

import (
	"GoConcurrencyExercise/api"
	"GoConcurrencyExercise/bank"
	"GoConcurrencyExercise/images"
)

func main() {
	///exercise 1
	images.Start()

	///exercise 2
	bank.ProcessTransactions()

	///exercise 3
	api.FetchData()
}
