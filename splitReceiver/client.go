package splitReceiver

import (
	"fmt"
	"log"
	"net/url"

	"github.com/xxxpay/xpay-go"
)

// Client 分账客户端
// 暂时只支持微信渠道特约商户
type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

// New 添加分账接收方
// ##### 微信分账接收方类型 type 说明
// - `MERCHANT_ID`：商户ID
// - `PERSONAL_WECHATID`：个人微信号
// - `PERSONAL_OPENID`：个人openid（由父商户APPID转换得到）
// - `PERSONAL_SUB_OPENID`: 个人sub_openid（由子商户APPID转换得到）
// ##### 微信分账接收方帐号 account 说明
// - 微信分账接收方类型是`MERCHANT_ID`时，是商户ID
// - 微信分账接收方类型是`PERSONAL_WECHATID`时，是个人微信号
// - 微信分账接收方类型是`PERSONAL_OPENID`时，是个人`openid`
// - 微信分账接收方类型是`PERSONAL_SUB_OPENID`时，是个人`sub_openid`
func (c Client) New(params *xpay.SplitReceiverParams) (*xpay.SplitReceiver, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("SplitReceiverParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of create SplitReceiverParams is :\n %v\n ", string(paramsString))
	}

	splitReceiver := &xpay.SplitReceiver{}
	err := c.backend.Call("POST", fmt.Sprintf("/split_receivers"), nil, paramsString, splitReceiver)
	return splitReceiver, err
}

// Get 查询分账接收方
func (c Client) Get(id string) (*xpay.SplitReceiver, error) {
	splitReceiver := &xpay.SplitReceiver{}

	err := c.backend.Call("GET", fmt.Sprintf("/split_receivers/%s", id), nil, nil, splitReceiver)
	return splitReceiver, err
}

// List 查询分账接收方列表
// | 参数 | 类型 | 长度/个数/范围 | 是否必需 | 默认值 | 说明
// | --- | --- | --- | --- | --- | ---
// | app | string | 20 | required | 无 | App ID。
// | type | string | [1~32] | optional | 无 | 分账接收方类型
// | channel | string | [`wx`、`wx_lite`、`wx_pub`、`wx_wap`、`wx_pub_qr`、`wx_pub_scan`] | optional | 无 | 暂时只支持微信渠道
func (c Client) List(app, typ, channel string, params *xpay.PagingParams) (xpay.SplitReceiverList, error) {
	values := &url.Values{}
	values.Add("app", app)
	if typ != "" {
		values.Add("type", typ)
	}
	if channel != "" {
		values.Add("channel", channel)
	}
	params.Filters.AppendTo(values)

	splitReceiverList := xpay.SplitReceiverList{}
	err := c.backend.Call("GET", "/split_receivers", values, nil, &splitReceiverList)
	return splitReceiverList, err
}

// Delete 删除分账接收方
func (c Client) Delete(id string) (*xpay.DeleteResult, error) {
	deleteResult := &xpay.DeleteResult{}

	err := c.backend.Call("DELETE", fmt.Sprintf("/split_receivers/%s", id), nil, nil, deleteResult)
	return deleteResult, err
}
