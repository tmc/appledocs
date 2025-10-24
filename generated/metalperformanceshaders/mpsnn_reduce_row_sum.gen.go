// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [ReduceRowSum] class.
type IReduceRowSum interface {
	IReduceUnary
	

	// properties:


	

	// methods:


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowsum/3197845-initwithcoder
func NewReduceRowSumWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceRowSum {
	instance := getReduceRowSumClass().Alloc()
	rv := objc.Send[ReduceRowSum](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowsum/2942536-initwithdevice
func NewReduceRowSumWithDevice(device unsafe.Pointer) ReduceRowSum {
	instance := getReduceRowSumClass().Alloc()
	rv := objc.Send[ReduceRowSum](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























