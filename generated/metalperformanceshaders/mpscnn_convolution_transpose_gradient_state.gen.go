// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSCNNConvolutionTransposeGradientState */


/* debug [class_header]: Header for MPSCNNConvolutionTransposeGradientState */
// The class instance for the [CNNConvolutionTransposeGradientState] class.
var (
	CNNConvolutionTransposeGradientStateClass     _CNNConvolutionTransposeGradientStateClass
	CNNConvolutionTransposeGradientStateClassOnce sync.Once
)

func getCNNConvolutionTransposeGradientStateClass() _CNNConvolutionTransposeGradientStateClass {
	CNNConvolutionTransposeGradientStateClassOnce.Do(func() {
		CNNConvolutionTransposeGradientStateClass = _CNNConvolutionTransposeGradientStateClass{objc.GetClass("MPSCNNConvolutionTransposeGradientState")}
	})
	return CNNConvolutionTransposeGradientStateClass
}

type _CNNConvolutionTransposeGradientStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNConvolutionTransposeGradientState */
// An interface definition for the [CNNConvolutionTransposeGradientState] class.
type ICNNConvolutionTransposeGradientState interface {
	ICNNConvolutionGradientState
	
/* debug [class_interface_properties]: Properties for CNNConvolutionTransposeGradientState */
	// properties:
	ConvolutionTranspose() IMPSCNNConvolutionTranspose
	SetConvolutionTranspose(value IMPSCNNConvolutionTranspose)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNConvolutionTransposeGradientState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNConvolutionTransposeGradientState */
// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionTransposeGradientStateClass) Alloc() CNNConvolutionTransposeGradientState {
	rv := objc.Send[CNNConvolutionTransposeGradientState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionTransposeGradientStateClass) New() CNNConvolutionTransposeGradientState {
	rv := objc.Send[CNNConvolutionTransposeGradientState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionTransposeGradientState) Init() CNNConvolutionTransposeGradientState {
	rv := objc.Send[CNNConvolutionTransposeGradientState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionTransposeGradientState) Autorelease() CNNConvolutionTransposeGradientState {
	rv := objc.Send[CNNConvolutionTransposeGradientState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionTransposeGradientState creates a new CNNConvolutionTransposeGradientState instance.
func NewCNNConvolutionTransposeGradientState() CNNConvolutionTransposeGradientState {
	return getCNNConvolutionTransposeGradientStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNConvolutionTransposeGradientState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientState
type CNNConvolutionTransposeGradientState struct {
	CNNConvolutionGradientState
}

// CNNConvolutionTransposeGradientStateFrom constructs a [CNNConvolutionTransposeGradientState] from an unsafe.Pointer.
func CNNConvolutionTransposeGradientStateFrom(ptr unsafe.Pointer) CNNConvolutionTransposeGradientState {
	return CNNConvolutionTransposeGradientState{
		CNNConvolutionGradientState: CNNConvolutionGradientStateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNConvolutionTransposeGradientState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNConvolutionTransposeGradientState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNConvolutionTransposeGradientState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNConvolutionTransposeGradientState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNConvolutionTransposeGradientState */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradientstate/3131790-convolutiontranspose
func (c_ CNNConvolutionTransposeGradientState) ConvolutionTranspose() IMPSCNNConvolutionTranspose {
	rv := objc.Send[CNNConvolutionTranspose](c_.ID, objc.Sel("convolutionTranspose"))
	return rv
}/* debug [instance_properties/getter]: convolutionTranspose */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradientstate/3131790-convolutiontranspose
func (c_ CNNConvolutionTransposeGradientState) SetConvolutionTranspose(value IMPSCNNConvolutionTranspose) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConvolutionTranspose:"), value)
}/* debug [instance_properties/setter]: convolutionTranspose */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNConvolutionTransposeGradientState */



