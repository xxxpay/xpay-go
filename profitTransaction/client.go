package profitTransaction

import (
	"fmt"
	"net/url"

	"github.com/xxxpay/xpay-go"
)

// Client 分账明细客户端
type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

// Get 查询分账明细

// Get 查询分账明细
func (c Client) Get(id string) (*xpay.ProfitTransaction, error) {
	profitTransaction := &xpay.ProfitTransaction{}

	err := c.backend.Call("GET", fmt.Sprintf("/profit_transactions/%s", id), nil, nil, profitTransaction)
	return profitTransaction, err
}

// List 查询分账明细列表
// | 参数 | 类型 | 长度/个数/范围 | 是否必需 | 默认值 | 说明
// | --- | --- | --- | --- | --- | ---
// | app | string | 20 | required | 无 | App ID。
// | split_profit| string | 17 | optional | 无 | 分账ID
// | split_receiver|  string | 19 | optional | 无 | 分账接收方ID
// | status | string | - | optional | 无 | 分账明细状态
func (c Client) List(app, splitProfit, splitReceiver, status string, params *xpay.PagingParams) (xpay.ProfitTransactionList, error) {
	values := &url.Values{}
	values.Add("app", app)
	if splitProfit != "" {
		values.Add("split_profit", splitProfit)
	}
	if splitReceiver != "" {
		values.Add("split_receiver", splitReceiver)
	}
	if status != "" {
		values.Add("status", status)
	}
	params.Filters.AppendTo(values)

	profitTransactionList := xpay.ProfitTransactionList{}
	err := c.backend.Call("GET", "/profit_transactions", values, nil, &profitTransactionList)
	return profitTransactionList, err
}
