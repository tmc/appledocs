// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [TensorExtents] class.
var (
	TensorExtentsClass     _TensorExtentsClass
	TensorExtentsClassOnce sync.Once
)

func getTensorExtentsClass() _TensorExtentsClass {
	TensorExtentsClassOnce.Do(func() {
		TensorExtentsClass = _TensorExtentsClass{objc.GetClass("MTLTensorExtents")}
	})
	return TensorExtentsClass
}

type _TensorExtentsClass struct {
	class objc.Class
}

// An interface definition for the [TensorExtents] class.
type ITensorExtents interface {
	objectivec.IObject
}

// An array of length matching the rank, holding the dimensions of a tensor.
//
// Supports rank up to .
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorExtents
type TensorExtents struct {
	objectivec.Object
}

// TensorExtentsFrom constructs a [TensorExtents] from an unsafe.Pointer.
//
// An array of length matching the rank, holding the dimensions of a tensor.
func TensorExtentsFrom(ptr unsafe.Pointer) TensorExtents {
	return TensorExtents{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TensorExtentsClass) Alloc() TensorExtents {
	rv := objc.Send[TensorExtents](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TensorExtentsClass) New() TensorExtents {
	rv := objc.Send[TensorExtents](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TensorExtents) Init() TensorExtents {
	rv := objc.Send[TensorExtents](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TensorExtents) Autorelease() TensorExtents {
	rv := objc.Send[TensorExtents](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTensorExtents creates a new TensorExtents instance.
func NewTensorExtents() TensorExtents {
	return getTensorExtentsClass().New()
}




