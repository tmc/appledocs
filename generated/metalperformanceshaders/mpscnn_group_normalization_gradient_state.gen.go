// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSCNNGroupNormalizationGradientState */


/* debug [class_header]: Header for MPSCNNGroupNormalizationGradientState */
// The class instance for the [CNNGroupNormalizationGradientState] class.
var (
	CNNGroupNormalizationGradientStateClass     _CNNGroupNormalizationGradientStateClass
	CNNGroupNormalizationGradientStateClassOnce sync.Once
)

func getCNNGroupNormalizationGradientStateClass() _CNNGroupNormalizationGradientStateClass {
	CNNGroupNormalizationGradientStateClassOnce.Do(func() {
		CNNGroupNormalizationGradientStateClass = _CNNGroupNormalizationGradientStateClass{objc.GetClass("MPSCNNGroupNormalizationGradientState")}
	})
	return CNNGroupNormalizationGradientStateClass
}

type _CNNGroupNormalizationGradientStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNGroupNormalizationGradientState */
// An interface definition for the [CNNGroupNormalizationGradientState] class.
type ICNNGroupNormalizationGradientState interface {
	IGradientState
	
/* debug [class_interface_properties]: Properties for CNNGroupNormalizationGradientState */
	// properties:
	Beta() Buffer get /* not a class type */
	SetBeta(value Buffer get /* not a class type */)
	Gamma() Buffer get /* not a class type */
	SetGamma(value Buffer get /* not a class type */)
	GradientForBeta() Buffer get /* not a class type */
	SetGradientForBeta(value Buffer get /* not a class type */)
	GradientForGamma() Buffer get /* not a class type */
	SetGradientForGamma(value Buffer get /* not a class type */)
	GroupNormalization() IMPSCNNGroupNormalization
	SetGroupNormalization(value IMPSCNNGroupNormalization)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNGroupNormalizationGradientState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNGroupNormalizationGradientState */
// Alloc allocates a new instance without initialization.
func (cc _CNNGroupNormalizationGradientStateClass) Alloc() CNNGroupNormalizationGradientState {
	rv := objc.Send[CNNGroupNormalizationGradientState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNGroupNormalizationGradientStateClass) New() CNNGroupNormalizationGradientState {
	rv := objc.Send[CNNGroupNormalizationGradientState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNGroupNormalizationGradientState) Init() CNNGroupNormalizationGradientState {
	rv := objc.Send[CNNGroupNormalizationGradientState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNGroupNormalizationGradientState) Autorelease() CNNGroupNormalizationGradientState {
	rv := objc.Send[CNNGroupNormalizationGradientState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNGroupNormalizationGradientState creates a new CNNGroupNormalizationGradientState instance.
func NewCNNGroupNormalizationGradientState() CNNGroupNormalizationGradientState {
	return getCNNGroupNormalizationGradientStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNGroupNormalizationGradientState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientState
type CNNGroupNormalizationGradientState struct {
	GradientState
}

// CNNGroupNormalizationGradientStateFrom constructs a [CNNGroupNormalizationGradientState] from an unsafe.Pointer.
func CNNGroupNormalizationGradientStateFrom(ptr unsafe.Pointer) CNNGroupNormalizationGradientState {
	return CNNGroupNormalizationGradientState{
		GradientState: GradientStateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNGroupNormalizationGradientState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNGroupNormalizationGradientState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNGroupNormalizationGradientState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNGroupNormalizationGradientState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNGroupNormalizationGradientState */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152557-beta
func (c_ CNNGroupNormalizationGradientState) Beta() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("beta"))
	return rv
}/* debug [instance_properties/getter]: beta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152557-beta
func (c_ CNNGroupNormalizationGradientState) SetBeta(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}/* debug [instance_properties/setter]: beta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152558-gamma
func (c_ CNNGroupNormalizationGradientState) Gamma() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gamma"))
	return rv
}/* debug [instance_properties/getter]: gamma */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152558-gamma
func (c_ CNNGroupNormalizationGradientState) SetGamma(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGamma:"), value)
}/* debug [instance_properties/setter]: gamma */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152559-gradientforbeta
func (c_ CNNGroupNormalizationGradientState) GradientForBeta() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gradientForBeta"))
	return rv
}/* debug [instance_properties/getter]: gradientForBeta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152559-gradientforbeta
func (c_ CNNGroupNormalizationGradientState) SetGradientForBeta(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientForBeta:"), value)
}/* debug [instance_properties/setter]: gradientForBeta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152560-gradientforgamma
func (c_ CNNGroupNormalizationGradientState) GradientForGamma() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gradientForGamma"))
	return rv
}/* debug [instance_properties/getter]: gradientForGamma */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152560-gradientforgamma
func (c_ CNNGroupNormalizationGradientState) SetGradientForGamma(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientForGamma:"), value)
}/* debug [instance_properties/setter]: gradientForGamma */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152561-groupnormalization
func (c_ CNNGroupNormalizationGradientState) GroupNormalization() IMPSCNNGroupNormalization {
	rv := objc.Send[CNNGroupNormalization](c_.ID, objc.Sel("groupNormalization"))
	return rv
}/* debug [instance_properties/getter]: groupNormalization */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152561-groupnormalization
func (c_ CNNGroupNormalizationGradientState) SetGroupNormalization(value IMPSCNNGroupNormalization) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroupNormalization:"), value)
}/* debug [instance_properties/setter]: groupNormalization */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNGroupNormalizationGradientState */



