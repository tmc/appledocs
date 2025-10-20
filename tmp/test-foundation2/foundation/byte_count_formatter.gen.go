// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var byteCountFormatterClass _ByteCountFormatterClass

func init() {
	byteCountFormatterClass = _ByteCountFormatterClass{objc.GetClass("NSByteCountFormatter")}
}

type _ByteCountFormatterClass struct {
	class objc.Class
}

type ByteCountFormatter struct {
	objc.ID
}

func ByteCountFormatterFrom(ptr unsafe.Pointer) ByteCountFormatter {
	return ByteCountFormatter{
		ID: objc.ID(ptr),
	}
}




