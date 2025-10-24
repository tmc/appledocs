// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNLocalCorrelation */


/* debug [class_header]: Header for MPSNNLocalCorrelation */
// The class instance for the [LocalCorrelation] class.
var (
	LocalCorrelationClass     _LocalCorrelationClass
	LocalCorrelationClassOnce sync.Once
)

func getLocalCorrelationClass() _LocalCorrelationClass {
	LocalCorrelationClassOnce.Do(func() {
		LocalCorrelationClass = _LocalCorrelationClass{objc.GetClass("MPSNNLocalCorrelation")}
	})
	return LocalCorrelationClass
}

type _LocalCorrelationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LocalCorrelation */
// An interface definition for the [LocalCorrelation] class.
type ILocalCorrelation interface {
	IReduceBinary
	
/* debug [class_interface_properties]: Properties for LocalCorrelation */
	// properties:
	StrideInX() objectivec.IObject
	SetStrideInX(value objectivec.IObject)
	StrideInY() objectivec.IObject
	SetStrideInY(value objectivec.IObject)
	WindowInX() objectivec.IObject
	SetWindowInX(value objectivec.IObject)
	WindowInY() objectivec.IObject
	SetWindowInY(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LocalCorrelation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LocalCorrelation */
// Alloc allocates a new instance without initialization.
func (lc _LocalCorrelationClass) Alloc() LocalCorrelation {
	rv := objc.Send[LocalCorrelation](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LocalCorrelationClass) New() LocalCorrelation {
	rv := objc.Send[LocalCorrelation](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LocalCorrelation) Init() LocalCorrelation {
	rv := objc.Send[LocalCorrelation](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LocalCorrelation) Autorelease() LocalCorrelation {
	rv := objc.Send[LocalCorrelation](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocalCorrelation creates a new LocalCorrelation instance.
func NewLocalCorrelation() LocalCorrelation {
	return getLocalCorrelationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LocalCorrelation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation
type LocalCorrelation struct {
	ReduceBinary
}

// LocalCorrelationFrom constructs a [LocalCorrelation] from an unsafe.Pointer.
func LocalCorrelationFrom(ptr unsafe.Pointer) LocalCorrelation {
	return LocalCorrelation{
		ReduceBinary: ReduceBinaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LocalCorrelation */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3197829-initwithcoder
func NewLocalCorrelationWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) LocalCorrelation {
	instance := getLocalCorrelationClass().Alloc()
	rv := objc.Send[LocalCorrelation](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLocalCorrelationWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131875-initwithdevice
func NewLocalCorrelationWithDevice(device unsafe.Pointer) LocalCorrelation {
	instance := getLocalCorrelationClass().Alloc()
	rv := objc.Send[LocalCorrelation](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLocalCorrelationWithDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131876-initwithdevice
func NewLocalCorrelationWithDeviceWindowInXWindowInYStrideInXStrideInY(device unsafe.Pointer, windowInX uint, windowInY uint, strideInX uint, strideInY uint) LocalCorrelation {
	instance := getLocalCorrelationClass().Alloc()
	rv := objc.Send[LocalCorrelation](instance.ID, objc.Sel("initWithDevice:windowInX:windowInY:strideInX:strideInY:"), device, windowInX, windowInY, strideInX, strideInY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLocalCorrelationWithDeviceWindowInXWindowInYStrideInXStrideInY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LocalCorrelation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LocalCorrelation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LocalCorrelation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LocalCorrelation */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131877-strideinx
func (l_ LocalCorrelation) StrideInX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("strideInX"))
	return rv
}/* debug [instance_properties/getter]: strideInX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131877-strideinx
func (l_ LocalCorrelation) SetStrideInX(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setStrideInX:"), value)
}/* debug [instance_properties/setter]: strideInX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131878-strideiny
func (l_ LocalCorrelation) StrideInY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("strideInY"))
	return rv
}/* debug [instance_properties/getter]: strideInY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131878-strideiny
func (l_ LocalCorrelation) SetStrideInY(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setStrideInY:"), value)
}/* debug [instance_properties/setter]: strideInY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131879-windowinx
func (l_ LocalCorrelation) WindowInX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("windowInX"))
	return rv
}/* debug [instance_properties/getter]: windowInX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131879-windowinx
func (l_ LocalCorrelation) SetWindowInX(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWindowInX:"), value)
}/* debug [instance_properties/setter]: windowInX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131880-windowiny
func (l_ LocalCorrelation) WindowInY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("windowInY"))
	return rv
}/* debug [instance_properties/getter]: windowInY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131880-windowiny
func (l_ LocalCorrelation) SetWindowInY(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWindowInY:"), value)
}/* debug [instance_properties/setter]: windowInY */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNLocalCorrelation */


