
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DatePickerCell] class.
var DatePickerCellClass _DatePickerCellClass

func init() {
	DatePickerCellClass = _DatePickerCellClass{objc.GetClass("NSDatePickerCell")}
}

type _DatePickerCellClass struct {
	objc.Class
}

// An interface definition for the [DatePickerCell] class.
type IDatePickerCell interface {
	ID() objc.ID
}

type DatePickerCell struct {
	id objc.ID
}

func DatePickerCellFrom(ptr unsafe.Pointer) DatePickerCell {
	return DatePickerCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ DatePickerCell) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _DatePickerCellClass) Alloc() DatePickerCell {
	rv := objc.Send[DatePickerCell](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _DatePickerCellClass) New() DatePickerCell {
	rv := objc.Send[DatePickerCell](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewDatePickerCell creates and returns a new initialized instance.
func NewDatePickerCell() DatePickerCell {
	return DatePickerCellClass.New()
}

// Init initializes the instance.
func (d_ DatePickerCell) Init() DatePickerCell {
	rv := objc.Send[DatePickerCell](d_.ID(), selInit)
	return rv
}
