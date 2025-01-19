package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/prosenhad/keykeep/account"
)

func main() {
Menu:
	for {
		fmt.Println(strings.Repeat("_", 10))
		variant := Menu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			fmt.Println("Тут будет поиск аккаунта")
		case 3:
			fmt.Println("Тут будет удаление аккаунта")
		default:
			break Menu
		}
		fmt.Println(strings.Repeat("_", 10))
	}

}

func Menu() int {
	var choice int
	fmt.Println("1. Создать вкладку")
	fmt.Println("2. Найти вкладку")
	fmt.Println("3. Удалить вкладку")
	fmt.Println("4. Выйти")
	fmt.Print("Ваш выбор: ")
	fmt.Scanln(&choice)
	return choice
}

func createAccount() {
	login := SetPrompt("Введите ваш Логин")
	password := SetPrompt("Введите ваш Пароль")
	url := SetPrompt("Введите url-адрес")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*myAccount)

	// data, err := vault.ToBytes()
	// if err != nil {
	// 	fmt.Println("Не удалось преобразовать в JSON")
	// 	return
	// }
	// files.WriteIntoFile(data, "data.json")
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
