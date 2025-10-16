
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [disableKeyEquivalentForDefaultButtonCell] class.
var disableKeyEquivalentForDefaultButtonCellClass _disableKeyEquivalentForDefaultButtonCellClass

func init() {
	disableKeyEquivalentForDefaultButtonCellClass = _disableKeyEquivalentForDefaultButtonCellClass{objc.GetClass("disableKeyEquivalentForDefaultButtonCell")}
}

type _disableKeyEquivalentForDefaultButtonCellClass struct {
	objc.Class
}

// An interface definition for the [disableKeyEquivalentForDefaultButtonCell] class.
type IdisableKeyEquivalentForDefaultButtonCell interface {
	ID() objc.ID
}

type disableKeyEquivalentForDefaultButtonCell struct {
	id objc.ID
}

func disableKeyEquivalentForDefaultButtonCellFrom(ptr unsafe.Pointer) disableKeyEquivalentForDefaultButtonCell {
	return disableKeyEquivalentForDefaultButtonCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ disableKeyEquivalentForDefaultButtonCell) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _disableKeyEquivalentForDefaultButtonCellClass) Alloc() disableKeyEquivalentForDefaultButtonCell {
	rv := objc.Send[disableKeyEquivalentForDefaultButtonCell](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _disableKeyEquivalentForDefaultButtonCellClass) New() disableKeyEquivalentForDefaultButtonCell {
	rv := objc.Send[disableKeyEquivalentForDefaultButtonCell](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdisableKeyEquivalentForDefaultButtonCell creates and returns a new initialized instance.
func NewdisableKeyEquivalentForDefaultButtonCell() disableKeyEquivalentForDefaultButtonCell {
	return disableKeyEquivalentForDefaultButtonCellClass.New()
}

// Init initializes the instance.
func (d_ disableKeyEquivalentForDefaultButtonCell) Init() disableKeyEquivalentForDefaultButtonCell {
	rv := objc.Send[disableKeyEquivalentForDefaultButtonCell](d_.ID(), selInit)
	return rv
}
