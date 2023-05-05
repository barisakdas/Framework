package main

import (
	"fmt"

	crypto "github.com/barisakdas/Framework/Cryptography/Extensions"
)

func main() {

	/************ CRYPTOGRAPHY ************/
	var key = "{my_private_key}"
	var text = "{my_text}"
	var _cryptoService = crypto.NewFileLoggerService([]byte(key))
	encryptedData, _ := _cryptoService.Encrypt(text)
	decryptedData, _ := _cryptoService.Decrypt(encryptedData)
	fmt.Println("EncryptedData: ", encryptedData)
	fmt.Println("DecryptedData: ", decryptedData)

}
