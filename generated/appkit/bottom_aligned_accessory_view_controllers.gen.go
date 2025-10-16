
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [bottomAlignedAccessoryViewControllers] class.
var bottomAlignedAccessoryViewControllersClass _bottomAlignedAccessoryViewControllersClass

func init() {
	bottomAlignedAccessoryViewControllersClass = _bottomAlignedAccessoryViewControllersClass{objc.GetClass("bottomAlignedAccessoryViewControllers")}
}

type _bottomAlignedAccessoryViewControllersClass struct {
	objc.Class
}

// An interface definition for the [bottomAlignedAccessoryViewControllers] class.
type IbottomAlignedAccessoryViewControllers interface {
	ID() objc.ID
}

type bottomAlignedAccessoryViewControllers struct {
	id objc.ID
}

func bottomAlignedAccessoryViewControllersFrom(ptr unsafe.Pointer) bottomAlignedAccessoryViewControllers {
	return bottomAlignedAccessoryViewControllers{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ bottomAlignedAccessoryViewControllers) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _bottomAlignedAccessoryViewControllersClass) Alloc() bottomAlignedAccessoryViewControllers {
	rv := objc.Send[bottomAlignedAccessoryViewControllers](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _bottomAlignedAccessoryViewControllersClass) New() bottomAlignedAccessoryViewControllers {
	rv := objc.Send[bottomAlignedAccessoryViewControllers](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbottomAlignedAccessoryViewControllers creates and returns a new initialized instance.
func NewbottomAlignedAccessoryViewControllers() bottomAlignedAccessoryViewControllers {
	return bottomAlignedAccessoryViewControllersClass.New()
}

// Init initializes the instance.
func (b_ bottomAlignedAccessoryViewControllers) Init() bottomAlignedAccessoryViewControllers {
	rv := objc.Send[bottomAlignedAccessoryViewControllers](b_.ID(), selInit)
	return rv
}
