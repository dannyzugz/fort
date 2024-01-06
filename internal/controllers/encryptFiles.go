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

	files, _ := os.ReadDir(path)
	fmt.Println("Encripting ....")

	for _, f := range files {
		if !f.IsDir() {
			fullPath := filepath.Join(path, f.Name())
			encryptFile(fullPath, pss)
		}
	}
}

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
