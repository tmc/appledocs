
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isBezeled] class.
var isBezeledClass _isBezeledClass

func init() {
	isBezeledClass = _isBezeledClass{objc.GetClass("isBezeled")}
}

type _isBezeledClass struct {
	objc.Class
}

// An interface definition for the [isBezeled] class.
type IisBezeled interface {
	ID() objc.ID
}

type isBezeled struct {
	id objc.ID
}

func isBezeledFrom(ptr unsafe.Pointer) isBezeled {
	return isBezeled{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isBezeled) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isBezeledClass) Alloc() isBezeled {
	rv := objc.Send[isBezeled](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isBezeledClass) New() isBezeled {
	rv := objc.Send[isBezeled](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisBezeled creates and returns a new initialized instance.
func NewisBezeled() isBezeled {
	return isBezeledClass.New()
}

// Init initializes the instance.
func (i_ isBezeled) Init() isBezeled {
	rv := objc.Send[isBezeled](i_.ID(), selInit)
	return rv
}
