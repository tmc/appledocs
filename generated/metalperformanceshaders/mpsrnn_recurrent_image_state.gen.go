// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [RNNRecurrentImageState] class.
var (
	RNNRecurrentImageStateClass     _RNNRecurrentImageStateClass
	RNNRecurrentImageStateClassOnce sync.Once
)

func getRNNRecurrentImageStateClass() _RNNRecurrentImageStateClass {
	RNNRecurrentImageStateClassOnce.Do(func() {
		RNNRecurrentImageStateClass = _RNNRecurrentImageStateClass{objc.GetClass("MPSRNNRecurrentImageState")}
	})
	return RNNRecurrentImageStateClass
}

type _RNNRecurrentImageStateClass struct {
	class objc.Class
}





// An interface definition for the [RNNRecurrentImageState] class.
type IRNNRecurrentImageState interface {
	IState
	

	// properties:


	

	// methods:
	GetMemoryCellImage()
	GetMemoryCellImageForLayerIndex(layerIndex uint) IImage
	GetRecurrentOutputImage()
	GetRecurrentOutputImageForLayerIndex(layerIndex uint) IImage


}





// Alloc allocates a new instance without initialization.
func (rc _RNNRecurrentImageStateClass) Alloc() RNNRecurrentImageState {
	rv := objc.Send[RNNRecurrentImageState](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RNNRecurrentImageStateClass) New() RNNRecurrentImageState {
	rv := objc.Send[RNNRecurrentImageState](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RNNRecurrentImageState) Init() RNNRecurrentImageState {
	rv := objc.Send[RNNRecurrentImageState](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RNNRecurrentImageState) Autorelease() RNNRecurrentImageState {
	rv := objc.Send[RNNRecurrentImageState](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRNNRecurrentImageState creates a new RNNRecurrentImageState instance.
func NewRNNRecurrentImageState() RNNRecurrentImageState {
	return getRNNRecurrentImageStateClass().New()
}





// A class that holds all the data that’s passed from one sequence iteration of the image-based recurrent neural network layer (stack) to the next.


// A class that holds all the data that’s passed from one sequence iteration of the image-based recurrent neural network layer (stack) to the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNRecurrentImageState
type RNNRecurrentImageState struct {
	State
}

// RNNRecurrentImageStateFrom constructs a [RNNRecurrentImageState] from an unsafe.Pointer.
//
// A class that holds all the data that’s passed from one sequence iteration of the image-based recurrent neural network layer (stack) to the next.
func RNNRecurrentImageStateFrom(ptr unsafe.Pointer) RNNRecurrentImageState {
	return RNNRecurrentImageState{
		State: StateFrom(ptr),
	}
}




















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentimagestate/2865740-getmemorycellimage
func (r_ RNNRecurrentImageState) GetMemoryCellImage() {
	objc.Send[objc.ID](r_.ID, objc.Sel("getMemoryCellImage"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentimagestate/2865740-getmemorycellimageforlayerindex
func (r_ RNNRecurrentImageState) GetMemoryCellImageForLayerIndex(layerIndex uint) IImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("getMemoryCellImageForLayerIndex:"), layerIndex)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentimagestate/2865742-getrecurrentoutputimage
func (r_ RNNRecurrentImageState) GetRecurrentOutputImage() {
	objc.Send[objc.ID](r_.ID, objc.Sel("getRecurrentOutputImage"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentimagestate/2865742-getrecurrentoutputimageforlayeri
func (r_ RNNRecurrentImageState) GetRecurrentOutputImageForLayerIndex(layerIndex uint) IImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("getRecurrentOutputImageForLayerIndex:"), layerIndex)
	return rv
}













