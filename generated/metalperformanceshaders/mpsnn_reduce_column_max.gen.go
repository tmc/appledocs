// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [ReduceColumnMax] class.
type IReduceColumnMax interface {
	IReduceUnary
	

	// properties:


	

	// methods:


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmax/3197830-initwithcoder
func NewReduceColumnMaxWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceColumnMax {
	instance := getReduceColumnMaxClass().Alloc()
	rv := objc.Send[ReduceColumnMax](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmax/2942541-initwithdevice
func NewReduceColumnMaxWithDevice(device unsafe.Pointer) ReduceColumnMax {
	instance := getReduceColumnMaxClass().Alloc()
	rv := objc.Send[ReduceColumnMax](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























