package mock

import (
	"fmt"

	"github.com/xm-chentl/go-code/noticeex"
)

type execFunc func(msg string) error

type mockImpl struct {
	call execFunc
}

func (m mockImpl) Sendf(format string, args ...interface{}) error {
	return m.call(fmt.Sprintf(format, args...))
}

func New(cb execFunc) noticeex.INotice {
	return &mockImpl{
		call: cb,
	}
}
