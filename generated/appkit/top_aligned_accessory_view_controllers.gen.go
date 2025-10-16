
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [topAlignedAccessoryViewControllers] class.
var topAlignedAccessoryViewControllersClass _topAlignedAccessoryViewControllersClass

func init() {
	topAlignedAccessoryViewControllersClass = _topAlignedAccessoryViewControllersClass{objc.GetClass("topAlignedAccessoryViewControllers")}
}

type _topAlignedAccessoryViewControllersClass struct {
	objc.Class
}

// An interface definition for the [topAlignedAccessoryViewControllers] class.
type ItopAlignedAccessoryViewControllers interface {
	ID() objc.ID
}

type topAlignedAccessoryViewControllers struct {
	id objc.ID
}

func topAlignedAccessoryViewControllersFrom(ptr unsafe.Pointer) topAlignedAccessoryViewControllers {
	return topAlignedAccessoryViewControllers{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ topAlignedAccessoryViewControllers) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _topAlignedAccessoryViewControllersClass) Alloc() topAlignedAccessoryViewControllers {
	rv := objc.Send[topAlignedAccessoryViewControllers](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _topAlignedAccessoryViewControllersClass) New() topAlignedAccessoryViewControllers {
	rv := objc.Send[topAlignedAccessoryViewControllers](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtopAlignedAccessoryViewControllers creates and returns a new initialized instance.
func NewtopAlignedAccessoryViewControllers() topAlignedAccessoryViewControllers {
	return topAlignedAccessoryViewControllersClass.New()
}

// Init initializes the instance.
func (t_ topAlignedAccessoryViewControllers) Init() topAlignedAccessoryViewControllers {
	rv := objc.Send[topAlignedAccessoryViewControllers](t_.ID(), selInit)
	return rv
}
