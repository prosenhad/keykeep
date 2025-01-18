package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"unicode/utf8"
)

const CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJLLMNOPQRSTUVWXYZ123456789!@#$%^&*"

type account struct {
	login    string
	password string
	url      string
}

func (a *account) GenPass(n int) {
	chars := []rune(CHARS)
	password := make([]rune, n)
	for i := range password {
		password[i] = chars[rand.IntN(len(chars))]
	}
	a.password = string(password)
}

func (a *account) Pass() {
	fmt.Println((*a).password)
}

func (a *account) Account() {
	fmt.Println(*a)
}

func NewAccount(userLogin string, userPass string, userURL string) (*account, error) {
	newAccount := account{}
	if utf8.RuneCountInString(userURL) == 0 {
		return nil, errors.New("пустой url")
	}
	_, err := url.ParseRequestURI(userURL)
	if err != nil {
		return nil, fmt.Errorf("неверный формат ссылки %s", userURL)
	}
	if utf8.RuneCountInString(userPass) != 0 && utf8.RuneCountInString(userPass) >= 8 {
		newAccount.password = userPass
	} else {
		fmt.Println("Пароль не безопасен. Был сгенерирован новый")
		utf8.RuneCountInString(userPass)
		newAccount.GenPass(10)
	}
	if utf8.RuneCountInString(userLogin) == 0 {
		return nil, errors.New("логин не может быть пустым")
	}
	newAccount.login = userLogin
	newAccount.url = userURL
	return &newAccount, nil
}

func SetPrompt(prompt string) string {
	fmt.Print(prompt + " :")
	var resp string
	fmt.Scanln(&resp)
	return resp
}

func main() {
	login := SetPrompt("Введите ваш Логин")
	password := SetPrompt("Введите ваш Пароль")
	url := SetPrompt("Введите url-адрес")
	account, err := NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())

	} else {
		fmt.Println(*account)
	}

}
