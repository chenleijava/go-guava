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
