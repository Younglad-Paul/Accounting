package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetUserInput(Prompt string, read *bufio.Reader) (string, error) {
	fmt.Print(Prompt)

	inpute, err := read.ReadString('\n')

	return strings.TrimSpace(inpute), err

}

func CreateAccount() Memory {
	Read := bufio.NewReader(os.Stdin)

	userInput, _ := GetUserInput("Create a new account Sheet: ", Read)

	Data := Query(userInput)

	return Data

}

func RunCommand(Data Memory) {
	Read := bufio.NewReader(os.Stdin)

	SelectType, _ := GetUserInput("Select Account Type: (C - Credit, D - Debit, S - to save, E - to exit and X to discard account. Alternatively you can force discard the process by holding ctrl + c)", Read)

	switch SelectType {
	case "C":
		name, _ := GetUserInput("Description: ", Read)

		if name == "E" {
			confirmReturn, _ := GetUserInput("Your input is 'E' which is a reserved function to return to discard current input. Returnig to root will discard this current Data. Press 'Y' to confirm or press 'N' to continue:   ", Read)

			switch confirmReturn {
			case "Y":
				fmt.Println("Returning to root...")
				RunCommand(Data)
			case "N":
				fmt.Println("Data retained.")
				name, _ := GetUserInput("Description: ", Read)
				amount, _ := GetUserInput("Amount: ", Read)

				Amount, err := strconv.ParseFloat(amount, 64)

				if err != nil {
					fmt.Println("Amount must be a number")
					RunCommand(Data)
				}

				Data.AddToDebt(name, Amount)
				RunCommand(Data)
			}
		}

		amount, _ := GetUserInput("Amount: ", Read)

		if amount == "E" {
			confirmReturn, _ := GetUserInput("Your input is 'E' which is a reserved function to return to discard current input. Returnig to root will discard this current Data. Press 'Y' to confirm or press 'N' to continue:   ", Read)

			switch confirmReturn {
			case "Y":
				fmt.Println("Returning to root...")
				RunCommand(Data)
			case "N":
				fmt.Println("Data retained.")
				name, _ := GetUserInput("Description: ", Read)
				amount, _ := GetUserInput("Amount: ", Read)

				Amount, err := strconv.ParseFloat(amount, 64)

				if err != nil {
					fmt.Println("Amount must be a number")
					RunCommand(Data)
				}

				Data.AddToDebt(name, Amount)
				RunCommand(Data)
			}
		}

		Amount, err := strconv.ParseFloat(amount, 64)

		if err != nil {
			fmt.Println("Amount must be a number")
			RunCommand(Data)
		}

		Data.AddToCredit(name, Amount)
		RunCommand(Data)

	case "D":
		name, _ := GetUserInput("Description: ", Read)

		if name == "E" {
			confirmReturn, _ := GetUserInput("Your input is 'E' which is a reserved function to return to discard current input. Returnig to root will discard this current Data. Press 'Y' to confirm or press 'N' to continue:   ", Read)

			switch confirmReturn {
			case "Y":
				fmt.Println("Returning to root...")
				RunCommand(Data)
			case "N":
				fmt.Println("Data retained.")
				name, _ := GetUserInput("Description: ", Read)
				amount, _ := GetUserInput("Amount: ", Read)

				Amount, err := strconv.ParseFloat(amount, 64)

				if err != nil {
					fmt.Println("Amount must be a number")
					RunCommand(Data)
				}

				Data.AddToDebt(name, Amount)
				RunCommand(Data)
			}
		}

		amount, _ := GetUserInput("Amount: ", Read)

		if amount == "E" {
			confirmReturn, _ := GetUserInput("Your input is 'E' which is a reserved function to return to root. Returnig to root will discard this current Data. Press 'Y' to confirm or press 'N' to continue", Read)

			switch confirmReturn {
			case "Y":
				fmt.Println("Returning to root...")
				RunCommand(Data)
			case "N":
				fmt.Println("Data retained.")
				name, _ := GetUserInput("Description: ", Read)
				amount, _ := GetUserInput("Amount: ", Read)

				Amount, err := strconv.ParseFloat(amount, 64)

				if err != nil {
					fmt.Println("Amount must be a number")
					RunCommand(Data)
				}

				Data.AddToDebt(name, Amount)
				RunCommand(Data)
			}
		}

		Amount, err := strconv.ParseFloat(amount, 64)

		if err != nil {
			fmt.Println("Amount must be a number")
			RunCommand(Data)
		}

		Data.AddToDebt(name, Amount)
		RunCommand(Data)

	case "S":
		Data.Save()
		fmt.Println("Data saved successfully")

	case "E":
		fmt.Println("Returning to root...")
		RunCommand(Data)

	case "X":
		ConfirmDiscard, _ := GetUserInput("Are you sure you want to discard Account? (Y/N): ", Read)

		switch ConfirmDiscard {
		case "Y":
			fmt.Println("Cancelled Account. Data discarded successfully")
		case "N":
			fmt.Println("Account Retained. Continue adding content.")
			RunCommand(Data)
		}

	default:
		fmt.Println("Invalid Operation. This may be because you need to use capital letters to select option (Turn on your Caps lock and try again).")
		RunCommand(Data)
	}

}
