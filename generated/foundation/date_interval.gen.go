// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DateInterval] class.
var DateIntervalClass = _DateIntervalClass{objc.GetClass("NSDateInterval")}

type _DateIntervalClass struct {
	class objc.Class
}

type DateInterval struct {
	objc.ID
}

func DateIntervalFrom(ptr unsafe.Pointer) DateInterval {
	return DateInterval{
		ID: objc.ID(ptr),
	}
}




