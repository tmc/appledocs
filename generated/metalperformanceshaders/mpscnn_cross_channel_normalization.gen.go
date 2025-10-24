// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNCrossChannelNormalization */


/* debug [class_header]: Header for MPSCNNCrossChannelNormalization */
// The class instance for the [CNNCrossChannelNormalization] class.
var (
	CNNCrossChannelNormalizationClass     _CNNCrossChannelNormalizationClass
	CNNCrossChannelNormalizationClassOnce sync.Once
)

func getCNNCrossChannelNormalizationClass() _CNNCrossChannelNormalizationClass {
	CNNCrossChannelNormalizationClassOnce.Do(func() {
		CNNCrossChannelNormalizationClass = _CNNCrossChannelNormalizationClass{objc.GetClass("MPSCNNCrossChannelNormalization")}
	})
	return CNNCrossChannelNormalizationClass
}

type _CNNCrossChannelNormalizationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNCrossChannelNormalization */
// An interface definition for the [CNNCrossChannelNormalization] class.
type ICNNCrossChannelNormalization interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for CNNCrossChannelNormalization */
	// properties:
	KernelSize() objectivec.IObject
	SetKernelSize(value objectivec.IObject)
	Beta() objectivec.IObject
	SetBeta(value objectivec.IObject)
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNCrossChannelNormalization */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNCrossChannelNormalization */
// Alloc allocates a new instance without initialization.
func (cc _CNNCrossChannelNormalizationClass) Alloc() CNNCrossChannelNormalization {
	rv := objc.Send[CNNCrossChannelNormalization](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNCrossChannelNormalizationClass) New() CNNCrossChannelNormalization {
	rv := objc.Send[CNNCrossChannelNormalization](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNCrossChannelNormalization) Init() CNNCrossChannelNormalization {
	rv := objc.Send[CNNCrossChannelNormalization](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNCrossChannelNormalization) Autorelease() CNNCrossChannelNormalization {
	rv := objc.Send[CNNCrossChannelNormalization](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNCrossChannelNormalization creates a new CNNCrossChannelNormalization instance.
func NewCNNCrossChannelNormalization() CNNCrossChannelNormalization {
	return getCNNCrossChannelNormalizationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNCrossChannelNormalization */
// A normalization kernel applied across feature channels.
//
// The normalization kernel applies the kernel to a local region across nearby feature channels, but with no spatial extent (i.e., they have the shape ). The normalized output is given by the function: Where the normalizing factor is: Where is the kernel size. The window itself is defined as: Where is the feature channel index (running from 0 to ) and is the number of feature channels, and the values of , , and are set via properties. It is your responsibility to ensure that the combination of the values of the and properties does not result in a situation where the denominator becomes zero - in such situations the resulting pixel-value is undefined.


// A normalization kernel applied across feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalization
type CNNCrossChannelNormalization struct {
	CNNKernel
}

// CNNCrossChannelNormalizationFrom constructs a [CNNCrossChannelNormalization] from an unsafe.Pointer.
//
// A normalization kernel applied across feature channels.
func CNNCrossChannelNormalizationFrom(ptr unsafe.Pointer) CNNCrossChannelNormalization {
	return CNNCrossChannelNormalization{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNCrossChannelNormalization */

// Initializes a normalization kernel in a channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalization/init(coder:device:)
func NewCNNCrossChannelNormalizationWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNCrossChannelNormalization {
	instance := getCNNCrossChannelNormalizationClass().Alloc()
	rv := objc.Send[CNNCrossChannelNormalization](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNCrossChannelNormalizationWithCoderDevice */


// Initializes a normalization kernel in a channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalization/init(device:kernelSize:)
func NewCNNCrossChannelNormalizationWithDeviceKernelSize(device unsafe.Pointer, kernelSize uint) CNNCrossChannelNormalization {
	instance := getCNNCrossChannelNormalizationClass().Alloc()
	rv := objc.Send[CNNCrossChannelNormalization](instance.ID, objc.Sel("initWithDevice:kernelSize:"), device, kernelSize)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNCrossChannelNormalizationWithDeviceKernelSize */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNCrossChannelNormalization */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNCrossChannelNormalization */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNCrossChannelNormalization */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNCrossChannelNormalization */

// The size of the square kernel window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648811-kernelsize
func (c_ CNNCrossChannelNormalization) KernelSize() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelSize"))
	return rv
}/* debug [instance_properties/getter]: kernelSize */


// The size of the square kernel window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648811-kernelsize
func (c_ CNNCrossChannelNormalization) SetKernelSize(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelSize:"), value)
}/* debug [instance_properties/setter]: kernelSize */


// The "beta" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648879-beta
func (c_ CNNCrossChannelNormalization) Beta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("beta"))
	return rv
}/* debug [instance_properties/getter]: beta */


// The "beta" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648879-beta
func (c_ CNNCrossChannelNormalization) SetBeta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}/* debug [instance_properties/setter]: beta */


// The "delta" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648881-delta
func (c_ CNNCrossChannelNormalization) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delta"))
	return rv
}/* debug [instance_properties/getter]: delta */


// The "delta" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648881-delta
func (c_ CNNCrossChannelNormalization) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelta:"), value)
}/* debug [instance_properties/setter]: delta */


// The "alpha" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648896-alpha
func (c_ CNNCrossChannelNormalization) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// The "alpha" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648896-alpha
func (c_ CNNCrossChannelNormalization) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}/* debug [instance_properties/setter]: alpha */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNCrossChannelNormalization */


