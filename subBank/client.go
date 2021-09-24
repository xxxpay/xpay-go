package subBank

import (
	"net/url"

	"github.com/xxxpay/xpay-go"
)

// Client 支行客户端
type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

// List 按银行编号和省市查询支行信息列表
// 参数 | 类型 | 长度/个数/范围 | 是否必须 | 默认值 | 描述
// app | string | 20 | required | 无 | App ID。
// open_bank_code | string | 4 | required | 无 | 银行编号
// prov | string | 1~20 | required | 无 | 省份。
// city | string | 1~40 | required | 无 | 城市。
// channel | string | [`chanpay`] | required | 无 | 渠道。
func (c Client) List(app, openBankCode, prov, city, channel string) (xpay.SubBankList, error) {
	values := &url.Values{}
	values.Add("app", app)
	values.Add("open_bank_code", openBankCode)
	values.Add("prov", prov)
	values.Add("city", city)
	values.Add("channel", channel)

	subBankList := xpay.SubBankList{}
	err := c.backend.Call("GET", "/sub_banks", values, nil, &subBankList)
	return subBankList, err
}
