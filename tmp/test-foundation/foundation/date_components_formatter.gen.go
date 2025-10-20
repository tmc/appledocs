// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var DateComponentsFormatterClass _DateComponentsFormatterClass

func init() {
	DateComponentsFormatterClass = _DateComponentsFormatterClass{objc.GetClass("NSDateComponentsFormatter")}
}

type _DateComponentsFormatterClass struct {
	class objc.Class
}

type DateComponentsFormatter struct {
	objc.ID
}

func DateComponentsFormatterFrom(ptr unsafe.Pointer) DateComponentsFormatter {
	return DateComponentsFormatter{
		ID: objc.ID(ptr),
	}
}




