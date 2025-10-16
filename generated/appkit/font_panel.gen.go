
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FontPanel] class.
var FontPanelClass _FontPanelClass

func init() {
	FontPanelClass = _FontPanelClass{objc.GetClass("NSFontPanel")}
}

type _FontPanelClass struct {
	objc.Class
}

// An interface definition for the [FontPanel] class.
type IFontPanel interface {
	ID() objc.ID
}

type FontPanel struct {
	id objc.ID
}

func FontPanelFrom(ptr unsafe.Pointer) FontPanel {
	return FontPanel{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ FontPanel) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _FontPanelClass) Alloc() FontPanel {
	rv := objc.Send[FontPanel](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _FontPanelClass) New() FontPanel {
	rv := objc.Send[FontPanel](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewFontPanel creates and returns a new initialized instance.
func NewFontPanel() FontPanel {
	return FontPanelClass.New()
}

// Init initializes the instance.
func (f_ FontPanel) Init() FontPanel {
	rv := objc.Send[FontPanel](f_.ID(), selInit)
	return rv
}
