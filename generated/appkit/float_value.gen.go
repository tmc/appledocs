
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [floatValue] class.
var floatValueClass _floatValueClass

func init() {
	floatValueClass = _floatValueClass{objc.GetClass("floatValue")}
}

type _floatValueClass struct {
	objc.Class
}

// An interface definition for the [floatValue] class.
type IfloatValue interface {
	ID() objc.ID
}

type floatValue struct {
	id objc.ID
}

func floatValueFrom(ptr unsafe.Pointer) floatValue {
	return floatValue{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ floatValue) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _floatValueClass) Alloc() floatValue {
	rv := objc.Send[floatValue](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _floatValueClass) New() floatValue {
	rv := objc.Send[floatValue](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewfloatValue creates and returns a new initialized instance.
func NewfloatValue() floatValue {
	return floatValueClass.New()
}

// Init initializes the instance.
func (f_ floatValue) Init() floatValue {
	rv := objc.Send[floatValue](f_.ID(), selInit)
	return rv
}
