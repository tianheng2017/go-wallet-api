package nft

import (
	"server/api/logic"
	"server/api/schemas/resp"
)

var NftService = nftService{}

// walletService NFT服务类
type nftService struct{}

// MintSeaDrop 铸造NFT
func (ws nftService) MintSeaDrop(minter string, quantity uint) resp.NftMintSeaDropResp {
	tx := logic.NftLogic.MintSeaDrop(minter, quantity)
	return resp.NftMintSeaDropResp{
		Tx: tx,
	}
}

// 其他
