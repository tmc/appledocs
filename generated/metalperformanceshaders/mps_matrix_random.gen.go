// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MatrixRandom] class.
var (
	MatrixRandomClass     _MatrixRandomClass
	MatrixRandomClassOnce sync.Once
)

func getMatrixRandomClass() _MatrixRandomClass {
	MatrixRandomClassOnce.Do(func() {
		MatrixRandomClass = _MatrixRandomClass{objc.GetClass("MPSMatrixRandom")}
	})
	return MatrixRandomClass
}

type _MatrixRandomClass struct {
	class objc.Class
}

// An interface definition for the [MatrixRandom] class.
type IMatrixRandom interface {
	IKernel
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandom
type MatrixRandom struct {
	Kernel
}

// MatrixRandomFrom constructs a [MatrixRandom] from an unsafe.Pointer.
func MatrixRandomFrom(ptr unsafe.Pointer) MatrixRandom {
	return MatrixRandom{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MatrixRandomClass) Alloc() MatrixRandom {
	rv := objc.Send[MatrixRandom](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MatrixRandomClass) New() MatrixRandom {
	rv := objc.Send[MatrixRandom](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixRandom) Init() MatrixRandom {
	rv := objc.Send[MatrixRandom](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixRandom) Autorelease() MatrixRandom {
	rv := objc.Send[MatrixRandom](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixRandom creates a new MatrixRandom instance.
func NewMatrixRandom() MatrixRandom {
	return getMatrixRandomClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandom/destinationDataType
func (m_ MatrixRandom) DestinationDataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("destinationDataType"))
	return rv
}



