
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [showsArrowButtons] class.
var showsArrowButtonsClass _showsArrowButtonsClass

func init() {
	showsArrowButtonsClass = _showsArrowButtonsClass{objc.GetClass("showsArrowButtons")}
}

type _showsArrowButtonsClass struct {
	objc.Class
}

// An interface definition for the [showsArrowButtons] class.
type IshowsArrowButtons interface {
	ID() objc.ID
}

type showsArrowButtons struct {
	id objc.ID
}

func showsArrowButtonsFrom(ptr unsafe.Pointer) showsArrowButtons {
	return showsArrowButtons{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ showsArrowButtons) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _showsArrowButtonsClass) Alloc() showsArrowButtons {
	rv := objc.Send[showsArrowButtons](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _showsArrowButtonsClass) New() showsArrowButtons {
	rv := objc.Send[showsArrowButtons](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewshowsArrowButtons creates and returns a new initialized instance.
func NewshowsArrowButtons() showsArrowButtons {
	return showsArrowButtonsClass.New()
}

// Init initializes the instance.
func (s_ showsArrowButtons) Init() showsArrowButtons {
	rv := objc.Send[showsArrowButtons](s_.ID(), selInit)
	return rv
}
