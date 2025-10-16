
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [integerValue] class.
var integerValueClass _integerValueClass

func init() {
	integerValueClass = _integerValueClass{objc.GetClass("integerValue")}
}

type _integerValueClass struct {
	objc.Class
}

// An interface definition for the [integerValue] class.
type IintegerValue interface {
	ID() objc.ID
}

type integerValue struct {
	id objc.ID
}

func integerValueFrom(ptr unsafe.Pointer) integerValue {
	return integerValue{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ integerValue) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _integerValueClass) Alloc() integerValue {
	rv := objc.Send[integerValue](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _integerValueClass) New() integerValue {
	rv := objc.Send[integerValue](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewintegerValue creates and returns a new initialized instance.
func NewintegerValue() integerValue {
	return integerValueClass.New()
}

// Init initializes the instance.
func (i_ integerValue) Init() integerValue {
	rv := objc.Send[integerValue](i_.ID(), selInit)
	return rv
}
