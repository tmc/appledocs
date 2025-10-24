// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ReduceFeatureChannelsMean] class.
var (
	ReduceFeatureChannelsMeanClass     _ReduceFeatureChannelsMeanClass
	ReduceFeatureChannelsMeanClassOnce sync.Once
)

func getReduceFeatureChannelsMeanClass() _ReduceFeatureChannelsMeanClass {
	ReduceFeatureChannelsMeanClassOnce.Do(func() {
		ReduceFeatureChannelsMeanClass = _ReduceFeatureChannelsMeanClass{objc.GetClass("MPSNNReduceFeatureChannelsMean")}
	})
	return ReduceFeatureChannelsMeanClass
}

type _ReduceFeatureChannelsMeanClass struct {
	class objc.Class
}





// An interface definition for the [ReduceFeatureChannelsMean] class.
type IReduceFeatureChannelsMean interface {
	IReduceUnary
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _ReduceFeatureChannelsMeanClass) Alloc() ReduceFeatureChannelsMean {
	rv := objc.Send[ReduceFeatureChannelsMean](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceFeatureChannelsMeanClass) New() ReduceFeatureChannelsMean {
	rv := objc.Send[ReduceFeatureChannelsMean](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceFeatureChannelsMean) Init() ReduceFeatureChannelsMean {
	rv := objc.Send[ReduceFeatureChannelsMean](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceFeatureChannelsMean) Autorelease() ReduceFeatureChannelsMean {
	rv := objc.Send[ReduceFeatureChannelsMean](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceFeatureChannelsMean creates a new ReduceFeatureChannelsMean instance.
func NewReduceFeatureChannelsMean() ReduceFeatureChannelsMean {
	return getReduceFeatureChannelsMeanClass().New()
}





// A reduction filter that returns the mean value for each feature channel in an image.


// A reduction filter that returns the mean value for each feature channel in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMean
type ReduceFeatureChannelsMean struct {
	ReduceUnary
}

// ReduceFeatureChannelsMeanFrom constructs a [ReduceFeatureChannelsMean] from an unsafe.Pointer.
//
// A reduction filter that returns the mean value for each feature channel in an image.
func ReduceFeatureChannelsMeanFrom(ptr unsafe.Pointer) ReduceFeatureChannelsMean {
	return ReduceFeatureChannelsMean{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmean/3197839-initwithcoder
func NewReduceFeatureChannelsMeanWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceFeatureChannelsMean {
	instance := getReduceFeatureChannelsMeanClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsMean](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmean/2942557-initwithdevice
func NewReduceFeatureChannelsMeanWithDevice(device unsafe.Pointer) ReduceFeatureChannelsMean {
	instance := getReduceFeatureChannelsMeanClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsMean](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























