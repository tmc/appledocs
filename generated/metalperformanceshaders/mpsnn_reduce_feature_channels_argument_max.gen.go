// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ReduceFeatureChannelsArgumentMax] class.
var (
	ReduceFeatureChannelsArgumentMaxClass     _ReduceFeatureChannelsArgumentMaxClass
	ReduceFeatureChannelsArgumentMaxClassOnce sync.Once
)

func getReduceFeatureChannelsArgumentMaxClass() _ReduceFeatureChannelsArgumentMaxClass {
	ReduceFeatureChannelsArgumentMaxClassOnce.Do(func() {
		ReduceFeatureChannelsArgumentMaxClass = _ReduceFeatureChannelsArgumentMaxClass{objc.GetClass("MPSNNReduceFeatureChannelsArgumentMax")}
	})
	return ReduceFeatureChannelsArgumentMaxClass
}

type _ReduceFeatureChannelsArgumentMaxClass struct {
	class objc.Class
}





// An interface definition for the [ReduceFeatureChannelsArgumentMax] class.
type IReduceFeatureChannelsArgumentMax interface {
	IReduceUnary
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _ReduceFeatureChannelsArgumentMaxClass) Alloc() ReduceFeatureChannelsArgumentMax {
	rv := objc.Send[ReduceFeatureChannelsArgumentMax](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceFeatureChannelsArgumentMaxClass) New() ReduceFeatureChannelsArgumentMax {
	rv := objc.Send[ReduceFeatureChannelsArgumentMax](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceFeatureChannelsArgumentMax) Init() ReduceFeatureChannelsArgumentMax {
	rv := objc.Send[ReduceFeatureChannelsArgumentMax](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceFeatureChannelsArgumentMax) Autorelease() ReduceFeatureChannelsArgumentMax {
	rv := objc.Send[ReduceFeatureChannelsArgumentMax](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceFeatureChannelsArgumentMax creates a new ReduceFeatureChannelsArgumentMax instance.
func NewReduceFeatureChannelsArgumentMax() ReduceFeatureChannelsArgumentMax {
	return getReduceFeatureChannelsArgumentMaxClass().New()
}





// A reduction filter that returns the index of the location of the maximum value for each feature channel in an image.


// A reduction filter that returns the index of the location of the maximum value for each feature channel in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsArgumentMax
type ReduceFeatureChannelsArgumentMax struct {
	ReduceUnary
}

// ReduceFeatureChannelsArgumentMaxFrom constructs a [ReduceFeatureChannelsArgumentMax] from an unsafe.Pointer.
//
// A reduction filter that returns the index of the location of the maximum value for each feature channel in an image.
func ReduceFeatureChannelsArgumentMaxFrom(ptr unsafe.Pointer) ReduceFeatureChannelsArgumentMax {
	return ReduceFeatureChannelsArgumentMax{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsargumentmax/3197836-initwithcoder
func NewReduceFeatureChannelsArgumentMaxWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceFeatureChannelsArgumentMax {
	instance := getReduceFeatureChannelsArgumentMaxClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsArgumentMax](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsargumentmax/2976518-initwithdevice
func NewReduceFeatureChannelsArgumentMaxWithDevice(device unsafe.Pointer) ReduceFeatureChannelsArgumentMax {
	instance := getReduceFeatureChannelsArgumentMaxClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsArgumentMax](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























