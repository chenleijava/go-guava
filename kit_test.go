package guava

import (
	"encoding/binary"
	"log"
	"sync/atomic"
	"testing"
)

//test into bytes
func TestIntToBytes(t *testing.T) {
	log.Printf("%+v", DataToBytes(int32(12), binary.BigEndian))
	log.Printf("%+v", DataToBytes(int32(12), binary.LittleEndian))
	log.Printf("%+v", []byte(Int2String(12)))
	var v int32 = 3
	addValue(&v)
	log.Printf("%d", v)

}

func addValue(value *int32) {
	for {
		// init
		v := *value
		//if v==*value ,swap old and new in value
		if atomic.CompareAndSwapInt32(value, v, *value+100) {
			break
		}
	}
}
