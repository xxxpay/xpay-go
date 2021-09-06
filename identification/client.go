package identification

import (
	"log"

	"github.com/xxxpay/xpay-go"
)

const (
	IDENTIFY_IDCARD   = "id_card"
	IDENTIFY_BANKCARD = "bank_card"
)

type Client struct {
	B   xpay.Backend
	Key string
}

func New(params *xpay.IdentificationParams) (*xpay.IdentificationResult, error) {
	return getC().New(params)
}

func (c Client) New(params *xpay.IdentificationParams) (*xpay.IdentificationResult, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("IdentificationParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of identification request to xpay is :\n %v\n ", string(paramsString))
	}
	identificationResult := &xpay.IdentificationResult{}

	err := c.B.Call("POST", "/identification", c.Key, nil, paramsString, identificationResult)
	return identificationResult, err
}

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}
