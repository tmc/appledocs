
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isModalPanel] class.
var isModalPanelClass _isModalPanelClass

func init() {
	isModalPanelClass = _isModalPanelClass{objc.GetClass("isModalPanel")}
}

type _isModalPanelClass struct {
	objc.Class
}

// An interface definition for the [isModalPanel] class.
type IisModalPanel interface {
	ID() objc.ID
}

type isModalPanel struct {
	id objc.ID
}

func isModalPanelFrom(ptr unsafe.Pointer) isModalPanel {
	return isModalPanel{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isModalPanel) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isModalPanelClass) Alloc() isModalPanel {
	rv := objc.Send[isModalPanel](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isModalPanelClass) New() isModalPanel {
	rv := objc.Send[isModalPanel](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisModalPanel creates and returns a new initialized instance.
func NewisModalPanel() isModalPanel {
	return isModalPanelClass.New()
}

// Init initializes the instance.
func (i_ isModalPanel) Init() isModalPanel {
	rv := objc.Send[isModalPanel](i_.ID(), selInit)
	return rv
}
