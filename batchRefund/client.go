package batchRefund

import (
	"fmt"
	"log"
	"net/url"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

/*
* 创建批量退款
* @param params BatchRefundParams
* @return BatchRefund
 */
func (c Client) New(params *xpay.BatchRefundParams) (*xpay.BatchRefund, error) {
	paramsString, _ := xpay.JsonEncode(params)
	batchRefund := &xpay.BatchRefund{}
	err := c.backend.Call("POST", "/batch_refunds",  nil, paramsString, batchRefund)
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
func (c Client) Get(Id string) (*xpay.BatchRefund, error) {
	batchRefund := &xpay.BatchRefund{}
	err := c.backend.Call("GET", fmt.Sprintf("/batch_refunds/%s", Id),  nil, nil, batchRefund)
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
func (c Client) List(params *xpay.PagingParams) (*xpay.BatchRefundlList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	batchRefundlList := &xpay.BatchRefundlList{}
	err := c.backend.Call("GET", "/batch_refunds",  body, nil, batchRefundlList)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get BatchRefunds List error: %v\n", err)
		}
	}
	return batchRefundlList, err
}
