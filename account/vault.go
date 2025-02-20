package account

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/prosenhad/keykeep/files"
)

// Новая структура для массива аккаунтов
type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// метод для перевода данных в байт-массив
func (v *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (v *Vault) AddAccount(acc Account) {
	v.Accounts = append(v.Accounts, acc)
	v.UpdatedAt = time.Now()
	data, err := v.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать в JSON")
	}
	files.WriteIntoFile(data, "data.json")
}

func (v *Vault) GetAccountByURL(url string) []Account {
	var foundAccounts []Account
	for _, acc := range v.Accounts {
		isMatched := strings.Contains(acc.Url, url)
		if isMatched {
			foundAccounts = append(foundAccounts, acc)
		}
	}
	return foundAccounts
}

func (v *Vault) DelAccountByURL(url string) bool {
	var untrackAccounts []Account
	isDeleted := false
	for _, acc := range v.Accounts {
		isMatched := strings.Contains(acc.Url, url)
		if !isMatched {
			untrackAccounts = append(untrackAccounts, acc)
			continue
		}
		isDeleted = true
	}
	v.Accounts = untrackAccounts
	v.UpdatedAt = time.Now()
	data, err := v.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать в JSON")
	}
	files.WriteIntoFile(data, "data.json")
	return isDeleted
}

// Дефолтное создание нового хранилища для МАССИВА аккаунтов
func NewVault() *Vault {
	file, err := files.ReadFromFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	// Возвращаем из файла байти и переделываем обратно в go-шные данные
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось преобразовать байты в Json. Файл был перезаписан")
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
		// log.Fatal(err)
	}
	return &vault
}
