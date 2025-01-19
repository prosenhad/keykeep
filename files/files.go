package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFromFile(filename string) {
	fn, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	fn += "/" + filename
	content, err := os.Open(fn)
	if err != nil {
		fmt.Println(err)
	}
	defer content.Close()
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		fmt.Println(scanner.Text(), "ENDING")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func WriteIntoFile(content string, fileName string) {
	fmt.Println(os.Getwd())
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

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись прошла")

}
