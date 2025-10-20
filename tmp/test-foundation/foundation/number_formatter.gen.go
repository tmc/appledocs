// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var NumberFormatterClass _NumberFormatterClass

func init() {
	NumberFormatterClass = _NumberFormatterClass{objc.GetClass("NSNumberFormatter")}
}

type _NumberFormatterClass struct {
	class objc.Class
}

type NumberFormatter struct {
	objc.ID
}

func NumberFormatterFrom(ptr unsafe.Pointer) NumberFormatter {
	return NumberFormatter{
		ID: objc.ID(ptr),
	}
}




