package service

import (
	"encoding/hex"
	. "github.com/irisnet/rainbow-sync/msgs"
	"github.com/irisnet/rainbow-sync/utils"
	"strings"
)

type (
	DocMsgPauseRequestContext struct {
		RequestContextId string `bson:"request_context_id" yaml:"request_context_id"`
		Consumer         string `bson:"consumer" yaml:"consumer"`
	}
)

func (m *DocMsgPauseRequestContext) GetType() string {
	return MsgTypePauseRequestContext
}

func (m *DocMsgPauseRequestContext) BuildMsg(v interface{}) {
	msg := v.(*MsgPauseRequestContext)

	m.RequestContextId = strings.ToUpper(hex.EncodeToString(msg.RequestContextId))
	m.Consumer = msg.Consumer.String()
}

func (m *DocMsgPauseRequestContext) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgPauseRequestContext
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Consumer.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}