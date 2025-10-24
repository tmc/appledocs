// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNUpsamplingBilinearGradient] class.
var (
	CNNUpsamplingBilinearGradientClass     _CNNUpsamplingBilinearGradientClass
	CNNUpsamplingBilinearGradientClassOnce sync.Once
)

func getCNNUpsamplingBilinearGradientClass() _CNNUpsamplingBilinearGradientClass {
	CNNUpsamplingBilinearGradientClassOnce.Do(func() {
		CNNUpsamplingBilinearGradientClass = _CNNUpsamplingBilinearGradientClass{objc.GetClass("MPSCNNUpsamplingBilinearGradient")}
	})
	return CNNUpsamplingBilinearGradientClass
}

type _CNNUpsamplingBilinearGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNUpsamplingBilinearGradient] class.
type ICNNUpsamplingBilinearGradient interface {
	ICNNUpsamplingGradient
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNUpsamplingBilinearGradientClass) Alloc() CNNUpsamplingBilinearGradient {
	rv := objc.Send[CNNUpsamplingBilinearGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNUpsamplingBilinearGradientClass) New() CNNUpsamplingBilinearGradient {
	rv := objc.Send[CNNUpsamplingBilinearGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNUpsamplingBilinearGradient) Init() CNNUpsamplingBilinearGradient {
	rv := objc.Send[CNNUpsamplingBilinearGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNUpsamplingBilinearGradient) Autorelease() CNNUpsamplingBilinearGradient {
	rv := objc.Send[CNNUpsamplingBilinearGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNUpsamplingBilinearGradient creates a new CNNUpsamplingBilinearGradient instance.
func NewCNNUpsamplingBilinearGradient() CNNUpsamplingBilinearGradient {
	return getCNNUpsamplingBilinearGradientClass().New()
}





// A gradient bilinear spatial upsampling filter.


// A gradient bilinear spatial upsampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearGradient
type CNNUpsamplingBilinearGradient struct {
	CNNUpsamplingGradient
}

// CNNUpsamplingBilinearGradientFrom constructs a [CNNUpsamplingBilinearGradient] from an unsafe.Pointer.
//
// A gradient bilinear spatial upsampling filter.
func CNNUpsamplingBilinearGradientFrom(ptr unsafe.Pointer) CNNUpsamplingBilinearGradient {
	return CNNUpsamplingBilinearGradient{
		CNNUpsamplingGradient: CNNUpsamplingGradientFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradient/2947918-initwithdevice
func NewCNNUpsamplingBilinearGradientWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device unsafe.Pointer, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingBilinearGradient {
	instance := getCNNUpsamplingBilinearGradientClass().Alloc()
	rv := objc.Send[CNNUpsamplingBilinearGradient](instance.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), device, integerScaleFactorX, integerScaleFactorY)
	rv.Autorelease()
	return rv
}



























