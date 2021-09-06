package cardInfo

import (
	"log"
	"time"

	"github.com/xxxpay/xpay-go"
)

// Client cardInfo 请求客户端
type Client struct {
	B   xpay.Backend
	Key string
}

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}

// New 发送 /card_info 请求
func New(params *xpay.CardInfoParams) (*xpay.CardInfo, error) {
	return getC().New(params)
}

// New 发送 /card_info 请求
func (c Client) New(params *xpay.CardInfoParams) (*xpay.CardInfo, error) {
	start := time.Now()
	paramsString, err := xpay.JsonEncode(params)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("CardInfoParams Marshall Errors is : %q\n", err)
		}
		return nil, err
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of cardInfo request to xpay is :\n %v\n ", string(paramsString))
	}

	cardInfo := &xpay.CardInfo{}
	errch := c.B.Call("POST", "/card_info", c.Key, nil, paramsString, cardInfo)
	if errch != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", errch)
		}
		return nil, errch
	}
	if xpay.LogLevel > 2 {
		log.Println("CardInfo completed in ", time.Since(start))
	}
	return cardInfo, nil
}
