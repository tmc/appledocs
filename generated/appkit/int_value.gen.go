
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [intValue] class.
var intValueClass _intValueClass

func init() {
	intValueClass = _intValueClass{objc.GetClass("intValue")}
}

type _intValueClass struct {
	objc.Class
}

// An interface definition for the [intValue] class.
type IintValue interface {
	ID() objc.ID
}

type intValue struct {
	id objc.ID
}

func intValueFrom(ptr unsafe.Pointer) intValue {
	return intValue{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ intValue) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _intValueClass) Alloc() intValue {
	rv := objc.Send[intValue](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _intValueClass) New() intValue {
	rv := objc.Send[intValue](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewintValue creates and returns a new initialized instance.
func NewintValue() intValue {
	return intValueClass.New()
}

// Init initializes the instance.
func (i_ intValue) Init() intValue {
	rv := objc.Send[intValue](i_.ID(), selInit)
	return rv
}
