
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PathControl] class.
var PathControlClass _PathControlClass

func init() {
	PathControlClass = _PathControlClass{objc.GetClass("NSPathControl")}
}

type _PathControlClass struct {
	objc.Class
}

// An interface definition for the [PathControl] class.
type IPathControl interface {
	ID() objc.ID
}

type PathControl struct {
	id objc.ID
}

func PathControlFrom(ptr unsafe.Pointer) PathControl {
	return PathControl{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PathControl) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PathControlClass) Alloc() PathControl {
	rv := objc.Send[PathControl](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PathControlClass) New() PathControl {
	rv := objc.Send[PathControl](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPathControl creates and returns a new initialized instance.
func NewPathControl() PathControl {
	return PathControlClass.New()
}

// Init initializes the instance.
func (p_ PathControl) Init() PathControl {
	rv := objc.Send[PathControl](p_.ID(), selInit)
	return rv
}
