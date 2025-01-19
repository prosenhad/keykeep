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

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJLLMNOPQRSTUVWXYZ123456789!@#$%^&*"

type accountUser struct {
	login    string `json:"login" xml:"xlogin"`
	password string `json:"password" xml:"xpassword"`
	url      string `json:"url" xml:"xurl"`
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	accountUser
}

func (a *accountWithTimeStamp) genPass(n int) {
	charsSlice := []rune(chars)
	password := make([]rune, n)
	for i := range password {
		password[i] = charsSlice[rand.IntN(len(charsSlice))]
	}
	a.password = string(password)
}

func (a *accountWithTimeStamp) Pass() {
	fmt.Println((*a).password)
}

func (a *accountWithTimeStamp) Account() {

	color.Cyan(a.login)
	color.Green(a.url)
	fmt.Println(a.createdAt.Date())
	fmt.Println(a.updatedAt.Date())
	color.Red(a.password)

}

func NewAccountWithTimestamp(userLogin string, userPass string, userURL string) (*accountWithTimeStamp, error) {
	newAccount := accountWithTimeStamp{}
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
		newAccount.genPass(10)
	}
	if utf8.RuneCountInString(userLogin) == 0 {
		return nil, errors.New("логин не может быть пустым")
	}
	newAccount.login = userLogin
	newAccount.url = userURL
	newAccount.createdAt = time.Now()
	newAccount.updatedAt = time.Now()
	return &newAccount, nil
}
