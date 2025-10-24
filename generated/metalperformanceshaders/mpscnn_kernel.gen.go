// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNNKernel] class.
var (
	CNNKernelClass     _CNNKernelClass
	CNNKernelClassOnce sync.Once
)

func getCNNKernelClass() _CNNKernelClass {
	CNNKernelClassOnce.Do(func() {
		CNNKernelClass = _CNNKernelClass{objc.GetClass("MPSCNNKernel")}
	})
	return CNNKernelClass
}

type _CNNKernelClass struct {
	class objc.Class
}

// An interface definition for the [CNNKernel] class.
type ICNNKernel interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other MetalPerformanceShaders classes.


// A parent class referenced by other MetalPerformanceShaders classes. [Full Topic]
type CNNKernel struct {
	objectivec.Object
}

// CNNKernelFrom constructs a [CNNKernel] from an unsafe.Pointer.
//
// A parent class referenced by other MetalPerformanceShaders classes.
func CNNKernelFrom(ptr unsafe.Pointer) CNNKernel {
	return CNNKernel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNKernelClass) Alloc() CNNKernel {
	rv := objc.Send[CNNKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNKernelClass) New() CNNKernel {
	rv := objc.Send[CNNKernel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNKernel) Init() CNNKernel {
	rv := objc.Send[CNNKernel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNKernel) Autorelease() CNNKernel {
	rv := objc.Send[CNNKernel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNKernel creates a new CNNKernel instance.
func NewCNNKernel() CNNKernel {
	return getCNNKernelClass().New()
}




