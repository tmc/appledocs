// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNCrossChannelNormalizationGradient */


/* debug [class_header]: Header for MPSCNNCrossChannelNormalizationGradient */
// The class instance for the [CNNCrossChannelNormalizationGradient] class.
var (
	CNNCrossChannelNormalizationGradientClass     _CNNCrossChannelNormalizationGradientClass
	CNNCrossChannelNormalizationGradientClassOnce sync.Once
)

func getCNNCrossChannelNormalizationGradientClass() _CNNCrossChannelNormalizationGradientClass {
	CNNCrossChannelNormalizationGradientClassOnce.Do(func() {
		CNNCrossChannelNormalizationGradientClass = _CNNCrossChannelNormalizationGradientClass{objc.GetClass("MPSCNNCrossChannelNormalizationGradient")}
	})
	return CNNCrossChannelNormalizationGradientClass
}

type _CNNCrossChannelNormalizationGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNCrossChannelNormalizationGradient */
// An interface definition for the [CNNCrossChannelNormalizationGradient] class.
type ICNNCrossChannelNormalizationGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for CNNCrossChannelNormalizationGradient */
	// properties:
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	KernelSize() objectivec.IObject
	SetKernelSize(value objectivec.IObject)
	Beta() objectivec.IObject
	SetBeta(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNCrossChannelNormalizationGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNCrossChannelNormalizationGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNCrossChannelNormalizationGradientClass) Alloc() CNNCrossChannelNormalizationGradient {
	rv := objc.Send[CNNCrossChannelNormalizationGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNCrossChannelNormalizationGradientClass) New() CNNCrossChannelNormalizationGradient {
	rv := objc.Send[CNNCrossChannelNormalizationGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNCrossChannelNormalizationGradient) Init() CNNCrossChannelNormalizationGradient {
	rv := objc.Send[CNNCrossChannelNormalizationGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNCrossChannelNormalizationGradient) Autorelease() CNNCrossChannelNormalizationGradient {
	rv := objc.Send[CNNCrossChannelNormalizationGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNCrossChannelNormalizationGradient creates a new CNNCrossChannelNormalizationGradient instance.
func NewCNNCrossChannelNormalizationGradient() CNNCrossChannelNormalizationGradient {
	return getCNNCrossChannelNormalizationGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNCrossChannelNormalizationGradient */
// A gradient normalization kernel applied across feature channels.


// A gradient normalization kernel applied across feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradient
type CNNCrossChannelNormalizationGradient struct {
	CNNGradientKernel
}

// CNNCrossChannelNormalizationGradientFrom constructs a [CNNCrossChannelNormalizationGradient] from an unsafe.Pointer.
//
// A gradient normalization kernel applied across feature channels.
func CNNCrossChannelNormalizationGradientFrom(ptr unsafe.Pointer) CNNCrossChannelNormalizationGradient {
	return CNNCrossChannelNormalizationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNCrossChannelNormalizationGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942476-initwithcoder
func NewCNNCrossChannelNormalizationGradientWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) CNNCrossChannelNormalizationGradient {
	instance := getCNNCrossChannelNormalizationGradientClass().Alloc()
	rv := objc.Send[CNNCrossChannelNormalizationGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNCrossChannelNormalizationGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942463-initwithdevice
func NewCNNCrossChannelNormalizationGradientWithDeviceKernelSize(device unsafe.Pointer, kernelSize uint) CNNCrossChannelNormalizationGradient {
	instance := getCNNCrossChannelNormalizationGradientClass().Alloc()
	rv := objc.Send[CNNCrossChannelNormalizationGradient](instance.ID, objc.Sel("initWithDevice:kernelSize:"), device, kernelSize)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNCrossChannelNormalizationGradientWithDeviceKernelSize */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNCrossChannelNormalizationGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNCrossChannelNormalizationGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNCrossChannelNormalizationGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNCrossChannelNormalizationGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942464-alpha
func (c_ CNNCrossChannelNormalizationGradient) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942464-alpha
func (c_ CNNCrossChannelNormalizationGradient) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}/* debug [instance_properties/setter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942465-delta
func (c_ CNNCrossChannelNormalizationGradient) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delta"))
	return rv
}/* debug [instance_properties/getter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942465-delta
func (c_ CNNCrossChannelNormalizationGradient) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelta:"), value)
}/* debug [instance_properties/setter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942468-kernelsize
func (c_ CNNCrossChannelNormalizationGradient) KernelSize() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelSize"))
	return rv
}/* debug [instance_properties/getter]: kernelSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942468-kernelsize
func (c_ CNNCrossChannelNormalizationGradient) SetKernelSize(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelSize:"), value)
}/* debug [instance_properties/setter]: kernelSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942477-beta
func (c_ CNNCrossChannelNormalizationGradient) Beta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("beta"))
	return rv
}/* debug [instance_properties/getter]: beta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942477-beta
func (c_ CNNCrossChannelNormalizationGradient) SetBeta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}/* debug [instance_properties/setter]: beta */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNCrossChannelNormalizationGradient */


