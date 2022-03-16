package netex

import "github.com/xm-chentl/go-code/netex/nettype"

type INetFactory interface {
	Build(nettype.Value) INetService
}
