// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DateComponents] class.
var DateComponentsClass = _DateComponentsClass{objc.GetClass("NSDateComponents")}

type _DateComponentsClass struct {
	class objc.Class
}

type DateComponents struct {
	objc.ID
}

func DateComponentsFrom(ptr unsafe.Pointer) DateComponents {
	return DateComponents{
		ID: objc.ID(ptr),
	}
}




