// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DateFormatter] class.
var DateFormatterClass objc.Class

func init() {
	DateFormatterClass = objc.GetClass("NSDateFormatter")
}

type DateFormatter struct {
	objc.ID
}

func DateFormatterFrom(ptr unsafe.Pointer) DateFormatter {
	return DateFormatter{
		ID: objc.ID(ptr),
	}
}




