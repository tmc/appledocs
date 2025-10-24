// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNNGradientKernel] class.
var (
	CNNGradientKernelClass     _CNNGradientKernelClass
	CNNGradientKernelClassOnce sync.Once
)

func getCNNGradientKernelClass() _CNNGradientKernelClass {
	CNNGradientKernelClassOnce.Do(func() {
		CNNGradientKernelClass = _CNNGradientKernelClass{objc.GetClass("MPSCNNGradientKernel")}
	})
	return CNNGradientKernelClass
}

type _CNNGradientKernelClass struct {
	class objc.Class
}

// An interface definition for the [CNNGradientKernel] class.
type ICNNGradientKernel interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other MetalPerformanceShaders classes.


// A parent class referenced by other MetalPerformanceShaders classes. [Full Topic]
type CNNGradientKernel struct {
	objectivec.Object
}

// CNNGradientKernelFrom constructs a [CNNGradientKernel] from an unsafe.Pointer.
//
// A parent class referenced by other MetalPerformanceShaders classes.
func CNNGradientKernelFrom(ptr unsafe.Pointer) CNNGradientKernel {
	return CNNGradientKernel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNGradientKernelClass) Alloc() CNNGradientKernel {
	rv := objc.Send[CNNGradientKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNGradientKernelClass) New() CNNGradientKernel {
	rv := objc.Send[CNNGradientKernel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNGradientKernel) Init() CNNGradientKernel {
	rv := objc.Send[CNNGradientKernel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNGradientKernel) Autorelease() CNNGradientKernel {
	rv := objc.Send[CNNGradientKernel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNGradientKernel creates a new CNNGradientKernel instance.
func NewCNNGradientKernel() CNNGradientKernel {
	return getCNNGradientKernelClass().New()
}




