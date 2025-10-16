
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [showsToolbarButton] class.
var showsToolbarButtonClass _showsToolbarButtonClass

func init() {
	showsToolbarButtonClass = _showsToolbarButtonClass{objc.GetClass("showsToolbarButton")}
}

type _showsToolbarButtonClass struct {
	objc.Class
}

// An interface definition for the [showsToolbarButton] class.
type IshowsToolbarButton interface {
	ID() objc.ID
}

type showsToolbarButton struct {
	id objc.ID
}

func showsToolbarButtonFrom(ptr unsafe.Pointer) showsToolbarButton {
	return showsToolbarButton{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ showsToolbarButton) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _showsToolbarButtonClass) Alloc() showsToolbarButton {
	rv := objc.Send[showsToolbarButton](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _showsToolbarButtonClass) New() showsToolbarButton {
	rv := objc.Send[showsToolbarButton](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewshowsToolbarButton creates and returns a new initialized instance.
func NewshowsToolbarButton() showsToolbarButton {
	return showsToolbarButtonClass.New()
}

// Init initializes the instance.
func (s_ showsToolbarButton) Init() showsToolbarButton {
	rv := objc.Send[showsToolbarButton](s_.ID(), selInit)
	return rv
}
