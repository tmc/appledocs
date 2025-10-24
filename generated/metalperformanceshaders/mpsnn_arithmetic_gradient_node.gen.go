// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNArithmeticGradientNode */


/* debug [class_header]: Header for MPSNNArithmeticGradientNode */
// The class instance for the [ArithmeticGradientNode] class.
var (
	ArithmeticGradientNodeClass     _ArithmeticGradientNodeClass
	ArithmeticGradientNodeClassOnce sync.Once
)

func getArithmeticGradientNodeClass() _ArithmeticGradientNodeClass {
	ArithmeticGradientNodeClassOnce.Do(func() {
		ArithmeticGradientNodeClass = _ArithmeticGradientNodeClass{objc.GetClass("MPSNNArithmeticGradientNode")}
	})
	return ArithmeticGradientNodeClass
}

type _ArithmeticGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ArithmeticGradientNode */
// An interface definition for the [ArithmeticGradientNode] class.
type IArithmeticGradientNode interface {
	IGradientFilterNode
	
/* debug [class_interface_properties]: Properties for ArithmeticGradientNode */
	// properties:
	SecondaryStrideInPixelsX() objectivec.IObject
	SetSecondaryStrideInPixelsX(value objectivec.IObject)
	SecondaryScale() objectivec.IObject
	SetSecondaryScale(value objectivec.IObject)
	SecondaryStrideInFeatureChannels() objectivec.IObject
	SetSecondaryStrideInFeatureChannels(value objectivec.IObject)
	MaximumValue() objectivec.IObject
	SetMaximumValue(value objectivec.IObject)
	IsSecondarySourceFilter() objectivec.IObject
	SetIsSecondarySourceFilter(value objectivec.IObject)
	Bias() objectivec.IObject
	SetBias(value objectivec.IObject)
	MinimumValue() objectivec.IObject
	SetMinimumValue(value objectivec.IObject)
	PrimaryScale() objectivec.IObject
	SetPrimaryScale(value objectivec.IObject)
	SecondaryStrideInPixelsY() objectivec.IObject
	SetSecondaryStrideInPixelsY(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ArithmeticGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ArithmeticGradientNode */
// Alloc allocates a new instance without initialization.
func (ac _ArithmeticGradientNodeClass) Alloc() ArithmeticGradientNode {
	rv := objc.Send[ArithmeticGradientNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _ArithmeticGradientNodeClass) New() ArithmeticGradientNode {
	rv := objc.Send[ArithmeticGradientNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ArithmeticGradientNode) Init() ArithmeticGradientNode {
	rv := objc.Send[ArithmeticGradientNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ArithmeticGradientNode) Autorelease() ArithmeticGradientNode {
	rv := objc.Send[ArithmeticGradientNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArithmeticGradientNode creates a new ArithmeticGradientNode instance.
func NewArithmeticGradientNode() ArithmeticGradientNode {
	return getArithmeticGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ArithmeticGradientNode */
// A representation of the base class for gradient arithmetic operators.


// A representation of the base class for gradient arithmetic operators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode
type ArithmeticGradientNode struct {
	GradientFilterNode
}

// ArithmeticGradientNodeFrom constructs a [ArithmeticGradientNode] from an unsafe.Pointer.
//
// A representation of the base class for gradient arithmetic operators.
func ArithmeticGradientNodeFrom(ptr unsafe.Pointer) ArithmeticGradientNode {
	return ArithmeticGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ArithmeticGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952980-initwithgradientimages
func NewArithmeticGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages unsafe.Pointer, filter IFilterNode, isSecondarySourceFilter bool) ArithmeticGradientNode {
	instance := getArithmeticGradientNodeClass().Alloc()
	rv := objc.Send[ArithmeticGradientNode](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), gradientImages, filter, isSecondarySourceFilter)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewArithmeticGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2956166-initwithsourcegradient
func NewArithmeticGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient IImageNode, sourceImage IImageNode, gradientState IBinaryGradientStateNode, isSecondarySourceFilter bool) ArithmeticGradientNode {
	instance := getArithmeticGradientNodeClass().Alloc()
	rv := objc.Send[ArithmeticGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewArithmeticGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ArithmeticGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2956167-nodewithsourcegradient
func (ac _ArithmeticGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient IImageNode, sourceImage IImageNode, gradientState IBinaryGradientStateNode, isSecondarySourceFilter bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ArithmeticGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ArithmeticGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ArithmeticGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952968-secondarystrideinpixelsx
func (a_ ArithmeticGradientNode) SecondaryStrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}/* debug [instance_properties/getter]: secondaryStrideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952968-secondarystrideinpixelsx
func (a_ ArithmeticGradientNode) SetSecondaryStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}/* debug [instance_properties/setter]: secondaryStrideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952981-secondaryscale
func (a_ ArithmeticGradientNode) SecondaryScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("secondaryScale"))
	return rv
}/* debug [instance_properties/getter]: secondaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952981-secondaryscale
func (a_ ArithmeticGradientNode) SetSecondaryScale(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSecondaryScale:"), value)
}/* debug [instance_properties/setter]: secondaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952984-secondarystrideinfeaturechannels
func (a_ ArithmeticGradientNode) SecondaryStrideInFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: secondaryStrideInFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952984-secondarystrideinfeaturechannels
func (a_ ArithmeticGradientNode) SetSecondaryStrideInFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}/* debug [instance_properties/setter]: secondaryStrideInFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952986-maximumvalue
func (a_ ArithmeticGradientNode) MaximumValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("maximumValue"))
	return rv
}/* debug [instance_properties/getter]: maximumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952986-maximumvalue
func (a_ ArithmeticGradientNode) SetMaximumValue(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumValue:"), value)
}/* debug [instance_properties/setter]: maximumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952987-issecondarysourcefilter
func (a_ ArithmeticGradientNode) IsSecondarySourceFilter() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("isSecondarySourceFilter"))
	return rv
}/* debug [instance_properties/getter]: isSecondarySourceFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952987-issecondarysourcefilter
func (a_ ArithmeticGradientNode) SetIsSecondarySourceFilter(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsSecondarySourceFilter:"), value)
}/* debug [instance_properties/setter]: isSecondarySourceFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952988-bias
func (a_ ArithmeticGradientNode) Bias() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("bias"))
	return rv
}/* debug [instance_properties/getter]: bias */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952988-bias
func (a_ ArithmeticGradientNode) SetBias(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBias:"), value)
}/* debug [instance_properties/setter]: bias */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952989-minimumvalue
func (a_ ArithmeticGradientNode) MinimumValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("minimumValue"))
	return rv
}/* debug [instance_properties/getter]: minimumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952989-minimumvalue
func (a_ ArithmeticGradientNode) SetMinimumValue(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMinimumValue:"), value)
}/* debug [instance_properties/setter]: minimumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952993-primaryscale
func (a_ ArithmeticGradientNode) PrimaryScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("primaryScale"))
	return rv
}/* debug [instance_properties/getter]: primaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952993-primaryscale
func (a_ ArithmeticGradientNode) SetPrimaryScale(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimaryScale:"), value)
}/* debug [instance_properties/setter]: primaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952994-secondarystrideinpixelsy
func (a_ ArithmeticGradientNode) SecondaryStrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}/* debug [instance_properties/getter]: secondaryStrideInPixelsY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952994-secondarystrideinpixelsy
func (a_ ArithmeticGradientNode) SetSecondaryStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}/* debug [instance_properties/setter]: secondaryStrideInPixelsY */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNArithmeticGradientNode */


