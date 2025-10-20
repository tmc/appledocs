// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var DateIntervalFormatterClass _DateIntervalFormatterClass

func init() {
	DateIntervalFormatterClass = _DateIntervalFormatterClass{objc.GetClass("NSDateIntervalFormatter")}
}

type _DateIntervalFormatterClass struct {
	class objc.Class
}

type DateIntervalFormatter struct {
	objc.ID
}

func DateIntervalFormatterFrom(ptr unsafe.Pointer) DateIntervalFormatter {
	return DateIntervalFormatter{
		ID: objc.ID(ptr),
	}
}




