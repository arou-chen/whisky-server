package tcp

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"reflect"
	"whisky-server/server"
	proto2 "whisky-server/server/tcp/proto"
	"whisky-server/server/tcp/util"
)

type Server struct {
	Coder  Code
	handle map[uint]Method
}

type UnPackData struct {
	list [][]byte
}

func NewServer() server.Server {
	code := NewCode()
	return &Server{
		Coder:   code,
		handle: make(map[uint]Method),
	}
}

func ProtoTest(in *proto2.Test) (out proto2.Test, err error) {
	fmt.Println(in.No)
	fmt.Println(in.Text)
	return proto2.Test{}, nil
}

func (s Server) Server() error {
	l, err := net.Listen("tcp",":3207")
	if err != nil {
		return err
	}

	s.Register(ProtoTest, 1)

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go s.handleConnection(conn)
	}

}

func (s Server) Stop() error {
	return nil
}

func (s Server) handleConnection(conn net.Conn) {
	var count int32
	for {
		count++
		n, err := conn.Read(s.Coder.Buffer)
		if err != nil {
			return
		}
		s.Coder.MsgLen = n
		unpackData := s.Coder.Unpack()
		for _, data := range unpackData.list {
			s.handleData(data)
		}
	}
}

func (s Server) handleData(data []byte) {
	id := BytesToUInt(data[:DataID])
	method, ok := s.handle[id]
	if !ok {
		return
	}
	protoData := data[DataID:]
	//fmt.Println(reflect.TypeOf(v))
	pt := reflect.Indirect(reflect.New(method.In.Elem()).Elem()).Addr().Interface().(proto.Message)
	err := proto.Unmarshal(protoData, pt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	method.Handler(pt)
}

func (s Server) Register(fn interface{}, id uint) {
	rv := reflect.ValueOf(fn)
	rt := rv.Type()
	name, _ := util.GetFuncName(rv)
	inTyp := rt.In(0)
	outTyp := rt.Out(0)
	handler := func(in interface{}) (out interface{}, err error) {
		rv.Call([]reflect.Value{reflect.ValueOf(in)})
		return nil, nil
	}
	method := NewMethod(name, handler, inTyp, outTyp)
	s.handle[id] = method
}