// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNUpsamplingBilinear */


/* debug [class_header]: Header for MPSCNNUpsamplingBilinear */
// The class instance for the [CNNUpsamplingBilinear] class.
var (
	CNNUpsamplingBilinearClass     _CNNUpsamplingBilinearClass
	CNNUpsamplingBilinearClassOnce sync.Once
)

func getCNNUpsamplingBilinearClass() _CNNUpsamplingBilinearClass {
	CNNUpsamplingBilinearClassOnce.Do(func() {
		CNNUpsamplingBilinearClass = _CNNUpsamplingBilinearClass{objc.GetClass("MPSCNNUpsamplingBilinear")}
	})
	return CNNUpsamplingBilinearClass
}

type _CNNUpsamplingBilinearClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNUpsamplingBilinear */
// An interface definition for the [CNNUpsamplingBilinear] class.
type ICNNUpsamplingBilinear interface {
	ICNNUpsampling
	
/* debug [class_interface_properties]: Properties for CNNUpsamplingBilinear */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNUpsamplingBilinear */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNUpsamplingBilinear */
// Alloc allocates a new instance without initialization.
func (cc _CNNUpsamplingBilinearClass) Alloc() CNNUpsamplingBilinear {
	rv := objc.Send[CNNUpsamplingBilinear](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNUpsamplingBilinearClass) New() CNNUpsamplingBilinear {
	rv := objc.Send[CNNUpsamplingBilinear](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNUpsamplingBilinear) Init() CNNUpsamplingBilinear {
	rv := objc.Send[CNNUpsamplingBilinear](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNUpsamplingBilinear) Autorelease() CNNUpsamplingBilinear {
	rv := objc.Send[CNNUpsamplingBilinear](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNUpsamplingBilinear creates a new CNNUpsamplingBilinear instance.
func NewCNNUpsamplingBilinear() CNNUpsamplingBilinear {
	return getCNNUpsamplingBilinearClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNUpsamplingBilinear */
// A bilinear spatial upsampling filter.
//
// This filter can be used to resample an existing using a different sampling frequency for the and dimensions with the purpose of enlarging the size of an image. The number of output feature channels remains the same as the number of input feature channels. The must be an integer value . The default value is .


// A bilinear spatial upsampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinear
type CNNUpsamplingBilinear struct {
	CNNUpsampling
}

// CNNUpsamplingBilinearFrom constructs a [CNNUpsamplingBilinear] from an unsafe.Pointer.
//
// A bilinear spatial upsampling filter.
func CNNUpsamplingBilinearFrom(ptr unsafe.Pointer) CNNUpsamplingBilinear {
	return CNNUpsamplingBilinear{
		CNNUpsampling: CNNUpsamplingFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNUpsamplingBilinear */

// Initializes a bilinear spatial upsampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinear/2875160-initwithdevice
func NewCNNUpsamplingBilinearWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device unsafe.Pointer, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingBilinear {
	instance := getCNNUpsamplingBilinearClass().Alloc()
	rv := objc.Send[CNNUpsamplingBilinear](instance.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), device, integerScaleFactorX, integerScaleFactorY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNUpsamplingBilinearWithDeviceIntegerScaleFactorXIntegerScaleFactorY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinear/2966661-initwithdevice
func NewCNNUpsamplingBilinearWithDeviceIntegerScaleFactorXIntegerScaleFactorYAlignCorners(device unsafe.Pointer, integerScaleFactorX uint, integerScaleFactorY uint, alignCorners bool) CNNUpsamplingBilinear {
	instance := getCNNUpsamplingBilinearClass().Alloc()
	rv := objc.Send[CNNUpsamplingBilinear](instance.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:alignCorners:"), device, integerScaleFactorX, integerScaleFactorY, alignCorners)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNUpsamplingBilinearWithDeviceIntegerScaleFactorXIntegerScaleFactorYAlignCorners */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNUpsamplingBilinear */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNUpsamplingBilinear */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNUpsamplingBilinear */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNUpsamplingBilinear */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNUpsamplingBilinear */


