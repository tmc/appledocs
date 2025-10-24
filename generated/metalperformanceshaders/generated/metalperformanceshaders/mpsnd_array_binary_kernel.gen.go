// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayBinaryKernel */


/* debug [class_header]: Header for MPSNDArrayBinaryKernel */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayBinaryKernel */
// An interface definition for the [NDArrayBinaryKernel] class.
type INDArrayBinaryKernel interface {
	INDArrayMultiaryKernel
	
/* debug [class_interface_properties]: Properties for NDArrayBinaryKernel */
	// properties:
	PrimaryDilationRates() NDArraySizes get /* not a class type */
	SetPrimaryDilationRates(value NDArraySizes get /* not a class type */)
	PrimaryEdgeMode() ImageEdgeMode get /* not a class type */
	SetPrimaryEdgeMode(value ImageEdgeMode get /* not a class type */)
	PrimaryKernelSizes() NDArraySizes get /* not a class type */
	SetPrimaryKernelSizes(value NDArraySizes get /* not a class type */)
	PrimaryOffsets() NDArrayOffsets get /* not a class type */
	SetPrimaryOffsets(value NDArrayOffsets get /* not a class type */)
	PrimaryStrides() NDArrayOffsets get /* not a class type */
	SetPrimaryStrides(value NDArrayOffsets get /* not a class type */)
	SecondaryDilationRates() NDArraySizes get /* not a class type */
	SetSecondaryDilationRates(value NDArraySizes get /* not a class type */)
	SecondaryEdgeMode() ImageEdgeMode get /* not a class type */
	SetSecondaryEdgeMode(value ImageEdgeMode get /* not a class type */)
	SecondaryKernelSizes() NDArraySizes get /* not a class type */
	SetSecondaryKernelSizes(value NDArraySizes get /* not a class type */)
	SecondaryOffsets() NDArrayOffsets get /* not a class type */
	SetSecondaryOffsets(value NDArrayOffsets get /* not a class type */)
	SecondaryStrides() NDArrayOffsets get /* not a class type */
	SetSecondaryStrides(value NDArrayOffsets get /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayBinaryKernel */
	// methods:
	Encode()
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArray(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray) INDArray
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayDestinationArray(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray, destination INDArray)
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateDestinationArray(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray, outGradientState IState, destination INDArray)
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateOutputStateIsTemporary(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray, outGradientState objectivec.IObject, outputStateIsTemporary bool) INDArray
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayBinaryKernel */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayBinaryKernelClass) Alloc() NDArrayBinaryKernel {
	rv := objc.Send[NDArrayBinaryKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayBinaryKernel */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayBinaryKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3175005-initwithcoder
func NewNDArrayBinaryKernelWithCoderDevice(coder Coder /* not a class type */, device unsafe.Pointer) NDArrayBinaryKernel {
	instance := getNDArrayBinaryKernelClass().Alloc()
	rv := objc.Send[NDArrayBinaryKernel](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayBinaryKernelWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143501-initwithdevice
func NewNDArrayBinaryKernelWithDevice(device unsafe.Pointer) NDArrayBinaryKernel {
	instance := getNDArrayBinaryKernelClass().Alloc()
	rv := objc.Send[NDArrayBinaryKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayBinaryKernelWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayBinaryKernel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayBinaryKernel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayBinaryKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143497-encode
func (n_ NDArrayBinaryKernel) Encode() {
	objc.Send[objc.ID](n_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143497-encodetocommandbuffer
func (n_ NDArrayBinaryKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArray(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:"), cmdBuf, primarySourceArray, secondarySourceArray)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferPrimarySourceArraySecondarySourceArray */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143498-encodetocommandbuffer
func (n_ NDArrayBinaryKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayDestinationArray(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:destinationArray:"), cmdBuf, primarySourceArray, secondarySourceArray, destination)
}/* debug [instance_methods/method]: EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayDestinationArray */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143499-encodetocommandbuffer
func (n_ NDArrayBinaryKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateDestinationArray(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray, outGradientState IState, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:resultState:destinationArray:"), cmdBuf, primarySourceArray, secondarySourceArray, outGradientState, destination)
}/* debug [instance_methods/method]: EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateDestinationArray */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143500-encodetocommandbuffer
func (n_ NDArrayBinaryKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateOutputStateIsTemporary(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray, outGradientState objectivec.IObject, outputStateIsTemporary bool) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:resultState:outputStateIsTemporary:"), cmdBuf, primarySourceArray, secondarySourceArray, outGradientState, outputStateIsTemporary)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateOutputStateIsTemporary */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayBinaryKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143502-primarydilationrates
func (n_ NDArrayBinaryKernel) PrimaryDilationRates() NDArraySizes get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("primaryDilationRates"))
	return rv
}/* debug [instance_properties/getter]: primaryDilationRates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143502-primarydilationrates
func (n_ NDArrayBinaryKernel) SetPrimaryDilationRates(value NDArraySizes get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryDilationRates:"), value)
}/* debug [instance_properties/setter]: primaryDilationRates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143503-primaryedgemode
func (n_ NDArrayBinaryKernel) PrimaryEdgeMode() ImageEdgeMode get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("primaryEdgeMode"))
	return rv
}/* debug [instance_properties/getter]: primaryEdgeMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143503-primaryedgemode
func (n_ NDArrayBinaryKernel) SetPrimaryEdgeMode(value ImageEdgeMode get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryEdgeMode:"), value)
}/* debug [instance_properties/setter]: primaryEdgeMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143504-primarykernelsizes
func (n_ NDArrayBinaryKernel) PrimaryKernelSizes() NDArraySizes get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("primaryKernelSizes"))
	return rv
}/* debug [instance_properties/getter]: primaryKernelSizes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143504-primarykernelsizes
func (n_ NDArrayBinaryKernel) SetPrimaryKernelSizes(value NDArraySizes get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryKernelSizes:"), value)
}/* debug [instance_properties/setter]: primaryKernelSizes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143505-primaryoffsets
func (n_ NDArrayBinaryKernel) PrimaryOffsets() NDArrayOffsets get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("primaryOffsets"))
	return rv
}/* debug [instance_properties/getter]: primaryOffsets */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143505-primaryoffsets
func (n_ NDArrayBinaryKernel) SetPrimaryOffsets(value NDArrayOffsets get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryOffsets:"), value)
}/* debug [instance_properties/setter]: primaryOffsets */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143506-primarystrides
func (n_ NDArrayBinaryKernel) PrimaryStrides() NDArrayOffsets get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("primaryStrides"))
	return rv
}/* debug [instance_properties/getter]: primaryStrides */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143506-primarystrides
func (n_ NDArrayBinaryKernel) SetPrimaryStrides(value NDArrayOffsets get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPrimaryStrides:"), value)
}/* debug [instance_properties/setter]: primaryStrides */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143507-secondarydilationrates
func (n_ NDArrayBinaryKernel) SecondaryDilationRates() NDArraySizes get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("secondaryDilationRates"))
	return rv
}/* debug [instance_properties/getter]: secondaryDilationRates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143507-secondarydilationrates
func (n_ NDArrayBinaryKernel) SetSecondaryDilationRates(value NDArraySizes get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryDilationRates:"), value)
}/* debug [instance_properties/setter]: secondaryDilationRates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143508-secondaryedgemode
func (n_ NDArrayBinaryKernel) SecondaryEdgeMode() ImageEdgeMode get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("secondaryEdgeMode"))
	return rv
}/* debug [instance_properties/getter]: secondaryEdgeMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143508-secondaryedgemode
func (n_ NDArrayBinaryKernel) SetSecondaryEdgeMode(value ImageEdgeMode get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryEdgeMode:"), value)
}/* debug [instance_properties/setter]: secondaryEdgeMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143509-secondarykernelsizes
func (n_ NDArrayBinaryKernel) SecondaryKernelSizes() NDArraySizes get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("secondaryKernelSizes"))
	return rv
}/* debug [instance_properties/getter]: secondaryKernelSizes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143509-secondarykernelsizes
func (n_ NDArrayBinaryKernel) SetSecondaryKernelSizes(value NDArraySizes get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryKernelSizes:"), value)
}/* debug [instance_properties/setter]: secondaryKernelSizes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143510-secondaryoffsets
func (n_ NDArrayBinaryKernel) SecondaryOffsets() NDArrayOffsets get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("secondaryOffsets"))
	return rv
}/* debug [instance_properties/getter]: secondaryOffsets */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143510-secondaryoffsets
func (n_ NDArrayBinaryKernel) SetSecondaryOffsets(value NDArrayOffsets get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryOffsets:"), value)
}/* debug [instance_properties/setter]: secondaryOffsets */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143511-secondarystrides
func (n_ NDArrayBinaryKernel) SecondaryStrides() NDArrayOffsets get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("secondaryStrides"))
	return rv
}/* debug [instance_properties/getter]: secondaryStrides */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143511-secondarystrides
func (n_ NDArrayBinaryKernel) SetSecondaryStrides(value NDArrayOffsets get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryStrides:"), value)
}/* debug [instance_properties/setter]: secondaryStrides */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayBinaryKernel */


