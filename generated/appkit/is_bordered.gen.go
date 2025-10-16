
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isBordered] class.
var isBorderedClass _isBorderedClass

func init() {
	isBorderedClass = _isBorderedClass{objc.GetClass("isBordered")}
}

type _isBorderedClass struct {
	objc.Class
}

// An interface definition for the [isBordered] class.
type IisBordered interface {
	ID() objc.ID
}

type isBordered struct {
	id objc.ID
}

func isBorderedFrom(ptr unsafe.Pointer) isBordered {
	return isBordered{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isBordered) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isBorderedClass) Alloc() isBordered {
	rv := objc.Send[isBordered](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isBorderedClass) New() isBordered {
	rv := objc.Send[isBordered](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisBordered creates and returns a new initialized instance.
func NewisBordered() isBordered {
	return isBorderedClass.New()
}

// Init initializes the instance.
func (i_ isBordered) Init() isBordered {
	rv := objc.Send[isBordered](i_.ID(), selInit)
	return rv
}
