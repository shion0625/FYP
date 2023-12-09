package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

func Encrypt(stringToEncrypt string, keyString string) (encryptedString string) {
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)
	key = key[:32]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return hex.EncodeToString(ciphertext)
}

func Decrypt(encryptedString string, keyString string) (decryptedString string) {
	enc, _ := hex.DecodeString(encryptedString)
	key, _ := hex.DecodeString(keyString)
	key = key[:32]
	fmt.Print("淀川")
	block, err := aes.NewCipher(key)
	fmt.Print("淀川")

	if err != nil {
		panic(err.Error())
	}
	fmt.Print("淀川")

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	fmt.Print("淀川")

	nonceSize := aesGCM.NonceSize()
	fmt.Print("淀川")

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	fmt.Print("淀川")

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Print("淀川だよーーー")
		fmt.Print(err)
		panic(err.Error())
	}
	fmt.Print("淀川")

	return string(plaintext)
}

func GetCardIssuer(cardNumber string) string {
	switch {
	case strings.HasPrefix(cardNumber, "4"):
		return "Visa"
	case strings.HasPrefix(cardNumber, "5"):
		return "MasterCard"
	case strings.HasPrefix(cardNumber, "34") || strings.HasPrefix(cardNumber, "37"):
		return "American Express"
	case strings.HasPrefix(cardNumber, "6"):
		return "Discover"
	case strings.HasPrefix(cardNumber, "36"):
		return "Diners Club"
	default:
		return "Unknown"
	}
}
