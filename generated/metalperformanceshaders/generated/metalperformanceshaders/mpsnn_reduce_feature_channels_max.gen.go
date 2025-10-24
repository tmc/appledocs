// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReduceFeatureChannelsMax */


/* debug [class_header]: Header for MPSNNReduceFeatureChannelsMax */
// The class instance for the [ReduceFeatureChannelsMax] class.
var (
	ReduceFeatureChannelsMaxClass     _ReduceFeatureChannelsMaxClass
	ReduceFeatureChannelsMaxClassOnce sync.Once
)

func getReduceFeatureChannelsMaxClass() _ReduceFeatureChannelsMaxClass {
	ReduceFeatureChannelsMaxClassOnce.Do(func() {
		ReduceFeatureChannelsMaxClass = _ReduceFeatureChannelsMaxClass{objc.GetClass("MPSNNReduceFeatureChannelsMax")}
	})
	return ReduceFeatureChannelsMaxClass
}

type _ReduceFeatureChannelsMaxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceFeatureChannelsMax */
// An interface definition for the [ReduceFeatureChannelsMax] class.
type IReduceFeatureChannelsMax interface {
	IReduceUnary
	
/* debug [class_interface_properties]: Properties for ReduceFeatureChannelsMax */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceFeatureChannelsMax */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceFeatureChannelsMax */
// Alloc allocates a new instance without initialization.
func (rc _ReduceFeatureChannelsMaxClass) Alloc() ReduceFeatureChannelsMax {
	rv := objc.Send[ReduceFeatureChannelsMax](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceFeatureChannelsMaxClass) New() ReduceFeatureChannelsMax {
	rv := objc.Send[ReduceFeatureChannelsMax](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceFeatureChannelsMax) Init() ReduceFeatureChannelsMax {
	rv := objc.Send[ReduceFeatureChannelsMax](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceFeatureChannelsMax) Autorelease() ReduceFeatureChannelsMax {
	rv := objc.Send[ReduceFeatureChannelsMax](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceFeatureChannelsMax creates a new ReduceFeatureChannelsMax instance.
func NewReduceFeatureChannelsMax() ReduceFeatureChannelsMax {
	return getReduceFeatureChannelsMaxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceFeatureChannelsMax */
// A reduction filter that returns the maximum value for each feature channel in an image.


// A reduction filter that returns the maximum value for each feature channel in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMax
type ReduceFeatureChannelsMax struct {
	ReduceUnary
}

// ReduceFeatureChannelsMaxFrom constructs a [ReduceFeatureChannelsMax] from an unsafe.Pointer.
//
// A reduction filter that returns the maximum value for each feature channel in an image.
func ReduceFeatureChannelsMaxFrom(ptr unsafe.Pointer) ReduceFeatureChannelsMax {
	return ReduceFeatureChannelsMax{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceFeatureChannelsMax */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmax/3197838-initwithcoder
func NewReduceFeatureChannelsMaxWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ReduceFeatureChannelsMax {
	instance := getReduceFeatureChannelsMaxClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsMax](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceFeatureChannelsMaxWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmax/2942532-initwithdevice
func NewReduceFeatureChannelsMaxWithDevice(device unsafe.Pointer) ReduceFeatureChannelsMax {
	instance := getReduceFeatureChannelsMaxClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsMax](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceFeatureChannelsMaxWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceFeatureChannelsMax */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceFeatureChannelsMax */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceFeatureChannelsMax */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceFeatureChannelsMax */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceFeatureChannelsMax */


