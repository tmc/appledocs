// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReduceColumnMax */


/* debug [class_header]: Header for MPSNNReduceColumnMax */
// The class instance for the [ReduceColumnMax] class.
var (
	ReduceColumnMaxClass     _ReduceColumnMaxClass
	ReduceColumnMaxClassOnce sync.Once
)

func getReduceColumnMaxClass() _ReduceColumnMaxClass {
	ReduceColumnMaxClassOnce.Do(func() {
		ReduceColumnMaxClass = _ReduceColumnMaxClass{objc.GetClass("MPSNNReduceColumnMax")}
	})
	return ReduceColumnMaxClass
}

type _ReduceColumnMaxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceColumnMax */
// An interface definition for the [ReduceColumnMax] class.
type IReduceColumnMax interface {
	IReduceUnary
	
/* debug [class_interface_properties]: Properties for ReduceColumnMax */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceColumnMax */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceColumnMax */
// Alloc allocates a new instance without initialization.
func (rc _ReduceColumnMaxClass) Alloc() ReduceColumnMax {
	rv := objc.Send[ReduceColumnMax](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceColumnMaxClass) New() ReduceColumnMax {
	rv := objc.Send[ReduceColumnMax](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceColumnMax) Init() ReduceColumnMax {
	rv := objc.Send[ReduceColumnMax](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceColumnMax) Autorelease() ReduceColumnMax {
	rv := objc.Send[ReduceColumnMax](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceColumnMax creates a new ReduceColumnMax instance.
func NewReduceColumnMax() ReduceColumnMax {
	return getReduceColumnMaxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceColumnMax */
// A reduction filter that returns the maximum value for each column in an image.


// A reduction filter that returns the maximum value for each column in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMax
type ReduceColumnMax struct {
	ReduceUnary
}

// ReduceColumnMaxFrom constructs a [ReduceColumnMax] from an unsafe.Pointer.
//
// A reduction filter that returns the maximum value for each column in an image.
func ReduceColumnMaxFrom(ptr unsafe.Pointer) ReduceColumnMax {
	return ReduceColumnMax{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceColumnMax */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmax/3197830-initwithcoder
func NewReduceColumnMaxWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceColumnMax {
	instance := getReduceColumnMaxClass().Alloc()
	rv := objc.Send[ReduceColumnMax](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceColumnMaxWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmax/2942541-initwithdevice
func NewReduceColumnMaxWithDevice(device unsafe.Pointer) ReduceColumnMax {
	instance := getReduceColumnMaxClass().Alloc()
	rv := objc.Send[ReduceColumnMax](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceColumnMaxWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceColumnMax */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceColumnMax */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceColumnMax */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceColumnMax */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceColumnMax */


