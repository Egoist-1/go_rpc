package codec

import "io"

type Header struct {
	//服务名的方法名
	ServiceMethod string
	//请求的序列号
	Seq uint64
	//错误的信息
	Error error
}

// 编解码接口
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(any) error
	Write(*Header, any) error
}
type Type string
type NewCodecFunc func(closer io.ReadWriteCloser) Codec

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func Init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
