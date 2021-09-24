package royaltyTemplate

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

// 创建分润模板
func (c Client) New(params *xpay.RoyaltyTmplParams) (*xpay.RoyaltyTmpl, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of create royalty_template request to xpay is :\n %v\n ", string(paramsString))
	}

	royaltyTemplate := &xpay.RoyaltyTmpl{}

	err := c.backend.Call("POST", "/royalty_templates", nil, paramsString, royaltyTemplate)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return royaltyTemplate, err

}

//查询指定的分润模板
func (c Client) Get(royaltyTmplId string) (*xpay.RoyaltyTmpl, error) {
	var body *url.Values
	body = &url.Values{}
	royaltyTmpl := &xpay.RoyaltyTmpl{}

	err := c.backend.Call("GET", fmt.Sprintf("/royalty_templates/%s", royaltyTmplId), body, nil, royaltyTmpl)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get royalty Template error: %v\n", err)
		}
	}
	return royaltyTmpl, err
}

//更新分润模板
func (c Client) Update(royaltyTmplId string, params *xpay.RoyaltyTmplUpdateParams) (*xpay.RoyaltyTmpl, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of update royalty template to xpay is :\n %v\n ", string(paramsString))
	}

	royaltyTmpl := &xpay.RoyaltyTmpl{}

	err := c.backend.Call("PUT", fmt.Sprintf("/royalty_templates/%s", royaltyTmplId), nil, paramsString, royaltyTmpl)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return royaltyTmpl, err
}

//删除分润模板
func (c Client) Delete(royaltyTmplId string) (*xpay.DeleteResult, error) {
	result := &xpay.DeleteResult{}

	err := c.backend.Call("DELETE", fmt.Sprintf("/royalty_templates/%s", royaltyTmplId), nil, nil, result)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Delete Royalty Template error: %v\n", err)
		}
	}
	return result, err
}

//查询分润模板列表
func (c Client) List(params *xpay.PagingParams) (*xpay.RoyaltyTmplList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	royaltyTmplList := &xpay.RoyaltyTmplList{}
	err := c.backend.Call("GET", "/royalty_templates", body, nil, royaltyTmplList)
	return royaltyTmplList, err
}
