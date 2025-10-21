// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CYOLOLossLayer] class.
var (
	CYOLOLossLayerClass     _CYOLOLossLayerClass
	CYOLOLossLayerClassOnce sync.Once
)

func getCYOLOLossLayerClass() _CYOLOLossLayerClass {
	CYOLOLossLayerClassOnce.Do(func() {
		CYOLOLossLayerClass = _CYOLOLossLayerClass{objc.GetClass("MLCYOLOLossLayer")}
	})
	return CYOLOLossLayerClass
}

type _CYOLOLossLayerClass struct {
	class objc.Class
}

// An interface definition for the [CYOLOLossLayer] class.
type ICYOLOLossLayer interface {
	ICLossLayer
}

// A layer that estimates loss for the YOLO algorithm.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCYOLOLossLayer
type CYOLOLossLayer struct {
	CLossLayer
}

// CYOLOLossLayerFrom constructs a [CYOLOLossLayer] from an unsafe.Pointer.
//
// A layer that estimates loss for the YOLO algorithm.
func CYOLOLossLayerFrom(ptr unsafe.Pointer) CYOLOLossLayer {
	return CYOLOLossLayer{
		CLossLayer: CLossLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CYOLOLossLayerClass) Alloc() CYOLOLossLayer {
	rv := objc.Send[CYOLOLossLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CYOLOLossLayerClass) New() CYOLOLossLayer {
	rv := objc.Send[CYOLOLossLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CYOLOLossLayer) Init() CYOLOLossLayer {
	rv := objc.Send[CYOLOLossLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CYOLOLossLayer) Autorelease() CYOLOLossLayer {
	rv := objc.Send[CYOLOLossLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCYOLOLossLayer creates a new CYOLOLossLayer instance.
func NewCYOLOLossLayer() CYOLOLossLayer {
	return getCYOLOLossLayerClass().New()
}


// The configuration object you use to create the YOLO loss layer.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcyololosslayer/yololossdescriptor
func (c_ CYOLOLossLayer) YoloLossDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("yoloLossDescriptor"))
	return rv
}


// SetYoloLossDescriptor sets the value of the yoloLossDescriptor property.
// The configuration object you use to create the YOLO loss layer.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcyololosslayer/yololossdescriptor
func (c_ CYOLOLossLayer) SetYoloLossDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setYoloLossDescriptor:"), value)
}




