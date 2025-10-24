// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNDropoutGradient */


/* debug [class_header]: Header for MPSCNNDropoutGradient */
// The class instance for the [CNNDropoutGradient] class.
var (
	CNNDropoutGradientClass     _CNNDropoutGradientClass
	CNNDropoutGradientClassOnce sync.Once
)

func getCNNDropoutGradientClass() _CNNDropoutGradientClass {
	CNNDropoutGradientClassOnce.Do(func() {
		CNNDropoutGradientClass = _CNNDropoutGradientClass{objc.GetClass("MPSCNNDropoutGradient")}
	})
	return CNNDropoutGradientClass
}

type _CNNDropoutGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNDropoutGradient */
// An interface definition for the [CNNDropoutGradient] class.
type ICNNDropoutGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for CNNDropoutGradient */
	// properties:
	MaskStrideInPixels() Size get /* not a class type */
	SetMaskStrideInPixels(value Size get /* not a class type */)
	KeepProbability() objectivec.IObject
	SetKeepProbability(value objectivec.IObject)
	Seed() objectivec.IObject
	SetSeed(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNDropoutGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNDropoutGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNDropoutGradientClass) Alloc() CNNDropoutGradient {
	rv := objc.Send[CNNDropoutGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNDropoutGradientClass) New() CNNDropoutGradient {
	rv := objc.Send[CNNDropoutGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNDropoutGradient) Init() CNNDropoutGradient {
	rv := objc.Send[CNNDropoutGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNDropoutGradient) Autorelease() CNNDropoutGradient {
	rv := objc.Send[CNNDropoutGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNDropoutGradient creates a new CNNDropoutGradient instance.
func NewCNNDropoutGradient() CNNDropoutGradient {
	return getCNNDropoutGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNDropoutGradient */
// A gradient dropout filter.


// A gradient dropout filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradient
type CNNDropoutGradient struct {
	CNNGradientKernel
}

// CNNDropoutGradientFrom constructs a [CNNDropoutGradient] from an unsafe.Pointer.
//
// A gradient dropout filter.
func CNNDropoutGradientFrom(ptr unsafe.Pointer) CNNDropoutGradient {
	return CNNDropoutGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNDropoutGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradient/2942521-initwithcoder
func NewCNNDropoutGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNDropoutGradient {
	instance := getCNNDropoutGradientClass().Alloc()
	rv := objc.Send[CNNDropoutGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNDropoutGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradient/2942518-initwithdevice
func NewCNNDropoutGradientWithDeviceKeepProbabilitySeedMaskStrideInPixels(device unsafe.Pointer, keepProbability float32, seed uint, maskStrideInPixels objc.IObject /* cross-framework: MTLSize */) CNNDropoutGradient {
	instance := getCNNDropoutGradientClass().Alloc()
	rv := objc.Send[CNNDropoutGradient](instance.ID, objc.Sel("initWithDevice:keepProbability:seed:maskStrideInPixels:"), device, keepProbability, seed, maskStrideInPixels)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNDropoutGradientWithDeviceKeepProbabilitySeedMaskStrideInPixels */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNDropoutGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNDropoutGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNDropoutGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNDropoutGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradient/2942515-maskstrideinpixels
func (c_ CNNDropoutGradient) MaskStrideInPixels() Size get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("maskStrideInPixels"))
	return rv
}/* debug [instance_properties/getter]: maskStrideInPixels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradient/2942515-maskstrideinpixels
func (c_ CNNDropoutGradient) SetMaskStrideInPixels(value Size get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaskStrideInPixels:"), value)
}/* debug [instance_properties/setter]: maskStrideInPixels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradient/2942520-keepprobability
func (c_ CNNDropoutGradient) KeepProbability() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("keepProbability"))
	return rv
}/* debug [instance_properties/getter]: keepProbability */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradient/2942520-keepprobability
func (c_ CNNDropoutGradient) SetKeepProbability(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeepProbability:"), value)
}/* debug [instance_properties/setter]: keepProbability */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradient/2942528-seed
func (c_ CNNDropoutGradient) Seed() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("seed"))
	return rv
}/* debug [instance_properties/getter]: seed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradient/2942528-seed
func (c_ CNNDropoutGradient) SetSeed(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSeed:"), value)
}/* debug [instance_properties/setter]: seed */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNDropoutGradient */


