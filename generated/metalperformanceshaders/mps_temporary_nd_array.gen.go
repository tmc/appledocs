// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TemporaryNDArray] class.
var (
	TemporaryNDArrayClass     _TemporaryNDArrayClass
	TemporaryNDArrayClassOnce sync.Once
)

func getTemporaryNDArrayClass() _TemporaryNDArrayClass {
	TemporaryNDArrayClassOnce.Do(func() {
		TemporaryNDArrayClass = _TemporaryNDArrayClass{objc.GetClass("MPSTemporaryNDArray")}
	})
	return TemporaryNDArrayClass
}

type _TemporaryNDArrayClass struct {
	class objc.Class
}

// An interface definition for the [TemporaryNDArray] class.
type ITemporaryNDArray interface {
	INDArray
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryNDArray
type TemporaryNDArray struct {
	NDArray
}

// TemporaryNDArrayFrom constructs a [TemporaryNDArray] from an unsafe.Pointer.
func TemporaryNDArrayFrom(ptr unsafe.Pointer) TemporaryNDArray {
	return TemporaryNDArray{
		NDArray: NDArrayFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TemporaryNDArrayClass) Alloc() TemporaryNDArray {
	rv := objc.Send[TemporaryNDArray](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TemporaryNDArrayClass) New() TemporaryNDArray {
	rv := objc.Send[TemporaryNDArray](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TemporaryNDArray) Init() TemporaryNDArray {
	rv := objc.Send[TemporaryNDArray](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TemporaryNDArray) Autorelease() TemporaryNDArray {
	rv := objc.Send[TemporaryNDArray](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTemporaryNDArray creates a new TemporaryNDArray instance.
func NewTemporaryNDArray() TemporaryNDArray {
	return getTemporaryNDArrayClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryNDArray/readCount
func (t_ TemporaryNDArray) ReadCount() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("readCount"))
	return rv
}

// SetReadCount sets the value of the readCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryNDArray/readCount
func (t_ TemporaryNDArray) SetReadCount(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setReadCount:"), value)
}
