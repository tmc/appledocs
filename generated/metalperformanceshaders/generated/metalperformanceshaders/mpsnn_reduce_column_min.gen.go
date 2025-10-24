// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReduceColumnMin */


/* debug [class_header]: Header for MPSNNReduceColumnMin */
// The class instance for the [ReduceColumnMin] class.
var (
	ReduceColumnMinClass     _ReduceColumnMinClass
	ReduceColumnMinClassOnce sync.Once
)

func getReduceColumnMinClass() _ReduceColumnMinClass {
	ReduceColumnMinClassOnce.Do(func() {
		ReduceColumnMinClass = _ReduceColumnMinClass{objc.GetClass("MPSNNReduceColumnMin")}
	})
	return ReduceColumnMinClass
}

type _ReduceColumnMinClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceColumnMin */
// An interface definition for the [ReduceColumnMin] class.
type IReduceColumnMin interface {
	IReduceUnary
	
/* debug [class_interface_properties]: Properties for ReduceColumnMin */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceColumnMin */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceColumnMin */
// Alloc allocates a new instance without initialization.
func (rc _ReduceColumnMinClass) Alloc() ReduceColumnMin {
	rv := objc.Send[ReduceColumnMin](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceColumnMinClass) New() ReduceColumnMin {
	rv := objc.Send[ReduceColumnMin](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceColumnMin) Init() ReduceColumnMin {
	rv := objc.Send[ReduceColumnMin](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceColumnMin) Autorelease() ReduceColumnMin {
	rv := objc.Send[ReduceColumnMin](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceColumnMin creates a new ReduceColumnMin instance.
func NewReduceColumnMin() ReduceColumnMin {
	return getReduceColumnMinClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceColumnMin */
// A reduction filter that returns the minimum value for each column in an image.


// A reduction filter that returns the minimum value for each column in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMin
type ReduceColumnMin struct {
	ReduceUnary
}

// ReduceColumnMinFrom constructs a [ReduceColumnMin] from an unsafe.Pointer.
//
// A reduction filter that returns the minimum value for each column in an image.
func ReduceColumnMinFrom(ptr unsafe.Pointer) ReduceColumnMin {
	return ReduceColumnMin{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceColumnMin */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmin/3197832-initwithcoder
func NewReduceColumnMinWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ReduceColumnMin {
	instance := getReduceColumnMinClass().Alloc()
	rv := objc.Send[ReduceColumnMin](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceColumnMinWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmin/2942542-initwithdevice
func NewReduceColumnMinWithDevice(device unsafe.Pointer) ReduceColumnMin {
	instance := getReduceColumnMinClass().Alloc()
	rv := objc.Send[ReduceColumnMin](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceColumnMinWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceColumnMin */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceColumnMin */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceColumnMin */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceColumnMin */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceColumnMin */


