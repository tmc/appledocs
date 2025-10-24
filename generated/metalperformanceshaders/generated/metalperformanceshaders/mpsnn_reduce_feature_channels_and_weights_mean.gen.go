// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReduceFeatureChannelsAndWeightsMean */


/* debug [class_header]: Header for MPSNNReduceFeatureChannelsAndWeightsMean */
// The class instance for the [ReduceFeatureChannelsAndWeightsMean] class.
var (
	ReduceFeatureChannelsAndWeightsMeanClass     _ReduceFeatureChannelsAndWeightsMeanClass
	ReduceFeatureChannelsAndWeightsMeanClassOnce sync.Once
)

func getReduceFeatureChannelsAndWeightsMeanClass() _ReduceFeatureChannelsAndWeightsMeanClass {
	ReduceFeatureChannelsAndWeightsMeanClassOnce.Do(func() {
		ReduceFeatureChannelsAndWeightsMeanClass = _ReduceFeatureChannelsAndWeightsMeanClass{objc.GetClass("MPSNNReduceFeatureChannelsAndWeightsMean")}
	})
	return ReduceFeatureChannelsAndWeightsMeanClass
}

type _ReduceFeatureChannelsAndWeightsMeanClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceFeatureChannelsAndWeightsMean */
// An interface definition for the [ReduceFeatureChannelsAndWeightsMean] class.
type IReduceFeatureChannelsAndWeightsMean interface {
	IReduceBinary
	
/* debug [class_interface_properties]: Properties for ReduceFeatureChannelsAndWeightsMean */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceFeatureChannelsAndWeightsMean */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceFeatureChannelsAndWeightsMean */
// Alloc allocates a new instance without initialization.
func (rc _ReduceFeatureChannelsAndWeightsMeanClass) Alloc() ReduceFeatureChannelsAndWeightsMean {
	rv := objc.Send[ReduceFeatureChannelsAndWeightsMean](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceFeatureChannelsAndWeightsMeanClass) New() ReduceFeatureChannelsAndWeightsMean {
	rv := objc.Send[ReduceFeatureChannelsAndWeightsMean](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceFeatureChannelsAndWeightsMean) Init() ReduceFeatureChannelsAndWeightsMean {
	rv := objc.Send[ReduceFeatureChannelsAndWeightsMean](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceFeatureChannelsAndWeightsMean) Autorelease() ReduceFeatureChannelsAndWeightsMean {
	rv := objc.Send[ReduceFeatureChannelsAndWeightsMean](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceFeatureChannelsAndWeightsMean creates a new ReduceFeatureChannelsAndWeightsMean instance.
func NewReduceFeatureChannelsAndWeightsMean() ReduceFeatureChannelsAndWeightsMean {
	return getReduceFeatureChannelsAndWeightsMeanClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceFeatureChannelsAndWeightsMean */
// A reduction filter that returns the weighted sum for each feature channel in an image.


// A reduction filter that returns the weighted sum for each feature channel in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsAndWeightsMean
type ReduceFeatureChannelsAndWeightsMean struct {
	ReduceBinary
}

// ReduceFeatureChannelsAndWeightsMeanFrom constructs a [ReduceFeatureChannelsAndWeightsMean] from an unsafe.Pointer.
//
// A reduction filter that returns the weighted sum for each feature channel in an image.
func ReduceFeatureChannelsAndWeightsMeanFrom(ptr unsafe.Pointer) ReduceFeatureChannelsAndWeightsMean {
	return ReduceFeatureChannelsAndWeightsMean{
		ReduceBinary: ReduceBinaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceFeatureChannelsAndWeightsMean */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightsmean/3197834-initwithcoder
func NewReduceFeatureChannelsAndWeightsMeanWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ReduceFeatureChannelsAndWeightsMean {
	instance := getReduceFeatureChannelsAndWeightsMeanClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsAndWeightsMean](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceFeatureChannelsAndWeightsMeanWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightsmean/2942537-initwithdevice
func NewReduceFeatureChannelsAndWeightsMeanWithDevice(device unsafe.Pointer) ReduceFeatureChannelsAndWeightsMean {
	instance := getReduceFeatureChannelsAndWeightsMeanClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsAndWeightsMean](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceFeatureChannelsAndWeightsMeanWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceFeatureChannelsAndWeightsMean */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceFeatureChannelsAndWeightsMean */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceFeatureChannelsAndWeightsMean */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceFeatureChannelsAndWeightsMean */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceFeatureChannelsAndWeightsMean */


