package nft

import (
	"server/api/logic"
	"server/api/schemas/resp"
)

var NftService = nftService{}

// walletService NFT服务类
type nftService struct{}

// MintSeaDrop 铸造NFT
func (ws nftService) MintSeaDrop(minter string, quantity uint) resp.NftTxResp {
	tx := logic.NftLogic.MintSeaDrop(minter, quantity)
	return resp.NftTxResp{
		Tx: tx,
	}
}

// TransferFrom 授权NFT转出
func (ws nftService) TransferFrom(from string, tokenId uint) resp.NftTxResp {
	tx := logic.NftLogic.TransferFrom(from, tokenId)
	return resp.NftTxResp{
		Tx: tx,
	}
}

// 其他
