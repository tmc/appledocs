
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [maximumPotentialExtendedDynamicRangeColorComponentValue] class.
var maximumPotentialExtendedDynamicRangeColorComponentValueClass _maximumPotentialExtendedDynamicRangeColorComponentValueClass

func init() {
	maximumPotentialExtendedDynamicRangeColorComponentValueClass = _maximumPotentialExtendedDynamicRangeColorComponentValueClass{objc.GetClass("maximumPotentialExtendedDynamicRangeColorComponentValue")}
}

type _maximumPotentialExtendedDynamicRangeColorComponentValueClass struct {
	objc.Class
}

// An interface definition for the [maximumPotentialExtendedDynamicRangeColorComponentValue] class.
type ImaximumPotentialExtendedDynamicRangeColorComponentValue interface {
	ID() objc.ID
}

type maximumPotentialExtendedDynamicRangeColorComponentValue struct {
	id objc.ID
}

func maximumPotentialExtendedDynamicRangeColorComponentValueFrom(ptr unsafe.Pointer) maximumPotentialExtendedDynamicRangeColorComponentValue {
	return maximumPotentialExtendedDynamicRangeColorComponentValue{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ maximumPotentialExtendedDynamicRangeColorComponentValue) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _maximumPotentialExtendedDynamicRangeColorComponentValueClass) Alloc() maximumPotentialExtendedDynamicRangeColorComponentValue {
	rv := objc.Send[maximumPotentialExtendedDynamicRangeColorComponentValue](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _maximumPotentialExtendedDynamicRangeColorComponentValueClass) New() maximumPotentialExtendedDynamicRangeColorComponentValue {
	rv := objc.Send[maximumPotentialExtendedDynamicRangeColorComponentValue](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmaximumPotentialExtendedDynamicRangeColorComponentValue creates and returns a new initialized instance.
func NewmaximumPotentialExtendedDynamicRangeColorComponentValue() maximumPotentialExtendedDynamicRangeColorComponentValue {
	return maximumPotentialExtendedDynamicRangeColorComponentValueClass.New()
}

// Init initializes the instance.
func (m_ maximumPotentialExtendedDynamicRangeColorComponentValue) Init() maximumPotentialExtendedDynamicRangeColorComponentValue {
	rv := objc.Send[maximumPotentialExtendedDynamicRangeColorComponentValue](m_.ID(), selInit)
	return rv
}
