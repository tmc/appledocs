// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DateComponents] class.
var DateComponentsClass objc.Class

func init() {
	DateComponentsClass = objc.GetClass("NSDateComponents")
}

type DateComponents struct {
	objc.ID
}

func DateComponentsFrom(ptr unsafe.Pointer) DateComponents {
	return DateComponents{
		ID: objc.ID(ptr),
	}
}




