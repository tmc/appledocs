
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [removeFromSuperviewWithoutNeedingDisplay] class.
var removeFromSuperviewWithoutNeedingDisplayClass _removeFromSuperviewWithoutNeedingDisplayClass

func init() {
	removeFromSuperviewWithoutNeedingDisplayClass = _removeFromSuperviewWithoutNeedingDisplayClass{objc.GetClass("removeFromSuperviewWithoutNeedingDisplay")}
}

type _removeFromSuperviewWithoutNeedingDisplayClass struct {
	objc.Class
}

// An interface definition for the [removeFromSuperviewWithoutNeedingDisplay] class.
type IremoveFromSuperviewWithoutNeedingDisplay interface {
	ID() objc.ID
}

type removeFromSuperviewWithoutNeedingDisplay struct {
	id objc.ID
}

func removeFromSuperviewWithoutNeedingDisplayFrom(ptr unsafe.Pointer) removeFromSuperviewWithoutNeedingDisplay {
	return removeFromSuperviewWithoutNeedingDisplay{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ removeFromSuperviewWithoutNeedingDisplay) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _removeFromSuperviewWithoutNeedingDisplayClass) Alloc() removeFromSuperviewWithoutNeedingDisplay {
	rv := objc.Send[removeFromSuperviewWithoutNeedingDisplay](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _removeFromSuperviewWithoutNeedingDisplayClass) New() removeFromSuperviewWithoutNeedingDisplay {
	rv := objc.Send[removeFromSuperviewWithoutNeedingDisplay](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewremoveFromSuperviewWithoutNeedingDisplay creates and returns a new initialized instance.
func NewremoveFromSuperviewWithoutNeedingDisplay() removeFromSuperviewWithoutNeedingDisplay {
	return removeFromSuperviewWithoutNeedingDisplayClass.New()
}

// Init initializes the instance.
func (r_ removeFromSuperviewWithoutNeedingDisplay) Init() removeFromSuperviewWithoutNeedingDisplay {
	rv := objc.Send[removeFromSuperviewWithoutNeedingDisplay](r_.ID(), selInit)
	return rv
}
