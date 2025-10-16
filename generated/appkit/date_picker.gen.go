
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DatePicker] class.
var DatePickerClass _DatePickerClass

func init() {
	DatePickerClass = _DatePickerClass{objc.GetClass("NSDatePicker")}
}

type _DatePickerClass struct {
	objc.Class
}

// An interface definition for the [DatePicker] class.
type IDatePicker interface {
	ID() objc.ID
}

type DatePicker struct {
	id objc.ID
}

func DatePickerFrom(ptr unsafe.Pointer) DatePicker {
	return DatePicker{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ DatePicker) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _DatePickerClass) Alloc() DatePicker {
	rv := objc.Send[DatePicker](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _DatePickerClass) New() DatePicker {
	rv := objc.Send[DatePicker](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewDatePicker creates and returns a new initialized instance.
func NewDatePicker() DatePicker {
	return DatePickerClass.New()
}

// Init initializes the instance.
func (d_ DatePicker) Init() DatePicker {
	rv := objc.Send[DatePicker](d_.ID(), selInit)
	return rv
}
