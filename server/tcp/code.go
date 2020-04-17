package tcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	Head = "arou"
	HeadLen = 4
	DataID  = 4
	DataLen = 4
)

type Code struct {
	Buffer 		[]byte
	LeftBuffer	[]byte
	MsgLen      int
}

//arou 后续buffer长度可能需改变？
func NewCode() Code {
	return Code{
		Buffer:     make([]byte, 1024),
		LeftBuffer: make([]byte, 1024),
		MsgLen:		0,
	}
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

func BytesToUInt(b []byte) uint {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	err := binary.Read(bytesBuffer, binary.BigEndian, &x)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return uint(x)
}

//arou 避免粘包？
func (c Code) Unpack() UnPackData {
	var data UnPackData
	data.list = make([][]byte, 0)

	buffer := append(c.LeftBuffer, c.Buffer[:c.MsgLen]...)
	length := len(buffer)

	var i int
	for i = 0; i < length; i++ {
		if length < i + HeadLen + DataLen {
			break
		}
		if string(buffer[i : i+HeadLen]) == Head {
			messageLen := BytesToInt(buffer[i+HeadLen:i+HeadLen+DataLen])
			if length < i + HeadLen + DataLen + messageLen {
				break
			}
			info := buffer[i+HeadLen+DataLen:i+HeadLen+DataLen+messageLen]
			data.list = append(data.list, info)

			i += HeadLen + DataLen + messageLen - 1
		}
	}

	c.LeftBuffer = make([]byte, 0)
	if i != length {
		c.LeftBuffer =  buffer[i:]
	}

	return data
}