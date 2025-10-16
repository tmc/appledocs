
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isOnActiveSpace] class.
var isOnActiveSpaceClass _isOnActiveSpaceClass

func init() {
	isOnActiveSpaceClass = _isOnActiveSpaceClass{objc.GetClass("isOnActiveSpace")}
}

type _isOnActiveSpaceClass struct {
	objc.Class
}

// An interface definition for the [isOnActiveSpace] class.
type IisOnActiveSpace interface {
	ID() objc.ID
}

type isOnActiveSpace struct {
	id objc.ID
}

func isOnActiveSpaceFrom(ptr unsafe.Pointer) isOnActiveSpace {
	return isOnActiveSpace{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isOnActiveSpace) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isOnActiveSpaceClass) Alloc() isOnActiveSpace {
	rv := objc.Send[isOnActiveSpace](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isOnActiveSpaceClass) New() isOnActiveSpace {
	rv := objc.Send[isOnActiveSpace](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisOnActiveSpace creates and returns a new initialized instance.
func NewisOnActiveSpace() isOnActiveSpace {
	return isOnActiveSpaceClass.New()
}

// Init initializes the instance.
func (i_ isOnActiveSpace) Init() isOnActiveSpace {
	rv := objc.Send[isOnActiveSpace](i_.ID(), selInit)
	return rv
}
