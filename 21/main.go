package main

import "fmt"

// BankCard - интерфейс, который представляет банковскую карту
type BankCard interface {
	InsertCard()
	Withdraw(amount float64)
}

// VisaCard - интерфейс, представляющий карту Visa
type VisaCard interface {
	ConnectToVisaSystem()
	Debit(amount float64)
}

// MasterCard - интерфейс, представляющий карту MasterCard
type MasterCard interface {
	ConnectToMasterCardSystem()
	Charge(amount float64)
}

// Visa - конкретная реализация карты Visa
type Visa struct{}

func (v *Visa) ConnectToVisaSystem() {
	fmt.Println("Connecting to Visa system...")
}

func (v *Visa) Debit(amount float64) {
	fmt.Printf("Debiting %f from Visa card\n", amount)
}

// MasterCard - конкретная реализация карты MasterCard
type Master struct{}

func (m *Master) ConnectToMasterCardSystem() {
	fmt.Println("Connecting to MasterCard system...")
}

func (m *Master) Charge(amount float64) {
	fmt.Printf("Charging %f from MasterCard\n", amount)
}

// VisaAdapter - адаптер для работы с Visa через интерфейс BankCard
type VisaAdapter struct {
	visaCard *Visa
}

func (va *VisaAdapter) InsertCard() {
	va.visaCard.ConnectToVisaSystem()
}

func (va *VisaAdapter) Withdraw(amount float64) {
	va.visaCard.Debit(amount)
}

// MasterCardAdapter - адаптер для работы с MasterCard через интерфейс BankCard
type MasterCardAdapter struct {
	masterCard *Master
}

func (ma *MasterCardAdapter) InsertCard() {
	ma.masterCard.ConnectToMasterCardSystem()
}

func (ma *MasterCardAdapter) Withdraw(amount float64) {
	ma.masterCard.Charge(amount)
}

// ATM - класс банкомата, использующий BankCard для выполнения операций
type ATM struct {
	card BankCard
}

func (a *ATM) InsertCard(card BankCard) {
	fmt.Println("Card inserted.")
	a.card = card
	a.card.InsertCard()
}

func (a *ATM) WithdrawCash(amount float64) {
	if a.card == nil {
		fmt.Println("No card inserted.")
		return
	}
	a.card.Withdraw(amount)
}

func main() {
	atm := &ATM{}

	// Работа с картой Visa
	visaCard := &Visa{}
	visaAdapter := &VisaAdapter{visaCard: visaCard}

	atm.InsertCard(visaAdapter)
	atm.WithdrawCash(100)

	fmt.Println("")

	// Работа с картой MasterCard
	masterCard := &Master{}
	masterCardAdapter := &MasterCardAdapter{masterCard: masterCard}

	atm.InsertCard(masterCardAdapter)
	atm.WithdrawCash(200)
}
