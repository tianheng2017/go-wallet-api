## 全局公共参数
#### 全局Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 全局Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 全局Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 全局认证方式
```text
noauth
```
#### 全局预执行脚本
```javascript
暂无预执行脚本
```
#### 全局后执行脚本
```javascript
暂无后执行脚本
```
## /wallet
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /wallet/批量生成钱包
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/wallet/create?count=1

#### 请求方式
> GET

#### Content-Type
> none

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
count | 1 | String | 否 | 范围1-100，不填默认为1
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 1,
	"msg": "成功",
	"data": [
		{
			"address": "0x3C9FBb56e8592ca216C0a23CB2049296d61B1d8a",
			"privateKey": "b1a124f1a9afc015cb52226a4ceab4abc35f52e4196085b9f5f321f4703f169c",
			"mnemonic": "estate vocal hockey miss usage thrive deputy weird risk similar cluster thank"
		}
	]
}
```
#### 错误响应示例
```javascript
{
	"code": 0,
	"msg": "Count最小只能为1",
	"data": []
}
```
## /wallet/助记词解锁
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/wallet/mnemonicUnlock

#### 请求方式
> POST

#### Content-Type
> urlencoded

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
mnemonic | blast fish near connect rocket clump early caught ginger multiply gain cousin | String | 是 | 助记词
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 1,
	"msg": "成功",
	"data": {
		"address": "0x248F6e9cf37ad1A1973E77F567beD6174A38AC24",
		"privateKey": "23b205763780c5d317f7073eac2ed14d37e90ede65e94841bede16fba26d475d"
	}
}
```
#### 错误响应示例
```javascript
{
	"code": 0,
	"msg": "助记词无效",
	"data": []
}
```
## /wallet/私钥解锁
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/wallet/privateKeyUnlock

#### 请求方式
> POST

#### Content-Type
> urlencoded

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
privateKey | 0xf6e95363c25239e847a261b6e166d5e33463f571055646d337df4ec89e2f87d5 | String | 是 | 私钥
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 1,
	"msg": "成功",
	"data": {
		"address": "0x2F6C2764912C1b02D644076d28e42a87000D54Bf"
	}
}
```
#### 错误响应示例
```javascript
{
	"code": 0,
	"msg": "私钥无效",
	"data": []
}
```
## /wallet/铸造NFT
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/nft/mintSeaDrop

#### 请求方式
> POST

#### Content-Type
> urlencoded

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
minter | 0xe29b14121A099384385BbDd184Ef63A862E03B24 | - | 是 | -
quantity | 1 | - | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 1,
	"msg": "成功",
	"data": {
		"tx": "0x2aedb3e7de91ade271bacbcfe11a09caadefdfe69384e79bfe1e6ee9ff6ea8a9"
	}
}
```
#### 错误响应示例
```javascript
{
	"code": 0,
	"msg": "Quantity必须大于0",
	"data": []
}
```
## /wallet/授权转出NFT
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/nft/transferFrom

#### 请求方式
> POST

#### Content-Type
> urlencoded

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
from | 0xe29b14121A099384385BbDd184Ef63A862E03B24 | String | 是 | -
tokenId | 1 | Integer | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /wallet/查询主网币余额
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/wallet/getBalance

#### 请求方式
> POST

#### Content-Type
> urlencoded

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
address | 0x03abab67eDaa3582aa777cFe94767c87fD942240 | String | 是 | 钱包地址
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 1,
	"msg": "成功",
	"data": {
		"balance": "0.489058"
	}
}
```
#### 错误响应示例
```javascript
{
	"code": 0,
	"msg": "钱包地址无效",
	"data": []
}
```
## /wallet/主网币转账
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/wallet/transfer

#### 请求方式
> POST

#### Content-Type
> urlencoded

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
to | 0xe29b14121A099384385BbDd184Ef63A862E03B24 | String | 是 | 接收者钱包地址
amount | 1 | Float | 是 | 转账数量，需要大于0
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 1,
	"msg": "成功",
	"data": {
		"tx": "0x1989cb5f0497425e096ab014c05648e7dc7d4032090a0f824c39a8d3d0ed41fb"
	}
}
```
#### 错误响应示例
```javascript
{
	"code": 0,
	"msg": "钱包余额不足, 当前余额: 0.498790",
	"data": []
}
```
## /wallet/查询USDT余额
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/wallet/getUsdtBalance

#### 请求方式
> POST

#### Content-Type
> urlencoded

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
address | 0x03abab67eDaa3582aa777cFe94767c87fD942240 | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 1,
	"msg": "成功",
	"data": {
		"balance": "29999000"
	}
}
```
#### 错误响应示例
```javascript
{
	"code": 0,
	"msg": "钱包格式无效",
	"data": []
}
```
## /wallet/USDT转账
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/wallet/usdtTransfer

#### 请求方式
> POST

#### Content-Type
> urlencoded

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
to | 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 | String | 是 | -
amount | 1 | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 1,
	"msg": "成功",
	"data": {
		"tx": "0xb78b652ce616d06b6ce231b72a5c5329a82a8d049e8c9dcefef0f1a7d4c8408d"
	}
}
```
#### 错误响应示例
```javascript
{
	"code": 0,
	"msg": "Amount必须大于0",
	"data": []
}
```
## /wallet/查询代币余额
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/wallet/getTokenBalance

#### 请求方式
> POST

#### Content-Type
> urlencoded

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
address | 0x03abab67eDaa3582aa777cFe94767c87fD942240 | - | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 1,
	"msg": "成功",
	"data": {
		"balance": "100000000"
	}
}
```
#### 错误响应示例
```javascript
{
	"code": 0,
	"msg": "钱包格式无效",
	"data": []
}
```
## /wallet/代币转账
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/wallet/tokenTransfer

#### 请求方式
> POST

#### Content-Type
> urlencoded

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
to | 0xe29b14121A099384385BbDd184Ef63A862E03B24 | String | 是 | -
amount | 1 | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 1,
	"msg": "成功",
	"data": {
		"tx": "0xb2d5d56bd74c85ad5190f602383d13b5598982e51d88f3fcab2ae1cc65788d7d"
	}
}
```
#### 错误响应示例
```javascript
{
	"code": 0,
	"msg": "Amount必须大于0",
	"data": []
}
```
## /wallet/触发每日解锁代币
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/wallet/unlock

#### 请求方式
> POST

#### Content-Type
> none

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{"code":1,"msg":"成功","data":{"tx":"0xf35cdda51082a3ba98d385ae16de697a3ea3ce1c8cf9e7809ab0faa59542ec3f"}}
```
#### 错误响应示例
```javascript
{"code":0,"msg":"解锁失败: VM Exception while processing transaction: revert The time since the last unlock is less than one day.","data":[]}
```
## /wallet/获取可用代币数量(已解锁)
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/wallet/getUnlockToken

#### 请求方式
> GET

#### Content-Type
> none

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{"code":1,"msg":"成功","data":{"num":"21600"}}
```
## /wallet/获取上次解锁时间戳
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/wallet/getLastUnlockTimestamp

#### 请求方式
> GET

#### Content-Type
> none

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{"code":1,"msg":"成功","data":{"timestamp":1681473538}}
```
## /wallet/获取合约启动时间戳
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:8000/api/wallet/getStartTimestamp

#### 请求方式
> GET

#### Content-Type
> none

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
apiKey | TzIyVyUC | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{"code":1,"msg":"成功","data":{"timestamp":1681387075}}
```