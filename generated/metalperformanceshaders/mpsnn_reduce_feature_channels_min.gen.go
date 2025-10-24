// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReduceFeatureChannelsMin */


/* debug [class_header]: Header for MPSNNReduceFeatureChannelsMin */
// The class instance for the [ReduceFeatureChannelsMin] class.
var (
	ReduceFeatureChannelsMinClass     _ReduceFeatureChannelsMinClass
	ReduceFeatureChannelsMinClassOnce sync.Once
)

func getReduceFeatureChannelsMinClass() _ReduceFeatureChannelsMinClass {
	ReduceFeatureChannelsMinClassOnce.Do(func() {
		ReduceFeatureChannelsMinClass = _ReduceFeatureChannelsMinClass{objc.GetClass("MPSNNReduceFeatureChannelsMin")}
	})
	return ReduceFeatureChannelsMinClass
}

type _ReduceFeatureChannelsMinClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceFeatureChannelsMin */
// An interface definition for the [ReduceFeatureChannelsMin] class.
type IReduceFeatureChannelsMin interface {
	IReduceUnary
	
/* debug [class_interface_properties]: Properties for ReduceFeatureChannelsMin */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceFeatureChannelsMin */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceFeatureChannelsMin */
// Alloc allocates a new instance without initialization.
func (rc _ReduceFeatureChannelsMinClass) Alloc() ReduceFeatureChannelsMin {
	rv := objc.Send[ReduceFeatureChannelsMin](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceFeatureChannelsMinClass) New() ReduceFeatureChannelsMin {
	rv := objc.Send[ReduceFeatureChannelsMin](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceFeatureChannelsMin) Init() ReduceFeatureChannelsMin {
	rv := objc.Send[ReduceFeatureChannelsMin](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceFeatureChannelsMin) Autorelease() ReduceFeatureChannelsMin {
	rv := objc.Send[ReduceFeatureChannelsMin](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceFeatureChannelsMin creates a new ReduceFeatureChannelsMin instance.
func NewReduceFeatureChannelsMin() ReduceFeatureChannelsMin {
	return getReduceFeatureChannelsMinClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceFeatureChannelsMin */
// A reduction filter that returns the minimum value for each feature channel in an image.


// A reduction filter that returns the minimum value for each feature channel in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMin
type ReduceFeatureChannelsMin struct {
	ReduceUnary
}

// ReduceFeatureChannelsMinFrom constructs a [ReduceFeatureChannelsMin] from an unsafe.Pointer.
//
// A reduction filter that returns the minimum value for each feature channel in an image.
func ReduceFeatureChannelsMinFrom(ptr unsafe.Pointer) ReduceFeatureChannelsMin {
	return ReduceFeatureChannelsMin{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceFeatureChannelsMin */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmin/3197840-initwithcoder
func NewReduceFeatureChannelsMinWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceFeatureChannelsMin {
	instance := getReduceFeatureChannelsMinClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsMin](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceFeatureChannelsMinWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmin/2942565-initwithdevice
func NewReduceFeatureChannelsMinWithDevice(device unsafe.Pointer) ReduceFeatureChannelsMin {
	instance := getReduceFeatureChannelsMinClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsMin](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceFeatureChannelsMinWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceFeatureChannelsMin */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceFeatureChannelsMin */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceFeatureChannelsMin */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceFeatureChannelsMin */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceFeatureChannelsMin */


