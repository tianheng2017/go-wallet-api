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

// TransferFrom 授权NFT转出
func (ws nftService) TransferFrom(from string, tokenId uint) resp.NftTransferFromResp {
	tx := logic.NftLogic.TransferFrom(from, tokenId)
	return resp.NftTransferFromResp{
		Tx: tx,
	}
}

// 其他
