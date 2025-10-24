// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNBatchNormalizationStatistics] class.
var (
	CNNBatchNormalizationStatisticsClass     _CNNBatchNormalizationStatisticsClass
	CNNBatchNormalizationStatisticsClassOnce sync.Once
)

func getCNNBatchNormalizationStatisticsClass() _CNNBatchNormalizationStatisticsClass {
	CNNBatchNormalizationStatisticsClassOnce.Do(func() {
		CNNBatchNormalizationStatisticsClass = _CNNBatchNormalizationStatisticsClass{objc.GetClass("MPSCNNBatchNormalizationStatistics")}
	})
	return CNNBatchNormalizationStatisticsClass
}

type _CNNBatchNormalizationStatisticsClass struct {
	class objc.Class
}





// An interface definition for the [CNNBatchNormalizationStatistics] class.
type ICNNBatchNormalizationStatistics interface {
	ICNNKernel
	

	// properties:


	

	// methods:
	EncodeBatch()
	EncodeBatchToCommandBufferSourceImagesBatchNormalizationState(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, batchNormalizationState ICNNBatchNormalizationState)


}





// Alloc allocates a new instance without initialization.
func (cc _CNNBatchNormalizationStatisticsClass) Alloc() CNNBatchNormalizationStatistics {
	rv := objc.Send[CNNBatchNormalizationStatistics](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNBatchNormalizationStatisticsClass) New() CNNBatchNormalizationStatistics {
	rv := objc.Send[CNNBatchNormalizationStatistics](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNBatchNormalizationStatistics) Init() CNNBatchNormalizationStatistics {
	rv := objc.Send[CNNBatchNormalizationStatistics](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNBatchNormalizationStatistics) Autorelease() CNNBatchNormalizationStatistics {
	rv := objc.Send[CNNBatchNormalizationStatistics](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNBatchNormalizationStatistics creates a new CNNBatchNormalizationStatistics instance.
func NewCNNBatchNormalizationStatistics() CNNBatchNormalizationStatistics {
	return getCNNBatchNormalizationStatisticsClass().New()
}





// An object that stores statistics required to execute batch normalization.


// An object that stores statistics required to execute batch normalization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationStatistics
type CNNBatchNormalizationStatistics struct {
	CNNKernel
}

// CNNBatchNormalizationStatisticsFrom constructs a [CNNBatchNormalizationStatistics] from an unsafe.Pointer.
//
// An object that stores statistics required to execute batch normalization.
func CNNBatchNormalizationStatisticsFrom(ptr unsafe.Pointer) CNNBatchNormalizationStatistics {
	return CNNBatchNormalizationStatistics{
		CNNKernel: CNNKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatistics/2942578-initwithcoder
func NewCNNBatchNormalizationStatisticsWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNBatchNormalizationStatistics {
	instance := getCNNBatchNormalizationStatisticsClass().Alloc()
	rv := objc.Send[CNNBatchNormalizationStatistics](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatistics/2953968-initwithdevice
func NewCNNBatchNormalizationStatisticsWithDevice(device unsafe.Pointer) CNNBatchNormalizationStatistics {
	instance := getCNNBatchNormalizationStatisticsClass().Alloc()
	rv := objc.Send[CNNBatchNormalizationStatistics](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatistics/2942584-encodebatch
func (c_ CNNBatchNormalizationStatistics) EncodeBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatistics/2942584-encodebatchtocommandbuffer
func (c_ CNNBatchNormalizationStatistics) EncodeBatchToCommandBufferSourceImagesBatchNormalizationState(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, batchNormalizationState ICNNBatchNormalizationState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:batchNormalizationState:"), commandBuffer, sourceImages, batchNormalizationState)
}












