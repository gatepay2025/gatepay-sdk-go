# 最新地址已迁移到：
https://github.com/gate/gateapi-go

# GatePay API GO SDK

GatePay API官方GO语言客户端代码库。

## 功能介绍

1. GatePay GO SDK适用于Go 1.6及以上版本。
2. 在开始使用GatePay SD之前需要提前去GatePay注册中心注册账号，申请secretKey密钥SK信息请妥善保管，如果遗失可能会造成非法用户使用此信息操作您在的资源，给你造成数据和财产损失。
3. 接口SDK。详见接口介绍
4. 客户端core.Client，支持HTTP签名请求和应答。



## 下载和安装

1.下载地址:

https://github.com/gatepay2025/gatepay-sdk-go 

2.您也可以使用以下命令获取安装包，代码会被下载到GOPATH环境变量中第一个路径src目录中。

```
go get https://github.com/gatepay2025/gatepay-sdk-go
```



## 调用SDK

**业务侧SDK的调用主要分为如下步骤：**

1. 设置secretKey

2. 创建业务Client

3. 设置业务请求参数

4. 调用SDK API接口

5. 处理得到响应

   大致代码如下：

   ```go
   package main
   
   import (
   	"context"
   	"github.com/gatepay2025/gatepay-sdk-go/core"
   	"github.com/gatepay2025/gatepay-sdk-go/core/stringutillib"
   	"github.com/gatepay2025/gatepay-sdk-go/services/common"
   	"github.com/shopspring/decimal"
   	"log"
   )
   
   func main() {
   		cfg := core.NewConfig()
       //设置商户注册获取的秘钥
   		credentials := core.NewCredentials("secret-key")
   		client, err := core.NewClient(cfg, credentials)
   		if err != nil {
   			return
   		}
   
       ctx := context.Background()
       service := &AddressApiService{Client: client}
       req := QueryAddressOrderRequest{MerchantTradeNo: "merchant-trade-no", PrepayID: "order-id"}
       req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
   
       resp, result, err := service.QueryAddressOrder(ctx, req)
       if err != nil {
         log.Printf("call QueryAddressOrder err:%s", err.Error())
       } else {
         log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
       }
   }
   ```

   

