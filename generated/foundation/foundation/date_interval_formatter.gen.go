// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DateIntervalFormatter] class.
var DateIntervalFormatterClass objc.Class

func init() {
	DateIntervalFormatterClass = objc.GetClass("NSDateIntervalFormatter")
}

type DateIntervalFormatter struct {
	objc.ID
}

func DateIntervalFormatterFrom(ptr unsafe.Pointer) DateIntervalFormatter {
	return DateIntervalFormatter{
		ID: objc.ID(ptr),
	}
}




