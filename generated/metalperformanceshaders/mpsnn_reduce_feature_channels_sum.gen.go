// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReduceFeatureChannelsSum */


/* debug [class_header]: Header for MPSNNReduceFeatureChannelsSum */
// The class instance for the [ReduceFeatureChannelsSum] class.
var (
	ReduceFeatureChannelsSumClass     _ReduceFeatureChannelsSumClass
	ReduceFeatureChannelsSumClassOnce sync.Once
)

func getReduceFeatureChannelsSumClass() _ReduceFeatureChannelsSumClass {
	ReduceFeatureChannelsSumClassOnce.Do(func() {
		ReduceFeatureChannelsSumClass = _ReduceFeatureChannelsSumClass{objc.GetClass("MPSNNReduceFeatureChannelsSum")}
	})
	return ReduceFeatureChannelsSumClass
}

type _ReduceFeatureChannelsSumClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceFeatureChannelsSum */
// An interface definition for the [ReduceFeatureChannelsSum] class.
type IReduceFeatureChannelsSum interface {
	IReduceUnary
	
/* debug [class_interface_properties]: Properties for ReduceFeatureChannelsSum */
	// properties:
	Weight() objectivec.IObject
	SetWeight(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceFeatureChannelsSum */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceFeatureChannelsSum */
// Alloc allocates a new instance without initialization.
func (rc _ReduceFeatureChannelsSumClass) Alloc() ReduceFeatureChannelsSum {
	rv := objc.Send[ReduceFeatureChannelsSum](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceFeatureChannelsSumClass) New() ReduceFeatureChannelsSum {
	rv := objc.Send[ReduceFeatureChannelsSum](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceFeatureChannelsSum) Init() ReduceFeatureChannelsSum {
	rv := objc.Send[ReduceFeatureChannelsSum](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceFeatureChannelsSum) Autorelease() ReduceFeatureChannelsSum {
	rv := objc.Send[ReduceFeatureChannelsSum](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceFeatureChannelsSum creates a new ReduceFeatureChannelsSum instance.
func NewReduceFeatureChannelsSum() ReduceFeatureChannelsSum {
	return getReduceFeatureChannelsSumClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceFeatureChannelsSum */
// A reduction filter that returns the sum of all values for each feature channel in an image.


// A reduction filter that returns the sum of all values for each feature channel in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsSum
type ReduceFeatureChannelsSum struct {
	ReduceUnary
}

// ReduceFeatureChannelsSumFrom constructs a [ReduceFeatureChannelsSum] from an unsafe.Pointer.
//
// A reduction filter that returns the sum of all values for each feature channel in an image.
func ReduceFeatureChannelsSumFrom(ptr unsafe.Pointer) ReduceFeatureChannelsSum {
	return ReduceFeatureChannelsSum{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceFeatureChannelsSum */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelssum/3197841-initwithcoder
func NewReduceFeatureChannelsSumWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceFeatureChannelsSum {
	instance := getReduceFeatureChannelsSumClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsSum](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceFeatureChannelsSumWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelssum/2942538-initwithdevice
func NewReduceFeatureChannelsSumWithDevice(device unsafe.Pointer) ReduceFeatureChannelsSum {
	instance := getReduceFeatureChannelsSumClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsSum](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceFeatureChannelsSumWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceFeatureChannelsSum */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceFeatureChannelsSum */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceFeatureChannelsSum */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceFeatureChannelsSum */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelssum/2942545-weight
func (r_ ReduceFeatureChannelsSum) Weight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("weight"))
	return rv
}/* debug [instance_properties/getter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelssum/2942545-weight
func (r_ ReduceFeatureChannelsSum) SetWeight(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setWeight:"), value)
}/* debug [instance_properties/setter]: weight */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceFeatureChannelsSum */


