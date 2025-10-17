// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DateInterval] class.
var DateIntervalClass objc.Class

func init() {
	DateIntervalClass = objc.GetClass("NSDateInterval")
}

type DateInterval struct {
	objc.ID
}

func DateIntervalFrom(ptr unsafe.Pointer) DateInterval {
	return DateInterval{
		ID: objc.ID(ptr),
	}
}




