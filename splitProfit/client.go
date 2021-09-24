package splitProfit

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

// New 请求分账
func (c Client) New(params *xpay.SplitProfitParams) (*xpay.SplitProfit, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("SplitProfitParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of create SplitProfitParams is :\n %v\n ", string(paramsString))
	}

	splitProfit := &xpay.SplitProfit{}
	err := c.backend.Call("POST", fmt.Sprintf("/split_profits"), nil, paramsString, splitProfit)
	return splitProfit, err
}

// Get 查询分账
func (c Client) Get(id string) (*xpay.SplitProfit, error) {
	splitProfit := &xpay.SplitProfit{}

	err := c.backend.Call("GET", fmt.Sprintf("/split_profits/%s", id), nil, nil, splitProfit)
	return splitProfit, err
}

// List 查询分账列表
// | 参数 | 类型 | 长度/个数/范围 | 是否必需 | 默认值 | 说明
// | --- | --- | --- | --- | --- | ---
// | app | string | 20 | required | 无 | App ID。
// | payment | string |  | optional | 无 | xpay 交易成功的 payment ID
// | type | string | optional | 无 | 分账类型: `split_normal` 为普通分账,`split_return` 为完结分账
// | channel | string | [`wx`、`wx_lite`、`wx_pub`、`wx_wap`、`wx_pub_qr`、`wx_pub_scan`] | optional | 无 | 暂时只支持微信渠道
func (c Client) List(app, payment, typ, channel string, params *xpay.PagingParams) (xpay.SplitProfitList, error) {
	values := &url.Values{}
	values.Add("app", app)
	if payment != "" {
		values.Add("payment", payment)
	}
	if typ != "" {
		values.Add("type", typ)
	}
	if channel != "" {
		values.Add("channel", channel)
	}
	params.Filters.AppendTo(values)

	splitProfitList := xpay.SplitProfitList{}
	err := c.backend.Call("GET", "/split_profits", values, nil, &splitProfitList)
	return splitProfitList, err
}
