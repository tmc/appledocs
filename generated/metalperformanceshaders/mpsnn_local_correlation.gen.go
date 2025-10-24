// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [LocalCorrelation] class.
var (
	LocalCorrelationClass     _LocalCorrelationClass
	LocalCorrelationClassOnce sync.Once
)

func getLocalCorrelationClass() _LocalCorrelationClass {
	LocalCorrelationClassOnce.Do(func() {
		LocalCorrelationClass = _LocalCorrelationClass{objc.GetClass("MPSNNLocalCorrelation")}
	})
	return LocalCorrelationClass
}

type _LocalCorrelationClass struct {
	class objc.Class
}





// An interface definition for the [LocalCorrelation] class.
type ILocalCorrelation interface {
	IReduceBinary
	

	// properties:
	StrideInX() objectivec.IObject
	SetStrideInX(value objectivec.IObject)
	StrideInY() objectivec.IObject
	SetStrideInY(value objectivec.IObject)
	WindowInX() objectivec.IObject
	SetWindowInX(value objectivec.IObject)
	WindowInY() objectivec.IObject
	SetWindowInY(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (lc _LocalCorrelationClass) Alloc() LocalCorrelation {
	rv := objc.Send[LocalCorrelation](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LocalCorrelationClass) New() LocalCorrelation {
	rv := objc.Send[LocalCorrelation](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LocalCorrelation) Init() LocalCorrelation {
	rv := objc.Send[LocalCorrelation](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LocalCorrelation) Autorelease() LocalCorrelation {
	rv := objc.Send[LocalCorrelation](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocalCorrelation creates a new LocalCorrelation instance.
func NewLocalCorrelation() LocalCorrelation {
	return getLocalCorrelationClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation
type LocalCorrelation struct {
	ReduceBinary
}

// LocalCorrelationFrom constructs a [LocalCorrelation] from an unsafe.Pointer.
func LocalCorrelationFrom(ptr unsafe.Pointer) LocalCorrelation {
	return LocalCorrelation{
		ReduceBinary: ReduceBinaryFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3197829-initwithcoder
func NewLocalCorrelationWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) LocalCorrelation {
	instance := getLocalCorrelationClass().Alloc()
	rv := objc.Send[LocalCorrelation](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131875-initwithdevice
func NewLocalCorrelationWithDevice(device unsafe.Pointer) LocalCorrelation {
	instance := getLocalCorrelationClass().Alloc()
	rv := objc.Send[LocalCorrelation](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131876-initwithdevice
func NewLocalCorrelationWithDeviceWindowInXWindowInYStrideInXStrideInY(device unsafe.Pointer, windowInX uint, windowInY uint, strideInX uint, strideInY uint) LocalCorrelation {
	instance := getLocalCorrelationClass().Alloc()
	rv := objc.Send[LocalCorrelation](instance.ID, objc.Sel("initWithDevice:windowInX:windowInY:strideInX:strideInY:"), device, windowInX, windowInY, strideInX, strideInY)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131877-strideinx
func (l_ LocalCorrelation) StrideInX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("strideInX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131877-strideinx
func (l_ LocalCorrelation) SetStrideInX(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setStrideInX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131878-strideiny
func (l_ LocalCorrelation) StrideInY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("strideInY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131878-strideiny
func (l_ LocalCorrelation) SetStrideInY(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setStrideInY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131879-windowinx
func (l_ LocalCorrelation) WindowInX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("windowInX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131879-windowinx
func (l_ LocalCorrelation) SetWindowInX(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWindowInX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131880-windowiny
func (l_ LocalCorrelation) WindowInY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("windowInY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131880-windowiny
func (l_ LocalCorrelation) SetWindowInY(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWindowInY:"), value)
}







