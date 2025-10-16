
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [boxType] class.
var boxTypeClass _boxTypeClass

func init() {
	boxTypeClass = _boxTypeClass{objc.GetClass("boxType")}
}

type _boxTypeClass struct {
	objc.Class
}

// An interface definition for the [boxType] class.
type IboxType interface {
	ID() objc.ID
}

type boxType struct {
	id objc.ID
}

func boxTypeFrom(ptr unsafe.Pointer) boxType {
	return boxType{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ boxType) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _boxTypeClass) Alloc() boxType {
	rv := objc.Send[boxType](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _boxTypeClass) New() boxType {
	rv := objc.Send[boxType](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewboxType creates and returns a new initialized instance.
func NewboxType() boxType {
	return boxTypeClass.New()
}

// Init initializes the instance.
func (b_ boxType) Init() boxType {
	rv := objc.Send[boxType](b_.ID(), selInit)
	return rv
}
