// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNBinaryConvolution */


/* debug [class_header]: Header for MPSCNNBinaryConvolution */
// The class instance for the [CNNBinaryConvolution] class.
var (
	CNNBinaryConvolutionClass     _CNNBinaryConvolutionClass
	CNNBinaryConvolutionClassOnce sync.Once
)

func getCNNBinaryConvolutionClass() _CNNBinaryConvolutionClass {
	CNNBinaryConvolutionClassOnce.Do(func() {
		CNNBinaryConvolutionClass = _CNNBinaryConvolutionClass{objc.GetClass("MPSCNNBinaryConvolution")}
	})
	return CNNBinaryConvolutionClass
}

type _CNNBinaryConvolutionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNBinaryConvolution */
// An interface definition for the [CNNBinaryConvolution] class.
type ICNNBinaryConvolution interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for CNNBinaryConvolution */
	// properties:
	OutputFeatureChannels() objectivec.IObject
	SetOutputFeatureChannels(value objectivec.IObject)
	InputFeatureChannels() objectivec.IObject
	SetInputFeatureChannels(value objectivec.IObject)
	Groups() int
	SetGroups(value int)
	KernelHeight() int
	SetKernelHeight(value int)
	KernelWidth() int
	SetKernelWidth(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNBinaryConvolution */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNBinaryConvolution */
// Alloc allocates a new instance without initialization.
func (cc _CNNBinaryConvolutionClass) Alloc() CNNBinaryConvolution {
	rv := objc.Send[CNNBinaryConvolution](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNBinaryConvolutionClass) New() CNNBinaryConvolution {
	rv := objc.Send[CNNBinaryConvolution](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNBinaryConvolution) Init() CNNBinaryConvolution {
	rv := objc.Send[CNNBinaryConvolution](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNBinaryConvolution) Autorelease() CNNBinaryConvolution {
	rv := objc.Send[CNNBinaryConvolution](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNBinaryConvolution creates a new CNNBinaryConvolution instance.
func NewCNNBinaryConvolution() CNNBinaryConvolution {
	return getCNNBinaryConvolutionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNBinaryConvolution */
// A convolution kernel with binary weights and an input image using binary approximations.
//
// The optionally first binarizes the input image and then convolves the result with a set of binary-valued filters, each producing one feature map in the output image (which is a normal image). The output is computed as follows: where the is over the spatial filter kernel window defined by and , is over the input feature channel indices within group, contains the binary weights, interpreted as or and is the array and bias is the array. Above is the image index in batch the sum over input channels runs through the group indices. The convolution operator ⊗ is defined by passed in at initialization time of the filter: and scaled according to the optional scaling operations. Note that we output the values of the bitwise convolutions to interval , which means that the output of the XNOR-operator is scaled implicitly as follows: This means that for a dot-product of two 32-bit words the result is: and scaled according to the optional scaling operations. Note that we output the values of the AND-operation is assumed to lie in interval and hence no more implicit scaling takes place. This means that for a dot-product of two 32-bit words the result is: The input data can be pre-offset and scaled by providing the and parameters for the initialization functions and this can be used for example to accomplish batch normalization of the data. The scaling of input values happens before possible beta-image computation. The parameter above is an optional image which is used to compute scaling factors for each spatial position and image index. For the XNOR-Net based networks this is computed as follows: where are summed over the convolution filter window. where is the original input image (in full precision) and is the number of input channels in the input image. Parameter is not passed as input and to enable beta-scaling the user can provide in the flags parameter in the initialization functions. Finally the normal activation neuron is applied and the result is written to the output image.


// A convolution kernel with binary weights and an input image using binary approximations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolution
type CNNBinaryConvolution struct {
	CNNKernel
}

// CNNBinaryConvolutionFrom constructs a [CNNBinaryConvolution] from an unsafe.Pointer.
//
// A convolution kernel with binary weights and an input image using binary approximations.
func CNNBinaryConvolutionFrom(ptr unsafe.Pointer) CNNBinaryConvolution {
	return CNNBinaryConvolution{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNBinaryConvolution */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolution/2867194-initwithcoder
func NewCNNBinaryConvolutionWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNBinaryConvolution {
	instance := getCNNBinaryConvolutionClass().Alloc()
	rv := objc.Send[CNNBinaryConvolution](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBinaryConvolutionWithCoderDevice */


// Initializes a binary convolution kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolution/2866978-initwithdevice
func NewCNNBinaryConvolutionWithDeviceConvolutionDataOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(device unsafe.Pointer, convolutionData unsafe.Pointer, outputBiasTerms objectivec.IObject, outputScaleTerms objectivec.IObject, inputBiasTerms objectivec.IObject, inputScaleTerms objectivec.IObject, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryConvolution {
	instance := getCNNBinaryConvolutionClass().Alloc()
	rv := objc.Send[CNNBinaryConvolution](instance.ID, objc.Sel("initWithDevice:convolutionData:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:"), device, convolutionData, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBinaryConvolutionWithDeviceConvolutionDataOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags */


// Initializes a binary convolution kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolution/2866981-initwithdevice
func NewCNNBinaryConvolutionWithDeviceConvolutionDataScaleValueTypeFlags(device unsafe.Pointer, convolutionData unsafe.Pointer, scaleValue float32, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryConvolution {
	instance := getCNNBinaryConvolutionClass().Alloc()
	rv := objc.Send[CNNBinaryConvolution](instance.ID, objc.Sel("initWithDevice:convolutionData:scaleValue:type:flags:"), device, convolutionData, scaleValue, type_, flags)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBinaryConvolutionWithDeviceConvolutionDataScaleValueTypeFlags */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNBinaryConvolution */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNBinaryConvolution */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNBinaryConvolution */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNBinaryConvolution */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolution/2866959-outputfeaturechannels
func (c_ CNNBinaryConvolution) OutputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: outputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolution/2866959-outputfeaturechannels
func (c_ CNNBinaryConvolution) SetOutputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: outputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolution/2867126-inputfeaturechannels
func (c_ CNNBinaryConvolution) InputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: inputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolution/2867126-inputfeaturechannels
func (c_ CNNBinaryConvolution) SetInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: inputFeatureChannels */


// The number of groups that the input and output channels are divided into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/groups
func (c_ CNNBinaryConvolution) Groups() int {
	rv := objc.Send[int](c_.ID, objc.Sel("groups"))
	return rv
}/* debug [instance_properties/getter]: groups */


// The number of groups that the input and output channels are divided into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/groups
func (c_ CNNBinaryConvolution) SetGroups(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroups:"), value)
}/* debug [instance_properties/setter]: groups */


// The height of the kernel window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/kernelheight
func (c_ CNNBinaryConvolution) KernelHeight() int {
	rv := objc.Send[int](c_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// The height of the kernel window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/kernelheight
func (c_ CNNBinaryConvolution) SetKernelHeight(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelHeight:"), value)
}/* debug [instance_properties/setter]: kernelHeight */


// The width of the kernel window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/kernelwidth
func (c_ CNNBinaryConvolution) KernelWidth() int {
	rv := objc.Send[int](c_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// The width of the kernel window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/kernelwidth
func (c_ CNNBinaryConvolution) SetKernelWidth(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelWidth:"), value)
}/* debug [instance_properties/setter]: kernelWidth */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNBinaryConvolution */


