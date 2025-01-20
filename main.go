package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/prosenhad/keykeep/account"
)

func main() {
	vault := account.NewVault()

Menu:
	for {
		fmt.Println(strings.Repeat("_", 10))
		variant := Menu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
		fmt.Println(strings.Repeat("_", 10))
	}

}

func Menu() int {
	fmt.Println("Hello")
	var choice int
	fmt.Println("1. Создать вкладку")
	fmt.Println("2. Найти вкладку")
	fmt.Println("3. Удалить вкладку")
	fmt.Println("4. Выйти")
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&choice)
	return choice
}

func createAccount(vault *account.Vault) {
	login := SetPrompt("Введите ваш Логин")
	password := SetPrompt("Введите ваш Пароль")
	url := SetPrompt("Введите url-адрес")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.Vault) {
	url := SetPrompt("Введите url или его часть")
	accounts := vault.GetAccountByURL(url)
	for _, acc := range accounts {
		acc.GetAccount()
	}

}

func deleteAccount(vault *account.Vault) {
	url := SetPrompt("Введите url или его часть")
	isDeleted := vault.DelAccountByURL(url)
	if isDeleted {
		color.Green("Ссылка(и) удалена(ы)")
	} else {
		color.Red("Не найдено ссылок с таким именем")
	}

}

func SetPrompt(prompt string) string {
	var resp string
	fmt.Print(prompt + " :")
	reader := bufio.NewReader(os.Stdin)
	ud, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	resp = strings.TrimSpace(ud)
	return resp
}
