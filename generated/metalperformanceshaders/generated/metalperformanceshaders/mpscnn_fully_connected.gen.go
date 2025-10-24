// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNFullyConnected */


/* debug [class_header]: Header for MPSCNNFullyConnected */
// The class instance for the [CNNFullyConnected] class.
var (
	CNNFullyConnectedClass     _CNNFullyConnectedClass
	CNNFullyConnectedClassOnce sync.Once
)

func getCNNFullyConnectedClass() _CNNFullyConnectedClass {
	CNNFullyConnectedClassOnce.Do(func() {
		CNNFullyConnectedClass = _CNNFullyConnectedClass{objc.GetClass("MPSCNNFullyConnected")}
	})
	return CNNFullyConnectedClass
}

type _CNNFullyConnectedClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNFullyConnected */
// An interface definition for the [CNNFullyConnected] class.
type ICNNFullyConnected interface {
	ICNNConvolution
	
/* debug [class_interface_properties]: Properties for CNNFullyConnected */
	// properties:
	Groups() int
	SetGroups(value int)
	StrideInPixelsY() int
	SetStrideInPixelsY(value int)
	ClipRect() MTLRegion /* not a class type */
	SetClipRect(value MTLRegion /* not a class type */)
	Offset() objc.IObject /* cross-framework: MPSOffset */
	SetOffset(value objc.IObject /* cross-framework: MPSOffset */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNFullyConnected */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNFullyConnected */
// Alloc allocates a new instance without initialization.
func (cc _CNNFullyConnectedClass) Alloc() CNNFullyConnected {
	rv := objc.Send[CNNFullyConnected](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNFullyConnectedClass) New() CNNFullyConnected {
	rv := objc.Send[CNNFullyConnected](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNFullyConnected) Init() CNNFullyConnected {
	rv := objc.Send[CNNFullyConnected](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNFullyConnected) Autorelease() CNNFullyConnected {
	rv := objc.Send[CNNFullyConnected](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNFullyConnected creates a new CNNFullyConnected instance.
func NewCNNFullyConnected() CNNFullyConnected {
	return getCNNFullyConnectedClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNFullyConnected */
// A fully connected convolution layer, also known as an inner product layer.
//
// A fully connected layer in a Convolutional Neural Network (CNN) is one where every input channel is connected to every output channel. The kernel width is equal to the width of the source image, and the kernel height is equal to the height of the source image. The width and height of the output is . A fully connected layer takes an object with dimensions , convolves it with , and produces a output. Thus, the following conditions must be true: You can think of a fully connected layer as a matrix multiplication where the image is flattened into a vector of length , and the weights are arranged in a matrix of dimension to produce an output vector of length . The value of the , , and properties must be . The property is not applicable and it is ignored. Because the clip rectangle is clamped to the destination image bounds, if the destination is , you do not need to set the property.


// A fully connected convolution layer, also known as an inner product layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnected
type CNNFullyConnected struct {
	CNNConvolution
}

// CNNFullyConnectedFrom constructs a [CNNFullyConnected] from an unsafe.Pointer.
//
// A fully connected convolution layer, also known as an inner product layer.
func CNNFullyConnectedFrom(ptr unsafe.Pointer) CNNFullyConnected {
	return CNNFullyConnected{
		CNNConvolution: CNNConvolutionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNFullyConnected */

// Initializes a fully connected convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnected/init(coder:device:)
func NewCNNFullyConnectedWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) CNNFullyConnected {
	instance := getCNNFullyConnectedClass().Alloc()
	rv := objc.Send[CNNFullyConnected](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNFullyConnectedWithCoderDevice */


// Initializes a fully connected convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnected/init(device:convolutionDescriptor:kernelWeights:biasTerms:flags:)
func NewCNNFullyConnectedWithDeviceConvolutionDescriptorKernelWeightsBiasTermsFlags(device unsafe.Pointer, convolutionDescriptor IMPSCNNConvolutionDescriptor, kernelWeights objectivec.IObject, biasTerms objectivec.IObject, flags CNNConvolutionFlags) CNNFullyConnected {
	instance := getCNNFullyConnectedClass().Alloc()
	rv := objc.Send[CNNFullyConnected](instance.ID, objc.Sel("initWithDevice:convolutionDescriptor:kernelWeights:biasTerms:flags:"), device, convolutionDescriptor, kernelWeights, biasTerms, flags)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNFullyConnectedWithDeviceConvolutionDescriptorKernelWeightsBiasTermsFlags */


// Initializes a fully connected convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnected/init(device:weights:)
func NewCNNFullyConnectedWithDeviceWeights(device unsafe.Pointer, weights unsafe.Pointer) CNNFullyConnected {
	instance := getCNNFullyConnectedClass().Alloc()
	rv := objc.Send[CNNFullyConnected](instance.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNFullyConnectedWithDeviceWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNFullyConnected */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNFullyConnected */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNFullyConnected */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNFullyConnected */

// The number of groups that the input and output channels are divided into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/groups
func (c_ CNNFullyConnected) Groups() int {
	rv := objc.Send[int](c_.ID, objc.Sel("groups"))
	return rv
}/* debug [instance_properties/getter]: groups */


// The number of groups that the input and output channels are divided into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/groups
func (c_ CNNFullyConnected) SetGroups(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroups:"), value)
}/* debug [instance_properties/setter]: groups */


// The output stride (downsampling factor) in the y dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/strideinpixelsy
func (c_ CNNFullyConnected) StrideInPixelsY() int {
	rv := objc.Send[int](c_.ID, objc.Sel("strideInPixelsY"))
	return rv
}/* debug [instance_properties/getter]: strideInPixelsY */


// The output stride (downsampling factor) in the y dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/strideinpixelsy
func (c_ CNNFullyConnected) SetStrideInPixelsY(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsY:"), value)
}/* debug [instance_properties/setter]: strideInPixelsY */


// An optional clip rectangle to use when writing data. Only the pixels in the clip rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/cliprect
func (c_ CNNFullyConnected) ClipRect() MTLRegion /* not a class type */ {
	rv := objc.Send[Region](c_.ID, objc.Sel("clipRect"))
	return rv
}/* debug [instance_properties/getter]: clipRect */


// An optional clip rectangle to use when writing data. Only the pixels in the clip rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/cliprect
func (c_ CNNFullyConnected) SetClipRect(value MTLRegion /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClipRect:"), value)
}/* debug [instance_properties/setter]: clipRect */


// The position of the destination image’s clip rectangle origin, relative to the source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/offset
func (c_ CNNFullyConnected) Offset() objc.IObject /* cross-framework: MPSOffset */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// The position of the destination image’s clip rectangle origin, relative to the source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/offset
func (c_ CNNFullyConnected) SetOffset(value objc.IObject /* cross-framework: MPSOffset */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNFullyConnected */


