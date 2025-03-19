package eth

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func SendEth(ctx *gin.Context, sender *models.Wallet, receiver string, amount float64) error {
	var baseUrl string = os.Getenv("INFURA_BASE_URL")
	client, err := ethclient.Dial(baseUrl + "v3/" + os.Getenv("INFURA_API_KEY"))
	if err != nil {
		log.Println(err)
		return err
	}
	privateKey, err := crypto.HexToECDSA(sender.PrivateKey)
	if err != nil {
		log.Println(err)
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
		return fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddr)

	if err != nil {
		log.Println(err)
		return err
	}

	value := big.NewInt(10000000000000000)
	gasLimit := uint64(21000)
	tip := big.NewInt(2000000000)
	feeCap := big.NewInt(20000000000)

	toAddr := common.HexToAddress(receiver)
	var data []byte
	chainID, err := client.NetworkID(ctx.Request.Context())

	if err != nil {
		log.Println(err)
		return err
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: feeCap,
		GasTipCap: tip,
		Gas:       gasLimit,
		To:        &toAddr,
		Value:     value,
		Data:      data,
	})

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		log.Println(err)
		return err
	}

	err = client.SendTransaction(ctx.Request.Context(), signedTx)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(signedTx.Hash().Hex())
	return nil
}
