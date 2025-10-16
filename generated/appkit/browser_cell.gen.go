
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [BrowserCell] class.
var BrowserCellClass _BrowserCellClass

func init() {
	BrowserCellClass = _BrowserCellClass{objc.GetClass("NSBrowserCell")}
}

type _BrowserCellClass struct {
	objc.Class
}

// An interface definition for the [BrowserCell] class.
type IBrowserCell interface {
	ID() objc.ID
}

type BrowserCell struct {
	id objc.ID
}

func BrowserCellFrom(ptr unsafe.Pointer) BrowserCell {
	return BrowserCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ BrowserCell) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _BrowserCellClass) Alloc() BrowserCell {
	rv := objc.Send[BrowserCell](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _BrowserCellClass) New() BrowserCell {
	rv := objc.Send[BrowserCell](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewBrowserCell creates and returns a new initialized instance.
func NewBrowserCell() BrowserCell {
	return BrowserCellClass.New()
}

// Init initializes the instance.
func (b_ BrowserCell) Init() BrowserCell {
	rv := objc.Send[BrowserCell](b_.ID(), selInit)
	return rv
}
