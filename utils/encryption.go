package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"os"
)

var bKey = os.Getenv("AES_KEY")
var bIV = RandomBytes(aes.BlockSize)

func EncryptAES256(data string) string {
	bPlaintext := PKCS5Padding([]byte(data), aes.BlockSize)
	block, _ := aes.NewCipher([]byte(bKey))
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func DecryptAES256(data string) string {
	cipherTextDecoded, _ := base64.StdEncoding.DecodeString(data)

	block, _ := aes.NewCipher([]byte(bKey))

	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks(cipherTextDecoded, cipherTextDecoded)
	return string(cipherTextDecoded)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func RandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

var hashes []string

func HashingCollision(data string) string {
	for _, hash := range hashes {
		if LevenshteinDistance(data, hash) < 4 {
			return hash
		}
	}
	hashes = append(hashes, data)
	return data
}
