
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isMiniaturizable] class.
var isMiniaturizableClass _isMiniaturizableClass

func init() {
	isMiniaturizableClass = _isMiniaturizableClass{objc.GetClass("isMiniaturizable")}
}

type _isMiniaturizableClass struct {
	objc.Class
}

// An interface definition for the [isMiniaturizable] class.
type IisMiniaturizable interface {
	ID() objc.ID
}

type isMiniaturizable struct {
	id objc.ID
}

func isMiniaturizableFrom(ptr unsafe.Pointer) isMiniaturizable {
	return isMiniaturizable{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isMiniaturizable) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isMiniaturizableClass) Alloc() isMiniaturizable {
	rv := objc.Send[isMiniaturizable](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isMiniaturizableClass) New() isMiniaturizable {
	rv := objc.Send[isMiniaturizable](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisMiniaturizable creates and returns a new initialized instance.
func NewisMiniaturizable() isMiniaturizable {
	return isMiniaturizableClass.New()
}

// Init initializes the instance.
func (i_ isMiniaturizable) Init() isMiniaturizable {
	rv := objc.Send[isMiniaturizable](i_.ID(), selInit)
	return rv
}
