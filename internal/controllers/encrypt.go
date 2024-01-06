package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(path string, pss []byte) {

	/* var seed string

	fmt.Print("Enter your seed: ")
	fmt.Scanln(&seed) */

	// dk := pbkdf2.Key(pss, nonce, 4096, 32, sha1.New)

	files, _ := os.ReadDir(path)
	fmt.Println("Encripting ....")

	for _, f := range files {
		if !f.IsDir() {
			fullPath := filepath.Join(path, f.Name())
			encryptFile(fullPath, pss)
		}
	}
}

/*func encryptFile(filename string) {

	 key := make([]byte, 16) // 128 bit key
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return err
	}
	fmt.Print("the key is: ")
	fmt.Println(key)

	// key := []byte("passphrasewhichneedstobe32bytes!") // replace with your key

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	nonce := make([]byte, 12) // 96 bit nonce
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return err
	}
	fmt.Print("the nonce is: ")
	fmt.Println(nonce)

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	ciphertext := aesgcm.Seal(nil, nonce, data, nil)

	fmt.Print("the ciphertext is: ")
	fmt.Println(ciphertext)

	err = os.WriteFile(filename, ciphertext, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("\nThe file %s was encrypted.\n", filename)

	return nil
}*/

func encryptFile(filename string, pss []byte) {

	plainText, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	}

	// key := []byte("passphrasewhichneedstobe32bytes!")

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce  err: %v", err.Error())
	}

	key := pbkdf2.Key(pss, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	cipherText := gcm.Seal(nonce, nonce, plainText, nil)

	err = os.WriteFile(filename, cipherText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}

}
