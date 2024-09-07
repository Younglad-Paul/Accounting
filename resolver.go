package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Memory struct {
	name       string
	CreditSide map[string]float64
	DebitSide  map[string]float64
	Sumation   map[string]float64
}

func Query(name string) Memory {
	query := Memory{
		name:       name,
		CreditSide: map[string]float64{},
		DebitSide:  map[string]float64{},
		Sumation:   map[string]float64{},
	}
	return query
}

func (query *Memory) AddToCredit(name string, Amount float64) {
	query.CreditSide[name] = Amount
}

func (query *Memory) SumCredit() string {
	var TotalCredit float64 = 0
	for _, value := range query.CreditSide {
		TotalCredit += value
	}
	return fmt.Sprint(TotalCredit)
}

func (query *Memory) AddToDebt(name string, Amount float64) {
	query.DebitSide[name] = Amount
}

func (query *Memory) SumDebit() string {
	var TotalDebit float64 = 0
	for _, value := range query.DebitSide {
		TotalDebit += value
	}
	return fmt.Sprint(TotalDebit)
}
func (query *Memory) TransferToSumation() {
	for name, amount := range query.CreditSide {
		query.Sumation[name] = amount
	}
	for name, amount := range query.DebitSide {
		query.Sumation[name] -= amount
	}
}
func (query *Memory) Format() string {
	Title := strings.ToUpper(query.name)

	query.TransferToSumation()

	for index, value := range query.Sumation {
		Title += fmt.Sprintf("\n %-30v  ...#%-30v\n", index, value)
	}

	SumCredits := query.SumCredit()
	SumTotalCredit, err := strconv.ParseFloat(SumCredits, 64)
	if err != nil {
		fmt.Println("There is an issue summing Credits, this is likely because the credit side is being summed incorrectly")
	}

	SumDebits := query.SumDebit()
	SumTotalDebit, err := strconv.ParseFloat(SumDebits, 64)
	if err != nil {
		fmt.Println("There is an issue summing Debit, this is likely because the debit side is being summed incorrectly")
	}

	Total := SumTotalCredit - SumTotalDebit

	Title += fmt.Sprintf("\n ________________________________________ \n %-30v  ...#%-30v \n ======================================== \n", "Balance:", Total)

	if SumTotalCredit > SumTotalDebit {
		Title += fmt.Sprintf("================> Profit <===============")
	} else if SumTotalDebit > SumTotalCredit {
		Title += fmt.Sprintf("=================> Loss <================")
	} else if SumTotalCredit == SumTotalDebit {
		Title += fmt.Sprintf("=================> Loss <================")
	}

	return Title
}

func (query *Memory) Save() {
	data := []byte(query.Format())
	err := os.WriteFile("Accounts/"+query.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("File saved successfully. Proceed to fetch account from the Accounts Folder")
}
