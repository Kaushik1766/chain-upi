package trx

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func AddWalletByPrivateKey(privateKey string) {
	pKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		// return nil, nil
		fmt.Println(err.Error())
	}
	// publicKey := pKey.Public()
	fmt.Println(string(crypto.FromECDSAPub(&pKey.PublicKey)))
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// fmt.Println(public)
	// if !ok {
	// 	fmt.Println("error")
	// }

}
