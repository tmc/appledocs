
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [canBecomeVisibleWithoutLogin] class.
var canBecomeVisibleWithoutLoginClass _canBecomeVisibleWithoutLoginClass

func init() {
	canBecomeVisibleWithoutLoginClass = _canBecomeVisibleWithoutLoginClass{objc.GetClass("canBecomeVisibleWithoutLogin")}
}

type _canBecomeVisibleWithoutLoginClass struct {
	objc.Class
}

// An interface definition for the [canBecomeVisibleWithoutLogin] class.
type IcanBecomeVisibleWithoutLogin interface {
	ID() objc.ID
}

type canBecomeVisibleWithoutLogin struct {
	id objc.ID
}

func canBecomeVisibleWithoutLoginFrom(ptr unsafe.Pointer) canBecomeVisibleWithoutLogin {
	return canBecomeVisibleWithoutLogin{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ canBecomeVisibleWithoutLogin) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _canBecomeVisibleWithoutLoginClass) Alloc() canBecomeVisibleWithoutLogin {
	rv := objc.Send[canBecomeVisibleWithoutLogin](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _canBecomeVisibleWithoutLoginClass) New() canBecomeVisibleWithoutLogin {
	rv := objc.Send[canBecomeVisibleWithoutLogin](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcanBecomeVisibleWithoutLogin creates and returns a new initialized instance.
func NewcanBecomeVisibleWithoutLogin() canBecomeVisibleWithoutLogin {
	return canBecomeVisibleWithoutLoginClass.New()
}

// Init initializes the instance.
func (c_ canBecomeVisibleWithoutLogin) Init() canBecomeVisibleWithoutLogin {
	rv := objc.Send[canBecomeVisibleWithoutLogin](c_.ID(), selInit)
	return rv
}
