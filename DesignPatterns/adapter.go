package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Paying using cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using bank account %d\n", bankAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (adapter *BankPaymentAdapter) Pay() {
	adapter.BankPayment.Pay(adapter.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)
	bpa := &BankPaymentAdapter{
		bankAccount: 42,
		BankPayment: &BankPayment{},
	}
	ProcessPayment(bpa)
}
