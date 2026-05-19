package protocol

import (
	"encoding/binary"
	"log"

	"github.com/panjf2000/gnet/v2"
)

const (
	HeaderLen = 8

	CmdCalculate uint32 = 1001
	CmdSmall     uint32 = 2001
	CmdMedium    uint32 = 3001
	CmdLarge     uint32 = 4001
	CmdShutDown  uint32 = 9999
)

// Conn 定义连接接口
// type Conn interface {
// 	RemoteAddr() net.Addr
// 	AsyncWrite([]byte, gnet.AsyncCallback) error
// }

func SendPacket(c gnet.Conn, cmdID uint32, body []byte) {
	var header [HeaderLen]byte
	binary.BigEndian.PutUint32(header[0:4], cmdID)
	binary.BigEndian.PutUint32(header[4:8], uint32(len(body)))

	err := c.AsyncWritev([][]byte{header[:], body}, nil)
	if err != nil {
		log.Printf("[发送失败] 无法投递回包给 %s: %v", c.RemoteAddr().String(), err)
	}
}
