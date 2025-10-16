
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isNavigational] class.
var isNavigationalClass _isNavigationalClass

func init() {
	isNavigationalClass = _isNavigationalClass{objc.GetClass("isNavigational")}
}

type _isNavigationalClass struct {
	objc.Class
}

// An interface definition for the [isNavigational] class.
type IisNavigational interface {
	ID() objc.ID
}

type isNavigational struct {
	id objc.ID
}

func isNavigationalFrom(ptr unsafe.Pointer) isNavigational {
	return isNavigational{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isNavigational) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isNavigationalClass) Alloc() isNavigational {
	rv := objc.Send[isNavigational](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isNavigationalClass) New() isNavigational {
	rv := objc.Send[isNavigational](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisNavigational creates and returns a new initialized instance.
func NewisNavigational() isNavigational {
	return isNavigationalClass.New()
}

// Init initializes the instance.
func (i_ isNavigational) Init() isNavigational {
	rv := objc.Send[isNavigational](i_.ID(), selInit)
	return rv
}
