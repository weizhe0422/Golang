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
		errors.New("ERROR: 创建一个AES 加密的builder")
		return
	}

	planttext := `故经之以五事，校之以计而索其情：一曰道，二曰天，三曰地，四曰将，五曰法。道者，令民与上同意也，故可与之死，可与之生，而不畏危。天者，阴阳、寒暑、时制也。地者，高下、远近、险易、广狭、死生也。将者，智、信、仁、勇、严也。法者，曲制、官道、主用也。凡此五者，将莫不闻，知之者胜，不知者不胜。故校之以计而索其情，曰：主孰有道？将孰有能？天地孰得？法令孰行？兵众孰强？士卒孰练？赏罚孰明？吾以此知胜负矣。`

	ciphertext := cipher.Encrypt([]byte(planttext))

	ciphertextWithBase64 := base64.URLEncoding.EncodeToString(ciphertext)

	fmt.Println(ciphertextWithBase64)

	ciphertext, err = base64.URLEncoding.DecodeString(ciphertextWithBase64)
	if err != nil {
		//t.Error(err)
	}
	planttextBytes := cipher.Decrypt(ciphertext)

	fmt.Println(string(planttextBytes))

}
