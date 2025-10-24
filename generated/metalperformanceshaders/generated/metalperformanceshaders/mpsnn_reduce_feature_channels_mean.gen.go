// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReduceFeatureChannelsMean */


/* debug [class_header]: Header for MPSNNReduceFeatureChannelsMean */
// The class instance for the [ReduceFeatureChannelsMean] class.
var (
	ReduceFeatureChannelsMeanClass     _ReduceFeatureChannelsMeanClass
	ReduceFeatureChannelsMeanClassOnce sync.Once
)

func getReduceFeatureChannelsMeanClass() _ReduceFeatureChannelsMeanClass {
	ReduceFeatureChannelsMeanClassOnce.Do(func() {
		ReduceFeatureChannelsMeanClass = _ReduceFeatureChannelsMeanClass{objc.GetClass("MPSNNReduceFeatureChannelsMean")}
	})
	return ReduceFeatureChannelsMeanClass
}

type _ReduceFeatureChannelsMeanClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceFeatureChannelsMean */
// An interface definition for the [ReduceFeatureChannelsMean] class.
type IReduceFeatureChannelsMean interface {
	IReduceUnary
	
/* debug [class_interface_properties]: Properties for ReduceFeatureChannelsMean */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceFeatureChannelsMean */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceFeatureChannelsMean */
// Alloc allocates a new instance without initialization.
func (rc _ReduceFeatureChannelsMeanClass) Alloc() ReduceFeatureChannelsMean {
	rv := objc.Send[ReduceFeatureChannelsMean](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceFeatureChannelsMeanClass) New() ReduceFeatureChannelsMean {
	rv := objc.Send[ReduceFeatureChannelsMean](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceFeatureChannelsMean) Init() ReduceFeatureChannelsMean {
	rv := objc.Send[ReduceFeatureChannelsMean](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceFeatureChannelsMean) Autorelease() ReduceFeatureChannelsMean {
	rv := objc.Send[ReduceFeatureChannelsMean](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceFeatureChannelsMean creates a new ReduceFeatureChannelsMean instance.
func NewReduceFeatureChannelsMean() ReduceFeatureChannelsMean {
	return getReduceFeatureChannelsMeanClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceFeatureChannelsMean */
// A reduction filter that returns the mean value for each feature channel in an image.


// A reduction filter that returns the mean value for each feature channel in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMean
type ReduceFeatureChannelsMean struct {
	ReduceUnary
}

// ReduceFeatureChannelsMeanFrom constructs a [ReduceFeatureChannelsMean] from an unsafe.Pointer.
//
// A reduction filter that returns the mean value for each feature channel in an image.
func ReduceFeatureChannelsMeanFrom(ptr unsafe.Pointer) ReduceFeatureChannelsMean {
	return ReduceFeatureChannelsMean{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceFeatureChannelsMean */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmean/3197839-initwithcoder
func NewReduceFeatureChannelsMeanWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ReduceFeatureChannelsMean {
	instance := getReduceFeatureChannelsMeanClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsMean](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceFeatureChannelsMeanWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmean/2942557-initwithdevice
func NewReduceFeatureChannelsMeanWithDevice(device unsafe.Pointer) ReduceFeatureChannelsMean {
	instance := getReduceFeatureChannelsMeanClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsMean](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceFeatureChannelsMeanWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceFeatureChannelsMean */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceFeatureChannelsMean */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceFeatureChannelsMean */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceFeatureChannelsMean */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceFeatureChannelsMean */


