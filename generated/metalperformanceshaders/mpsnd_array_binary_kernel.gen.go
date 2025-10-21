// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NDArrayBinaryKernel] class.
var (
	NDArrayBinaryKernelClass     _NDArrayBinaryKernelClass
	NDArrayBinaryKernelClassOnce sync.Once
)

func getNDArrayBinaryKernelClass() _NDArrayBinaryKernelClass {
	NDArrayBinaryKernelClassOnce.Do(func() {
		NDArrayBinaryKernelClass = _NDArrayBinaryKernelClass{objc.GetClass("MPSNDArrayBinaryKernel")}
	})
	return NDArrayBinaryKernelClass
}

type _NDArrayBinaryKernelClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayBinaryKernel] class.
type INDArrayBinaryKernel interface {
	INDArrayMultiaryKernel
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArray(cmdBuf objc.ID, primarySourceArray unsafe.Pointer, secondarySourceArray unsafe.Pointer) unsafe.Pointer
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel
type NDArrayBinaryKernel struct {
	NDArrayMultiaryKernel
}

// NDArrayBinaryKernelFrom constructs a [NDArrayBinaryKernel] from an unsafe.Pointer.
func NDArrayBinaryKernelFrom(ptr unsafe.Pointer) NDArrayBinaryKernel {
	return NDArrayBinaryKernel{
		NDArrayMultiaryKernel: NDArrayMultiaryKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayBinaryKernelClass) Alloc() NDArrayBinaryKernel {
	rv := objc.Send[NDArrayBinaryKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayBinaryKernelClass) New() NDArrayBinaryKernel {
	rv := objc.Send[NDArrayBinaryKernel](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayBinaryKernel) Init() NDArrayBinaryKernel {
	rv := objc.Send[NDArrayBinaryKernel](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayBinaryKernel) Autorelease() NDArrayBinaryKernel {
	rv := objc.Send[NDArrayBinaryKernel](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayBinaryKernel creates a new NDArrayBinaryKernel instance.
func NewNDArrayBinaryKernel() NDArrayBinaryKernel {
	return getNDArrayBinaryKernelClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel/init(device:)
func NewNDArrayBinaryKernelWithDevice(device objc.ID) NDArrayBinaryKernel {
	instance := getNDArrayBinaryKernelClass().Alloc()
	rv := objc.Send[NDArrayBinaryKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel/encode(to:primarySourceArray:secondarySourceArray:)
func (n_ NDArrayBinaryKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArray(cmdBuf objc.ID, primarySourceArray unsafe.Pointer, secondarySourceArray unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:"), cmdBuf, primarySourceArray, secondarySourceArray)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primarydilationrates
func (n_ NDArrayBinaryKernel) PrimaryDilationRates() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("primaryDilationRates"))
	return rv
}


// SetPrimaryDilationRates sets the value of the primaryDilationRates property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primarydilationrates
func (n_ NDArrayBinaryKernel) SetPrimaryDilationRates(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryDilationRates:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primaryedgemode
func (n_ NDArrayBinaryKernel) PrimaryEdgeMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("primaryEdgeMode"))
	return rv
}


// SetPrimaryEdgeMode sets the value of the primaryEdgeMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primaryedgemode
func (n_ NDArrayBinaryKernel) SetPrimaryEdgeMode(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryEdgeMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primarykernelsizes
func (n_ NDArrayBinaryKernel) PrimaryKernelSizes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("primaryKernelSizes"))
	return rv
}


// SetPrimaryKernelSizes sets the value of the primaryKernelSizes property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primarykernelsizes
func (n_ NDArrayBinaryKernel) SetPrimaryKernelSizes(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryKernelSizes:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primaryoffsets
func (n_ NDArrayBinaryKernel) PrimaryOffsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("primaryOffsets"))
	return rv
}


// SetPrimaryOffsets sets the value of the primaryOffsets property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primaryoffsets
func (n_ NDArrayBinaryKernel) SetPrimaryOffsets(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryOffsets:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primarystrides
func (n_ NDArrayBinaryKernel) PrimaryStrides() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("primaryStrides"))
	return rv
}


// SetPrimaryStrides sets the value of the primaryStrides property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primarystrides
func (n_ NDArrayBinaryKernel) SetPrimaryStrides(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryStrides:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondarydilationrates
func (n_ NDArrayBinaryKernel) SecondaryDilationRates() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("secondaryDilationRates"))
	return rv
}


// SetSecondaryDilationRates sets the value of the secondaryDilationRates property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondarydilationrates
func (n_ NDArrayBinaryKernel) SetSecondaryDilationRates(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryDilationRates:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondaryedgemode
func (n_ NDArrayBinaryKernel) SecondaryEdgeMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("secondaryEdgeMode"))
	return rv
}


// SetSecondaryEdgeMode sets the value of the secondaryEdgeMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondaryedgemode
func (n_ NDArrayBinaryKernel) SetSecondaryEdgeMode(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryEdgeMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondarykernelsizes
func (n_ NDArrayBinaryKernel) SecondaryKernelSizes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("secondaryKernelSizes"))
	return rv
}


// SetSecondaryKernelSizes sets the value of the secondaryKernelSizes property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondarykernelsizes
func (n_ NDArrayBinaryKernel) SetSecondaryKernelSizes(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryKernelSizes:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondaryoffsets
func (n_ NDArrayBinaryKernel) SecondaryOffsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("secondaryOffsets"))
	return rv
}


// SetSecondaryOffsets sets the value of the secondaryOffsets property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondaryoffsets
func (n_ NDArrayBinaryKernel) SetSecondaryOffsets(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryOffsets:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondarystrides
func (n_ NDArrayBinaryKernel) SecondaryStrides() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("secondaryStrides"))
	return rv
}


// SetSecondaryStrides sets the value of the secondaryStrides property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondarystrides
func (n_ NDArrayBinaryKernel) SetSecondaryStrides(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryStrides:"), value)
}


