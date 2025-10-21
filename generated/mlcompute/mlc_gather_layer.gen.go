// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CGatherLayer] class.
var (
	CGatherLayerClass     _CGatherLayerClass
	CGatherLayerClassOnce sync.Once
)

func getCGatherLayerClass() _CGatherLayerClass {
	CGatherLayerClassOnce.Do(func() {
		CGatherLayerClass = _CGatherLayerClass{objc.GetClass("MLCGatherLayer")}
	})
	return CGatherLayerClass
}

type _CGatherLayerClass struct {
	class objc.Class
}

// An interface definition for the [CGatherLayer] class.
type ICGatherLayer interface {
	ICLayer
}

// A layer that fetches data at the locations you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGatherLayer
type CGatherLayer struct {
	CLayer
}

// CGatherLayerFrom constructs a [CGatherLayer] from an unsafe.Pointer.
//
// A layer that fetches data at the locations you specify.
func CGatherLayerFrom(ptr unsafe.Pointer) CGatherLayer {
	return CGatherLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CGatherLayerClass) Alloc() CGatherLayer {
	rv := objc.Send[CGatherLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CGatherLayerClass) New() CGatherLayer {
	rv := objc.Send[CGatherLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CGatherLayer) Init() CGatherLayer {
	rv := objc.Send[CGatherLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CGatherLayer) Autorelease() CGatherLayer {
	rv := objc.Send[CGatherLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCGatherLayer creates a new CGatherLayer instance.
func NewCGatherLayer() CGatherLayer {
	return getCGatherLayerClass().New()
}


// The dimension to index.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgatherlayer/dimension
func (c_ CGatherLayer) Dimension() int {
	rv := objc.Send[int](c_.ID, objc.Sel("dimension"))
	return rv
}


// SetDimension sets the value of the dimension property.
// The dimension to index.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgatherlayer/dimension
func (c_ CGatherLayer) SetDimension(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDimension:"), value)
}



