package batchTransfer

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
* 创建批量转账
* @param params BatchTransferParams
* @return BatchTransfer
 */
func New(params *xpay.BatchTransferParams) (*xpay.BatchTransfer, error) {
	return getC().New(params)
}

func (c Client) New(params *xpay.BatchTransferParams) (*xpay.BatchTransfer, error) {
	paramsString, _ := xpay.JsonEncode(params)
	batchTransfer := &xpay.BatchTransfer{}
	err := c.B.Call("POST", "/batch_transfers", c.Key, nil, paramsString, batchTransfer)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("New BatchTransfer error: %v\n", err)
		}
	}
	return batchTransfer, err
}

/*
* 查询批量转账
* @param Id string
* @return BatchTransfer
 */
func Get(Id string) (*xpay.BatchTransfer, error) {
	return getC().Get(Id)
}

func (c Client) Get(Id string) (*xpay.BatchTransfer, error) {
	batchTransfer := &xpay.BatchTransfer{}
	err := c.B.Call("GET", fmt.Sprintf("/batch_transfers/%s", Id), c.Key, nil, nil, batchTransfer)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Batchtransfer error: %v\n", err)
		}
	}
	return batchTransfer, err
}

/*
* 查询批量转账列表
* @param params PagingParams
* @return BatchTransferlList
 */
func List(params *xpay.PagingParams) (*xpay.BatchTransferlList, error) {
	return getC().List(params)
}

func (c Client) List(params *xpay.PagingParams) (*xpay.BatchTransferlList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	batchTransferlList := &xpay.BatchTransferlList{}
	err := c.B.Call("GET", "/batch_transfers", c.Key, body, nil, batchTransferlList)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Batchtransfer List error: %v\n", err)
		}
	}
	return batchTransferlList, err
}

/*
* 取消批量转账
* @param batchTransferId string
* @return BatchTransfer
 */
func Cancel(batchTransferId string) (*xpay.BatchTransfer, error) {
	return getC().Cancel(batchTransferId)
}
func (c Client) Cancel(batchTransferId string) (*xpay.BatchTransfer, error) {
	cancelParams := struct {
		Status string `json:"status"`
	}{
		Status: "canceled",
	}
	paramsString, _ := xpay.JsonEncode(cancelParams)

	batchTransfer := &xpay.BatchTransfer{}
	err := c.B.Call("PUT", fmt.Sprintf("/batch_transfers/%s", batchTransferId), c.Key, nil, paramsString, batchTransfer)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf(" BatchTransfer error: %v\n", err)
		}
	}
	return batchTransfer, err
}
