## 助记词解锁

#### 接口URL
> http://127.0.0.1:8000/api/wallet/mnemonicToAddressAndPrivateKey

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
## 私钥解锁

#### 接口URL
> http://127.0.0.1:8000/api/wallet/privateKeyToAddress

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
## 批量生成钱包

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
count | 1 | Integer | 否 | 生成数量（范围1-1000，不传默认为1）

#### 成功响应示例
```javascript
{
	"code": 1,
	"msg": "成功",
	"data": [
		{
			"address": "0x7D915cf37A3AdB3fa3D15844043e41aE87179c27",
			"privateKey": "3f5cb203bd988dcfcfdd15344593b0326031da687970ba50f7fa31e20eac0979",
			"mnemonic": "zone tone access tenant swallow cat brush already chaos future orphan lonely"
		},
		{
			"address": "0xe4c92614Edd8b168484D3e8dFe84529c62e8d313",
			"privateKey": "ef10c657c1884bbe3021b29def9db582ee320251d5cca04d5fa461afa21ffa78",
			"mnemonic": "fog mobile earth old average improve ankle festival donkey join library disagree"
		},
		{
			"address": "0x7CcfF2d5CA749c1750E8c2297B53f5230bDED275",
			"privateKey": "a571ca9a6ce0b618b4457673c05603310e7a6021572005634b07d2f5e52cf9dc",
			"mnemonic": "road tree swing busy salute foot virtual staff escape sugar summer aisle"
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
## 查询主网币余额

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
address | 0x7d22FD81b542fd7B1Ae82f9b45f9bB69B6ae9eA01 | String | 是 | 钱包地址

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
## 主网币转账

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
toAddress | 0x2F6C2764912C1b02D644076d28e42a87000D54Bf | String | 是 | 接收者钱包地址
number | 0.0001 | Float | 是 | 转账数量，需要大于0
privateKey | ab31ee83f2b1c5b8c1f7071948c1a8ea187954085f179b55d3f7b26c4c2514d8 | String | 是 | 发送者私钥

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
## 铸造NFT

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
minter | 0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2 | String | 是 | 接收人钱包
quantity | 1 | uint | 是 | 数量，需要大于0

#### 成功响应示例
```javascript
{
	"code": 1,
	"msg": "成功",
	"data": {
        // 交易哈希，用于查询交易状态（交易失败了也有hash）
        // 交易状态：0：交易失败，1：交易成功，2：交易进行中
		"tx": "0x7a38b3d4501a152f26e089f3202265a5ce6148bbc5719ae9fba4463dfd48907a"
	}
}
```
#### 错误响应示例
```javascript
{
	"code": 0,
	"msg": "接口不通",
	"data": []
}
```