// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DateFormatter] class.
var DateFormatterClass = _DateFormatterClass{objc.GetClass("NSDateFormatter")}

type _DateFormatterClass struct {
	class objc.Class
}

type DateFormatter struct {
	objc.ID
}

func DateFormatterFrom(ptr unsafe.Pointer) DateFormatter {
	return DateFormatter{
		ID: objc.ID(ptr),
	}
}




