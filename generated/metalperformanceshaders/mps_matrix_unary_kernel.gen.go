// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MatrixUnaryKernel] class.
var (
	MatrixUnaryKernelClass     _MatrixUnaryKernelClass
	MatrixUnaryKernelClassOnce sync.Once
)

func getMatrixUnaryKernelClass() _MatrixUnaryKernelClass {
	MatrixUnaryKernelClassOnce.Do(func() {
		MatrixUnaryKernelClass = _MatrixUnaryKernelClass{objc.GetClass("MPSMatrixUnaryKernel")}
	})
	return MatrixUnaryKernelClass
}

type _MatrixUnaryKernelClass struct {
	class objc.Class
}

// An interface definition for the [MatrixUnaryKernel] class.
type IMatrixUnaryKernel interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other MetalPerformanceShaders classes.


// A parent class referenced by other MetalPerformanceShaders classes. [Full Topic]
type MatrixUnaryKernel struct {
	objectivec.Object
}

// MatrixUnaryKernelFrom constructs a [MatrixUnaryKernel] from an unsafe.Pointer.
//
// A parent class referenced by other MetalPerformanceShaders classes.
func MatrixUnaryKernelFrom(ptr unsafe.Pointer) MatrixUnaryKernel {
	return MatrixUnaryKernel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MatrixUnaryKernelClass) Alloc() MatrixUnaryKernel {
	rv := objc.Send[MatrixUnaryKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MatrixUnaryKernelClass) New() MatrixUnaryKernel {
	rv := objc.Send[MatrixUnaryKernel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixUnaryKernel) Init() MatrixUnaryKernel {
	rv := objc.Send[MatrixUnaryKernel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixUnaryKernel) Autorelease() MatrixUnaryKernel {
	rv := objc.Send[MatrixUnaryKernel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixUnaryKernel creates a new MatrixUnaryKernel instance.
func NewMatrixUnaryKernel() MatrixUnaryKernel {
	return getMatrixUnaryKernelClass().New()
}




