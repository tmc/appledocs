// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	PrimaryDilationRates() NDArraySizes /* not a class type */
	SetPrimaryDilationRates(value NDArraySizes /* not a class type */)
	PrimaryEdgeMode() ImageEdgeMode
	SetPrimaryEdgeMode(value ImageEdgeMode)
	PrimaryKernelSizes() NDArraySizes /* not a class type */
	SetPrimaryKernelSizes(value NDArraySizes /* not a class type */)
	PrimaryOffsets() NDArrayOffsets /* not a class type */
	SetPrimaryOffsets(value NDArrayOffsets /* not a class type */)
	PrimaryStrides() NDArrayOffsets /* not a class type */
	SetPrimaryStrides(value NDArrayOffsets /* not a class type */)
	SecondaryDilationRates() NDArraySizes /* not a class type */
	SetSecondaryDilationRates(value NDArraySizes /* not a class type */)
	SecondaryEdgeMode() ImageEdgeMode
	SetSecondaryEdgeMode(value ImageEdgeMode)
	SecondaryKernelSizes() NDArraySizes /* not a class type */
	SetSecondaryKernelSizes(value NDArraySizes /* not a class type */)
	SecondaryOffsets() NDArrayOffsets /* not a class type */
	SetSecondaryOffsets(value NDArrayOffsets /* not a class type */)
	SecondaryStrides() NDArrayOffsets /* not a class type */
	SetSecondaryStrides(value NDArrayOffsets /* not a class type */)
	// methods:
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArray(cmdBuf objectivec.IObject, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray) INDArray
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel/init(device:)
func NewNDArrayBinaryKernelWithDevice(device objectivec.IObject) NDArrayBinaryKernel {
	instance := getNDArrayBinaryKernelClass().Alloc()
	rv := objc.Send[NDArrayBinaryKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel/encode(to:primarySourceArray:secondarySourceArray:)
func (n_ NDArrayBinaryKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArray(cmdBuf objectivec.IObject, primarySourceArray IMPSNDArray, secondarySourceArray IMPSNDArray) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:"), cmdBuf, primarySourceArray, secondarySourceArray)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primarydilationrates
func (n_ NDArrayBinaryKernel) PrimaryDilationRates() NDArraySizes /* not a class type */ {
	rv := objc.Send[NDArraySizes](n_.ID, objc.Sel("primaryDilationRates"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primarydilationrates
func (n_ NDArrayBinaryKernel) SetPrimaryDilationRates(value NDArraySizes /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryDilationRates:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primaryedgemode
func (n_ NDArrayBinaryKernel) PrimaryEdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](n_.ID, objc.Sel("primaryEdgeMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primaryedgemode
func (n_ NDArrayBinaryKernel) SetPrimaryEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryEdgeMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primarykernelsizes
func (n_ NDArrayBinaryKernel) PrimaryKernelSizes() NDArraySizes /* not a class type */ {
	rv := objc.Send[NDArraySizes](n_.ID, objc.Sel("primaryKernelSizes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primarykernelsizes
func (n_ NDArrayBinaryKernel) SetPrimaryKernelSizes(value NDArraySizes /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryKernelSizes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primaryoffsets
func (n_ NDArrayBinaryKernel) PrimaryOffsets() NDArrayOffsets /* not a class type */ {
	rv := objc.Send[NDArrayOffsets](n_.ID, objc.Sel("primaryOffsets"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primaryoffsets
func (n_ NDArrayBinaryKernel) SetPrimaryOffsets(value NDArrayOffsets /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryOffsets:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primarystrides
func (n_ NDArrayBinaryKernel) PrimaryStrides() NDArrayOffsets /* not a class type */ {
	rv := objc.Send[NDArrayOffsets](n_.ID, objc.Sel("primaryStrides"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/primarystrides
func (n_ NDArrayBinaryKernel) SetPrimaryStrides(value NDArrayOffsets /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryStrides:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondarydilationrates
func (n_ NDArrayBinaryKernel) SecondaryDilationRates() NDArraySizes /* not a class type */ {
	rv := objc.Send[NDArraySizes](n_.ID, objc.Sel("secondaryDilationRates"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondarydilationrates
func (n_ NDArrayBinaryKernel) SetSecondaryDilationRates(value NDArraySizes /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryDilationRates:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondaryedgemode
func (n_ NDArrayBinaryKernel) SecondaryEdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](n_.ID, objc.Sel("secondaryEdgeMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondaryedgemode
func (n_ NDArrayBinaryKernel) SetSecondaryEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryEdgeMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondarykernelsizes
func (n_ NDArrayBinaryKernel) SecondaryKernelSizes() NDArraySizes /* not a class type */ {
	rv := objc.Send[NDArraySizes](n_.ID, objc.Sel("secondaryKernelSizes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondarykernelsizes
func (n_ NDArrayBinaryKernel) SetSecondaryKernelSizes(value NDArraySizes /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryKernelSizes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondaryoffsets
func (n_ NDArrayBinaryKernel) SecondaryOffsets() NDArrayOffsets /* not a class type */ {
	rv := objc.Send[NDArrayOffsets](n_.ID, objc.Sel("secondaryOffsets"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondaryoffsets
func (n_ NDArrayBinaryKernel) SetSecondaryOffsets(value NDArrayOffsets /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryOffsets:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondarystrides
func (n_ NDArrayBinaryKernel) SecondaryStrides() NDArrayOffsets /* not a class type */ {
	rv := objc.Send[NDArrayOffsets](n_.ID, objc.Sel("secondaryStrides"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/secondarystrides
func (n_ NDArrayBinaryKernel) SetSecondaryStrides(value NDArrayOffsets /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryStrides:"), value)
}


