// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Vector] class.
var (
	VectorClass     _VectorClass
	VectorClassOnce sync.Once
)

func getVectorClass() _VectorClass {
	VectorClassOnce.Do(func() {
		VectorClass = _VectorClass{objc.GetClass("VNVector")}
	})
	return VectorClass
}

type _VectorClass struct {
	class objc.Class
}

// An interface definition for the [Vector] class.
type IVector interface {
	objectivec.IObject
}

// An immutable 2D vector represented by its x-axis and y-axis projections.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector
type Vector struct {
	objectivec.Object
}

// VectorFrom constructs a [Vector] from an unsafe.Pointer.
//
// An immutable 2D vector represented by its x-axis and y-axis projections.
func VectorFrom(ptr unsafe.Pointer) Vector {
	return Vector{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VectorClass) Alloc() Vector {
	rv := objc.Send[Vector](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VectorClass) New() Vector {
	rv := objc.Send[Vector](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ Vector) Init() Vector {
	rv := objc.Send[Vector](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ Vector) Autorelease() Vector {
	rv := objc.Send[Vector](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVector creates a new Vector instance.
func NewVector() Vector {
	return getVectorClass().New()
}




