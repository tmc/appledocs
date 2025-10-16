
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [maxValue] class.
var maxValueClass _maxValueClass

func init() {
	maxValueClass = _maxValueClass{objc.GetClass("maxValue")}
}

type _maxValueClass struct {
	objc.Class
}

// An interface definition for the [maxValue] class.
type ImaxValue interface {
	ID() objc.ID
}

type maxValue struct {
	id objc.ID
}

func maxValueFrom(ptr unsafe.Pointer) maxValue {
	return maxValue{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ maxValue) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _maxValueClass) Alloc() maxValue {
	rv := objc.Send[maxValue](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _maxValueClass) New() maxValue {
	rv := objc.Send[maxValue](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmaxValue creates and returns a new initialized instance.
func NewmaxValue() maxValue {
	return maxValueClass.New()
}

// Init initializes the instance.
func (m_ maxValue) Init() maxValue {
	rv := objc.Send[maxValue](m_.ID(), selInit)
	return rv
}
