// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNBinaryArithmeticNode */


/* debug [class_header]: Header for MPSNNBinaryArithmeticNode */
// The class instance for the [BinaryArithmeticNode] class.
var (
	BinaryArithmeticNodeClass     _BinaryArithmeticNodeClass
	BinaryArithmeticNodeClassOnce sync.Once
)

func getBinaryArithmeticNodeClass() _BinaryArithmeticNodeClass {
	BinaryArithmeticNodeClassOnce.Do(func() {
		BinaryArithmeticNodeClass = _BinaryArithmeticNodeClass{objc.GetClass("MPSNNBinaryArithmeticNode")}
	})
	return BinaryArithmeticNodeClass
}

type _BinaryArithmeticNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BinaryArithmeticNode */
// An interface definition for the [BinaryArithmeticNode] class.
type IBinaryArithmeticNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for BinaryArithmeticNode */
	// properties:
	Bias() objectivec.IObject
	SetBias(value objectivec.IObject)
	PrimaryScale() objectivec.IObject
	SetPrimaryScale(value objectivec.IObject)
	MinimumValue() objectivec.IObject
	SetMinimumValue(value objectivec.IObject)
	SecondaryStrideInPixelsX() objectivec.IObject
	SetSecondaryStrideInPixelsX(value objectivec.IObject)
	PrimaryStrideInPixelsX() objectivec.IObject
	SetPrimaryStrideInPixelsX(value objectivec.IObject)
	SecondaryStrideInFeatureChannels() objectivec.IObject
	SetSecondaryStrideInFeatureChannels(value objectivec.IObject)
	SecondaryScale() objectivec.IObject
	SetSecondaryScale(value objectivec.IObject)
	MaximumValue() objectivec.IObject
	SetMaximumValue(value objectivec.IObject)
	PrimaryStrideInFeatureChannels() objectivec.IObject
	SetPrimaryStrideInFeatureChannels(value objectivec.IObject)
	SecondaryStrideInPixelsY() objectivec.IObject
	SetSecondaryStrideInPixelsY(value objectivec.IObject)
	PrimaryStrideInPixelsY() objectivec.IObject
	SetPrimaryStrideInPixelsY(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BinaryArithmeticNode */
	// methods:
	GradientFilters()
	GradientFiltersWithSources(gradientImages unsafe.Pointer) unsafe.Pointer
	GradientClass()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BinaryArithmeticNode */
// Alloc allocates a new instance without initialization.
func (bc _BinaryArithmeticNodeClass) Alloc() BinaryArithmeticNode {
	rv := objc.Send[BinaryArithmeticNode](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BinaryArithmeticNodeClass) New() BinaryArithmeticNode {
	rv := objc.Send[BinaryArithmeticNode](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BinaryArithmeticNode) Init() BinaryArithmeticNode {
	rv := objc.Send[BinaryArithmeticNode](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BinaryArithmeticNode) Autorelease() BinaryArithmeticNode {
	rv := objc.Send[BinaryArithmeticNode](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBinaryArithmeticNode creates a new BinaryArithmeticNode instance.
func NewBinaryArithmeticNode() BinaryArithmeticNode {
	return getBinaryArithmeticNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BinaryArithmeticNode */
// Virtual base class for basic arithmetic nodes.


// Virtual base class for basic arithmetic nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode
type BinaryArithmeticNode struct {
	FilterNode
}

// BinaryArithmeticNodeFrom constructs a [BinaryArithmeticNode] from an unsafe.Pointer.
//
// Virtual base class for basic arithmetic nodes.
func BinaryArithmeticNodeFrom(ptr unsafe.Pointer) BinaryArithmeticNode {
	return BinaryArithmeticNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BinaryArithmeticNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890825-initwithleftsource
func NewBinaryArithmeticNodeWithLeftSourceRightSource(left IImageNode, right IImageNode) BinaryArithmeticNode {
	instance := getBinaryArithmeticNodeClass().Alloc()
	rv := objc.Send[BinaryArithmeticNode](instance.ID, objc.Sel("initWithLeftSource:rightSource:"), left, right)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBinaryArithmeticNodeWithLeftSourceRightSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890820-initwithsources
func NewBinaryArithmeticNodeWithSources(sourceNodes unsafe.Pointer) BinaryArithmeticNode {
	instance := getBinaryArithmeticNodeClass().Alloc()
	rv := objc.Send[BinaryArithmeticNode](instance.ID, objc.Sel("initWithSources:"), sourceNodes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBinaryArithmeticNodeWithSources */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BinaryArithmeticNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890829-nodewithsources
func (bc _BinaryArithmeticNodeClass) NodeWithSources(sourceNodes unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("nodeWithSources:"), sourceNodes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSources) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890830-nodewithleftsource
func (bc _BinaryArithmeticNodeClass) NodeWithLeftSourceRightSource(left IImageNode, right IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("nodeWithLeftSource:rightSource:"), left, right)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithLeftSourceRightSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BinaryArithmeticNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BinaryArithmeticNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952967-gradientfilters
func (b_ BinaryArithmeticNode) GradientFilters() {
	objc.Send[objc.ID](b_.ID, objc.Sel("gradientFilters"))
}/* debug [instance_methods/method]: GradientFilters */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952967-gradientfilterswithsources
func (b_ BinaryArithmeticNode) GradientFiltersWithSources(gradientImages unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("gradientFiltersWithSources:"), gradientImages)
	return rv
}/* debug [instance_methods/method]: GradientFiltersWithSources */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952978-gradientclass
func (b_ BinaryArithmeticNode) GradientClass() {
	objc.Send[objc.ID](b_.ID, objc.Sel("gradientClass"))
}/* debug [instance_methods/method]: GradientClass */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BinaryArithmeticNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952964-bias
func (b_ BinaryArithmeticNode) Bias() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("bias"))
	return rv
}/* debug [instance_properties/getter]: bias */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952964-bias
func (b_ BinaryArithmeticNode) SetBias(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBias:"), value)
}/* debug [instance_properties/setter]: bias */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952966-primaryscale
func (b_ BinaryArithmeticNode) PrimaryScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("primaryScale"))
	return rv
}/* debug [instance_properties/getter]: primaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952966-primaryscale
func (b_ BinaryArithmeticNode) SetPrimaryScale(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryScale:"), value)
}/* debug [instance_properties/setter]: primaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952970-minimumvalue
func (b_ BinaryArithmeticNode) MinimumValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("minimumValue"))
	return rv
}/* debug [instance_properties/getter]: minimumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952970-minimumvalue
func (b_ BinaryArithmeticNode) SetMinimumValue(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMinimumValue:"), value)
}/* debug [instance_properties/setter]: minimumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952972-secondarystrideinpixelsx
func (b_ BinaryArithmeticNode) SecondaryStrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}/* debug [instance_properties/getter]: secondaryStrideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952972-secondarystrideinpixelsx
func (b_ BinaryArithmeticNode) SetSecondaryStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}/* debug [instance_properties/setter]: secondaryStrideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952973-primarystrideinpixelsx
func (b_ BinaryArithmeticNode) PrimaryStrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("primaryStrideInPixelsX"))
	return rv
}/* debug [instance_properties/getter]: primaryStrideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952973-primarystrideinpixelsx
func (b_ BinaryArithmeticNode) SetPrimaryStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryStrideInPixelsX:"), value)
}/* debug [instance_properties/setter]: primaryStrideInPixelsX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952974-secondarystrideinfeaturechannels
func (b_ BinaryArithmeticNode) SecondaryStrideInFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: secondaryStrideInFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952974-secondarystrideinfeaturechannels
func (b_ BinaryArithmeticNode) SetSecondaryStrideInFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}/* debug [instance_properties/setter]: secondaryStrideInFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952976-secondaryscale
func (b_ BinaryArithmeticNode) SecondaryScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("secondaryScale"))
	return rv
}/* debug [instance_properties/getter]: secondaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952976-secondaryscale
func (b_ BinaryArithmeticNode) SetSecondaryScale(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryScale:"), value)
}/* debug [instance_properties/setter]: secondaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952979-maximumvalue
func (b_ BinaryArithmeticNode) MaximumValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("maximumValue"))
	return rv
}/* debug [instance_properties/getter]: maximumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952979-maximumvalue
func (b_ BinaryArithmeticNode) SetMaximumValue(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMaximumValue:"), value)
}/* debug [instance_properties/setter]: maximumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952983-primarystrideinfeaturechannels
func (b_ BinaryArithmeticNode) PrimaryStrideInFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("primaryStrideInFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: primaryStrideInFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952983-primarystrideinfeaturechannels
func (b_ BinaryArithmeticNode) SetPrimaryStrideInFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryStrideInFeatureChannels:"), value)
}/* debug [instance_properties/setter]: primaryStrideInFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952985-secondarystrideinpixelsy
func (b_ BinaryArithmeticNode) SecondaryStrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}/* debug [instance_properties/getter]: secondaryStrideInPixelsY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952985-secondarystrideinpixelsy
func (b_ BinaryArithmeticNode) SetSecondaryStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}/* debug [instance_properties/setter]: secondaryStrideInPixelsY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952996-primarystrideinpixelsy
func (b_ BinaryArithmeticNode) PrimaryStrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("primaryStrideInPixelsY"))
	return rv
}/* debug [instance_properties/getter]: primaryStrideInPixelsY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952996-primarystrideinpixelsy
func (b_ BinaryArithmeticNode) SetPrimaryStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrimaryStrideInPixelsY:"), value)
}/* debug [instance_properties/setter]: primaryStrideInPixelsY */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNBinaryArithmeticNode */


