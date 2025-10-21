// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CUpsampleLayer] class.
var (
	CUpsampleLayerClass     _CUpsampleLayerClass
	CUpsampleLayerClassOnce sync.Once
)

func getCUpsampleLayerClass() _CUpsampleLayerClass {
	CUpsampleLayerClassOnce.Do(func() {
		CUpsampleLayerClass = _CUpsampleLayerClass{objc.GetClass("MLCUpsampleLayer")}
	})
	return CUpsampleLayerClass
}

type _CUpsampleLayerClass struct {
	class objc.Class
}

// An interface definition for the [CUpsampleLayer] class.
type ICUpsampleLayer interface {
	ICLayer
}

// A layer that applies upsampling with the shape you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCUpsampleLayer
type CUpsampleLayer struct {
	CLayer
}

// CUpsampleLayerFrom constructs a [CUpsampleLayer] from an unsafe.Pointer.
//
// A layer that applies upsampling with the shape you specify.
func CUpsampleLayerFrom(ptr unsafe.Pointer) CUpsampleLayer {
	return CUpsampleLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CUpsampleLayerClass) Alloc() CUpsampleLayer {
	rv := objc.Send[CUpsampleLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CUpsampleLayerClass) New() CUpsampleLayer {
	rv := objc.Send[CUpsampleLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CUpsampleLayer) Init() CUpsampleLayer {
	rv := objc.Send[CUpsampleLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CUpsampleLayer) Autorelease() CUpsampleLayer {
	rv := objc.Send[CUpsampleLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCUpsampleLayer creates a new CUpsampleLayer instance.
func NewCUpsampleLayer() CUpsampleLayer {
	return getCUpsampleLayerClass().New()
}


// A Boolean that indicates whether the layer aligns the corner pixels of the input and output tensors.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcupsamplelayer/alignscorners
func (c_ CUpsampleLayer) AlignsCorners() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("alignsCorners"))
	return rv
}


// SetAlignsCorners sets the value of the alignsCorners property.
// A Boolean that indicates whether the layer aligns the corner pixels of the input and output tensors.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcupsamplelayer/alignscorners
func (c_ CUpsampleLayer) SetAlignsCorners(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlignsCorners:"), value)
}

// An array that contains the dimensions of the result tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcupsamplelayer/shape-61n1u
func (c_ CUpsampleLayer) Shape() int {
	rv := objc.Send[int](c_.ID, objc.Sel("shape"))
	return rv
}


// SetShape sets the value of the shape property.
// An array that contains the dimensions of the result tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcupsamplelayer/shape-61n1u
func (c_ CUpsampleLayer) SetShape(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShape:"), value)
}

// The upsampling algorithm type.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcupsamplelayer/samplemode
func (c_ CUpsampleLayer) SampleMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sampleMode"))
	return rv
}


// SetSampleMode sets the value of the sampleMode property.
// The upsampling algorithm type.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcupsamplelayer/samplemode
func (c_ CUpsampleLayer) SetSampleMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleMode:"), value)
}



