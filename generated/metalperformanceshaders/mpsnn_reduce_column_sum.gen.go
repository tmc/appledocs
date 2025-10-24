// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [ReduceColumnSum] class.
type IReduceColumnSum interface {
	IReduceUnary
	

	// properties:


	

	// methods:


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnsum/3197833-initwithcoder
func NewReduceColumnSumWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceColumnSum {
	instance := getReduceColumnSumClass().Alloc()
	rv := objc.Send[ReduceColumnSum](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnsum/2942540-initwithdevice
func NewReduceColumnSumWithDevice(device unsafe.Pointer) ReduceColumnSum {
	instance := getReduceColumnSumClass().Alloc()
	rv := objc.Send[ReduceColumnSum](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























