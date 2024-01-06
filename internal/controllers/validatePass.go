package controllers

import (
	// "bufio"
	"fmt"
	"log"
	// "os"
	// "golang.org/x/term"
)

func ValidatePass() []byte {
	/* if bytes.Equal(password1, password2) {
		return true
	} else {
		return false
	} */

	fmt.Println("Enter password")
	// pss, _ := term.ReadPassword(0)
	pss, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return []byte(pss)
}
