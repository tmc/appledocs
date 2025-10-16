
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [alphaValue] class.
var alphaValueClass _alphaValueClass

func init() {
	alphaValueClass = _alphaValueClass{objc.GetClass("alphaValue")}
}

type _alphaValueClass struct {
	objc.Class
}

// An interface definition for the [alphaValue] class.
type IalphaValue interface {
	ID() objc.ID
}

type alphaValue struct {
	id objc.ID
}

func alphaValueFrom(ptr unsafe.Pointer) alphaValue {
	return alphaValue{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ alphaValue) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _alphaValueClass) Alloc() alphaValue {
	rv := objc.Send[alphaValue](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _alphaValueClass) New() alphaValue {
	rv := objc.Send[alphaValue](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewalphaValue creates and returns a new initialized instance.
func NewalphaValue() alphaValue {
	return alphaValueClass.New()
}

// Init initializes the instance.
func (a_ alphaValue) Init() alphaValue {
	rv := objc.Send[alphaValue](a_.ID(), selInit)
	return rv
}
