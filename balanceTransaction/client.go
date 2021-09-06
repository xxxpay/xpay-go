package balanceTransaction

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
* 查询用户明细对象
* @param appId string
* @param txnId string
* @return BalanceTransaction
 */
func Get(appId, txnId string) (*xpay.BalanceTransaction, error) {
	return getC().Get(appId, txnId)
}

func (c Client) Get(appId, txnId string) (*xpay.BalanceTransaction, error) {
	balanceTransactions := &xpay.BalanceTransaction{}
	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/balance_transactions/%s", appId, txnId), c.Key, nil, nil, balanceTransactions)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get BalanceTransactions error: %v\n", err)
		}
	}
	return balanceTransactions, err
}

/*
* 查询用户明细对象列表
* @param appId string
* @param params PagingParams
* @return BalanceTransactionList
 */
func List(appId string, params *xpay.PagingParams) (*xpay.BalanceTransactionList, error) {
	return getC().List(appId, params)
}

func (c Client) List(appId string, params *xpay.PagingParams) (*xpay.BalanceTransactionList, error) {
	balanceList := &xpay.BalanceTransactionList{}
	body := &url.Values{}
	params.Filters.AppendTo(body)

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/balance_transactions", appId), c.Key, body, nil, balanceList)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get BalanceTransactions List error: %v\n", err)
		}
	}
	return balanceList, err
}
