
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [resizeFlags] class.
var resizeFlagsClass _resizeFlagsClass

func init() {
	resizeFlagsClass = _resizeFlagsClass{objc.GetClass("resizeFlags")}
}

type _resizeFlagsClass struct {
	objc.Class
}

// An interface definition for the [resizeFlags] class.
type IresizeFlags interface {
	ID() objc.ID
}

type resizeFlags struct {
	id objc.ID
}

func resizeFlagsFrom(ptr unsafe.Pointer) resizeFlags {
	return resizeFlags{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ resizeFlags) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _resizeFlagsClass) Alloc() resizeFlags {
	rv := objc.Send[resizeFlags](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _resizeFlagsClass) New() resizeFlags {
	rv := objc.Send[resizeFlags](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewresizeFlags creates and returns a new initialized instance.
func NewresizeFlags() resizeFlags {
	return resizeFlagsClass.New()
}

// Init initializes the instance.
func (r_ resizeFlags) Init() resizeFlags {
	rv := objc.Send[resizeFlags](r_.ID(), selInit)
	return rv
}
