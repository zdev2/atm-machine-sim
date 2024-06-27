package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"atm-machine/helpers"
)

type Card struct {
	NumberCard string `json:"card-num"`
	Username   string `json:"card-user"`
	Password   string `json:"card-pass"`
	Balance    int    `json:"card-balance"`
}

func main() {
	filePath := "./card-slot/card.json"
	helpers.Clear()
	fmt.Println("Masukkan kartu ATM ke dalam slot...")
	helpers.WaitForEnter()

	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("Kartu berhasil dimasukkan ke slot")
	} else {
		fmt.Println("Kartu gagal dimasukkan ke slot")
		return
	}

	login(filePath)
}

func login(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	var passInput string
	var card Card
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&card); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("Nomor: %v\nNama: %v\nPassword: ", card.NumberCard, card.Username)
	fmt.Scanln(&passInput)
	helpers.Clear()

	if passInput == card.Password {
		menu(card)
	} else {
		fmt.Println("Login gagal")
		fmt.Println("Silahkan coba lagi")
		helpers.WaitForEnter()
		helpers.Clear()
		login(filePath) // Reopen the file and retry login
	}
}

func menu(card Card) {
	helpers.Clear()
	fmt.Printf("Halo, %v\n", card.Username)
	fmt.Println("1. Cek Saldo")
	fmt.Println("2. Tarik Tunai")
	fmt.Println("3. Setor Tunai")
	fmt.Println("4. Keluar")
	fmt.Print("Pilihan: ")
	var input string
	fmt.Scanln(&input)
	helpers.Clear()

	switch input {
	case "1":
		checkBalance(card)
	case "2":
		withdraw(card)
	case "3":
		deposit(card)
	case "4":
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak tersedia")
		helpers.WaitForEnter()
		menu(card)
	}
}

func checkBalance(card Card) {
	fmt.Printf("Saldo Anda: %v\n", formatBalance(card.Balance))
	helpers.WaitForEnter()
	menu(card)
}

func withdraw(card Card) {
	var amount int
	fmt.Print("Masukkan jumlah yang akan ditarik: ")
	fmt.Scanln(&amount)

	if amount > card.Balance {
		fmt.Println("Saldo tidak mencukupi")
	} else {
		card.Balance -= amount
		fmt.Printf("Anda telah menarik %v\n", formatBalance(amount))
		
		// Save the updated card details
		if err := saveCardDetails("./card-slot/card.json", card); err != nil {
			fmt.Println("Error saving card details:", err)
		}
	}

	helpers.WaitForEnter()
	menu(card)
}

func deposit(card Card) {
	var amount int
	fmt.Print("Masukkan jumlah yang akan disetor: ")
	fmt.Scanln(&amount)

	card.Balance += amount
	fmt.Printf("Anda telah menyetor %v\n", formatBalance(amount))
	
	// Save the updated card details
	if err := saveCardDetails("./card-slot/card.json", card); err != nil {
		fmt.Println("Error saving card details:", err)
	}

	helpers.WaitForEnter()
	menu(card)
}

func formatBalance(balance int) string {
	balanceStr := fmt.Sprintf("%d", balance)
	var result strings.Builder

	length := len(balanceStr)
	for i, count := length-1, 0; i >= 0; i, count = i-1, count+1 {
		if count > 0 && count%3 == 0 {
			result.WriteString(".")
		}
		result.WriteByte(balanceStr[i])
	}

	// Reverse the string
	formattedBalance := []rune(result.String())
	for i, j := 0, len(formattedBalance)-1; i < j; i, j = i+1, j-1 {
		formattedBalance[i], formattedBalance[j] = formattedBalance[j], formattedBalance[i]
	}

	return "Rp" + string(formattedBalance)
}

func saveCardDetails(filePath string, card Card) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(card, "", "  ")
	if err != nil {
		return err
	}

	if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
}
