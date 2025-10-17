// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DateComponentsFormatter] class.
var DateComponentsFormatterClass objc.Class

func init() {
	DateComponentsFormatterClass = objc.GetClass("NSDateComponentsFormatter")
}

type DateComponentsFormatter struct {
	objc.ID
}

func DateComponentsFormatterFrom(ptr unsafe.Pointer) DateComponentsFormatter {
	return DateComponentsFormatter{
		ID: objc.ID(ptr),
	}
}




