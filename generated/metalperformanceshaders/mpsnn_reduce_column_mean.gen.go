// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReduceColumnMean */


/* debug [class_header]: Header for MPSNNReduceColumnMean */
// The class instance for the [ReduceColumnMean] class.
var (
	ReduceColumnMeanClass     _ReduceColumnMeanClass
	ReduceColumnMeanClassOnce sync.Once
)

func getReduceColumnMeanClass() _ReduceColumnMeanClass {
	ReduceColumnMeanClassOnce.Do(func() {
		ReduceColumnMeanClass = _ReduceColumnMeanClass{objc.GetClass("MPSNNReduceColumnMean")}
	})
	return ReduceColumnMeanClass
}

type _ReduceColumnMeanClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceColumnMean */
// An interface definition for the [ReduceColumnMean] class.
type IReduceColumnMean interface {
	IReduceUnary
	
/* debug [class_interface_properties]: Properties for ReduceColumnMean */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceColumnMean */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceColumnMean */
// Alloc allocates a new instance without initialization.
func (rc _ReduceColumnMeanClass) Alloc() ReduceColumnMean {
	rv := objc.Send[ReduceColumnMean](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceColumnMeanClass) New() ReduceColumnMean {
	rv := objc.Send[ReduceColumnMean](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceColumnMean) Init() ReduceColumnMean {
	rv := objc.Send[ReduceColumnMean](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceColumnMean) Autorelease() ReduceColumnMean {
	rv := objc.Send[ReduceColumnMean](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceColumnMean creates a new ReduceColumnMean instance.
func NewReduceColumnMean() ReduceColumnMean {
	return getReduceColumnMeanClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceColumnMean */
// A reduction filter that returns the mean value for each column in an image.


// A reduction filter that returns the mean value for each column in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMean
type ReduceColumnMean struct {
	ReduceUnary
}

// ReduceColumnMeanFrom constructs a [ReduceColumnMean] from an unsafe.Pointer.
//
// A reduction filter that returns the mean value for each column in an image.
func ReduceColumnMeanFrom(ptr unsafe.Pointer) ReduceColumnMean {
	return ReduceColumnMean{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceColumnMean */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmean/3197831-initwithcoder
func NewReduceColumnMeanWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceColumnMean {
	instance := getReduceColumnMeanClass().Alloc()
	rv := objc.Send[ReduceColumnMean](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceColumnMeanWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmean/2942546-initwithdevice
func NewReduceColumnMeanWithDevice(device unsafe.Pointer) ReduceColumnMean {
	instance := getReduceColumnMeanClass().Alloc()
	rv := objc.Send[ReduceColumnMean](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceColumnMeanWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceColumnMean */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceColumnMean */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceColumnMean */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceColumnMean */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceColumnMean */


