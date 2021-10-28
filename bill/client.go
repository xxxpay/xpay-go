package payment

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

// 查询交易账单
func (c Client) Trades(appId string, params *xpay.TradeListParams) (*xpay.TradeList, error) {
	start := time.Now()
	tradeList := &xpay.TradeList{}

	qs := make(url.Values)
	qs.Add("channel", params.Channel)
	qs.Add("date", params.Date)

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%v/bills/trades", appId), &qs, []byte{}, tradeList)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	if xpay.LogLevel > 2 {
		log.Println("List trades completed in ", time.Since(start))
	}
	return tradeList, err
}
