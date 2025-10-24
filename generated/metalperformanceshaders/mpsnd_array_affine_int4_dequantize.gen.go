// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NDArrayAffineInt4Dequantize] class.
var (
	NDArrayAffineInt4DequantizeClass     _NDArrayAffineInt4DequantizeClass
	NDArrayAffineInt4DequantizeClassOnce sync.Once
)

func getNDArrayAffineInt4DequantizeClass() _NDArrayAffineInt4DequantizeClass {
	NDArrayAffineInt4DequantizeClassOnce.Do(func() {
		NDArrayAffineInt4DequantizeClass = _NDArrayAffineInt4DequantizeClass{objc.GetClass("MPSNDArrayAffineInt4Dequantize")}
	})
	return NDArrayAffineInt4DequantizeClass
}

type _NDArrayAffineInt4DequantizeClass struct {
	class objc.Class
}





// An interface definition for the [NDArrayAffineInt4Dequantize] class.
type INDArrayAffineInt4Dequantize interface {
	INDArrayMultiaryKernel
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NDArrayAffineInt4DequantizeClass) Alloc() NDArrayAffineInt4Dequantize {
	rv := objc.Send[NDArrayAffineInt4Dequantize](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayAffineInt4DequantizeClass) New() NDArrayAffineInt4Dequantize {
	rv := objc.Send[NDArrayAffineInt4Dequantize](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayAffineInt4Dequantize) Init() NDArrayAffineInt4Dequantize {
	rv := objc.Send[NDArrayAffineInt4Dequantize](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayAffineInt4Dequantize) Autorelease() NDArrayAffineInt4Dequantize {
	rv := objc.Send[NDArrayAffineInt4Dequantize](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayAffineInt4Dequantize creates a new NDArrayAffineInt4Dequantize instance.
func NewNDArrayAffineInt4Dequantize() NDArrayAffineInt4Dequantize {
	return getNDArrayAffineInt4DequantizeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineInt4Dequantize
type NDArrayAffineInt4Dequantize struct {
	NDArrayMultiaryKernel
}

// NDArrayAffineInt4DequantizeFrom constructs a [NDArrayAffineInt4Dequantize] from an unsafe.Pointer.
func NDArrayAffineInt4DequantizeFrom(ptr unsafe.Pointer) NDArrayAffineInt4Dequantize {
	return NDArrayAffineInt4Dequantize{
		NDArrayMultiaryKernel: NDArrayMultiaryKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayaffineint4dequantize/4446149-initwithdevice
func NewNDArrayAffineInt4DequantizeWithDeviceQuantizationDescriptor(device unsafe.Pointer, quantizationDescriptor INDArrayAffineQuantizationDescriptor) NDArrayAffineInt4Dequantize {
	instance := getNDArrayAffineInt4DequantizeClass().Alloc()
	rv := objc.Send[NDArrayAffineInt4Dequantize](instance.ID, objc.Sel("initWithDevice:quantizationDescriptor:"), device, quantizationDescriptor)
	rv.Autorelease()
	return rv
}



























