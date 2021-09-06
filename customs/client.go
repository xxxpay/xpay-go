package customs

import (
	"fmt"
	"log"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	B   xpay.Backend
	Key string
}

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}

func New(params *xpay.CustomsParams) (*xpay.Customs, error) {
	return getC().New(params)
}

func (c Client) New(params *xpay.CustomsParams) (*xpay.Customs, error) {
	paramsString, _ := xpay.JsonEncode(params)
	customs := &xpay.Customs{}
	err := c.B.Call("POST", "/customs", c.Key, nil, paramsString, customs)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("New Customs error: %v\n", err)
		}
	}
	return customs, err
}

func Get(Id string) (*xpay.Customs, error) {
	return getC().Get(Id)
}

func (c Client) Get(Id string) (*xpay.Customs, error) {
	customs := &xpay.Customs{}
	err := c.B.Call("GET", fmt.Sprintf("/customs/%s", Id), c.Key, nil, nil, customs)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Customs error: %v\n", err)
		}
	}
	return customs, err
}
