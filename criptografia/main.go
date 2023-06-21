package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// This should be in an env file in production
const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func main() {

	token := "WyMU4gwdQv1TMpmmJ3qBUB+XK8FFqR/ZTxMhsIXYrIZoSWXaXoat7IYIFVqk0EUxCApoJA=="

	decText, err := Decrypt(token, MySecret)
	if err != nil {
		fmt.Println("error decrypting your encrypted text: ", err)
	}
	fmt.Println(decText)

	// StringToEncrypt := "0c3ad1ae206543e1af32 asp42349803jasljl asjlasfjlaipi"
	// encText, err := Encrypt(StringToEncrypt, MySecret)
	// if err != nil {
	// 	fmt.Println("error encrypting your classified text: ", err)
	// }
	// fmt.Println(encText)
}
