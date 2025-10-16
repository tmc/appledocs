
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isTransparent] class.
var isTransparentClass _isTransparentClass

func init() {
	isTransparentClass = _isTransparentClass{objc.GetClass("isTransparent")}
}

type _isTransparentClass struct {
	objc.Class
}

// An interface definition for the [isTransparent] class.
type IisTransparent interface {
	ID() objc.ID
}

type isTransparent struct {
	id objc.ID
}

func isTransparentFrom(ptr unsafe.Pointer) isTransparent {
	return isTransparent{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isTransparent) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isTransparentClass) Alloc() isTransparent {
	rv := objc.Send[isTransparent](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isTransparentClass) New() isTransparent {
	rv := objc.Send[isTransparent](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisTransparent creates and returns a new initialized instance.
func NewisTransparent() isTransparent {
	return isTransparentClass.New()
}

// Init initializes the instance.
func (i_ isTransparent) Init() isTransparent {
	rv := objc.Send[isTransparent](i_.ID(), selInit)
	return rv
}
