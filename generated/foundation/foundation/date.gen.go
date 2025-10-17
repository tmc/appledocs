// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Date] class.
var DateClass objc.Class

func init() {
	DateClass = objc.GetClass("NSDate")
}

type Date struct {
	objc.ID
}

func DateFrom(ptr unsafe.Pointer) Date {
	return Date{
		ID: objc.ID(ptr),
	}
}




