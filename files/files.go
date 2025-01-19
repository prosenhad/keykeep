package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ReadFromFile(filename string) ([]byte, error) {
	fn, _ := os.UserHomeDir()
	// if err != nil {
	// 	color.Red(err.Error())
	// 	return nil, err
	// }
	fn += "/" + filename
	data, err := os.ReadFile(fn)
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}
	return data, nil

	// file, err := os.Open(fn)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text(), "ENDING")
	// }
	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

}

func WriteIntoFile(content []byte, fileName string) {
	hd, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Create(hd + "/" + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись прошла")

}
