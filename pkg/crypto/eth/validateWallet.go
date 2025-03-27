package eth

import "github.com/ethereum/go-ethereum/common"

func ValidateWallet(wallet string) bool {
	if common.IsHexAddress(wallet) {
		return true
	} else {
		return false
	}
}
