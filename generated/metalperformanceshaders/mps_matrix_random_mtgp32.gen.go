// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MatrixRandomMTGP32] class.
var (
	MatrixRandomMTGP32Class     _MatrixRandomMTGP32Class
	MatrixRandomMTGP32ClassOnce sync.Once
)

func getMatrixRandomMTGP32Class() _MatrixRandomMTGP32Class {
	MatrixRandomMTGP32ClassOnce.Do(func() {
		MatrixRandomMTGP32Class = _MatrixRandomMTGP32Class{objc.GetClass("MPSMatrixRandomMTGP32")}
	})
	return MatrixRandomMTGP32Class
}

type _MatrixRandomMTGP32Class struct {
	class objc.Class
}





// An interface definition for the [MatrixRandomMTGP32] class.
type IMatrixRandomMTGP32 interface {
	IMatrixRandom
	

	// properties:


	

	// methods:
	SynchronizeState()
	SynchronizeStateOnCommandBuffer(commandBuffer unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixRandomMTGP32Class) Alloc() MatrixRandomMTGP32 {
	rv := objc.Send[MatrixRandomMTGP32](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixRandomMTGP32Class) New() MatrixRandomMTGP32 {
	rv := objc.Send[MatrixRandomMTGP32](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixRandomMTGP32) Init() MatrixRandomMTGP32 {
	rv := objc.Send[MatrixRandomMTGP32](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixRandomMTGP32) Autorelease() MatrixRandomMTGP32 {
	rv := objc.Send[MatrixRandomMTGP32](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixRandomMTGP32 creates a new MatrixRandomMTGP32 instance.
func NewMatrixRandomMTGP32() MatrixRandomMTGP32 {
	return getMatrixRandomMTGP32Class().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomMTGP32
type MatrixRandomMTGP32 struct {
	MatrixRandom
}

// MatrixRandomMTGP32From constructs a [MatrixRandomMTGP32] from an unsafe.Pointer.
func MatrixRandomMTGP32From(ptr unsafe.Pointer) MatrixRandomMTGP32 {
	return MatrixRandomMTGP32{
		MatrixRandom: MatrixRandomFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242864-initwithcoder
func NewMatrixRandomMTGP32WithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixRandomMTGP32 {
	instance := getMatrixRandomMTGP32Class().Alloc()
	rv := objc.Send[MatrixRandomMTGP32](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242865-initwithdevice
func NewMatrixRandomMTGP32WithDevice(device unsafe.Pointer) MatrixRandomMTGP32 {
	instance := getMatrixRandomMTGP32Class().Alloc()
	rv := objc.Send[MatrixRandomMTGP32](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242866-initwithdevice
func NewMatrixRandomMTGP32WithDeviceDestinationDataTypeSeed(device unsafe.Pointer, destinationDataType DataType, seed uint) MatrixRandomMTGP32 {
	instance := getMatrixRandomMTGP32Class().Alloc()
	rv := objc.Send[MatrixRandomMTGP32](instance.ID, objc.Sel("initWithDevice:destinationDataType:seed:"), device, destinationDataType, seed)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242867-initwithdevice
func NewMatrixRandomMTGP32WithDeviceDestinationDataTypeSeedDistributionDescriptor(device unsafe.Pointer, destinationDataType DataType, seed uint, distributionDescriptor IMatrixRandomDistributionDescriptor) MatrixRandomMTGP32 {
	instance := getMatrixRandomMTGP32Class().Alloc()
	rv := objc.Send[MatrixRandomMTGP32](instance.ID, objc.Sel("initWithDevice:destinationDataType:seed:distributionDescriptor:"), device, destinationDataType, seed, distributionDescriptor)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242868-synchronizestate
func (m_ MatrixRandomMTGP32) SynchronizeState() {
	objc.Send[objc.ID](m_.ID, objc.Sel("synchronizeState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242868-synchronizestateoncommandbuffer
func (m_ MatrixRandomMTGP32) SynchronizeStateOnCommandBuffer(commandBuffer unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("synchronizeStateOnCommandBuffer:"), commandBuffer)
}












