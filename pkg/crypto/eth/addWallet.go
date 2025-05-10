package eth

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/sha3"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func PrivateKeyToWallet(PrivateKey string) (*models.Wallet, error) {
	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	walletAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	wallet := models.Wallet{
		Address:    walletAddress.String(),
		PrivateKey: PrivateKey,
		Chain:      "eth",
	}
	// err = db.AddWallet(&wallet)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }
	return &wallet, nil
}
func toChecksumAddress(address string) string {
	// Remove '0x' prefix and convert to lowercase
	addr := strings.ToLower(address[2:])

	// Compute Keccak-256 hash of the lowercase address
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(addr))
	hashBytes := hash.Sum(nil)

	// Build checksummed address
	result := "0x"
	for i, char := range addr {
		// Get the corresponding nibble from the hash
		nibble := hashBytes[i>>1]
		if i%2 == 0 {
			nibble = nibble >> 4 // High nibble
		} else {
			nibble = nibble & 0x0f // Low nibble
		}
		// If nibble >= 8 and char is a-f, convert to uppercase
		if nibble >= 8 && char >= 'a' && char <= 'f' {
			result += string(char - 32) // Convert to uppercase
		} else {
			result += string(char)
		}
	}
	return result
}

func GenEthereumWallet() (string, string, string, error) {
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

	// Step 5: Take the last 20 bytes and format as Ethereum address
	addressBytes := hashedPubKey[len(hashedPubKey)-20:]
	lowerAddress := "0x" + hex.EncodeToString(addressBytes)

	address := toChecksumAddress(lowerAddress)

	return privateKeyHex, hex.EncodeToString(pubKeyBytes), address, nil
}

func GenEthWallet() (string, string) {
	privKey, pubKey, addr, err := GenEthereumWallet()
	if err != nil {
		fmt.Println("Error:", err)
		return "", ""
	}
	fmt.Println("Private Key (hex):", privKey)
	fmt.Println("Public Key (hex):", pubKey)
	fmt.Println("Ethereum Address:", addr)

	return privKey, addr
}
