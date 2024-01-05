package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"os"
)

func Encrypt(path string) {

	plainText, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	}

	/* key, err := os.ReadFile("assets/key.txt")
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	} */

	key := []byte("passphrasewhichneedstobe32bytes!")

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce  err: %v", err.Error())
	}

	cipherText := gcm.Seal(nonce, nonce, plainText, nil)

	err = os.WriteFile(path, cipherText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}

	/* var seed string

	fmt.Print("Enter your seed: ")
	fmt.Scanln(&seed) */

	/* files, _ := os.ReadDir(path)
	fmt.Println("Encripting ....")

	for _, f := range files {
		if !f.IsDir() {
			fullPath := filepath.Join(path, f.Name())
			encryptFile(fullPath)
		}
	} */
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
