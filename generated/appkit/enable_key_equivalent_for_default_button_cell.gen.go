
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [enableKeyEquivalentForDefaultButtonCell] class.
var enableKeyEquivalentForDefaultButtonCellClass _enableKeyEquivalentForDefaultButtonCellClass

func init() {
	enableKeyEquivalentForDefaultButtonCellClass = _enableKeyEquivalentForDefaultButtonCellClass{objc.GetClass("enableKeyEquivalentForDefaultButtonCell")}
}

type _enableKeyEquivalentForDefaultButtonCellClass struct {
	objc.Class
}

// An interface definition for the [enableKeyEquivalentForDefaultButtonCell] class.
type IenableKeyEquivalentForDefaultButtonCell interface {
	ID() objc.ID
}

type enableKeyEquivalentForDefaultButtonCell struct {
	id objc.ID
}

func enableKeyEquivalentForDefaultButtonCellFrom(ptr unsafe.Pointer) enableKeyEquivalentForDefaultButtonCell {
	return enableKeyEquivalentForDefaultButtonCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ enableKeyEquivalentForDefaultButtonCell) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _enableKeyEquivalentForDefaultButtonCellClass) Alloc() enableKeyEquivalentForDefaultButtonCell {
	rv := objc.Send[enableKeyEquivalentForDefaultButtonCell](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _enableKeyEquivalentForDefaultButtonCellClass) New() enableKeyEquivalentForDefaultButtonCell {
	rv := objc.Send[enableKeyEquivalentForDefaultButtonCell](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewenableKeyEquivalentForDefaultButtonCell creates and returns a new initialized instance.
func NewenableKeyEquivalentForDefaultButtonCell() enableKeyEquivalentForDefaultButtonCell {
	return enableKeyEquivalentForDefaultButtonCellClass.New()
}

// Init initializes the instance.
func (e_ enableKeyEquivalentForDefaultButtonCell) Init() enableKeyEquivalentForDefaultButtonCell {
	rv := objc.Send[enableKeyEquivalentForDefaultButtonCell](e_.ID(), selInit)
	return rv
}
