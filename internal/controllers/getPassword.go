package controllers

import (
	// "bytes"
	"bufio"
	"fmt"
	"log"
	"os"
	// "golang.org/x/term"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var reader2 *bufio.Reader = bufio.NewReader(os.Stdin)

func GetPassword() []byte {

	var pss1, pss2 string

	fmt.Println("Enter password")
	// pss1, _ := term.ReadPassword(0)
	// fmt.Scanln(pss1)

	pss1, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Confirm the password")
	// pss2, _ := term.ReadPassword(0)
	// fmt.Scanln(pss2)

	pss2, err2 := reader2.ReadString('\n')
	if err2 != nil {
		log.Fatal(err2)
	}

	if pss1 == pss2 {
		return []byte(pss1)
	} else {
		fmt.Println("The passwords don t match. Try again")
		return GetPassword()
	}

}
