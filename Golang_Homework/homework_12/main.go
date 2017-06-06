package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"errors"

	"github.com/89hmdys/toast/cipher"
	"github.com/89hmdys/toast/crypto"
)

func KeyAES() []byte {
	var err error
	var key = make([]byte, 32)
	if _, err = rand.Read(key); err != nil {
		panic(fmt.Errorf("Error generate aes key\n\n%+v", err))
	}
	return key
}

func main() {
	mode := cipher.NewCBCMode() //加密工作模式，支持 CBC ECB CFB CTR 四种工作模式
	key := KeyAES()
	fmt.Println("KEY: ", key)
	cipher, err := crypto.NewAESWith([]byte(key), mode) //创建一个AES 加密的builder
	if err != nil {
		errors.New("ERROR: 建立AES加密的builder")
		return
	}

	planttext := `我愛台灣`

	ciphertext := cipher.Encrypt([]byte(planttext))

	ciphertextWithBase64 := base64.URLEncoding.EncodeToString(ciphertext)

	fmt.Println("Ciphertext:", ciphertextWithBase64)

	ciphertext, err = base64.URLEncoding.DecodeString(ciphertextWithBase64)
	if err != nil {
		//t.Error(err)
	}
	planttextBytes := cipher.Decrypt(ciphertext)

	fmt.Println(string(planttextBytes))

}
