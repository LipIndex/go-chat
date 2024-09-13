package utils

import (
	"unsafe"
)

// 无额外内存消耗 byte转string
func ByteToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func StringToByte(data string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&data))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
