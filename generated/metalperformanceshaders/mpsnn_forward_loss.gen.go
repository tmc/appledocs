// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ForwardLoss] class.
var (
	ForwardLossClass     _ForwardLossClass
	ForwardLossClassOnce sync.Once
)

func getForwardLossClass() _ForwardLossClass {
	ForwardLossClassOnce.Do(func() {
		ForwardLossClass = _ForwardLossClass{objc.GetClass("MPSNNForwardLoss")}
	})
	return ForwardLossClass
}

type _ForwardLossClass struct {
	class objc.Class
}

// An interface definition for the [ForwardLoss] class.
type IForwardLoss interface {
	objectivec.IObject
	EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages(commandBuffer objc.ID, sourceImages unsafe.Pointer, labels unsafe.Pointer, weights unsafe.Pointer, destinationStates unsafe.Pointer, destinationImages unsafe.Pointer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss
type ForwardLoss struct {
	objectivec.Object
}

// ForwardLossFrom constructs a [ForwardLoss] from an unsafe.Pointer.
func ForwardLossFrom(ptr unsafe.Pointer) ForwardLoss {
	return ForwardLoss{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _ForwardLossClass) Alloc() ForwardLoss {
	rv := objc.Send[ForwardLoss](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _ForwardLossClass) New() ForwardLoss {
	rv := objc.Send[ForwardLoss](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ ForwardLoss) Init() ForwardLoss {
	rv := objc.Send[ForwardLoss](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ ForwardLoss) Autorelease() ForwardLoss {
	rv := objc.Send[ForwardLoss](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewForwardLoss creates a new ForwardLoss instance.
func NewForwardLoss() ForwardLoss {
	return getForwardLossClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/init(coder:device:)
func NewForwardLossWithCoderDevice(aDecoder unsafe.Pointer, device objc.ID) ForwardLoss {
	instance := getForwardLossClass().Alloc()
	rv := objc.Send[ForwardLoss](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/init(device:lossDescriptor:)
func NewForwardLossWithDeviceLossDescriptor(device objc.ID, lossDescriptor unsafe.Pointer) ForwardLoss {
	instance := getForwardLossClass().Alloc()
	rv := objc.Send[ForwardLoss](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/encodeBatch(commandBuffer:sourceImages:labels:weights:destinationStates:destinationImages:)
func (f_ ForwardLoss) EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages(commandBuffer objc.ID, sourceImages unsafe.Pointer, labels unsafe.Pointer, weights unsafe.Pointer, destinationStates unsafe.Pointer, destinationImages unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:weights:destinationStates:destinationImages:"), commandBuffer, sourceImages, labels, weights, destinationStates, destinationImages)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/delta
func (f_ ForwardLoss) Delta() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("delta"))
	return rv
}

// SetDelta sets the value of the delta property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/delta
func (f_ ForwardLoss) SetDelta(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDelta:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/epsilon
func (f_ ForwardLoss) Epsilon() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("epsilon"))
	return rv
}

// SetEpsilon sets the value of the epsilon property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/epsilon
func (f_ ForwardLoss) SetEpsilon(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setEpsilon:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/reduceAcrossBatch
func (f_ ForwardLoss) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss/reductionType
func (f_ ForwardLoss) ReductionType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("reductionType"))
	return rv
}
