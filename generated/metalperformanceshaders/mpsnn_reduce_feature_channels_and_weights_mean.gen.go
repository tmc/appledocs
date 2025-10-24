// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ReduceFeatureChannelsAndWeightsMean] class.
var (
	ReduceFeatureChannelsAndWeightsMeanClass     _ReduceFeatureChannelsAndWeightsMeanClass
	ReduceFeatureChannelsAndWeightsMeanClassOnce sync.Once
)

func getReduceFeatureChannelsAndWeightsMeanClass() _ReduceFeatureChannelsAndWeightsMeanClass {
	ReduceFeatureChannelsAndWeightsMeanClassOnce.Do(func() {
		ReduceFeatureChannelsAndWeightsMeanClass = _ReduceFeatureChannelsAndWeightsMeanClass{objc.GetClass("MPSNNReduceFeatureChannelsAndWeightsMean")}
	})
	return ReduceFeatureChannelsAndWeightsMeanClass
}

type _ReduceFeatureChannelsAndWeightsMeanClass struct {
	class objc.Class
}





// An interface definition for the [ReduceFeatureChannelsAndWeightsMean] class.
type IReduceFeatureChannelsAndWeightsMean interface {
	IReduceBinary
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _ReduceFeatureChannelsAndWeightsMeanClass) Alloc() ReduceFeatureChannelsAndWeightsMean {
	rv := objc.Send[ReduceFeatureChannelsAndWeightsMean](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceFeatureChannelsAndWeightsMeanClass) New() ReduceFeatureChannelsAndWeightsMean {
	rv := objc.Send[ReduceFeatureChannelsAndWeightsMean](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceFeatureChannelsAndWeightsMean) Init() ReduceFeatureChannelsAndWeightsMean {
	rv := objc.Send[ReduceFeatureChannelsAndWeightsMean](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceFeatureChannelsAndWeightsMean) Autorelease() ReduceFeatureChannelsAndWeightsMean {
	rv := objc.Send[ReduceFeatureChannelsAndWeightsMean](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceFeatureChannelsAndWeightsMean creates a new ReduceFeatureChannelsAndWeightsMean instance.
func NewReduceFeatureChannelsAndWeightsMean() ReduceFeatureChannelsAndWeightsMean {
	return getReduceFeatureChannelsAndWeightsMeanClass().New()
}





// A reduction filter that returns the weighted sum for each feature channel in an image.


// A reduction filter that returns the weighted sum for each feature channel in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsAndWeightsMean
type ReduceFeatureChannelsAndWeightsMean struct {
	ReduceBinary
}

// ReduceFeatureChannelsAndWeightsMeanFrom constructs a [ReduceFeatureChannelsAndWeightsMean] from an unsafe.Pointer.
//
// A reduction filter that returns the weighted sum for each feature channel in an image.
func ReduceFeatureChannelsAndWeightsMeanFrom(ptr unsafe.Pointer) ReduceFeatureChannelsAndWeightsMean {
	return ReduceFeatureChannelsAndWeightsMean{
		ReduceBinary: ReduceBinaryFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightsmean/3197834-initwithcoder
func NewReduceFeatureChannelsAndWeightsMeanWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceFeatureChannelsAndWeightsMean {
	instance := getReduceFeatureChannelsAndWeightsMeanClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsAndWeightsMean](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightsmean/2942537-initwithdevice
func NewReduceFeatureChannelsAndWeightsMeanWithDevice(device unsafe.Pointer) ReduceFeatureChannelsAndWeightsMean {
	instance := getReduceFeatureChannelsAndWeightsMeanClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsAndWeightsMean](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























