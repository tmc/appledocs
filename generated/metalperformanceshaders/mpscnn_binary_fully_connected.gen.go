// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNBinaryFullyConnected */


/* debug [class_header]: Header for MPSCNNBinaryFullyConnected */
// The class instance for the [CNNBinaryFullyConnected] class.
var (
	CNNBinaryFullyConnectedClass     _CNNBinaryFullyConnectedClass
	CNNBinaryFullyConnectedClassOnce sync.Once
)

func getCNNBinaryFullyConnectedClass() _CNNBinaryFullyConnectedClass {
	CNNBinaryFullyConnectedClassOnce.Do(func() {
		CNNBinaryFullyConnectedClass = _CNNBinaryFullyConnectedClass{objc.GetClass("MPSCNNBinaryFullyConnected")}
	})
	return CNNBinaryFullyConnectedClass
}

type _CNNBinaryFullyConnectedClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNBinaryFullyConnected */
// An interface definition for the [CNNBinaryFullyConnected] class.
type ICNNBinaryFullyConnected interface {
	ICNNBinaryConvolution
	
/* debug [class_interface_properties]: Properties for CNNBinaryFullyConnected */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNBinaryFullyConnected */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNBinaryFullyConnected */
// Alloc allocates a new instance without initialization.
func (cc _CNNBinaryFullyConnectedClass) Alloc() CNNBinaryFullyConnected {
	rv := objc.Send[CNNBinaryFullyConnected](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNBinaryFullyConnectedClass) New() CNNBinaryFullyConnected {
	rv := objc.Send[CNNBinaryFullyConnected](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNBinaryFullyConnected) Init() CNNBinaryFullyConnected {
	rv := objc.Send[CNNBinaryFullyConnected](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNBinaryFullyConnected) Autorelease() CNNBinaryFullyConnected {
	rv := objc.Send[CNNBinaryFullyConnected](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNBinaryFullyConnected creates a new CNNBinaryFullyConnected instance.
func NewCNNBinaryFullyConnected() CNNBinaryFullyConnected {
	return getCNNBinaryFullyConnectedClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNBinaryFullyConnected */
// A fully connected convolution layer with binary weights and optionally binarized input image.


// A fully connected convolution layer with binary weights and optionally binarized input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryFullyConnected
type CNNBinaryFullyConnected struct {
	CNNBinaryConvolution
}

// CNNBinaryFullyConnectedFrom constructs a [CNNBinaryFullyConnected] from an unsafe.Pointer.
//
// A fully connected convolution layer with binary weights and optionally binarized input image.
func CNNBinaryFullyConnectedFrom(ptr unsafe.Pointer) CNNBinaryFullyConnected {
	return CNNBinaryFullyConnected{
		CNNBinaryConvolution: CNNBinaryConvolutionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNBinaryFullyConnected */

// Initializes a fully connected convolution layer with binary weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryfullyconnected/2867052-initwithcoder
func NewCNNBinaryFullyConnectedWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNBinaryFullyConnected {
	instance := getCNNBinaryFullyConnectedClass().Alloc()
	rv := objc.Send[CNNBinaryFullyConnected](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBinaryFullyConnectedWithCoderDevice */


// Initializes a fully connected convolution layer with binary weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryfullyconnected/2867059-initwithdevice
func NewCNNBinaryFullyConnectedWithDeviceConvolutionDataOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(device unsafe.Pointer, convolutionData unsafe.Pointer, outputBiasTerms objectivec.IObject, outputScaleTerms objectivec.IObject, inputBiasTerms objectivec.IObject, inputScaleTerms objectivec.IObject, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryFullyConnected {
	instance := getCNNBinaryFullyConnectedClass().Alloc()
	rv := objc.Send[CNNBinaryFullyConnected](instance.ID, objc.Sel("initWithDevice:convolutionData:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:"), device, convolutionData, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBinaryFullyConnectedWithDeviceConvolutionDataOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags */


// Initializes a fully connected convolution layer with binary weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryfullyconnected/2867196-initwithdevice
func NewCNNBinaryFullyConnectedWithDeviceConvolutionDataScaleValueTypeFlags(device unsafe.Pointer, convolutionData unsafe.Pointer, scaleValue float32, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryFullyConnected {
	instance := getCNNBinaryFullyConnectedClass().Alloc()
	rv := objc.Send[CNNBinaryFullyConnected](instance.ID, objc.Sel("initWithDevice:convolutionData:scaleValue:type:flags:"), device, convolutionData, scaleValue, type_, flags)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBinaryFullyConnectedWithDeviceConvolutionDataScaleValueTypeFlags */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNBinaryFullyConnected */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNBinaryFullyConnected */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNBinaryFullyConnected */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNBinaryFullyConnected */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNBinaryFullyConnected */


