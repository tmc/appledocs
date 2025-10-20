// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var calendarDateClass _CalendarDateClass

func init() {
	calendarDateClass = _CalendarDateClass{objc.GetClass("NSCalendarDate")}
}

type _CalendarDateClass struct {
	class objc.Class
}

type CalendarDate struct {
	objc.ID
}

func CalendarDateFrom(ptr unsafe.Pointer) CalendarDate {
	return CalendarDate{
		ID: objc.ID(ptr),
	}
}




