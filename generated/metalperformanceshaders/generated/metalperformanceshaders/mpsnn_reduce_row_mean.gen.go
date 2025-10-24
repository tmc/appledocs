// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReduceRowMean */


/* debug [class_header]: Header for MPSNNReduceRowMean */
// The class instance for the [ReduceRowMean] class.
var (
	ReduceRowMeanClass     _ReduceRowMeanClass
	ReduceRowMeanClassOnce sync.Once
)

func getReduceRowMeanClass() _ReduceRowMeanClass {
	ReduceRowMeanClassOnce.Do(func() {
		ReduceRowMeanClass = _ReduceRowMeanClass{objc.GetClass("MPSNNReduceRowMean")}
	})
	return ReduceRowMeanClass
}

type _ReduceRowMeanClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceRowMean */
// An interface definition for the [ReduceRowMean] class.
type IReduceRowMean interface {
	IReduceUnary
	
/* debug [class_interface_properties]: Properties for ReduceRowMean */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceRowMean */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceRowMean */
// Alloc allocates a new instance without initialization.
func (rc _ReduceRowMeanClass) Alloc() ReduceRowMean {
	rv := objc.Send[ReduceRowMean](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceRowMeanClass) New() ReduceRowMean {
	rv := objc.Send[ReduceRowMean](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceRowMean) Init() ReduceRowMean {
	rv := objc.Send[ReduceRowMean](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceRowMean) Autorelease() ReduceRowMean {
	rv := objc.Send[ReduceRowMean](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceRowMean creates a new ReduceRowMean instance.
func NewReduceRowMean() ReduceRowMean {
	return getReduceRowMeanClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceRowMean */
// A reduction filter that returns the mean value for each row in an image.


// A reduction filter that returns the mean value for each row in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMean
type ReduceRowMean struct {
	ReduceUnary
}

// ReduceRowMeanFrom constructs a [ReduceRowMean] from an unsafe.Pointer.
//
// A reduction filter that returns the mean value for each row in an image.
func ReduceRowMeanFrom(ptr unsafe.Pointer) ReduceRowMean {
	return ReduceRowMean{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceRowMean */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowmean/3197843-initwithcoder
func NewReduceRowMeanWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ReduceRowMean {
	instance := getReduceRowMeanClass().Alloc()
	rv := objc.Send[ReduceRowMean](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceRowMeanWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowmean/2942548-initwithdevice
func NewReduceRowMeanWithDevice(device unsafe.Pointer) ReduceRowMean {
	instance := getReduceRowMeanClass().Alloc()
	rv := objc.Send[ReduceRowMean](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceRowMeanWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceRowMean */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceRowMean */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceRowMean */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceRowMean */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceRowMean */


