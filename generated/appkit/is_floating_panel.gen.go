
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isFloatingPanel] class.
var isFloatingPanelClass _isFloatingPanelClass

func init() {
	isFloatingPanelClass = _isFloatingPanelClass{objc.GetClass("isFloatingPanel")}
}

type _isFloatingPanelClass struct {
	objc.Class
}

// An interface definition for the [isFloatingPanel] class.
type IisFloatingPanel interface {
	ID() objc.ID
}

type isFloatingPanel struct {
	id objc.ID
}

func isFloatingPanelFrom(ptr unsafe.Pointer) isFloatingPanel {
	return isFloatingPanel{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isFloatingPanel) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isFloatingPanelClass) Alloc() isFloatingPanel {
	rv := objc.Send[isFloatingPanel](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isFloatingPanelClass) New() isFloatingPanel {
	rv := objc.Send[isFloatingPanel](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisFloatingPanel creates and returns a new initialized instance.
func NewisFloatingPanel() isFloatingPanel {
	return isFloatingPanelClass.New()
}

// Init initializes the instance.
func (i_ isFloatingPanel) Init() isFloatingPanel {
	rv := objc.Send[isFloatingPanel](i_.ID(), selInit)
	return rv
}
