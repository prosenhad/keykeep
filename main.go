package main

import (
	"fmt"

	"github.com/prosenhad/keykeep/account"
)

func main() {
	login := SetPrompt("Введите ваш Логин")
	password := SetPrompt("Введите ваш Пароль")
	url := SetPrompt("Введите url-адрес")
	account, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())

	} else {
		account.Account()

	}

}

func SetPrompt(prompt string) string {
	fmt.Print(prompt + " :")
	var resp string
	fmt.Scanln(&resp)
	return resp
}
