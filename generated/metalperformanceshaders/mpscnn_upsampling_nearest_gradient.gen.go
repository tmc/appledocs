// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNUpsamplingNearestGradient] class.
var (
	CNNUpsamplingNearestGradientClass     _CNNUpsamplingNearestGradientClass
	CNNUpsamplingNearestGradientClassOnce sync.Once
)

func getCNNUpsamplingNearestGradientClass() _CNNUpsamplingNearestGradientClass {
	CNNUpsamplingNearestGradientClassOnce.Do(func() {
		CNNUpsamplingNearestGradientClass = _CNNUpsamplingNearestGradientClass{objc.GetClass("MPSCNNUpsamplingNearestGradient")}
	})
	return CNNUpsamplingNearestGradientClass
}

type _CNNUpsamplingNearestGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNUpsamplingNearestGradient] class.
type ICNNUpsamplingNearestGradient interface {
	ICNNUpsamplingGradient
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNUpsamplingNearestGradientClass) Alloc() CNNUpsamplingNearestGradient {
	rv := objc.Send[CNNUpsamplingNearestGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNUpsamplingNearestGradientClass) New() CNNUpsamplingNearestGradient {
	rv := objc.Send[CNNUpsamplingNearestGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNUpsamplingNearestGradient) Init() CNNUpsamplingNearestGradient {
	rv := objc.Send[CNNUpsamplingNearestGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNUpsamplingNearestGradient) Autorelease() CNNUpsamplingNearestGradient {
	rv := objc.Send[CNNUpsamplingNearestGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNUpsamplingNearestGradient creates a new CNNUpsamplingNearestGradient instance.
func NewCNNUpsamplingNearestGradient() CNNUpsamplingNearestGradient {
	return getCNNUpsamplingNearestGradientClass().New()
}





// A gradient upsampling filter that samples the pixel nearest to the source when upsampling to the destination pixel.


// A gradient upsampling filter that samples the pixel nearest to the source when upsampling to the destination pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestGradient
type CNNUpsamplingNearestGradient struct {
	CNNUpsamplingGradient
}

// CNNUpsamplingNearestGradientFrom constructs a [CNNUpsamplingNearestGradient] from an unsafe.Pointer.
//
// A gradient upsampling filter that samples the pixel nearest to the source when upsampling to the destination pixel.
func CNNUpsamplingNearestGradientFrom(ptr unsafe.Pointer) CNNUpsamplingNearestGradient {
	return CNNUpsamplingNearestGradient{
		CNNUpsamplingGradient: CNNUpsamplingGradientFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradient/2947920-initwithdevice
func NewCNNUpsamplingNearestGradientWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device unsafe.Pointer, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingNearestGradient {
	instance := getCNNUpsamplingNearestGradientClass().Alloc()
	rv := objc.Send[CNNUpsamplingNearestGradient](instance.ID, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), device, integerScaleFactorX, integerScaleFactorY)
	rv.Autorelease()
	return rv
}



























