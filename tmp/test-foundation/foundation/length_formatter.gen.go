// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var LengthFormatterClass _LengthFormatterClass

func init() {
	LengthFormatterClass = _LengthFormatterClass{objc.GetClass("NSLengthFormatter")}
}

type _LengthFormatterClass struct {
	class objc.Class
}

type LengthFormatter struct {
	objc.ID
}

func LengthFormatterFrom(ptr unsafe.Pointer) LengthFormatter {
	return LengthFormatter{
		ID: objc.ID(ptr),
	}
}




