
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [defaultButtonCell] class.
var defaultButtonCellClass _defaultButtonCellClass

func init() {
	defaultButtonCellClass = _defaultButtonCellClass{objc.GetClass("defaultButtonCell")}
}

type _defaultButtonCellClass struct {
	objc.Class
}

// An interface definition for the [defaultButtonCell] class.
type IdefaultButtonCell interface {
	ID() objc.ID
}

type defaultButtonCell struct {
	id objc.ID
}

func defaultButtonCellFrom(ptr unsafe.Pointer) defaultButtonCell {
	return defaultButtonCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ defaultButtonCell) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _defaultButtonCellClass) Alloc() defaultButtonCell {
	rv := objc.Send[defaultButtonCell](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _defaultButtonCellClass) New() defaultButtonCell {
	rv := objc.Send[defaultButtonCell](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdefaultButtonCell creates and returns a new initialized instance.
func NewdefaultButtonCell() defaultButtonCell {
	return defaultButtonCellClass.New()
}

// Init initializes the instance.
func (d_ defaultButtonCell) Init() defaultButtonCell {
	rv := objc.Send[defaultButtonCell](d_.ID(), selInit)
	return rv
}
