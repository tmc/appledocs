// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ReduceFeatureChannelsSum] class.
var (
	ReduceFeatureChannelsSumClass     _ReduceFeatureChannelsSumClass
	ReduceFeatureChannelsSumClassOnce sync.Once
)

func getReduceFeatureChannelsSumClass() _ReduceFeatureChannelsSumClass {
	ReduceFeatureChannelsSumClassOnce.Do(func() {
		ReduceFeatureChannelsSumClass = _ReduceFeatureChannelsSumClass{objc.GetClass("MPSNNReduceFeatureChannelsSum")}
	})
	return ReduceFeatureChannelsSumClass
}

type _ReduceFeatureChannelsSumClass struct {
	class objc.Class
}





// An interface definition for the [ReduceFeatureChannelsSum] class.
type IReduceFeatureChannelsSum interface {
	IReduceUnary
	

	// properties:
	Weight() objectivec.IObject
	SetWeight(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _ReduceFeatureChannelsSumClass) Alloc() ReduceFeatureChannelsSum {
	rv := objc.Send[ReduceFeatureChannelsSum](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceFeatureChannelsSumClass) New() ReduceFeatureChannelsSum {
	rv := objc.Send[ReduceFeatureChannelsSum](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceFeatureChannelsSum) Init() ReduceFeatureChannelsSum {
	rv := objc.Send[ReduceFeatureChannelsSum](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceFeatureChannelsSum) Autorelease() ReduceFeatureChannelsSum {
	rv := objc.Send[ReduceFeatureChannelsSum](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceFeatureChannelsSum creates a new ReduceFeatureChannelsSum instance.
func NewReduceFeatureChannelsSum() ReduceFeatureChannelsSum {
	return getReduceFeatureChannelsSumClass().New()
}





// A reduction filter that returns the sum of all values for each feature channel in an image.


// A reduction filter that returns the sum of all values for each feature channel in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsSum
type ReduceFeatureChannelsSum struct {
	ReduceUnary
}

// ReduceFeatureChannelsSumFrom constructs a [ReduceFeatureChannelsSum] from an unsafe.Pointer.
//
// A reduction filter that returns the sum of all values for each feature channel in an image.
func ReduceFeatureChannelsSumFrom(ptr unsafe.Pointer) ReduceFeatureChannelsSum {
	return ReduceFeatureChannelsSum{
		ReduceUnary: ReduceUnaryFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelssum/3197841-initwithcoder
func NewReduceFeatureChannelsSumWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceFeatureChannelsSum {
	instance := getReduceFeatureChannelsSumClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsSum](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelssum/2942538-initwithdevice
func NewReduceFeatureChannelsSumWithDevice(device unsafe.Pointer) ReduceFeatureChannelsSum {
	instance := getReduceFeatureChannelsSumClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsSum](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelssum/2942545-weight
func (r_ ReduceFeatureChannelsSum) Weight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("weight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelssum/2942545-weight
func (r_ ReduceFeatureChannelsSum) SetWeight(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setWeight:"), value)
}







