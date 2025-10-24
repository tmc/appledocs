// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNArithmeticGradient */


/* debug [class_header]: Header for MPSCNNArithmeticGradient */
// The class instance for the [CNNArithmeticGradient] class.
var (
	CNNArithmeticGradientClass     _CNNArithmeticGradientClass
	CNNArithmeticGradientClassOnce sync.Once
)

func getCNNArithmeticGradientClass() _CNNArithmeticGradientClass {
	CNNArithmeticGradientClassOnce.Do(func() {
		CNNArithmeticGradientClass = _CNNArithmeticGradientClass{objc.GetClass("MPSCNNArithmeticGradient")}
	})
	return CNNArithmeticGradientClass
}

type _CNNArithmeticGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNArithmeticGradient */
// An interface definition for the [CNNArithmeticGradient] class.
type ICNNArithmeticGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for CNNArithmeticGradient */
	// properties:
	IsSecondarySourceFilter() objectivec.IObject
	SetIsSecondarySourceFilter(value objectivec.IObject)
	SecondaryStrideInFeatureChannels() objectivec.IObject
	SetSecondaryStrideInFeatureChannels(value objectivec.IObject)
	MaximumValue() objectivec.IObject
	SetMaximumValue(value objectivec.IObject)
	Bias() objectivec.IObject
	SetBias(value objectivec.IObject)
	SecondaryScale() objectivec.IObject
	SetSecondaryScale(value objectivec.IObject)
	MinimumValue() objectivec.IObject
	SetMinimumValue(value objectivec.IObject)
	PrimaryScale() objectivec.IObject
	SetPrimaryScale(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNArithmeticGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNArithmeticGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNArithmeticGradientClass) Alloc() CNNArithmeticGradient {
	rv := objc.Send[CNNArithmeticGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNArithmeticGradientClass) New() CNNArithmeticGradient {
	rv := objc.Send[CNNArithmeticGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNArithmeticGradient) Init() CNNArithmeticGradient {
	rv := objc.Send[CNNArithmeticGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNArithmeticGradient) Autorelease() CNNArithmeticGradient {
	rv := objc.Send[CNNArithmeticGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNArithmeticGradient creates a new CNNArithmeticGradient instance.
func NewCNNArithmeticGradient() CNNArithmeticGradient {
	return getCNNArithmeticGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNArithmeticGradient */
// The base class for gradient arithmetic operators.


// The base class for gradient arithmetic operators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradient
type CNNArithmeticGradient struct {
	CNNGradientKernel
}

// CNNArithmeticGradientFrom constructs a [CNNArithmeticGradient] from an unsafe.Pointer.
//
// The base class for gradient arithmetic operators.
func CNNArithmeticGradientFrom(ptr unsafe.Pointer) CNNArithmeticGradient {
	return CNNArithmeticGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNArithmeticGradient *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNArithmeticGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNArithmeticGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNArithmeticGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNArithmeticGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951852-issecondarysourcefilter
func (c_ CNNArithmeticGradient) IsSecondarySourceFilter() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("isSecondarySourceFilter"))
	return rv
}/* debug [instance_properties/getter]: isSecondarySourceFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951852-issecondarysourcefilter
func (c_ CNNArithmeticGradient) SetIsSecondarySourceFilter(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSecondarySourceFilter:"), value)
}/* debug [instance_properties/setter]: isSecondarySourceFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951853-secondarystrideinfeaturechannels
func (c_ CNNArithmeticGradient) SecondaryStrideInFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: secondaryStrideInFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951853-secondarystrideinfeaturechannels
func (c_ CNNArithmeticGradient) SetSecondaryStrideInFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}/* debug [instance_properties/setter]: secondaryStrideInFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951854-maximumvalue
func (c_ CNNArithmeticGradient) MaximumValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("maximumValue"))
	return rv
}/* debug [instance_properties/getter]: maximumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951854-maximumvalue
func (c_ CNNArithmeticGradient) SetMaximumValue(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumValue:"), value)
}/* debug [instance_properties/setter]: maximumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951855-bias
func (c_ CNNArithmeticGradient) Bias() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("bias"))
	return rv
}/* debug [instance_properties/getter]: bias */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951855-bias
func (c_ CNNArithmeticGradient) SetBias(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBias:"), value)
}/* debug [instance_properties/setter]: bias */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951856-secondaryscale
func (c_ CNNArithmeticGradient) SecondaryScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("secondaryScale"))
	return rv
}/* debug [instance_properties/getter]: secondaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951856-secondaryscale
func (c_ CNNArithmeticGradient) SetSecondaryScale(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryScale:"), value)
}/* debug [instance_properties/setter]: secondaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951858-minimumvalue
func (c_ CNNArithmeticGradient) MinimumValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("minimumValue"))
	return rv
}/* debug [instance_properties/getter]: minimumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951858-minimumvalue
func (c_ CNNArithmeticGradient) SetMinimumValue(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumValue:"), value)
}/* debug [instance_properties/setter]: minimumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951861-primaryscale
func (c_ CNNArithmeticGradient) PrimaryScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("primaryScale"))
	return rv
}/* debug [instance_properties/getter]: primaryScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951861-primaryscale
func (c_ CNNArithmeticGradient) SetPrimaryScale(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryScale:"), value)
}/* debug [instance_properties/setter]: primaryScale */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNArithmeticGradient */



