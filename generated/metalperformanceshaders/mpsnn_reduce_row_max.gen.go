// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReduceRowMax */


/* debug [class_header]: Header for MPSNNReduceRowMax */
// The class instance for the [ReduceRowMax] class.
var (
	ReduceRowMaxClass     _ReduceRowMaxClass
	ReduceRowMaxClassOnce sync.Once
)

func getReduceRowMaxClass() _ReduceRowMaxClass {
	ReduceRowMaxClassOnce.Do(func() {
		ReduceRowMaxClass = _ReduceRowMaxClass{objc.GetClass("MPSNNReduceRowMax")}
	})
	return ReduceRowMaxClass
}

type _ReduceRowMaxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceRowMax */
// An interface definition for the [ReduceRowMax] class.
type IReduceRowMax interface {
	IReduceUnary
	
/* debug [class_interface_properties]: Properties for ReduceRowMax */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceRowMax */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceRowMax */
// Alloc allocates a new instance without initialization.
func (rc _ReduceRowMaxClass) Alloc() ReduceRowMax {
	rv := objc.Send[ReduceRowMax](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceRowMaxClass) New() ReduceRowMax {
	rv := objc.Send[ReduceRowMax](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceRowMax) Init() ReduceRowMax {
	rv := objc.Send[ReduceRowMax](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceRowMax) Autorelease() ReduceRowMax {
	rv := objc.Send[ReduceRowMax](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceRowMax creates a new ReduceRowMax instance.
func NewReduceRowMax() ReduceRowMax {
	return getReduceRowMaxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceRowMax */
// A reduction filter that returns the maximum value for each row in an image.


// A reduction filter that returns the maximum value for each row in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMax
type ReduceRowMax struct {
	ReduceUnary
}

// ReduceRowMaxFrom constructs a [ReduceRowMax] from an unsafe.Pointer.
//
// A reduction filter that returns the maximum value for each row in an image.
func ReduceRowMaxFrom(ptr unsafe.Pointer) ReduceRowMax {
	return ReduceRowMax{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceRowMax */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowmax/3197842-initwithcoder
func NewReduceRowMaxWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceRowMax {
	instance := getReduceRowMaxClass().Alloc()
	rv := objc.Send[ReduceRowMax](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceRowMaxWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowmax/2942559-initwithdevice
func NewReduceRowMaxWithDevice(device unsafe.Pointer) ReduceRowMax {
	instance := getReduceRowMaxClass().Alloc()
	rv := objc.Send[ReduceRowMax](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceRowMaxWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceRowMax */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceRowMax */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceRowMax */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceRowMax */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceRowMax */


