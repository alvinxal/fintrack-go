package main

import "fmt"

const (
	Income  = "income"
	Expense = "expense"
)

type Transaction struct {
	Nama   string
	Tipe   string
	Amount int
}

func applyTransaction(saldo int, t Transaction) int {
	switch t.Tipe {
	case Income:
		saldo += t.Amount
	case Expense:
		saldo -= t.Amount
	default:
		fmt.Println("Tipe transaksi tidak ada:", t.Tipe)
	}

	return saldo
}

func printTransactions(transactions []Transaction) {
	fmt.Println("----")
	for _, t := range transactions {
		fmt.Println("Transaksi:", t.Nama)
		fmt.Println("Jenis transaksi:", t.Tipe)
		fmt.Println("Jumlah:", t.Amount)
		fmt.Println("----")
	}

}

func main() {
	saldo := 0

	transactions := []Transaction{
		{Nama: "Gaji", Tipe: Income, Amount: 4500000},
		{Nama: "Makan", Tipe: Expense, Amount: 25000},
	}

	for _, t := range transactions {
		saldo = applyTransaction(saldo, t)
	}

	printTransactions(transactions)
	fmt.Println("Saldo akhir: ", saldo)
}
