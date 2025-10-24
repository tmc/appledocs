// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ReduceRowMean] class.
var (
	ReduceRowMeanClass     _ReduceRowMeanClass
	ReduceRowMeanClassOnce sync.Once
)

func getReduceRowMeanClass() _ReduceRowMeanClass {
	ReduceRowMeanClassOnce.Do(func() {
		ReduceRowMeanClass = _ReduceRowMeanClass{objc.GetClass("MPSNNReduceRowMean")}
	})
	return ReduceRowMeanClass
}

type _ReduceRowMeanClass struct {
	class objc.Class
}





// An interface definition for the [ReduceRowMean] class.
type IReduceRowMean interface {
	IReduceUnary
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _ReduceRowMeanClass) Alloc() ReduceRowMean {
	rv := objc.Send[ReduceRowMean](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceRowMeanClass) New() ReduceRowMean {
	rv := objc.Send[ReduceRowMean](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceRowMean) Init() ReduceRowMean {
	rv := objc.Send[ReduceRowMean](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceRowMean) Autorelease() ReduceRowMean {
	rv := objc.Send[ReduceRowMean](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceRowMean creates a new ReduceRowMean instance.
func NewReduceRowMean() ReduceRowMean {
	return getReduceRowMeanClass().New()
}





// A reduction filter that returns the mean value for each row in an image.


// A reduction filter that returns the mean value for each row in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMean
type ReduceRowMean struct {
	ReduceUnary
}

// ReduceRowMeanFrom constructs a [ReduceRowMean] from an unsafe.Pointer.
//
// A reduction filter that returns the mean value for each row in an image.
func ReduceRowMeanFrom(ptr unsafe.Pointer) ReduceRowMean {
	return ReduceRowMean{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowmean/3197843-initwithcoder
func NewReduceRowMeanWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceRowMean {
	instance := getReduceRowMeanClass().Alloc()
	rv := objc.Send[ReduceRowMean](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowmean/2942548-initwithdevice
func NewReduceRowMeanWithDevice(device unsafe.Pointer) ReduceRowMean {
	instance := getReduceRowMeanClass().Alloc()
	rv := objc.Send[ReduceRowMean](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























