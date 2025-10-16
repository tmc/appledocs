
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [backingType] class.
var backingTypeClass _backingTypeClass

func init() {
	backingTypeClass = _backingTypeClass{objc.GetClass("backingType")}
}

type _backingTypeClass struct {
	objc.Class
}

// An interface definition for the [backingType] class.
type IbackingType interface {
	ID() objc.ID
}

type backingType struct {
	id objc.ID
}

func backingTypeFrom(ptr unsafe.Pointer) backingType {
	return backingType{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ backingType) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _backingTypeClass) Alloc() backingType {
	rv := objc.Send[backingType](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _backingTypeClass) New() backingType {
	rv := objc.Send[backingType](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbackingType creates and returns a new initialized instance.
func NewbackingType() backingType {
	return backingTypeClass.New()
}

// Init initializes the instance.
func (b_ backingType) Init() backingType {
	rv := objc.Send[backingType](b_.ID(), selInit)
	return rv
}
