package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"os"
)

func Decrypt(path string) {

	cipherText, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
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

	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}

	err = os.WriteFile(path, plainText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
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
