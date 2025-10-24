// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLComputePlanCost */


/* debug [class_header]: Header for MLComputePlanCost */
// The class instance for the [ComputePlanCost] class.
var (
	ComputePlanCostClass     _ComputePlanCostClass
	ComputePlanCostClassOnce sync.Once
)

func getComputePlanCostClass() _ComputePlanCostClass {
	ComputePlanCostClassOnce.Do(func() {
		ComputePlanCostClass = _ComputePlanCostClass{objc.GetClass("MLComputePlanCost")}
	})
	return ComputePlanCostClass
}

type _ComputePlanCostClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ComputePlanCost */
// An interface definition for the [ComputePlanCost] class.
type IComputePlanCost interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ComputePlanCost */
	// properties:
	Weight() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ComputePlanCost */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ComputePlanCost */
// Alloc allocates a new instance without initialization.
func (cc _ComputePlanCostClass) Alloc() ComputePlanCost {
	rv := objc.Send[ComputePlanCost](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ComputePlanCostClass) New() ComputePlanCost {
	rv := objc.Send[ComputePlanCost](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComputePlanCost) Init() ComputePlanCost {
	rv := objc.Send[ComputePlanCost](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComputePlanCost) Autorelease() ComputePlanCost {
	rv := objc.Send[ComputePlanCost](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComputePlanCost creates a new ComputePlanCost instance.
func NewComputePlanCost() ComputePlanCost {
	return getComputePlanCostClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ComputePlanCost */
// A class that represents the estimated cost of executing a layer or operation.


// A class that represents the estimated cost of executing a layer or operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlanCost
type ComputePlanCost struct {
	objectivec.Object
}

// ComputePlanCostFrom constructs a [ComputePlanCost] from an unsafe.Pointer.
//
// A class that represents the estimated cost of executing a layer or operation.
func ComputePlanCostFrom(ptr unsafe.Pointer) ComputePlanCost {
	return ComputePlanCost{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ComputePlanCost *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ComputePlanCost */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ComputePlanCost */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ComputePlanCost */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ComputePlanCost */

// The estimated workload of executing the operation over the total model execution. The value is between [0.0, 1.0].
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlanCost/weight
func (c_ ComputePlanCost) Weight() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("weight"))
	return rv
}/* debug [instance_properties/getter]: weight */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLComputePlanCost */



