// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNUpsamplingNearestGradient */


/* debug [class_header]: Header for MPSCNNUpsamplingNearestGradient */
// The class instance for the [CNNUpsamplingNearestGradient] class.
var (
	CNNUpsamplingNearestGradientClass     _CNNUpsamplingNearestGradientClass
	CNNUpsamplingNearestGradientClassOnce sync.Once
)

func getCNNUpsamplingNearestGradientClass() _CNNUpsamplingNearestGradientClass {
	CNNUpsamplingNearestGradientClassOnce.Do(func() {
		CNNUpsamplingNearestGradientClass = _CNNUpsamplingNearestGradientClass{objc.GetClass("MPSCNNUpsamplingNearestGradient")}
	})
	return CNNUpsamplingNearestGradientClass
}

type _CNNUpsamplingNearestGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNUpsamplingNearestGradient */
// An interface definition for the [CNNUpsamplingNearestGradient] class.
type ICNNUpsamplingNearestGradient interface {
	ICNNUpsamplingGradient
	
/* debug [class_interface_properties]: Properties for CNNUpsamplingNearestGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNUpsamplingNearestGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNUpsamplingNearestGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNUpsamplingNearestGradientClass) Alloc() CNNUpsamplingNearestGradient {
	rv := objc.Send[CNNUpsamplingNearestGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNUpsamplingNearestGradientClass) New() CNNUpsamplingNearestGradient {
	rv := objc.Send[CNNUpsamplingNearestGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNUpsamplingNearestGradient) Init() CNNUpsamplingNearestGradient {
	rv := objc.Send[CNNUpsamplingNearestGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNUpsamplingNearestGradient) Autorelease() CNNUpsamplingNearestGradient {
	rv := objc.Send[CNNUpsamplingNearestGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNUpsamplingNearestGradient creates a new CNNUpsamplingNearestGradient instance.
func NewCNNUpsamplingNearestGradient() CNNUpsamplingNearestGradient {
	return getCNNUpsamplingNearestGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNUpsamplingNearestGradient */
// A gradient upsampling filter that samples the pixel nearest to the source when upsampling to the destination pixel.


// A gradient upsampling filter that samples the pixel nearest to the source when upsampling to the destination pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestGradient
type CNNUpsamplingNearestGradient struct {
	CNNUpsamplingGradient
}

// CNNUpsamplingNearestGradientFrom constructs a [CNNUpsamplingNearestGradient] from an unsafe.Pointer.
//
// A gradient upsampling filter that samples the pixel nearest to the source when upsampling to the destination pixel.
func CNNUpsamplingNearestGradientFrom(ptr unsafe.Pointer) CNNUpsamplingNearestGradient {
	return CNNUpsamplingNearestGradient{
		CNNUpsamplingGradient: CNNUpsamplingGradientFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNUpsamplingNearestGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradient/2947920-initwithdevice
func NewCNNUpsamplingNearestGradientWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device unsafe.Pointer, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingNearestGradient {
	instance := getCNNUpsamplingNearestGradientClass().Alloc()
	rv := objc.Send[CNNUpsamplingNearestGradient](instance.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), device, integerScaleFactorX, integerScaleFactorY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNUpsamplingNearestGradientWithDeviceIntegerScaleFactorXIntegerScaleFactorY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNUpsamplingNearestGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNUpsamplingNearestGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNUpsamplingNearestGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNUpsamplingNearestGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNUpsamplingNearestGradient */


