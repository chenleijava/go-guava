package guava

import (
	"encoding/binary"
	"github.com/chenleijava/go-guava/bufpool"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//获取当前执行文件的相对路径
func ExePath() string {
	execPath, _ := exec.LookPath(os.Args[0])
	return execPath[0:strings.LastIndex(execPath, "/")]
}

//获取source的子串,如果start小于0或者end大于source长度则返回""
//start:开始index，从0开始，包括0
//end:结束index，以end结束，但不包括end
func SubString(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)
	if start < 0 || end > length || start > end {
		return ""
	}
	if start == 0 && end == length {
		return source
	}
	return string(r[start:end])
}

//设置log库格式化
func LogFormatInit() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
}

//计算分页偏移量
func Offset(page, limit int) int {
	return (page - 1) * limit
}

//string to int
func String2Int(value string) int {
	v, e := strconv.Atoi(value)
	if e != nil {
		log.Printf("string2int error:%s", e)
		return -1
	}
	return v
}

//int to string
func Int2String(value int) string {
	return strconv.Itoa(value)
}

// remove from slice
func Remove(s []string, value string) []string {
	for i, p := range s {
		if p == value {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}
	return s
}

//clear int map
func ClearIntMap(mp *map[string]int) {
	for k := range *mp {
		delete(*mp, k)
	}
}

//clear string map
func ClearStringMap(mp *map[string]string) {
	for k := range *mp {
		delete(*mp, k)
	}
}

// int to string
func ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

//int to bytes
func DataToBytes(n interface{}, order binary.ByteOrder) []byte {
	buf := bufpool.GetBytesBuffer()
	defer bufpool.PutBytesBuffer(buf)

	err := binary.Write(buf, order, n)
	if err != nil {
		log.Fatalf("err:%s", err.Error())
	}
	return buf.Bytes()
}
