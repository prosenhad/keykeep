package main

import (
	"fmt"
	"reflect"

	"github.com/prosenhad/keykeep/account"
)

func main() {
	// files.WriteIntoFile("Test", "passwords.txt")
	login := SetPrompt("Введите ваш Логин")
	password := SetPrompt("Введите ваш Пароль")
	url := SetPrompt("Введите url-адрес")
	account, err := account.NewAccountWithTimestamp(login, password, url)
	if err != nil {
		fmt.Println(err.Error())

	} else {
		account.Account()
	}

	field, _ := reflect.TypeOf(account).Elem().FieldByName("password")
	fmt.Println(string(field.Tag))

}

func SetPrompt(prompt string) string {
	fmt.Print(prompt + " :")
	var resp string
	fmt.Scanln(&resp)
	return resp
}
