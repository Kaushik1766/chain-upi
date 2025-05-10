package middlware

// func ValidateWallet() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var reqBody wallet.WalletForm
// 		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
// 			return
// 		}
// 		switch strings.ToLower(reqBody.Chain) {
// 		case "trx":
// 			if trx.ValidateAddress(reqBody.PrivateKey) {
// 				ctx.Set("wallet", reqBody)
// 				ctx.Next()
// 			} else {
// 				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Address"})
// 				return
// 			}
// 		}
// 	}
// }
