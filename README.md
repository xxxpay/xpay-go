# XPay Go SDK

## 简介
xpay 文件夹里是 SDK 文件

## 版本要求
建议 Go 语言版本 1.4 以上 

## 安装

导入 xpay 模块

```go
go get github.com/xxxpay/xpay-go/xpay
```

导入后，在调用的时候需要

```go
import (xpay "github.com/xxxpay/xpay-go/xpay")
```
具体使用相应模块的话还需要

```go
import (xpay "github.com/xxxpay/xpay-go/xpay/xxx")
```

## 接入方法

### 初始化
   
```go    
// 设置 API-KEY 
xpay.Key= "YOUR-KEY"
```

### 支付
```go
//获得的第一个参数即是 Payment 对象
payment, err := payment.New(&paymentParams)
```

### Payment查询
```go
//查询单个 Payment 对象
payment, err := payment.Get(ch_id)
```

```go
//查询 Payment 列表
payments, err := payment.List(&paymentListParams)

```

### 退款
``` go
//payment_id为待退款的Payment的ID
refund, err := refund.New(payment_id, refundParams)
```

### 退款查询
```go
//查询单个Refund对象
refund, err := refund.Get(ch_id, re_id)
```

```go
//查询Refund对象列表
refunds, err := refund.List(ch_id, &refundListParams)
```


### 微信红包
```go
//获得的第一个参数即是 RedEnvelope 对象
redenvelope, err := redEnvelope.New(&redEnvelopeParams)
```

### 红包查询
```go
//查询单个 RedEnvelope 对象
redenvelope, err := redEnvelope.Get(red_id)
```

```go
//查询 RedEnvelope 列表
redenvelope, err := redEnvelope.List(&redEnvelopeListParams)
```

### event查询
```go
//查询单个 event 对象
event, err := event.Get(red_id)
```

### 身份认证
```go
//鉴别用户身份证、银行卡信息的真伪
result, err := identification.New(&identificationParams)
```

### 批量退款
```go
//发起批量退款
batch_refund, err := batchRefund.New(params)
```

### 批量企业付款
```go
//发起批量企业付款
 batch_transfer, err := batchTransfer.New(params)
```

## Debug
SDK 提供了 debug 模式。只需要更改 xpay.go 文件中的 LogLevel 变量值，即可触发相应级别的 log，代码中对级别有注释。默认的级别是 2

## 版本号
调用

```go
xpay.Version()
```
会返回 sdk 版本号

## 中文报错信息
XPay 支持中文和英文两种语言的报错信息。SDK 默认的 Accept-Language 是英文的，如果您想要接收到的错误提示是中文的，只需要设置一下即可：

```go
xpay.AcceptLanguage = "zh-CN"
```

商户私钥 java pkcs8 格式的 的在 golang 要转成 pkcs1 格式的
```bash
 openssl rsa -in your_pkcs8.pem -out your_pkcs1.pem
```

**详细信息请参考 [API 文档](https://xpay.lucfish.com/document/api?go)**。


