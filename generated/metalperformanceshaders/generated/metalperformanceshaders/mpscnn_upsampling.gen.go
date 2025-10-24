// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNUpsampling */


/* debug [class_header]: Header for MPSCNNUpsampling */
// The class instance for the [CNNUpsampling] class.
var (
	CNNUpsamplingClass     _CNNUpsamplingClass
	CNNUpsamplingClassOnce sync.Once
)

func getCNNUpsamplingClass() _CNNUpsamplingClass {
	CNNUpsamplingClassOnce.Do(func() {
		CNNUpsamplingClass = _CNNUpsamplingClass{objc.GetClass("MPSCNNUpsampling")}
	})
	return CNNUpsamplingClass
}

type _CNNUpsamplingClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNUpsampling */
// An interface definition for the [CNNUpsampling] class.
type ICNNUpsampling interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for CNNUpsampling */
	// properties:
	ScaleFactorY() objectivec.IObject
	SetScaleFactorY(value objectivec.IObject)
	ScaleFactorX() objectivec.IObject
	SetScaleFactorX(value objectivec.IObject)
	AlignCorners() objectivec.IObject
	SetAlignCorners(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNUpsampling */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNUpsampling */
// Alloc allocates a new instance without initialization.
func (cc _CNNUpsamplingClass) Alloc() CNNUpsampling {
	rv := objc.Send[CNNUpsampling](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNUpsamplingClass) New() CNNUpsampling {
	rv := objc.Send[CNNUpsampling](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNUpsampling) Init() CNNUpsampling {
	rv := objc.Send[CNNUpsampling](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNUpsampling) Autorelease() CNNUpsampling {
	rv := objc.Send[CNNUpsampling](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNUpsampling creates a new CNNUpsampling instance.
func NewCNNUpsampling() CNNUpsampling {
	return getCNNUpsamplingClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNUpsampling */
// A filter that resamples an existing MPS image.
//
// This filter can be used to resample an existing using a different sampling frequency for the and dimensions with the purpose of enlarging the size of an image. The number of output feature channels remains the same as the number of input feature channels. The must be an integer value . The default value is . Nearest and bilinear variants are supported.


// A filter that resamples an existing MPS image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsampling
type CNNUpsampling struct {
	CNNKernel
}

// CNNUpsamplingFrom constructs a [CNNUpsampling] from an unsafe.Pointer.
//
// A filter that resamples an existing MPS image.
func CNNUpsamplingFrom(ptr unsafe.Pointer) CNNUpsampling {
	return CNNUpsampling{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNUpsampling *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNUpsampling */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNUpsampling */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNUpsampling */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNUpsampling */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsampling/2875154-scalefactory
func (c_ CNNUpsampling) ScaleFactorY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorY"))
	return rv
}/* debug [instance_properties/getter]: scaleFactorY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsampling/2875154-scalefactory
func (c_ CNNUpsampling) SetScaleFactorY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorY:"), value)
}/* debug [instance_properties/setter]: scaleFactorY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsampling/2875206-scalefactorx
func (c_ CNNUpsampling) ScaleFactorX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorX"))
	return rv
}/* debug [instance_properties/getter]: scaleFactorX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsampling/2875206-scalefactorx
func (c_ CNNUpsampling) SetScaleFactorX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorX:"), value)
}/* debug [instance_properties/setter]: scaleFactorX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsampling/2966660-aligncorners
func (c_ CNNUpsampling) AlignCorners() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("alignCorners"))
	return rv
}/* debug [instance_properties/getter]: alignCorners */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsampling/2966660-aligncorners
func (c_ CNNUpsampling) SetAlignCorners(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlignCorners:"), value)
}/* debug [instance_properties/setter]: alignCorners */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNUpsampling */



