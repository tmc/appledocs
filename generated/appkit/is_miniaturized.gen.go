
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isMiniaturized] class.
var isMiniaturizedClass _isMiniaturizedClass

func init() {
	isMiniaturizedClass = _isMiniaturizedClass{objc.GetClass("isMiniaturized")}
}

type _isMiniaturizedClass struct {
	objc.Class
}

// An interface definition for the [isMiniaturized] class.
type IisMiniaturized interface {
	ID() objc.ID
}

type isMiniaturized struct {
	id objc.ID
}

func isMiniaturizedFrom(ptr unsafe.Pointer) isMiniaturized {
	return isMiniaturized{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isMiniaturized) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isMiniaturizedClass) Alloc() isMiniaturized {
	rv := objc.Send[isMiniaturized](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isMiniaturizedClass) New() isMiniaturized {
	rv := objc.Send[isMiniaturized](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisMiniaturized creates and returns a new initialized instance.
func NewisMiniaturized() isMiniaturized {
	return isMiniaturizedClass.New()
}

// Init initializes the instance.
func (i_ isMiniaturized) Init() isMiniaturized {
	rv := objc.Send[isMiniaturized](i_.ID(), selInit)
	return rv
}
