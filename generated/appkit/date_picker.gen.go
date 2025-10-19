// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DatePicker] class.
var datePickerClass = _DatePickerClass{objc.GetClass("NSDatePicker")}

type _DatePickerClass struct {
	class objc.Class
}

// An interface definition for the [DatePicker] class.
type IDatePicker interface {
	IControl
}

// A display of a calendar date with controls for editing the date value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker

type DatePicker struct {
	Control
}

// DatePickerFrom constructs a [DatePicker] from an unsafe.Pointer.
//
// A display of a calendar date with controls for editing the date value.
func DatePickerFrom(ptr unsafe.Pointer) DatePicker {
	return DatePicker{
		Control: ControlFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (dc _DatePickerClass) Alloc() DatePicker {
	rv := objc.Send[DatePicker](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (dc _DatePickerClass) New() DatePicker {
	rv := objc.Send[DatePicker](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DatePicker) Init() DatePicker {
	rv := objc.Send[DatePicker](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DatePicker) Autorelease() DatePicker {
	rv := objc.Send[DatePicker](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDatePicker creates a new DatePicker instance.
func NewDatePicker() DatePicker {
	return datePickerClass.New()
}




