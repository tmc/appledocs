// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSCNNInstanceNormalizationGradientState */


/* debug [class_header]: Header for MPSCNNInstanceNormalizationGradientState */
// The class instance for the [CNNInstanceNormalizationGradientState] class.
var (
	CNNInstanceNormalizationGradientStateClass     _CNNInstanceNormalizationGradientStateClass
	CNNInstanceNormalizationGradientStateClassOnce sync.Once
)

func getCNNInstanceNormalizationGradientStateClass() _CNNInstanceNormalizationGradientStateClass {
	CNNInstanceNormalizationGradientStateClassOnce.Do(func() {
		CNNInstanceNormalizationGradientStateClass = _CNNInstanceNormalizationGradientStateClass{objc.GetClass("MPSCNNInstanceNormalizationGradientState")}
	})
	return CNNInstanceNormalizationGradientStateClass
}

type _CNNInstanceNormalizationGradientStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNInstanceNormalizationGradientState */
// An interface definition for the [CNNInstanceNormalizationGradientState] class.
type ICNNInstanceNormalizationGradientState interface {
	IGradientState
	
/* debug [class_interface_properties]: Properties for CNNInstanceNormalizationGradientState */
	// properties:
	InstanceNormalization() IMPSCNNInstanceNormalization
	SetInstanceNormalization(value IMPSCNNInstanceNormalization)
	GradientForGamma() Buffer get /* not a class type */
	SetGradientForGamma(value Buffer get /* not a class type */)
	GradientForBeta() Buffer get /* not a class type */
	SetGradientForBeta(value Buffer get /* not a class type */)
	Gamma() Buffer get /* not a class type */
	SetGamma(value Buffer get /* not a class type */)
	Beta() Buffer get /* not a class type */
	SetBeta(value Buffer get /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNInstanceNormalizationGradientState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNInstanceNormalizationGradientState */
// Alloc allocates a new instance without initialization.
func (cc _CNNInstanceNormalizationGradientStateClass) Alloc() CNNInstanceNormalizationGradientState {
	rv := objc.Send[CNNInstanceNormalizationGradientState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNInstanceNormalizationGradientStateClass) New() CNNInstanceNormalizationGradientState {
	rv := objc.Send[CNNInstanceNormalizationGradientState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNInstanceNormalizationGradientState) Init() CNNInstanceNormalizationGradientState {
	rv := objc.Send[CNNInstanceNormalizationGradientState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNInstanceNormalizationGradientState) Autorelease() CNNInstanceNormalizationGradientState {
	rv := objc.Send[CNNInstanceNormalizationGradientState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNInstanceNormalizationGradientState creates a new CNNInstanceNormalizationGradientState instance.
func NewCNNInstanceNormalizationGradientState() CNNInstanceNormalizationGradientState {
	return getCNNInstanceNormalizationGradientStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNInstanceNormalizationGradientState */
// An object that stores information required to execute a gradient pass for instance normalization.


// An object that stores information required to execute a gradient pass for instance normalization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientState
type CNNInstanceNormalizationGradientState struct {
	GradientState
}

// CNNInstanceNormalizationGradientStateFrom constructs a [CNNInstanceNormalizationGradientState] from an unsafe.Pointer.
//
// An object that stores information required to execute a gradient pass for instance normalization.
func CNNInstanceNormalizationGradientStateFrom(ptr unsafe.Pointer) CNNInstanceNormalizationGradientState {
	return CNNInstanceNormalizationGradientState{
		GradientState: GradientStateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNInstanceNormalizationGradientState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNInstanceNormalizationGradientState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNInstanceNormalizationGradientState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNInstanceNormalizationGradientState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNInstanceNormalizationGradientState */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953924-instancenormalization
func (c_ CNNInstanceNormalizationGradientState) InstanceNormalization() IMPSCNNInstanceNormalization {
	rv := objc.Send[CNNInstanceNormalization](c_.ID, objc.Sel("instanceNormalization"))
	return rv
}/* debug [instance_properties/getter]: instanceNormalization */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953924-instancenormalization
func (c_ CNNInstanceNormalizationGradientState) SetInstanceNormalization(value IMPSCNNInstanceNormalization) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstanceNormalization:"), value)
}/* debug [instance_properties/setter]: instanceNormalization */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953928-gradientforgamma
func (c_ CNNInstanceNormalizationGradientState) GradientForGamma() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gradientForGamma"))
	return rv
}/* debug [instance_properties/getter]: gradientForGamma */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953928-gradientforgamma
func (c_ CNNInstanceNormalizationGradientState) SetGradientForGamma(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientForGamma:"), value)
}/* debug [instance_properties/setter]: gradientForGamma */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953930-gradientforbeta
func (c_ CNNInstanceNormalizationGradientState) GradientForBeta() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gradientForBeta"))
	return rv
}/* debug [instance_properties/getter]: gradientForBeta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953930-gradientforbeta
func (c_ CNNInstanceNormalizationGradientState) SetGradientForBeta(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientForBeta:"), value)
}/* debug [instance_properties/setter]: gradientForBeta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2956161-gamma
func (c_ CNNInstanceNormalizationGradientState) Gamma() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gamma"))
	return rv
}/* debug [instance_properties/getter]: gamma */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2956161-gamma
func (c_ CNNInstanceNormalizationGradientState) SetGamma(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGamma:"), value)
}/* debug [instance_properties/setter]: gamma */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2956162-beta
func (c_ CNNInstanceNormalizationGradientState) Beta() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("beta"))
	return rv
}/* debug [instance_properties/getter]: beta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2956162-beta
func (c_ CNNInstanceNormalizationGradientState) SetBeta(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}/* debug [instance_properties/setter]: beta */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNInstanceNormalizationGradientState */



