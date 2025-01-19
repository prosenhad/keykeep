package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
	"unicode/utf8"

	"github.com/fatih/color"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJLLMNOPQRSTUVWXYZ123456789!#*"

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (a *Account) genPass(n int) {
	charsSlice := []rune(chars)
	password := make([]rune, n)
	for i := range password {
		password[i] = charsSlice[rand.IntN(len(charsSlice))]
	}
	a.Password = string(password)
}

func (a *Account) Pass() {
	fmt.Println((*a).Password)
}

func (a *Account) GetAccount() {
	if a == nil {
		fmt.Println("НИЧЕГО НЕ ЗАПИСАЛОСЬ")
		return
	}
	color.Cyan(a.Login)
	color.Green(a.Url)
	fmt.Println(a.CreatedAt.Date())
	// fmt.Println(a.UpdatedAt.Date())
	color.Red(a.Password)
}

func NewAccount(userLogin string, userPass string, userURL string) (*Account, error) {
	newAccount := Account{}
	if utf8.RuneCountInString(userURL) == 0 {
		return nil, errors.New("пустой url")
	}
	_, err := url.ParseRequestURI(userURL)
	if err != nil {
		return nil, fmt.Errorf("неверный формат ссылки %s", userURL)
	}
	if utf8.RuneCountInString(userPass) != 0 && utf8.RuneCountInString(userPass) >= 8 {
		newAccount.Password = userPass
	} else {
		fmt.Println("Пароль не безопасен. Был сгенерирован новый")
		utf8.RuneCountInString(userPass)
		newAccount.genPass(10)
	}
	if utf8.RuneCountInString(userLogin) == 0 {
		return nil, errors.New("логин не может быть пустым")
	}
	newAccount.Login = userLogin
	newAccount.Url = userURL
	newAccount.CreatedAt = time.Now()
	newAccount.UpdatedAt = time.Now()
	return &newAccount, nil
}
