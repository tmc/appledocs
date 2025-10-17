
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TreeController] class.
var TreeControllerClass _TreeControllerClass

func init() {
	TreeControllerClass = _TreeControllerClass{objc.GetClass("NSTreeController")}
}

type _TreeControllerClass struct {
	objc.Class
}

// An interface definition for the [TreeController] class.
type ITreeController interface {
	ID() objc.ID
}

type TreeController struct {
	id objc.ID
}

func TreeControllerFrom(ptr unsafe.Pointer) TreeController {
	return TreeController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TreeController) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TreeControllerClass) Alloc() TreeController {
	rv := objc.Send[TreeController](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TreeControllerClass) New() TreeController {
	rv := objc.Send[TreeController](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTreeController creates and returns a new initialized instance.
func NewTreeController() TreeController {
	return TreeControllerClass.New()
}

// Init initializes the instance.
func (t_ TreeController) Init() TreeController {
	rv := objc.Send[TreeController](t_.ID(), selInit)
	return rv
}
