package handler

import (
	"gnet_test1/internal/protocol"

	"github.com/panjf2000/gnet/v2"
)

// CalculateHandler 处理计算请求
type CalculateHandler struct{}

func (h *CalculateHandler) Handle(c gnet.Conn, cmdID uint16, body []byte) {
	result := "calculated result"
	protocol.SendPacket(c, uint32(cmdID), []byte(result))
}

// SmallHandler 处理小数据包请求
type SmallHandler struct{}

func (h *SmallHandler) Handle(c gnet.Conn, cmdID uint16, body []byte) {
	response := "small response"
	protocol.SendPacket(c, uint32(cmdID), []byte(response))
}

// MediumHandler 处理中等数据包请求
type MediumHandler struct{}

func (h *MediumHandler) Handle(c gnet.Conn, cmdID uint16, body []byte) {
	response := "medium response"
	protocol.SendPacket(c, uint32(cmdID), []byte(response))
}

// LargeHandler 处理大数据包请求
type LargeHandler struct{}

func (h *LargeHandler) Handle(c gnet.Conn, cmdID uint16, body []byte) {
	response := "large response"
	protocol.SendPacket(c, uint32(cmdID), []byte(response))
}
