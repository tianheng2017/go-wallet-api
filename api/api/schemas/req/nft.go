package req

// NftMintSeaDropReq 铸造NFT参数
type NftMintSeaDropReq struct {
	Minter   string `form:"minter" binding:"required"` // 接收者
	Quantity uint   `form:"quantity" binding:"gt=0"`   // 数量
}

// NftTransferFromReq 铸造NFT参数
type NftTransferFromReq struct {
	From    string `form:"from" binding:"required"` // 发送人
	TokenId uint   `form:"tokenId" binding:"gt=0"`  // 数量
}
