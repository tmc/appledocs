// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNFullyConnectedGradient */


/* debug [class_header]: Header for MPSCNNFullyConnectedGradient */
// The class instance for the [CNNFullyConnectedGradient] class.
var (
	CNNFullyConnectedGradientClass     _CNNFullyConnectedGradientClass
	CNNFullyConnectedGradientClassOnce sync.Once
)

func getCNNFullyConnectedGradientClass() _CNNFullyConnectedGradientClass {
	CNNFullyConnectedGradientClassOnce.Do(func() {
		CNNFullyConnectedGradientClass = _CNNFullyConnectedGradientClass{objc.GetClass("MPSCNNFullyConnectedGradient")}
	})
	return CNNFullyConnectedGradientClass
}

type _CNNFullyConnectedGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNFullyConnectedGradient */
// An interface definition for the [CNNFullyConnectedGradient] class.
type ICNNFullyConnectedGradient interface {
	ICNNConvolutionGradient
	
/* debug [class_interface_properties]: Properties for CNNFullyConnectedGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNFullyConnectedGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNFullyConnectedGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNFullyConnectedGradientClass) Alloc() CNNFullyConnectedGradient {
	rv := objc.Send[CNNFullyConnectedGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNFullyConnectedGradientClass) New() CNNFullyConnectedGradient {
	rv := objc.Send[CNNFullyConnectedGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNFullyConnectedGradient) Init() CNNFullyConnectedGradient {
	rv := objc.Send[CNNFullyConnectedGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNFullyConnectedGradient) Autorelease() CNNFullyConnectedGradient {
	rv := objc.Send[CNNFullyConnectedGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNFullyConnectedGradient creates a new CNNFullyConnectedGradient instance.
func NewCNNFullyConnectedGradient() CNNFullyConnectedGradient {
	return getCNNFullyConnectedGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNFullyConnectedGradient */
// A gradient fully connected convolution layer.


// A gradient fully connected convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedGradient
type CNNFullyConnectedGradient struct {
	CNNConvolutionGradient
}

// CNNFullyConnectedGradientFrom constructs a [CNNFullyConnectedGradient] from an unsafe.Pointer.
//
// A gradient fully connected convolution layer.
func CNNFullyConnectedGradientFrom(ptr unsafe.Pointer) CNNFullyConnectedGradient {
	return CNNFullyConnectedGradient{
		CNNConvolutionGradient: CNNConvolutionGradientFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNFullyConnectedGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectedgradient/2951923-initwithcoder
func NewCNNFullyConnectedGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNFullyConnectedGradient {
	instance := getCNNFullyConnectedGradientClass().Alloc()
	rv := objc.Send[CNNFullyConnectedGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNFullyConnectedGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectedgradient/2951921-initwithdevice
func NewCNNFullyConnectedGradientWithDeviceWeights(device unsafe.Pointer, weights unsafe.Pointer) CNNFullyConnectedGradient {
	instance := getCNNFullyConnectedGradientClass().Alloc()
	rv := objc.Send[CNNFullyConnectedGradient](instance.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNFullyConnectedGradientWithDeviceWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNFullyConnectedGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNFullyConnectedGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNFullyConnectedGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNFullyConnectedGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNFullyConnectedGradient */


