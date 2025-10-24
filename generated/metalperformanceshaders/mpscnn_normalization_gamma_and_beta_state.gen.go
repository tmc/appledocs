// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNormalizationGammaAndBetaState */


/* debug [class_header]: Header for MPSCNNNormalizationGammaAndBetaState */
// The class instance for the [CNNNormalizationGammaAndBetaState] class.
var (
	CNNNormalizationGammaAndBetaStateClass     _CNNNormalizationGammaAndBetaStateClass
	CNNNormalizationGammaAndBetaStateClassOnce sync.Once
)

func getCNNNormalizationGammaAndBetaStateClass() _CNNNormalizationGammaAndBetaStateClass {
	CNNNormalizationGammaAndBetaStateClassOnce.Do(func() {
		CNNNormalizationGammaAndBetaStateClass = _CNNNormalizationGammaAndBetaStateClass{objc.GetClass("MPSCNNNormalizationGammaAndBetaState")}
	})
	return CNNNormalizationGammaAndBetaStateClass
}

type _CNNNormalizationGammaAndBetaStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNormalizationGammaAndBetaState */
// An interface definition for the [CNNNormalizationGammaAndBetaState] class.
type ICNNNormalizationGammaAndBetaState interface {
	IState
	
/* debug [class_interface_properties]: Properties for CNNNormalizationGammaAndBetaState */
	// properties:
	Gamma() Buffer get /* not a class type */
	SetGamma(value Buffer get /* not a class type */)
	Beta() Buffer get /* not a class type */
	SetBeta(value Buffer get /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNormalizationGammaAndBetaState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNormalizationGammaAndBetaState */
// Alloc allocates a new instance without initialization.
func (cc _CNNNormalizationGammaAndBetaStateClass) Alloc() CNNNormalizationGammaAndBetaState {
	rv := objc.Send[CNNNormalizationGammaAndBetaState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNormalizationGammaAndBetaStateClass) New() CNNNormalizationGammaAndBetaState {
	rv := objc.Send[CNNNormalizationGammaAndBetaState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNormalizationGammaAndBetaState) Init() CNNNormalizationGammaAndBetaState {
	rv := objc.Send[CNNNormalizationGammaAndBetaState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNormalizationGammaAndBetaState) Autorelease() CNNNormalizationGammaAndBetaState {
	rv := objc.Send[CNNNormalizationGammaAndBetaState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNormalizationGammaAndBetaState creates a new CNNNormalizationGammaAndBetaState instance.
func NewCNNNormalizationGammaAndBetaState() CNNNormalizationGammaAndBetaState {
	return getCNNNormalizationGammaAndBetaStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNormalizationGammaAndBetaState */
// An object that stores gamma and beta terms used to apply a scale and bias in instance- or batch-normalization operations.


// An object that stores gamma and beta terms used to apply a scale and bias in instance- or batch-normalization operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationGammaAndBetaState
type CNNNormalizationGammaAndBetaState struct {
	State
}

// CNNNormalizationGammaAndBetaStateFrom constructs a [CNNNormalizationGammaAndBetaState] from an unsafe.Pointer.
//
// An object that stores gamma and beta terms used to apply a scale and bias in instance- or batch-normalization operations.
func CNNNormalizationGammaAndBetaStateFrom(ptr unsafe.Pointer) CNNNormalizationGammaAndBetaState {
	return CNNNormalizationGammaAndBetaState{
		State: StateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNormalizationGammaAndBetaState */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953936-initwithgamma
func NewCNNNormalizationGammaAndBetaStateWithGammaBeta(gamma unsafe.Pointer, beta unsafe.Pointer) CNNNormalizationGammaAndBetaState {
	instance := getCNNNormalizationGammaAndBetaStateClass().Alloc()
	rv := objc.Send[CNNNormalizationGammaAndBetaState](instance.ID, objc.Sel("initWithGamma:beta:"), gamma, beta)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNormalizationGammaAndBetaStateWithGammaBeta */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNormalizationGammaAndBetaState */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953937-temporarystate
func (cc _CNNNormalizationGammaAndBetaStateClass) TemporaryState() {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("temporaryState"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TemporaryState) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953937-temporarystatewithcommandbuffer
func (cc _CNNNormalizationGammaAndBetaStateClass) TemporaryStateWithCommandBufferNumberOfFeatureChannels(commandBuffer unsafe.Pointer, numberOfFeatureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("temporaryStateWithCommandBuffer:numberOfFeatureChannels:"), commandBuffer, numberOfFeatureChannels)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TemporaryStateWithCommandBufferNumberOfFeatureChannels) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNormalizationGammaAndBetaState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNormalizationGammaAndBetaState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNormalizationGammaAndBetaState */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953934-gamma
func (c_ CNNNormalizationGammaAndBetaState) Gamma() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gamma"))
	return rv
}/* debug [instance_properties/getter]: gamma */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953934-gamma
func (c_ CNNNormalizationGammaAndBetaState) SetGamma(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGamma:"), value)
}/* debug [instance_properties/setter]: gamma */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953938-beta
func (c_ CNNNormalizationGammaAndBetaState) Beta() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("beta"))
	return rv
}/* debug [instance_properties/getter]: beta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953938-beta
func (c_ CNNNormalizationGammaAndBetaState) SetBeta(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}/* debug [instance_properties/setter]: beta */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNormalizationGammaAndBetaState */


