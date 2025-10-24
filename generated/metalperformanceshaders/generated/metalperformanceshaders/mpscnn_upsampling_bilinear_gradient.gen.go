// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNUpsamplingBilinearGradient */


/* debug [class_header]: Header for MPSCNNUpsamplingBilinearGradient */
// The class instance for the [CNNUpsamplingBilinearGradient] class.
var (
	CNNUpsamplingBilinearGradientClass     _CNNUpsamplingBilinearGradientClass
	CNNUpsamplingBilinearGradientClassOnce sync.Once
)

func getCNNUpsamplingBilinearGradientClass() _CNNUpsamplingBilinearGradientClass {
	CNNUpsamplingBilinearGradientClassOnce.Do(func() {
		CNNUpsamplingBilinearGradientClass = _CNNUpsamplingBilinearGradientClass{objc.GetClass("MPSCNNUpsamplingBilinearGradient")}
	})
	return CNNUpsamplingBilinearGradientClass
}

type _CNNUpsamplingBilinearGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNUpsamplingBilinearGradient */
// An interface definition for the [CNNUpsamplingBilinearGradient] class.
type ICNNUpsamplingBilinearGradient interface {
	ICNNUpsamplingGradient
	
/* debug [class_interface_properties]: Properties for CNNUpsamplingBilinearGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNUpsamplingBilinearGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNUpsamplingBilinearGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNUpsamplingBilinearGradientClass) Alloc() CNNUpsamplingBilinearGradient {
	rv := objc.Send[CNNUpsamplingBilinearGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNUpsamplingBilinearGradientClass) New() CNNUpsamplingBilinearGradient {
	rv := objc.Send[CNNUpsamplingBilinearGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNUpsamplingBilinearGradient) Init() CNNUpsamplingBilinearGradient {
	rv := objc.Send[CNNUpsamplingBilinearGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNUpsamplingBilinearGradient) Autorelease() CNNUpsamplingBilinearGradient {
	rv := objc.Send[CNNUpsamplingBilinearGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNUpsamplingBilinearGradient creates a new CNNUpsamplingBilinearGradient instance.
func NewCNNUpsamplingBilinearGradient() CNNUpsamplingBilinearGradient {
	return getCNNUpsamplingBilinearGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNUpsamplingBilinearGradient */
// A gradient bilinear spatial upsampling filter.


// A gradient bilinear spatial upsampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearGradient
type CNNUpsamplingBilinearGradient struct {
	CNNUpsamplingGradient
}

// CNNUpsamplingBilinearGradientFrom constructs a [CNNUpsamplingBilinearGradient] from an unsafe.Pointer.
//
// A gradient bilinear spatial upsampling filter.
func CNNUpsamplingBilinearGradientFrom(ptr unsafe.Pointer) CNNUpsamplingBilinearGradient {
	return CNNUpsamplingBilinearGradient{
		CNNUpsamplingGradient: CNNUpsamplingGradientFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNUpsamplingBilinearGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradient/2947918-initwithdevice
func NewCNNUpsamplingBilinearGradientWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device unsafe.Pointer, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingBilinearGradient {
	instance := getCNNUpsamplingBilinearGradientClass().Alloc()
	rv := objc.Send[CNNUpsamplingBilinearGradient](instance.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), device, integerScaleFactorX, integerScaleFactorY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNUpsamplingBilinearGradientWithDeviceIntegerScaleFactorXIntegerScaleFactorY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNUpsamplingBilinearGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNUpsamplingBilinearGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNUpsamplingBilinearGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNUpsamplingBilinearGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNUpsamplingBilinearGradient */


