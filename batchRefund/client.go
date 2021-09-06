package batchRefund

import (
	"fmt"
	"log"
	"net/url"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	B   xpay.Backend
	Key string
}

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}

/*
* 创建批量退款
* @param params BatchRefundParams
* @return BatchRefund
 */
func New(params *xpay.BatchRefundParams) (*xpay.BatchRefund, error) {
	return getC().New(params)
}

func (c Client) New(params *xpay.BatchRefundParams) (*xpay.BatchRefund, error) {
	paramsString, _ := xpay.JsonEncode(params)
	batchRefund := &xpay.BatchRefund{}
	err := c.B.Call("POST", "/batch_refunds", c.Key, nil, paramsString, batchRefund)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("New BatchRefunds error: %v\n", err)
		}
	}
	return batchRefund, err
}

/*
* 查询批量退款
* @param Id string
* @return BatchRefund
 */
func Get(Id string) (*xpay.BatchRefund, error) {
	return getC().Get(Id)
}

func (c Client) Get(Id string) (*xpay.BatchRefund, error) {
	batchRefund := &xpay.BatchRefund{}
	err := c.B.Call("GET", fmt.Sprintf("/batch_refunds/%s", Id), c.Key, nil, nil, batchRefund)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get BatchRefunds error: %v\n", err)
		}
	}
	return batchRefund, err
}

/*
* 查询批量退款列表
* @param params PagingParams
* @return BatchRefundlList
 */
func List(params *xpay.PagingParams) (*xpay.BatchRefundlList, error) {
	return getC().List(params)
}

func (c Client) List(params *xpay.PagingParams) (*xpay.BatchRefundlList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	batchRefundlList := &xpay.BatchRefundlList{}
	err := c.B.Call("GET", "/batch_refunds", c.Key, body, nil, batchRefundlList)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get BatchRefunds List error: %v\n", err)
		}
	}
	return batchRefundlList, err
}
