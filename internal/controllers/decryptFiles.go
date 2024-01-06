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

func Decrypt(path string, pss []byte) {

	files, _ := os.ReadDir(path)
	fmt.Println("Decripting ....")

	for _, f := range files {
		if !f.IsDir() {
			fullPath := filepath.Join(path, f.Name())
			decryptFile(fullPath, pss)
		}
	}

}

func decryptFile(filename string, pss []byte) {
	cipherText, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	nonce := cipherText[:12]

	key := pbkdf2.Key(pss, nonce, 4096, 32, sha1.New)

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
