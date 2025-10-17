
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ObjectController] class.
var ObjectControllerClass _ObjectControllerClass

func init() {
	ObjectControllerClass = _ObjectControllerClass{objc.GetClass("NSObjectController")}
}

type _ObjectControllerClass struct {
	objc.Class
}

// An interface definition for the [ObjectController] class.
type IObjectController interface {
	ID() objc.ID
}

type ObjectController struct {
	id objc.ID
}

func ObjectControllerFrom(ptr unsafe.Pointer) ObjectController {
	return ObjectController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ ObjectController) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _ObjectControllerClass) Alloc() ObjectController {
	rv := objc.Send[ObjectController](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _ObjectControllerClass) New() ObjectController {
	rv := objc.Send[ObjectController](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewObjectController creates and returns a new initialized instance.
func NewObjectController() ObjectController {
	return ObjectControllerClass.New()
}

// Init initializes the instance.
func (o_ ObjectController) Init() ObjectController {
	rv := objc.Send[ObjectController](o_.ID(), selInit)
	return rv
}
