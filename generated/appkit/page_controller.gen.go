
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PageController] class.
var PageControllerClass _PageControllerClass

func init() {
	PageControllerClass = _PageControllerClass{objc.GetClass("NSPageController")}
}

type _PageControllerClass struct {
	objc.Class
}

// An interface definition for the [PageController] class.
type IPageController interface {
	ID() objc.ID
}

type PageController struct {
	id objc.ID
}

func PageControllerFrom(ptr unsafe.Pointer) PageController {
	return PageController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PageController) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PageControllerClass) Alloc() PageController {
	rv := objc.Send[PageController](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PageControllerClass) New() PageController {
	rv := objc.Send[PageController](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPageController creates and returns a new initialized instance.
func NewPageController() PageController {
	return PageControllerClass.New()
}

// Init initializes the instance.
func (p_ PageController) Init() PageController {
	rv := objc.Send[PageController](p_.ID(), selInit)
	return rv
}
