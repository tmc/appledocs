// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ReduceColumnMean] class.
var (
	ReduceColumnMeanClass     _ReduceColumnMeanClass
	ReduceColumnMeanClassOnce sync.Once
)

func getReduceColumnMeanClass() _ReduceColumnMeanClass {
	ReduceColumnMeanClassOnce.Do(func() {
		ReduceColumnMeanClass = _ReduceColumnMeanClass{objc.GetClass("MPSNNReduceColumnMean")}
	})
	return ReduceColumnMeanClass
}

type _ReduceColumnMeanClass struct {
	class objc.Class
}





// An interface definition for the [ReduceColumnMean] class.
type IReduceColumnMean interface {
	IReduceUnary
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _ReduceColumnMeanClass) Alloc() ReduceColumnMean {
	rv := objc.Send[ReduceColumnMean](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceColumnMeanClass) New() ReduceColumnMean {
	rv := objc.Send[ReduceColumnMean](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceColumnMean) Init() ReduceColumnMean {
	rv := objc.Send[ReduceColumnMean](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceColumnMean) Autorelease() ReduceColumnMean {
	rv := objc.Send[ReduceColumnMean](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceColumnMean creates a new ReduceColumnMean instance.
func NewReduceColumnMean() ReduceColumnMean {
	return getReduceColumnMeanClass().New()
}





// A reduction filter that returns the mean value for each column in an image.


// A reduction filter that returns the mean value for each column in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMean
type ReduceColumnMean struct {
	ReduceUnary
}

// ReduceColumnMeanFrom constructs a [ReduceColumnMean] from an unsafe.Pointer.
//
// A reduction filter that returns the mean value for each column in an image.
func ReduceColumnMeanFrom(ptr unsafe.Pointer) ReduceColumnMean {
	return ReduceColumnMean{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmean/3197831-initwithcoder
func NewReduceColumnMeanWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceColumnMean {
	instance := getReduceColumnMeanClass().Alloc()
	rv := objc.Send[ReduceColumnMean](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmean/2942546-initwithdevice
func NewReduceColumnMeanWithDevice(device unsafe.Pointer) ReduceColumnMean {
	instance := getReduceColumnMeanClass().Alloc()
	rv := objc.Send[ReduceColumnMean](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























