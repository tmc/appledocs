// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ReduceFeatureChannelsArgumentMin] class.
var (
	ReduceFeatureChannelsArgumentMinClass     _ReduceFeatureChannelsArgumentMinClass
	ReduceFeatureChannelsArgumentMinClassOnce sync.Once
)

func getReduceFeatureChannelsArgumentMinClass() _ReduceFeatureChannelsArgumentMinClass {
	ReduceFeatureChannelsArgumentMinClassOnce.Do(func() {
		ReduceFeatureChannelsArgumentMinClass = _ReduceFeatureChannelsArgumentMinClass{objc.GetClass("MPSNNReduceFeatureChannelsArgumentMin")}
	})
	return ReduceFeatureChannelsArgumentMinClass
}

type _ReduceFeatureChannelsArgumentMinClass struct {
	class objc.Class
}





// An interface definition for the [ReduceFeatureChannelsArgumentMin] class.
type IReduceFeatureChannelsArgumentMin interface {
	IReduceUnary
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _ReduceFeatureChannelsArgumentMinClass) Alloc() ReduceFeatureChannelsArgumentMin {
	rv := objc.Send[ReduceFeatureChannelsArgumentMin](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceFeatureChannelsArgumentMinClass) New() ReduceFeatureChannelsArgumentMin {
	rv := objc.Send[ReduceFeatureChannelsArgumentMin](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceFeatureChannelsArgumentMin) Init() ReduceFeatureChannelsArgumentMin {
	rv := objc.Send[ReduceFeatureChannelsArgumentMin](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceFeatureChannelsArgumentMin) Autorelease() ReduceFeatureChannelsArgumentMin {
	rv := objc.Send[ReduceFeatureChannelsArgumentMin](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceFeatureChannelsArgumentMin creates a new ReduceFeatureChannelsArgumentMin instance.
func NewReduceFeatureChannelsArgumentMin() ReduceFeatureChannelsArgumentMin {
	return getReduceFeatureChannelsArgumentMinClass().New()
}





// A reduction filter that returns the index of the location of the minimum value for each feature channel in an image.


// A reduction filter that returns the index of the location of the minimum value for each feature channel in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsArgumentMin
type ReduceFeatureChannelsArgumentMin struct {
	ReduceUnary
}

// ReduceFeatureChannelsArgumentMinFrom constructs a [ReduceFeatureChannelsArgumentMin] from an unsafe.Pointer.
//
// A reduction filter that returns the index of the location of the minimum value for each feature channel in an image.
func ReduceFeatureChannelsArgumentMinFrom(ptr unsafe.Pointer) ReduceFeatureChannelsArgumentMin {
	return ReduceFeatureChannelsArgumentMin{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsargumentmin/3197837-initwithcoder
func NewReduceFeatureChannelsArgumentMinWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceFeatureChannelsArgumentMin {
	instance := getReduceFeatureChannelsArgumentMinClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsArgumentMin](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsargumentmin/2976520-initwithdevice
func NewReduceFeatureChannelsArgumentMinWithDevice(device unsafe.Pointer) ReduceFeatureChannelsArgumentMin {
	instance := getReduceFeatureChannelsArgumentMinClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsArgumentMin](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























