package tcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	proto2 "whisky-server/server/tcp/proto"
)

const (
	Head = "arou"
	HeadLen = 4
	DataID  = 4
	DataLen = 4
)

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

func Unpack(buffer []byte, readerChannel chan []byte) []byte {
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
			data := buffer[i+HeadLen+DataLen:i+HeadLen+DataLen+messageLen]
			test := proto2.Test{}
			projectId := BytesToUInt(data[:DataID])
			data = data[DataID:]
			proto.Unmarshal(data, &test)
			fmt.Println(projectId)
			fmt.Println(test.Text)
			fmt.Println(test.No)

			i += HeadLen + DataLen + messageLen - 1
		}
	}

	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}