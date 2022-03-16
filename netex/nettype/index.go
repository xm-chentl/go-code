package nettype

type Value int

const (
	Http Value = iota + 1
	MessageQueue
	RPC
)
