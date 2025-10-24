// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNLocalContrastNormalization */


/* debug [class_header]: Header for MPSCNNLocalContrastNormalization */
// The class instance for the [CNNLocalContrastNormalization] class.
var (
	CNNLocalContrastNormalizationClass     _CNNLocalContrastNormalizationClass
	CNNLocalContrastNormalizationClassOnce sync.Once
)

func getCNNLocalContrastNormalizationClass() _CNNLocalContrastNormalizationClass {
	CNNLocalContrastNormalizationClassOnce.Do(func() {
		CNNLocalContrastNormalizationClass = _CNNLocalContrastNormalizationClass{objc.GetClass("MPSCNNLocalContrastNormalization")}
	})
	return CNNLocalContrastNormalizationClass
}

type _CNNLocalContrastNormalizationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNLocalContrastNormalization */
// An interface definition for the [CNNLocalContrastNormalization] class.
type ICNNLocalContrastNormalization interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for CNNLocalContrastNormalization */
	// properties:
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	Beta() objectivec.IObject
	SetBeta(value objectivec.IObject)
	Pm() objectivec.IObject
	SetPm(value objectivec.IObject)
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	Ps() objectivec.IObject
	SetPs(value objectivec.IObject)
	P0() objectivec.IObject
	SetP0(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNLocalContrastNormalization */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNLocalContrastNormalization */
// Alloc allocates a new instance without initialization.
func (cc _CNNLocalContrastNormalizationClass) Alloc() CNNLocalContrastNormalization {
	rv := objc.Send[CNNLocalContrastNormalization](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLocalContrastNormalizationClass) New() CNNLocalContrastNormalization {
	rv := objc.Send[CNNLocalContrastNormalization](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLocalContrastNormalization) Init() CNNLocalContrastNormalization {
	rv := objc.Send[CNNLocalContrastNormalization](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLocalContrastNormalization) Autorelease() CNNLocalContrastNormalization {
	rv := objc.Send[CNNLocalContrastNormalization](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLocalContrastNormalization creates a new CNNLocalContrastNormalization instance.
func NewCNNLocalContrastNormalization() CNNLocalContrastNormalization {
	return getCNNLocalContrastNormalizationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNLocalContrastNormalization */
// A local-contrast normalization kernel.
//
// The local contrast normalization kernel is quite similar to the spatial normalization kernel, described in the class, in that it applies the kernel over local regions which extend spatially, but are in separate feature channels (i.e., they have the shape ). However, instead of dividing by the local “energy” of the feature, the denominator uses the local variance of the feature - effectively the mean value of the feature is subtracted from the signal. For each feature channel, the function computes the variance and mean of inside each rectangle around the spatial point . Then the result is computed for each element of as follows: Where and are the values of the and the properties, respectively, and the values of the , , and properties can be used to offset and scale the result in various ways. For example setting , , , , and scales input data so that the result has unit variance and zero mean, provided that input variance is positive. It is your responsibility to ensure that the combination of the values of the and properties does not result in a situation where the denominator becomes zero - in such situations the resulting pixel-value is undefined. A good way to guard against tiny variances is to regulate the expression with a small delta value, for example .


// A local-contrast normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalization
type CNNLocalContrastNormalization struct {
	CNNKernel
}

// CNNLocalContrastNormalizationFrom constructs a [CNNLocalContrastNormalization] from an unsafe.Pointer.
//
// A local-contrast normalization kernel.
func CNNLocalContrastNormalizationFrom(ptr unsafe.Pointer) CNNLocalContrastNormalization {
	return CNNLocalContrastNormalization{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNLocalContrastNormalization */

// Initializes a local contrast normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalization/init(coder:device:)
func NewCNNLocalContrastNormalizationWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNLocalContrastNormalization {
	instance := getCNNLocalContrastNormalizationClass().Alloc()
	rv := objc.Send[CNNLocalContrastNormalization](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNLocalContrastNormalizationWithCoderDevice */


// Initializes a local contrast normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648924-initwithdevice
func NewCNNLocalContrastNormalizationWithDeviceKernelWidthKernelHeight(device unsafe.Pointer, kernelWidth uint, kernelHeight uint) CNNLocalContrastNormalization {
	instance := getCNNLocalContrastNormalizationClass().Alloc()
	rv := objc.Send[CNNLocalContrastNormalization](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNLocalContrastNormalizationWithDeviceKernelWidthKernelHeight */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNLocalContrastNormalization */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNLocalContrastNormalization */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNLocalContrastNormalization */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNLocalContrastNormalization */

// The "delta" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648812-delta
func (c_ CNNLocalContrastNormalization) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delta"))
	return rv
}/* debug [instance_properties/getter]: delta */


// The "delta" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648812-delta
func (c_ CNNLocalContrastNormalization) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelta:"), value)
}/* debug [instance_properties/setter]: delta */


// The "beta" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648905-beta
func (c_ CNNLocalContrastNormalization) Beta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("beta"))
	return rv
}/* debug [instance_properties/getter]: beta */


// The "beta" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648905-beta
func (c_ CNNLocalContrastNormalization) SetBeta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}/* debug [instance_properties/setter]: beta */


// The "pm" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648907-pm
func (c_ CNNLocalContrastNormalization) Pm() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("pm"))
	return rv
}/* debug [instance_properties/getter]: pm */


// The "pm" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648907-pm
func (c_ CNNLocalContrastNormalization) SetPm(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPm:"), value)
}/* debug [instance_properties/setter]: pm */


// The "alpha" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648923-alpha
func (c_ CNNLocalContrastNormalization) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// The "alpha" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648923-alpha
func (c_ CNNLocalContrastNormalization) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}/* debug [instance_properties/setter]: alpha */


// The "ps" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648942-ps
func (c_ CNNLocalContrastNormalization) Ps() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("ps"))
	return rv
}/* debug [instance_properties/getter]: ps */


// The "ps" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648942-ps
func (c_ CNNLocalContrastNormalization) SetPs(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPs:"), value)
}/* debug [instance_properties/setter]: ps */


// The "p0" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648953-p0
func (c_ CNNLocalContrastNormalization) P0() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("p0"))
	return rv
}/* debug [instance_properties/getter]: p0 */


// The "p0" variable of the kernel function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648953-p0
func (c_ CNNLocalContrastNormalization) SetP0(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setP0:"), value)
}/* debug [instance_properties/setter]: p0 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNLocalContrastNormalization */


