// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [ReduceColumnMin] class.
type IReduceColumnMin interface {
	IReduceUnary
	

	// properties:


	

	// methods:


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmin/3197832-initwithcoder
func NewReduceColumnMinWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceColumnMin {
	instance := getReduceColumnMinClass().Alloc()
	rv := objc.Send[ReduceColumnMin](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmin/2942542-initwithdevice
func NewReduceColumnMinWithDevice(device unsafe.Pointer) ReduceColumnMin {
	instance := getReduceColumnMinClass().Alloc()
	rv := objc.Send[ReduceColumnMin](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























