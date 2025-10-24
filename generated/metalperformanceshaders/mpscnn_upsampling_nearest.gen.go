// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNUpsamplingNearest */


/* debug [class_header]: Header for MPSCNNUpsamplingNearest */
// The class instance for the [CNNUpsamplingNearest] class.
var (
	CNNUpsamplingNearestClass     _CNNUpsamplingNearestClass
	CNNUpsamplingNearestClassOnce sync.Once
)

func getCNNUpsamplingNearestClass() _CNNUpsamplingNearestClass {
	CNNUpsamplingNearestClassOnce.Do(func() {
		CNNUpsamplingNearestClass = _CNNUpsamplingNearestClass{objc.GetClass("MPSCNNUpsamplingNearest")}
	})
	return CNNUpsamplingNearestClass
}

type _CNNUpsamplingNearestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNUpsamplingNearest */
// An interface definition for the [CNNUpsamplingNearest] class.
type ICNNUpsamplingNearest interface {
	ICNNUpsampling
	
/* debug [class_interface_properties]: Properties for CNNUpsamplingNearest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNUpsamplingNearest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNUpsamplingNearest */
// Alloc allocates a new instance without initialization.
func (cc _CNNUpsamplingNearestClass) Alloc() CNNUpsamplingNearest {
	rv := objc.Send[CNNUpsamplingNearest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNUpsamplingNearestClass) New() CNNUpsamplingNearest {
	rv := objc.Send[CNNUpsamplingNearest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNUpsamplingNearest) Init() CNNUpsamplingNearest {
	rv := objc.Send[CNNUpsamplingNearest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNUpsamplingNearest) Autorelease() CNNUpsamplingNearest {
	rv := objc.Send[CNNUpsamplingNearest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNUpsamplingNearest creates a new CNNUpsamplingNearest instance.
func NewCNNUpsamplingNearest() CNNUpsamplingNearest {
	return getCNNUpsamplingNearestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNUpsamplingNearest */
// A nearest spatial upsampling filter.
//
// This filter can be used to resample an existing using a different sampling frequency for the and dimensions with the purpose of enlarging the size of an image. The number of output feature channels remains the same as the number of input feature channels. The must be an integer value . The default value is .


// A nearest spatial upsampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearest
type CNNUpsamplingNearest struct {
	CNNUpsampling
}

// CNNUpsamplingNearestFrom constructs a [CNNUpsamplingNearest] from an unsafe.Pointer.
//
// A nearest spatial upsampling filter.
func CNNUpsamplingNearestFrom(ptr unsafe.Pointer) CNNUpsamplingNearest {
	return CNNUpsamplingNearest{
		CNNUpsampling: CNNUpsamplingFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNUpsamplingNearest */

// Initializes a nearest spatial upsampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearest/2875223-initwithdevice
func NewCNNUpsamplingNearestWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device unsafe.Pointer, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingNearest {
	instance := getCNNUpsamplingNearestClass().Alloc()
	rv := objc.Send[CNNUpsamplingNearest](instance.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), device, integerScaleFactorX, integerScaleFactorY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNUpsamplingNearestWithDeviceIntegerScaleFactorXIntegerScaleFactorY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNUpsamplingNearest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNUpsamplingNearest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNUpsamplingNearest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNUpsamplingNearest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNUpsamplingNearest */


