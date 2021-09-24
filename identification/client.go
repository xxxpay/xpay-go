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
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
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

	err := c.backend.Call("POST", "/identification", nil, paramsString, identificationResult)
	return identificationResult, err
}
