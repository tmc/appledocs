
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [removeFromSuperview] class.
var removeFromSuperviewClass _removeFromSuperviewClass

func init() {
	removeFromSuperviewClass = _removeFromSuperviewClass{objc.GetClass("removeFromSuperview")}
}

type _removeFromSuperviewClass struct {
	objc.Class
}

// An interface definition for the [removeFromSuperview] class.
type IremoveFromSuperview interface {
	ID() objc.ID
}

type removeFromSuperview struct {
	id objc.ID
}

func removeFromSuperviewFrom(ptr unsafe.Pointer) removeFromSuperview {
	return removeFromSuperview{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ removeFromSuperview) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _removeFromSuperviewClass) Alloc() removeFromSuperview {
	rv := objc.Send[removeFromSuperview](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _removeFromSuperviewClass) New() removeFromSuperview {
	rv := objc.Send[removeFromSuperview](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewremoveFromSuperview creates and returns a new initialized instance.
func NewremoveFromSuperview() removeFromSuperview {
	return removeFromSuperviewClass.New()
}

// Init initializes the instance.
func (r_ removeFromSuperview) Init() removeFromSuperview {
	rv := objc.Send[removeFromSuperview](r_.ID(), selInit)
	return rv
}
