package channel

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

func New(appId, subAppId string, params *xpay.ChannelParams) (*xpay.Channel, error) {
	return getC().New(appId, subAppId, params)
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
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/sub_apps/%s/channels", appId, subAppId), c.Key, nil, paramsString, channel)
	return channel, err
}

func Get(appId, subAppId, channel string) (*xpay.Channel, error) {
	return getC().Get(appId, subAppId, channel)
}

func (c Client) Get(appId, subAppId, channelName string) (*xpay.Channel, error) {
	channel := &xpay.Channel{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/sub_apps/%s/channels/%s", appId, subAppId, channelName), c.Key, nil, nil, channel)
	return channel, err
}

func Delete(appId, subAppId, channelName string) (*xpay.ChannelDeleteResult, error) {
	return getC().Delete(appId, subAppId, channelName)
}

func (c Client) Delete(appId, subAppId, channelName string) (*xpay.ChannelDeleteResult, error) {
	result := &xpay.ChannelDeleteResult{}

	err := c.B.Call("DELETE", fmt.Sprintf("/apps/%s/sub_apps/%s/channels/%s", appId, subAppId, channelName), c.Key, nil, nil, result)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return result, err
}

func Update(appId, subAppId, channelName string, params xpay.ChannelUpdateParams) (*xpay.Channel, error) {
	return getC().Update(appId, subAppId, channelName, params)
}

func (c Client) Update(appId, subAppId, channelName string, params xpay.ChannelUpdateParams) (*xpay.Channel, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of update Channel  to xpay is :\n %v\n ", string(paramsString))
	}

	channel := &xpay.Channel{}

	err := c.B.Call("PUT", fmt.Sprintf("/apps/%s/sub_apps/%s/channels/%s", appId, subAppId, channelName), c.Key, nil, paramsString, channel)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return channel, err
}
