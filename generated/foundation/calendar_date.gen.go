// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CalendarDate] class.
var CalendarDateClass objc.Class

func init() {
	CalendarDateClass = objc.GetClass("NSCalendarDate")
}

type CalendarDate struct {
	objc.ID
}

func CalendarDateFrom(ptr unsafe.Pointer) CalendarDate {
	return CalendarDate{
		ID: objc.ID(ptr),
	}
}



