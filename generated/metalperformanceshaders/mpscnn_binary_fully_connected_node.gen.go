// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNBinaryFullyConnectedNode */


/* debug [class_header]: Header for MPSCNNBinaryFullyConnectedNode */
// The class instance for the [CNNBinaryFullyConnectedNode] class.
var (
	CNNBinaryFullyConnectedNodeClass     _CNNBinaryFullyConnectedNodeClass
	CNNBinaryFullyConnectedNodeClassOnce sync.Once
)

func getCNNBinaryFullyConnectedNodeClass() _CNNBinaryFullyConnectedNodeClass {
	CNNBinaryFullyConnectedNodeClassOnce.Do(func() {
		CNNBinaryFullyConnectedNodeClass = _CNNBinaryFullyConnectedNodeClass{objc.GetClass("MPSCNNBinaryFullyConnectedNode")}
	})
	return CNNBinaryFullyConnectedNodeClass
}

type _CNNBinaryFullyConnectedNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNBinaryFullyConnectedNode */
// An interface definition for the [CNNBinaryFullyConnectedNode] class.
type ICNNBinaryFullyConnectedNode interface {
	ICNNBinaryConvolutionNode
	
/* debug [class_interface_properties]: Properties for CNNBinaryFullyConnectedNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNBinaryFullyConnectedNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNBinaryFullyConnectedNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNBinaryFullyConnectedNodeClass) Alloc() CNNBinaryFullyConnectedNode {
	rv := objc.Send[CNNBinaryFullyConnectedNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNBinaryFullyConnectedNodeClass) New() CNNBinaryFullyConnectedNode {
	rv := objc.Send[CNNBinaryFullyConnectedNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNBinaryFullyConnectedNode) Init() CNNBinaryFullyConnectedNode {
	rv := objc.Send[CNNBinaryFullyConnectedNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNBinaryFullyConnectedNode) Autorelease() CNNBinaryFullyConnectedNode {
	rv := objc.Send[CNNBinaryFullyConnectedNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNBinaryFullyConnectedNode creates a new CNNBinaryFullyConnectedNode instance.
func NewCNNBinaryFullyConnectedNode() CNNBinaryFullyConnectedNode {
	return getCNNBinaryFullyConnectedNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNBinaryFullyConnectedNode */
// A representation of a fully connected convolution layer with binary weights and optionally binarized input image.


// A representation of a fully connected convolution layer with binary weights and optionally binarized input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryFullyConnectedNode
type CNNBinaryFullyConnectedNode struct {
	CNNBinaryConvolutionNode
}

// CNNBinaryFullyConnectedNodeFrom constructs a [CNNBinaryFullyConnectedNode] from an unsafe.Pointer.
//
// A representation of a fully connected convolution layer with binary weights and optionally binarized input image.
func CNNBinaryFullyConnectedNodeFrom(ptr unsafe.Pointer) CNNBinaryFullyConnectedNode {
	return CNNBinaryFullyConnectedNode{
		CNNBinaryConvolutionNode: CNNBinaryConvolutionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNBinaryFullyConnectedNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryfullyconnectednode/2942637-initwithsource
func NewCNNBinaryFullyConnectedNodeWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(sourceNode IImageNode, weights unsafe.Pointer, outputBiasTerms objectivec.IObject, outputScaleTerms objectivec.IObject, inputBiasTerms objectivec.IObject, inputScaleTerms objectivec.IObject, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryFullyConnectedNode {
	instance := getCNNBinaryFullyConnectedNodeClass().Alloc()
	rv := objc.Send[CNNBinaryFullyConnectedNode](instance.ID, objc.Sel("initWithSource:weights:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:"), sourceNode, weights, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBinaryFullyConnectedNodeWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryfullyconnectednode/2866443-initwithsource
func NewCNNBinaryFullyConnectedNodeWithSourceWeightsScaleValueTypeFlags(sourceNode IImageNode, weights unsafe.Pointer, scaleValue float32, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryFullyConnectedNode {
	instance := getCNNBinaryFullyConnectedNodeClass().Alloc()
	rv := objc.Send[CNNBinaryFullyConnectedNode](instance.ID, objc.Sel("initWithSource:weights:scaleValue:type:flags:"), sourceNode, weights, scaleValue, type_, flags)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBinaryFullyConnectedNodeWithSourceWeightsScaleValueTypeFlags */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNBinaryFullyConnectedNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryfullyconnectednode/2866469-nodewithsource
func (cc _CNNBinaryFullyConnectedNodeClass) NodeWithSourceWeightsScaleValueTypeFlags(sourceNode IImageNode, weights unsafe.Pointer, scaleValue float32, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:weights:scaleValue:type:flags:"), sourceNode, weights, scaleValue, type_, flags)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceWeightsScaleValueTypeFlags) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryfullyconnectednode/2942635-nodewithsource
func (cc _CNNBinaryFullyConnectedNodeClass) NodeWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(sourceNode IImageNode, weights unsafe.Pointer, outputBiasTerms objectivec.IObject, outputScaleTerms objectivec.IObject, inputBiasTerms objectivec.IObject, inputScaleTerms objectivec.IObject, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:weights:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:"), sourceNode, weights, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNBinaryFullyConnectedNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNBinaryFullyConnectedNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNBinaryFullyConnectedNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNBinaryFullyConnectedNode */


