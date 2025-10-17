
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Controller] class.
var ControllerClass _ControllerClass

func init() {
	ControllerClass = _ControllerClass{objc.GetClass("NSController")}
}

type _ControllerClass struct {
	objc.Class
}

// An interface definition for the [Controller] class.
type IController interface {
	ID() objc.ID
}

type Controller struct {
	id objc.ID
}

func ControllerFrom(ptr unsafe.Pointer) Controller {
	return Controller{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ Controller) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ControllerClass) Alloc() Controller {
	rv := objc.Send[Controller](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ControllerClass) New() Controller {
	rv := objc.Send[Controller](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewController creates and returns a new initialized instance.
func NewController() Controller {
	return ControllerClass.New()
}

// Init initializes the instance.
func (c_ Controller) Init() Controller {
	rv := objc.Send[Controller](c_.ID(), selInit)
	return rv
}
