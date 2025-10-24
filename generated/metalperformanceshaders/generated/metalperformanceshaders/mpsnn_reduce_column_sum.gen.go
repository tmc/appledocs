// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReduceColumnSum */


/* debug [class_header]: Header for MPSNNReduceColumnSum */
// The class instance for the [ReduceColumnSum] class.
var (
	ReduceColumnSumClass     _ReduceColumnSumClass
	ReduceColumnSumClassOnce sync.Once
)

func getReduceColumnSumClass() _ReduceColumnSumClass {
	ReduceColumnSumClassOnce.Do(func() {
		ReduceColumnSumClass = _ReduceColumnSumClass{objc.GetClass("MPSNNReduceColumnSum")}
	})
	return ReduceColumnSumClass
}

type _ReduceColumnSumClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceColumnSum */
// An interface definition for the [ReduceColumnSum] class.
type IReduceColumnSum interface {
	IReduceUnary
	
/* debug [class_interface_properties]: Properties for ReduceColumnSum */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceColumnSum */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceColumnSum */
// Alloc allocates a new instance without initialization.
func (rc _ReduceColumnSumClass) Alloc() ReduceColumnSum {
	rv := objc.Send[ReduceColumnSum](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceColumnSumClass) New() ReduceColumnSum {
	rv := objc.Send[ReduceColumnSum](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceColumnSum) Init() ReduceColumnSum {
	rv := objc.Send[ReduceColumnSum](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceColumnSum) Autorelease() ReduceColumnSum {
	rv := objc.Send[ReduceColumnSum](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceColumnSum creates a new ReduceColumnSum instance.
func NewReduceColumnSum() ReduceColumnSum {
	return getReduceColumnSumClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceColumnSum */
// A reduction filter that returns the sum of all values for each column in an image.


// A reduction filter that returns the sum of all values for each column in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnSum
type ReduceColumnSum struct {
	ReduceUnary
}

// ReduceColumnSumFrom constructs a [ReduceColumnSum] from an unsafe.Pointer.
//
// A reduction filter that returns the sum of all values for each column in an image.
func ReduceColumnSumFrom(ptr unsafe.Pointer) ReduceColumnSum {
	return ReduceColumnSum{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceColumnSum */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnsum/3197833-initwithcoder
func NewReduceColumnSumWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ReduceColumnSum {
	instance := getReduceColumnSumClass().Alloc()
	rv := objc.Send[ReduceColumnSum](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceColumnSumWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnsum/2942540-initwithdevice
func NewReduceColumnSumWithDevice(device unsafe.Pointer) ReduceColumnSum {
	instance := getReduceColumnSumClass().Alloc()
	rv := objc.Send[ReduceColumnSum](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceColumnSumWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceColumnSum */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceColumnSum */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceColumnSum */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceColumnSum */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceColumnSum */


