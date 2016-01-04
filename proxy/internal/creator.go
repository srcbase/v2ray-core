package internal

import (
	"github.com/v2ray/v2ray-core/app"
	"github.com/v2ray/v2ray-core/proxy"
)

type InboundConnectionHandlerCreator func(space app.Space, config interface{}) (proxy.InboundConnectionHandler, error)
type OutboundConnectionHandlerCreator func(space app.Space, config interface{}) (proxy.OutboundConnectionHandler, error)