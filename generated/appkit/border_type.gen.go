
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [borderType] class.
var borderTypeClass _borderTypeClass

func init() {
	borderTypeClass = _borderTypeClass{objc.GetClass("borderType")}
}

type _borderTypeClass struct {
	objc.Class
}

// An interface definition for the [borderType] class.
type IborderType interface {
	ID() objc.ID
}

type borderType struct {
	id objc.ID
}

func borderTypeFrom(ptr unsafe.Pointer) borderType {
	return borderType{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ borderType) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _borderTypeClass) Alloc() borderType {
	rv := objc.Send[borderType](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _borderTypeClass) New() borderType {
	rv := objc.Send[borderType](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewborderType creates and returns a new initialized instance.
func NewborderType() borderType {
	return borderTypeClass.New()
}

// Init initializes the instance.
func (b_ borderType) Init() borderType {
	rv := objc.Send[borderType](b_.ID(), selInit)
	return rv
}
