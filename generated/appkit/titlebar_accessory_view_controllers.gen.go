
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [titlebarAccessoryViewControllers] class.
var titlebarAccessoryViewControllersClass _titlebarAccessoryViewControllersClass

func init() {
	titlebarAccessoryViewControllersClass = _titlebarAccessoryViewControllersClass{objc.GetClass("titlebarAccessoryViewControllers")}
}

type _titlebarAccessoryViewControllersClass struct {
	objc.Class
}

// An interface definition for the [titlebarAccessoryViewControllers] class.
type ItitlebarAccessoryViewControllers interface {
	ID() objc.ID
}

type titlebarAccessoryViewControllers struct {
	id objc.ID
}

func titlebarAccessoryViewControllersFrom(ptr unsafe.Pointer) titlebarAccessoryViewControllers {
	return titlebarAccessoryViewControllers{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ titlebarAccessoryViewControllers) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _titlebarAccessoryViewControllersClass) Alloc() titlebarAccessoryViewControllers {
	rv := objc.Send[titlebarAccessoryViewControllers](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _titlebarAccessoryViewControllersClass) New() titlebarAccessoryViewControllers {
	rv := objc.Send[titlebarAccessoryViewControllers](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtitlebarAccessoryViewControllers creates and returns a new initialized instance.
func NewtitlebarAccessoryViewControllers() titlebarAccessoryViewControllers {
	return titlebarAccessoryViewControllersClass.New()
}

// Init initializes the instance.
func (t_ titlebarAccessoryViewControllers) Init() titlebarAccessoryViewControllers {
	rv := objc.Send[titlebarAccessoryViewControllers](t_.ID(), selInit)
	return rv
}
