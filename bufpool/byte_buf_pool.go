package bufpool

import (
	"bytes"
	"sync"
)

//bytes buf pool
var byteBufPool = &sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer([]byte{})
	},
}

//get bytes buffer
func GetBytesBuffer() *bytes.Buffer {
	return byteBufPool.Get().(*bytes.Buffer)
}

//put bytes buffer
func PutBytesBuffer(buf *bytes.Buffer) {
	buf.Reset()
	byteBufPool.Put(buf)
}

//string buf pool
var stringBufPool = &sync.Pool{
	New: func() interface{} {
		return bytes.NewBufferString("")
	},
}

//get string buf
func GetStringBuffer() *bytes.Buffer {
	return stringBufPool.Get().(*bytes.Buffer)
}

//put string buf
func PutStringBuffer(buf *bytes.Buffer) {
	buf.Reset()
	stringBufPool.Put(buf)
}
