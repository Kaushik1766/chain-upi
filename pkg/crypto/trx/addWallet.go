package trx

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58/base58"
	"golang.org/x/crypto/sha3"
)

func privateKeyToAdress(privateKey string) (string, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	var arr []byte

	x := key.X
	arr = append(arr, x.Bytes()...)

	y := key.Y
	arr = append(arr, y.Bytes()...)

	kec256 := crypto.Keccak256(arr)
	last40 := kec256[12:]
	append41, err := hex.DecodeString("41")
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	append41 = append(append41, last40...)
	hash1 := sha256.Sum256(append41)
	hash2 := sha256.Sum256(hash1[:])
	checksum := hash2[:4]
	append41 = append(append41, checksum...)
	// fmt.Println(hex.EncodeToString(append41))
	address := base58.Encode(append41)
	// fmt.Println(address)
	return address, nil
}

func PrivateKeyToWallet(privateKey string) (*string, error) {
	address, err := privateKeyToAdress(privateKey)
	if err != nil {
		return nil, err
	}
	// wallet.Address = address
	// wallet.Chain = "trx"
	// wallet.PrivateKey = privateKey
	ok := ValidateAddress(address)
	if !ok {
		return nil, fmt.Errorf("INVALID WALLET ADDRESS")
	}
	return &address, nil
}

const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func encodeBase58Check(input []byte) string {

	h := sha256.Sum256(input)
	h2 := sha256.Sum256(h[:])
	checksum := h2[:4]

	withChecksum := append(input, checksum...)

	num := new(big.Int).SetBytes(withChecksum)
	base := big.NewInt(58)
	result := ""

	for num.Sign() > 0 {
		mod := new(big.Int)
		num.DivMod(num, base, mod)
		result = string(base58Alphabet[mod.Int64()]) + result
	}

	for _, b := range input {
		if b == 0 {
			result = "1" + result
		} else {
			break
		}
	}

	return result
}

func GenWallet() (string, string, string, error) {
	// Step 1: Generate secp256k1 private key
	privKey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to generate private key: %v", err)
	}

	// Step 2: Convert private key to hex
	privateKeyHex := hex.EncodeToString(privKey.D.Bytes())

	// Step 3: Serialize public key (Uncompressed format: 0x04 + X + Y)
	pubKeyBytes := append([]byte{0x04}, append(privKey.PublicKey.X.Bytes(), privKey.PublicKey.Y.Bytes()...)...)

	// Step 4: Hash public key using SHA3 Keccak256 (last 20 bytes)
	keccakHash := sha3.NewLegacyKeccak256()
	keccakHash.Write(pubKeyBytes[1:]) // Exclude 0x04 prefix
	hashedPubKey := keccakHash.Sum(nil)

	// Step 5: Create raw address with TRON prefix (0x41)
	rawAddress := append([]byte{0x41}, hashedPubKey[len(hashedPubKey)-20:]...)

	// Step 6: Encode address in Base58Check format
	address := encodeBase58Check(rawAddress)

	return privateKeyHex, hex.EncodeToString(pubKeyBytes), address, nil
}

func GenerateWallet() (string, string) {
	privKey, pubKey, address, err := GenWallet()
	if err != nil {
		fmt.Println("Error:", err)
		return "", ""
	}

	fmt.Println("Private Key:", privKey)
	fmt.Println("Public Key:", pubKey)
	fmt.Println("TRON Address:", address)

	return privKey, address
}
