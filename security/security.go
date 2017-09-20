package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/robjporter/go-functions/format/as"
)

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func bufferSecurityKey(key string) string {
	if len(key) < 16 {
		return RightPad2Len(key, "0", 16)
	} else if len(key) > 16 && len(key) < 24 {
		return RightPad2Len(key, "0", 24)
	} else if len(key) > 24 && len(key) < 32 {
		return RightPad2Len(key, "0", 32)
	} else if len(key) > 32 {
		return key[:31]
	}
	return key
}

func Encrypt(key, text []byte) string {
	tmpKey := as.ToString(key)
	if len(tmpKey) != 16 && len(tmpKey) != 24 && len(tmpKey) != 32 {
		tmpKey = bufferSecurityKey(tmpKey)
	}
	key = []byte(tmpKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)
	return encodeBase64(ciphertext)
}

func Decrypt(key []byte, b64 string) string {
	tmpKey := as.ToString(key)
	if len(tmpKey) != 16 && len(tmpKey) != 24 && len(tmpKey) != 32 {
		tmpKey = bufferSecurityKey(tmpKey)
	}
	key = []byte(tmpKey)
	text := decodeBase64(b64)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(text) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return string(text)
}
