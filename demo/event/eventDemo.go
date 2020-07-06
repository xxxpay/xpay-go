package event

import (
	//"encoding/json"

	"github.com/xxxpay/xpay-go/xpay"
	"github.com/xxxpay/xpay-go/xpay/event"
)

var Demo = new(EventDemo)

type EventDemo struct {
	demoAppID string
}

func (c *EventDemo) Setup(app string) {
	c.demoAppID = app
}

// 查询指定的 event 对象，通过 event 对象的 id 查询一个已创建的 event 对象
func (c *EventDemo) Get() (*xpay.Event, error) {
	return event.Get("evt_zRFRk6ekazsH7t7yCqEeovhk")
}

func (c *EventDemo) Run() {
	c.Get()
}
