// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var ListFormatterClass _ListFormatterClass

func init() {
	ListFormatterClass = _ListFormatterClass{objc.GetClass("NSListFormatter")}
}

type _ListFormatterClass struct {
	class objc.Class
}

type ListFormatter struct {
	objc.ID
}

func ListFormatterFrom(ptr unsafe.Pointer) ListFormatter {
	return ListFormatter{
		ID: objc.ID(ptr),
	}
}




