
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [doubleValue] class.
var doubleValueClass _doubleValueClass

func init() {
	doubleValueClass = _doubleValueClass{objc.GetClass("doubleValue")}
}

type _doubleValueClass struct {
	objc.Class
}

// An interface definition for the [doubleValue] class.
type IdoubleValue interface {
	ID() objc.ID
}

type doubleValue struct {
	id objc.ID
}

func doubleValueFrom(ptr unsafe.Pointer) doubleValue {
	return doubleValue{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ doubleValue) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _doubleValueClass) Alloc() doubleValue {
	rv := objc.Send[doubleValue](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _doubleValueClass) New() doubleValue {
	rv := objc.Send[doubleValue](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdoubleValue creates and returns a new initialized instance.
func NewdoubleValue() doubleValue {
	return doubleValueClass.New()
}

// Init initializes the instance.
func (d_ doubleValue) Init() doubleValue {
	rv := objc.Send[doubleValue](d_.ID(), selInit)
	return rv
}
