// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNBinaryConvolutionNode */


/* debug [class_header]: Header for MPSCNNBinaryConvolutionNode */
// The class instance for the [CNNBinaryConvolutionNode] class.
var (
	CNNBinaryConvolutionNodeClass     _CNNBinaryConvolutionNodeClass
	CNNBinaryConvolutionNodeClassOnce sync.Once
)

func getCNNBinaryConvolutionNodeClass() _CNNBinaryConvolutionNodeClass {
	CNNBinaryConvolutionNodeClassOnce.Do(func() {
		CNNBinaryConvolutionNodeClass = _CNNBinaryConvolutionNodeClass{objc.GetClass("MPSCNNBinaryConvolutionNode")}
	})
	return CNNBinaryConvolutionNodeClass
}

type _CNNBinaryConvolutionNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNBinaryConvolutionNode */
// An interface definition for the [CNNBinaryConvolutionNode] class.
type ICNNBinaryConvolutionNode interface {
	ICNNConvolutionNode
	
/* debug [class_interface_properties]: Properties for CNNBinaryConvolutionNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNBinaryConvolutionNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNBinaryConvolutionNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNBinaryConvolutionNodeClass) Alloc() CNNBinaryConvolutionNode {
	rv := objc.Send[CNNBinaryConvolutionNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNBinaryConvolutionNodeClass) New() CNNBinaryConvolutionNode {
	rv := objc.Send[CNNBinaryConvolutionNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNBinaryConvolutionNode) Init() CNNBinaryConvolutionNode {
	rv := objc.Send[CNNBinaryConvolutionNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNBinaryConvolutionNode) Autorelease() CNNBinaryConvolutionNode {
	rv := objc.Send[CNNBinaryConvolutionNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNBinaryConvolutionNode creates a new CNNBinaryConvolutionNode instance.
func NewCNNBinaryConvolutionNode() CNNBinaryConvolutionNode {
	return getCNNBinaryConvolutionNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNBinaryConvolutionNode */
// A representation of a convolution kernel with binary weights and an input image using binary approximations.


// A representation of a convolution kernel with binary weights and an input image using binary approximations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryConvolutionNode
type CNNBinaryConvolutionNode struct {
	CNNConvolutionNode
}

// CNNBinaryConvolutionNodeFrom constructs a [CNNBinaryConvolutionNode] from an unsafe.Pointer.
//
// A representation of a convolution kernel with binary weights and an input image using binary approximations.
func CNNBinaryConvolutionNodeFrom(ptr unsafe.Pointer) CNNBinaryConvolutionNode {
	return CNNBinaryConvolutionNode{
		CNNConvolutionNode: CNNConvolutionNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNBinaryConvolutionNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolutionnode/2942631-initwithsource
func NewCNNBinaryConvolutionNodeWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(sourceNode IImageNode, weights unsafe.Pointer, outputBiasTerms objectivec.IObject, outputScaleTerms objectivec.IObject, inputBiasTerms objectivec.IObject, inputScaleTerms objectivec.IObject, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryConvolutionNode {
	instance := getCNNBinaryConvolutionNodeClass().Alloc()
	rv := objc.Send[CNNBinaryConvolutionNode](instance.ID, objc.Sel("initWithSource:weights:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:"), sourceNode, weights, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBinaryConvolutionNodeWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolutionnode/2866509-initwithsource
func NewCNNBinaryConvolutionNodeWithSourceWeightsScaleValueTypeFlags(sourceNode IImageNode, weights unsafe.Pointer, scaleValue float32, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryConvolutionNode {
	instance := getCNNBinaryConvolutionNodeClass().Alloc()
	rv := objc.Send[CNNBinaryConvolutionNode](instance.ID, objc.Sel("initWithSource:weights:scaleValue:type:flags:"), sourceNode, weights, scaleValue, type_, flags)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBinaryConvolutionNodeWithSourceWeightsScaleValueTypeFlags */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNBinaryConvolutionNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolutionnode/2866487-nodewithsource
func (cc _CNNBinaryConvolutionNodeClass) NodeWithSourceWeightsScaleValueTypeFlags(sourceNode IImageNode, weights unsafe.Pointer, scaleValue float32, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:weights:scaleValue:type:flags:"), sourceNode, weights, scaleValue, type_, flags)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceWeightsScaleValueTypeFlags) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolutionnode/2942632-nodewithsource
func (cc _CNNBinaryConvolutionNodeClass) NodeWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(sourceNode IImageNode, weights unsafe.Pointer, outputBiasTerms objectivec.IObject, outputScaleTerms objectivec.IObject, inputBiasTerms objectivec.IObject, inputScaleTerms objectivec.IObject, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:weights:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:"), sourceNode, weights, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNBinaryConvolutionNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNBinaryConvolutionNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNBinaryConvolutionNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNBinaryConvolutionNode */


