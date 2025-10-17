
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ArrayController] class.
var ArrayControllerClass _ArrayControllerClass

func init() {
	ArrayControllerClass = _ArrayControllerClass{objc.GetClass("NSArrayController")}
}

type _ArrayControllerClass struct {
	objc.Class
}

// An interface definition for the [ArrayController] class.
type IArrayController interface {
	ID() objc.ID
}

type ArrayController struct {
	id objc.ID
}

func ArrayControllerFrom(ptr unsafe.Pointer) ArrayController {
	return ArrayController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ ArrayController) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _ArrayControllerClass) Alloc() ArrayController {
	rv := objc.Send[ArrayController](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _ArrayControllerClass) New() ArrayController {
	rv := objc.Send[ArrayController](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewArrayController creates and returns a new initialized instance.
func NewArrayController() ArrayController {
	return ArrayControllerClass.New()
}

// Init initializes the instance.
func (a_ ArrayController) Init() ArrayController {
	rv := objc.Send[ArrayController](a_.ID(), selInit)
	return rv
}
