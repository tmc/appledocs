// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ReduceFeatureChannelsMax] class.
var (
	ReduceFeatureChannelsMaxClass     _ReduceFeatureChannelsMaxClass
	ReduceFeatureChannelsMaxClassOnce sync.Once
)

func getReduceFeatureChannelsMaxClass() _ReduceFeatureChannelsMaxClass {
	ReduceFeatureChannelsMaxClassOnce.Do(func() {
		ReduceFeatureChannelsMaxClass = _ReduceFeatureChannelsMaxClass{objc.GetClass("MPSNNReduceFeatureChannelsMax")}
	})
	return ReduceFeatureChannelsMaxClass
}

type _ReduceFeatureChannelsMaxClass struct {
	class objc.Class
}





// An interface definition for the [ReduceFeatureChannelsMax] class.
type IReduceFeatureChannelsMax interface {
	IReduceUnary
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _ReduceFeatureChannelsMaxClass) Alloc() ReduceFeatureChannelsMax {
	rv := objc.Send[ReduceFeatureChannelsMax](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceFeatureChannelsMaxClass) New() ReduceFeatureChannelsMax {
	rv := objc.Send[ReduceFeatureChannelsMax](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceFeatureChannelsMax) Init() ReduceFeatureChannelsMax {
	rv := objc.Send[ReduceFeatureChannelsMax](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceFeatureChannelsMax) Autorelease() ReduceFeatureChannelsMax {
	rv := objc.Send[ReduceFeatureChannelsMax](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceFeatureChannelsMax creates a new ReduceFeatureChannelsMax instance.
func NewReduceFeatureChannelsMax() ReduceFeatureChannelsMax {
	return getReduceFeatureChannelsMaxClass().New()
}





// A reduction filter that returns the maximum value for each feature channel in an image.


// A reduction filter that returns the maximum value for each feature channel in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMax
type ReduceFeatureChannelsMax struct {
	ReduceUnary
}

// ReduceFeatureChannelsMaxFrom constructs a [ReduceFeatureChannelsMax] from an unsafe.Pointer.
//
// A reduction filter that returns the maximum value for each feature channel in an image.
func ReduceFeatureChannelsMaxFrom(ptr unsafe.Pointer) ReduceFeatureChannelsMax {
	return ReduceFeatureChannelsMax{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmax/3197838-initwithcoder
func NewReduceFeatureChannelsMaxWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceFeatureChannelsMax {
	instance := getReduceFeatureChannelsMaxClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsMax](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmax/2942532-initwithdevice
func NewReduceFeatureChannelsMaxWithDevice(device unsafe.Pointer) ReduceFeatureChannelsMax {
	instance := getReduceFeatureChannelsMaxClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsMax](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























