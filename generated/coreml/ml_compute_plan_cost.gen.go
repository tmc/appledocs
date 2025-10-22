// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ComputePlanCost] class.
type IComputePlanCost interface {
	objectivec.IObject
	Weight() float64
}

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

// Alloc allocates a new instance without initialization.
func (cc _ComputePlanCostClass) Alloc() ComputePlanCost {
	rv := objc.Send[ComputePlanCost](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The estimated workload of executing the operation over the total model execution. The value is between [0.0, 1.0].
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlanCost/weight

func (c_ ComputePlanCost) Weight() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("weight"))
	return rv
}



