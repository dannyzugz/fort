package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/crypto/pbkdf2"
)

func Decrypt(path string) {

	pss := ValidatePass()
	files, _ := os.ReadDir(path)
	fmt.Println("Decripting ....")

	for _, f := range files {
		if !f.IsDir() {
			fullPath := filepath.Join(path, f.Name())
			decryptFile(fullPath, pss)
		}
	}

	/* var seed string

	fmt.Print("Enter your seed: ")
	fmt.Scanln(&seed)

	key := []byte(seed) */

	/* ciphertext, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	// key := []byte("passphrasewhichneedstobe32bytes!") // replace with your key
	key := []byte{95, 215, 224, 6, 46, 92, 209, 244, 24, 211, 21, 191, 16, 128, 33, 166} // replace with your key
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(ciphertext))
	// fmt.Println(len(nonce))
	// nonceSize := gcm.NonceSize()
	if len(ciphertext) < 12 {
		fmt.Println(err)
	}
	// nonce, ciphertext := ciphertext[:12], ciphertext[12:]
	nonce := []byte{202, 38, 69, 52, 200, 208, 175, 111, 90, 78, 151, 142}
	ciphertext = ciphertext[12:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(plaintext)) */

}

func decryptFile(filename string, pss []byte) {
	cipherText, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	nonce := cipherText[:12]

	key := pbkdf2.Key(pss, nonce, 4096, 32, sha1.New)
	/* key, err := os.ReadFile("assets/key.txt")
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	} */

	// key := []byte("passphrasewhichneedstobe32bytes!")

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	cipherText = cipherText[12:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}

	err = os.WriteFile(filename, plainText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
}
