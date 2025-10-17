// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Calendar] class.
var CalendarClass = _CalendarClass{objc.GetClass("NSCalendar")}

type _CalendarClass struct {
	class objc.Class
}

type Calendar struct {
	objc.ID
}

func CalendarFrom(ptr unsafe.Pointer) Calendar {
	return Calendar{
		ID: objc.ID(ptr),
	}
}




