// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNUpsamplingGradient */


/* debug [class_header]: Header for MPSCNNUpsamplingGradient */
// The class instance for the [CNNUpsamplingGradient] class.
var (
	CNNUpsamplingGradientClass     _CNNUpsamplingGradientClass
	CNNUpsamplingGradientClassOnce sync.Once
)

func getCNNUpsamplingGradientClass() _CNNUpsamplingGradientClass {
	CNNUpsamplingGradientClassOnce.Do(func() {
		CNNUpsamplingGradientClass = _CNNUpsamplingGradientClass{objc.GetClass("MPSCNNUpsamplingGradient")}
	})
	return CNNUpsamplingGradientClass
}

type _CNNUpsamplingGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNUpsamplingGradient */
// An interface definition for the [CNNUpsamplingGradient] class.
type ICNNUpsamplingGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for CNNUpsamplingGradient */
	// properties:
	ScaleFactorY() objectivec.IObject
	SetScaleFactorY(value objectivec.IObject)
	ScaleFactorX() objectivec.IObject
	SetScaleFactorX(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNUpsamplingGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNUpsamplingGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNUpsamplingGradientClass) Alloc() CNNUpsamplingGradient {
	rv := objc.Send[CNNUpsamplingGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNUpsamplingGradientClass) New() CNNUpsamplingGradient {
	rv := objc.Send[CNNUpsamplingGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNUpsamplingGradient) Init() CNNUpsamplingGradient {
	rv := objc.Send[CNNUpsamplingGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNUpsamplingGradient) Autorelease() CNNUpsamplingGradient {
	rv := objc.Send[CNNUpsamplingGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNUpsamplingGradient creates a new CNNUpsamplingGradient instance.
func NewCNNUpsamplingGradient() CNNUpsamplingGradient {
	return getCNNUpsamplingGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNUpsamplingGradient */
// A gradient filter that upsamples an existing Metal Performance Shaders image.


// A gradient filter that upsamples an existing Metal Performance Shaders image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingGradient
type CNNUpsamplingGradient struct {
	CNNGradientKernel
}

// CNNUpsamplingGradientFrom constructs a [CNNUpsamplingGradient] from an unsafe.Pointer.
//
// A gradient filter that upsamples an existing Metal Performance Shaders image.
func CNNUpsamplingGradientFrom(ptr unsafe.Pointer) CNNUpsamplingGradient {
	return CNNUpsamplingGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNUpsamplingGradient *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNUpsamplingGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNUpsamplingGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNUpsamplingGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNUpsamplingGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplinggradient/2942628-scalefactory
func (c_ CNNUpsamplingGradient) ScaleFactorY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorY"))
	return rv
}/* debug [instance_properties/getter]: scaleFactorY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplinggradient/2942628-scalefactory
func (c_ CNNUpsamplingGradient) SetScaleFactorY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorY:"), value)
}/* debug [instance_properties/setter]: scaleFactorY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplinggradient/2942630-scalefactorx
func (c_ CNNUpsamplingGradient) ScaleFactorX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("scaleFactorX"))
	return rv
}/* debug [instance_properties/getter]: scaleFactorX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplinggradient/2942630-scalefactorx
func (c_ CNNUpsamplingGradient) SetScaleFactorX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScaleFactorX:"), value)
}/* debug [instance_properties/setter]: scaleFactorX */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNUpsamplingGradient */



