// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MatrixRandomPhilox] class.
var (
	MatrixRandomPhiloxClass     _MatrixRandomPhiloxClass
	MatrixRandomPhiloxClassOnce sync.Once
)

func getMatrixRandomPhiloxClass() _MatrixRandomPhiloxClass {
	MatrixRandomPhiloxClassOnce.Do(func() {
		MatrixRandomPhiloxClass = _MatrixRandomPhiloxClass{objc.GetClass("MPSMatrixRandomPhilox")}
	})
	return MatrixRandomPhiloxClass
}

type _MatrixRandomPhiloxClass struct {
	class objc.Class
}





// An interface definition for the [MatrixRandomPhilox] class.
type IMatrixRandomPhilox interface {
	IMatrixRandom
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixRandomPhiloxClass) Alloc() MatrixRandomPhilox {
	rv := objc.Send[MatrixRandomPhilox](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixRandomPhiloxClass) New() MatrixRandomPhilox {
	rv := objc.Send[MatrixRandomPhilox](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixRandomPhilox) Init() MatrixRandomPhilox {
	rv := objc.Send[MatrixRandomPhilox](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixRandomPhilox) Autorelease() MatrixRandomPhilox {
	rv := objc.Send[MatrixRandomPhilox](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixRandomPhilox creates a new MatrixRandomPhilox instance.
func NewMatrixRandomPhilox() MatrixRandomPhilox {
	return getMatrixRandomPhiloxClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomPhilox
type MatrixRandomPhilox struct {
	MatrixRandom
}

// MatrixRandomPhiloxFrom constructs a [MatrixRandomPhilox] from an unsafe.Pointer.
func MatrixRandomPhiloxFrom(ptr unsafe.Pointer) MatrixRandomPhilox {
	return MatrixRandomPhilox{
		MatrixRandom: MatrixRandomFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomphilox/3242870-initwithcoder
func NewMatrixRandomPhiloxWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixRandomPhilox {
	instance := getMatrixRandomPhiloxClass().Alloc()
	rv := objc.Send[MatrixRandomPhilox](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomphilox/3242871-initwithdevice
func NewMatrixRandomPhiloxWithDevice(device unsafe.Pointer) MatrixRandomPhilox {
	instance := getMatrixRandomPhiloxClass().Alloc()
	rv := objc.Send[MatrixRandomPhilox](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomphilox/3242872-initwithdevice
func NewMatrixRandomPhiloxWithDeviceDestinationDataTypeSeed(device unsafe.Pointer, destinationDataType DataType, seed uint) MatrixRandomPhilox {
	instance := getMatrixRandomPhiloxClass().Alloc()
	rv := objc.Send[MatrixRandomPhilox](instance.ID, objc.Sel("initWithDevice:destinationDataType:seed:"), device, destinationDataType, seed)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomphilox/3242873-initwithdevice
func NewMatrixRandomPhiloxWithDeviceDestinationDataTypeSeedDistributionDescriptor(device unsafe.Pointer, destinationDataType DataType, seed uint, distributionDescriptor IMatrixRandomDistributionDescriptor) MatrixRandomPhilox {
	instance := getMatrixRandomPhiloxClass().Alloc()
	rv := objc.Send[MatrixRandomPhilox](instance.ID, objc.Sel("initWithDevice:destinationDataType:seed:distributionDescriptor:"), device, destinationDataType, seed, distributionDescriptor)
	rv.Autorelease()
	return rv
}



























