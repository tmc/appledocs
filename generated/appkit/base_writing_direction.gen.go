
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [baseWritingDirection] class.
var baseWritingDirectionClass _baseWritingDirectionClass

func init() {
	baseWritingDirectionClass = _baseWritingDirectionClass{objc.GetClass("baseWritingDirection")}
}

type _baseWritingDirectionClass struct {
	objc.Class
}

// An interface definition for the [baseWritingDirection] class.
type IbaseWritingDirection interface {
	ID() objc.ID
}

type baseWritingDirection struct {
	id objc.ID
}

func baseWritingDirectionFrom(ptr unsafe.Pointer) baseWritingDirection {
	return baseWritingDirection{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ baseWritingDirection) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _baseWritingDirectionClass) Alloc() baseWritingDirection {
	rv := objc.Send[baseWritingDirection](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _baseWritingDirectionClass) New() baseWritingDirection {
	rv := objc.Send[baseWritingDirection](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbaseWritingDirection creates and returns a new initialized instance.
func NewbaseWritingDirection() baseWritingDirection {
	return baseWritingDirectionClass.New()
}

// Init initializes the instance.
func (b_ baseWritingDirection) Init() baseWritingDirection {
	rv := objc.Send[baseWritingDirection](b_.ID(), selInit)
	return rv
}
