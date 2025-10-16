
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [minValue] class.
var minValueClass _minValueClass

func init() {
	minValueClass = _minValueClass{objc.GetClass("minValue")}
}

type _minValueClass struct {
	objc.Class
}

// An interface definition for the [minValue] class.
type IminValue interface {
	ID() objc.ID
}

type minValue struct {
	id objc.ID
}

func minValueFrom(ptr unsafe.Pointer) minValue {
	return minValue{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ minValue) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _minValueClass) Alloc() minValue {
	rv := objc.Send[minValue](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _minValueClass) New() minValue {
	rv := objc.Send[minValue](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewminValue creates and returns a new initialized instance.
func NewminValue() minValue {
	return minValueClass.New()
}

// Init initializes the instance.
func (m_ minValue) Init() minValue {
	rv := objc.Send[minValue](m_.ID(), selInit)
	return rv
}
