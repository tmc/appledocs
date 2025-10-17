// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Calendar] class.
var CalendarClass objc.Class

func init() {
	CalendarClass = objc.GetClass("NSCalendar")
}

type Calendar struct {
	objc.ID
}

func CalendarFrom(ptr unsafe.Pointer) Calendar {
	return Calendar{
		ID: objc.ID(ptr),
	}
}



