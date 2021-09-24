package batchTransfer

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
* 创建批量转账
* @param params BatchTransferParams
* @return BatchTransfer
 */
func (c Client) New(params *xpay.BatchTransferParams) (*xpay.BatchTransfer, error) {
	paramsString, _ := xpay.JsonEncode(params)
	batchTransfer := &xpay.BatchTransfer{}
	err := c.backend.Call("POST", "/batch_transfers", nil, paramsString, batchTransfer)
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
func (c Client) Get(Id string) (*xpay.BatchTransfer, error) {
	batchTransfer := &xpay.BatchTransfer{}
	err := c.backend.Call("GET", fmt.Sprintf("/batch_transfers/%s", Id), nil, nil, batchTransfer)
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
func (c Client) List(params *xpay.PagingParams) (*xpay.BatchTransferlList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	batchTransferlList := &xpay.BatchTransferlList{}
	err := c.backend.Call("GET", "/batch_transfers", body, nil, batchTransferlList)
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
func (c Client) Cancel(batchTransferId string) (*xpay.BatchTransfer, error) {
	cancelParams := struct {
		Status string `json:"status"`
	}{
		Status: "canceled",
	}
	paramsString, _ := xpay.JsonEncode(cancelParams)

	batchTransfer := &xpay.BatchTransfer{}
	err := c.backend.Call("PUT", fmt.Sprintf("/batch_transfers/%s", batchTransferId), nil, paramsString, batchTransfer)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf(" BatchTransfer error: %v\n", err)
		}
	}
	return batchTransfer, err
}
