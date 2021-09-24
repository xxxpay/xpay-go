package channel

import (
	"fmt"
	"log"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

func (c Client) New(appId, subAppId string, params *xpay.ChannelParams) (*xpay.Channel, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("ChannelParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of create user is :\n %v\n ", string(paramsString))
	}

	channel := &xpay.Channel{}
	err := c.backend.Call("POST", fmt.Sprintf("/apps/%s/sub_apps/%s/channels", appId, subAppId), nil, paramsString, channel)
	return channel, err
}

func (c Client) Get(appId, subAppId, channelName string) (*xpay.Channel, error) {
	channel := &xpay.Channel{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/sub_apps/%s/channels/%s", appId, subAppId, channelName), nil, nil, channel)
	return channel, err
}

func (c Client) Delete(appId, subAppId, channelName string) (*xpay.ChannelDeleteResult, error) {
	result := &xpay.ChannelDeleteResult{}

	err := c.backend.Call("DELETE", fmt.Sprintf("/apps/%s/sub_apps/%s/channels/%s", appId, subAppId, channelName), nil, nil, result)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return result, err
}

func (c Client) Update(appId, subAppId, channelName string, params xpay.ChannelUpdateParams) (*xpay.Channel, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of update Channel  to xpay is :\n %v\n ", string(paramsString))
	}

	channel := &xpay.Channel{}

	err := c.backend.Call("PUT", fmt.Sprintf("/apps/%s/sub_apps/%s/channels/%s", appId, subAppId, channelName), nil, paramsString, channel)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return channel, err
}
