// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReduceRowMin */


/* debug [class_header]: Header for MPSNNReduceRowMin */
// The class instance for the [ReduceRowMin] class.
var (
	ReduceRowMinClass     _ReduceRowMinClass
	ReduceRowMinClassOnce sync.Once
)

func getReduceRowMinClass() _ReduceRowMinClass {
	ReduceRowMinClassOnce.Do(func() {
		ReduceRowMinClass = _ReduceRowMinClass{objc.GetClass("MPSNNReduceRowMin")}
	})
	return ReduceRowMinClass
}

type _ReduceRowMinClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceRowMin */
// An interface definition for the [ReduceRowMin] class.
type IReduceRowMin interface {
	IReduceUnary
	
/* debug [class_interface_properties]: Properties for ReduceRowMin */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceRowMin */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceRowMin */
// Alloc allocates a new instance without initialization.
func (rc _ReduceRowMinClass) Alloc() ReduceRowMin {
	rv := objc.Send[ReduceRowMin](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceRowMinClass) New() ReduceRowMin {
	rv := objc.Send[ReduceRowMin](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceRowMin) Init() ReduceRowMin {
	rv := objc.Send[ReduceRowMin](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceRowMin) Autorelease() ReduceRowMin {
	rv := objc.Send[ReduceRowMin](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceRowMin creates a new ReduceRowMin instance.
func NewReduceRowMin() ReduceRowMin {
	return getReduceRowMinClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceRowMin */
// A reduction filter that returns the minimum value for each row in an image.


// A reduction filter that returns the minimum value for each row in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMin
type ReduceRowMin struct {
	ReduceUnary
}

// ReduceRowMinFrom constructs a [ReduceRowMin] from an unsafe.Pointer.
//
// A reduction filter that returns the minimum value for each row in an image.
func ReduceRowMinFrom(ptr unsafe.Pointer) ReduceRowMin {
	return ReduceRowMin{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceRowMin */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowmin/3197844-initwithcoder
func NewReduceRowMinWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ReduceRowMin {
	instance := getReduceRowMinClass().Alloc()
	rv := objc.Send[ReduceRowMin](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceRowMinWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowmin/2942555-initwithdevice
func NewReduceRowMinWithDevice(device unsafe.Pointer) ReduceRowMin {
	instance := getReduceRowMinClass().Alloc()
	rv := objc.Send[ReduceRowMin](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceRowMinWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceRowMin */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceRowMin */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceRowMin */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceRowMin */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceRowMin */


