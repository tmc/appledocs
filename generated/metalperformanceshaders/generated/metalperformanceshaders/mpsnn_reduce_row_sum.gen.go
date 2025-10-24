// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReduceRowSum */


/* debug [class_header]: Header for MPSNNReduceRowSum */
// The class instance for the [ReduceRowSum] class.
var (
	ReduceRowSumClass     _ReduceRowSumClass
	ReduceRowSumClassOnce sync.Once
)

func getReduceRowSumClass() _ReduceRowSumClass {
	ReduceRowSumClassOnce.Do(func() {
		ReduceRowSumClass = _ReduceRowSumClass{objc.GetClass("MPSNNReduceRowSum")}
	})
	return ReduceRowSumClass
}

type _ReduceRowSumClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceRowSum */
// An interface definition for the [ReduceRowSum] class.
type IReduceRowSum interface {
	IReduceUnary
	
/* debug [class_interface_properties]: Properties for ReduceRowSum */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceRowSum */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceRowSum */
// Alloc allocates a new instance without initialization.
func (rc _ReduceRowSumClass) Alloc() ReduceRowSum {
	rv := objc.Send[ReduceRowSum](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceRowSumClass) New() ReduceRowSum {
	rv := objc.Send[ReduceRowSum](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceRowSum) Init() ReduceRowSum {
	rv := objc.Send[ReduceRowSum](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceRowSum) Autorelease() ReduceRowSum {
	rv := objc.Send[ReduceRowSum](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceRowSum creates a new ReduceRowSum instance.
func NewReduceRowSum() ReduceRowSum {
	return getReduceRowSumClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceRowSum */
// A reduction filter that returns the sum of all values for each row in an image.


// A reduction filter that returns the sum of all values for each row in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowSum
type ReduceRowSum struct {
	ReduceUnary
}

// ReduceRowSumFrom constructs a [ReduceRowSum] from an unsafe.Pointer.
//
// A reduction filter that returns the sum of all values for each row in an image.
func ReduceRowSumFrom(ptr unsafe.Pointer) ReduceRowSum {
	return ReduceRowSum{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceRowSum */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowsum/3197845-initwithcoder
func NewReduceRowSumWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ReduceRowSum {
	instance := getReduceRowSumClass().Alloc()
	rv := objc.Send[ReduceRowSum](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceRowSumWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowsum/2942536-initwithdevice
func NewReduceRowSumWithDevice(device unsafe.Pointer) ReduceRowSum {
	instance := getReduceRowSumClass().Alloc()
	rv := objc.Send[ReduceRowSum](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceRowSumWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceRowSum */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceRowSum */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceRowSum */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceRowSum */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceRowSum */


