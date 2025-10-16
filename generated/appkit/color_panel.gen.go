
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ColorPanel] class.
var ColorPanelClass _ColorPanelClass

func init() {
	ColorPanelClass = _ColorPanelClass{objc.GetClass("NSColorPanel")}
}

type _ColorPanelClass struct {
	objc.Class
}

// An interface definition for the [ColorPanel] class.
type IColorPanel interface {
	ID() objc.ID
}

type ColorPanel struct {
	id objc.ID
}

func ColorPanelFrom(ptr unsafe.Pointer) ColorPanel {
	return ColorPanel{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ ColorPanel) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ColorPanelClass) Alloc() ColorPanel {
	rv := objc.Send[ColorPanel](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ColorPanelClass) New() ColorPanel {
	rv := objc.Send[ColorPanel](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewColorPanel creates and returns a new initialized instance.
func NewColorPanel() ColorPanel {
	return ColorPanelClass.New()
}

// Init initializes the instance.
func (c_ ColorPanel) Init() ColorPanel {
	rv := objc.Send[ColorPanel](c_.ID(), selInit)
	return rv
}
