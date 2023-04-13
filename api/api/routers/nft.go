package routers

import (
	"server/api/schemas/req"
	"server/api/service/nft"
	"server/core"
	"server/core/response"
	"server/util"

	"github.com/gin-gonic/gin"
)

var NftGroup = core.Group("/nft")

func init() {
	group := NftGroup
	group.AddPOST("/mintSeaDrop", mintSeaDrop)
}

// 铸造NFT
// 参数：minter: 接收者, quantity: 数量
func mintSeaDrop(c *gin.Context) {
	var mintReq req.NftMintSeaDropReq
	util.VerifyUtil.VerifyBody(c, &mintReq)
	response.OkWithData(c, nft.NftService.MintSeaDrop(mintReq.Minter, mintReq.Quantity))
}
