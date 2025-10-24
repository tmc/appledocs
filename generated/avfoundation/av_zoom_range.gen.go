// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ZoomRange] class.
var (
	ZoomRangeClass     _ZoomRangeClass
	ZoomRangeClassOnce sync.Once
)

func getZoomRangeClass() _ZoomRangeClass {
	ZoomRangeClassOnce.Do(func() {
		ZoomRangeClass = _ZoomRangeClass{objc.GetClass("AVZoomRange")}
	})
	return ZoomRangeClass
}

type _ZoomRangeClass struct {
	class objc.Class
}





// An interface definition for the [ZoomRange] class.
type IZoomRange interface {
	objectivec.IObject
	

	// properties:
	MaxZoomFactor() float64
	MinZoomFactor() float64
	VideoMaxZoomFactor() float64
	SetVideoMaxZoomFactor(value float64)
	VideoZoomFactorUpscaleThreshold() float64
	SetVideoZoomFactorUpscaleThreshold(value float64)
	ZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported() bool
	SetZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported(value bool)


	

	// methods:
	ContainsZoomFactor(zoomFactor float64) bool


}





// Alloc allocates a new instance without initialization.
func (zc _ZoomRangeClass) Alloc() ZoomRange {
	rv := objc.Send[ZoomRange](objc.ID(zc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (zc _ZoomRangeClass) New() ZoomRange {
	rv := objc.Send[ZoomRange](objc.ID(zc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (z_ ZoomRange) Init() ZoomRange {
	rv := objc.Send[ZoomRange](z_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (z_ ZoomRange) Autorelease() ZoomRange {
	rv := objc.Send[ZoomRange](z_.ID, objc.Sel("autorelease"))
	return rv
}

// NewZoomRange creates a new ZoomRange instance.
func NewZoomRange() ZoomRange {
	return getZoomRangeClass().New()
}





// An object that defines an inclusive range of zoom values.


// An object that defines an inclusive range of zoom values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVZoomRange
type ZoomRange struct {
	objectivec.Object
}

// ZoomRangeFrom constructs a [ZoomRange] from an unsafe.Pointer.
//
// An object that defines an inclusive range of zoom values.
func ZoomRangeFrom(ptr unsafe.Pointer) ZoomRange {
	return ZoomRange{objectivec.Object{objc.ID(ptr)}}
}




















// Returns a Boolean value that indicates whether the specified zoom factor exists in the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVZoomRange/containsZoomFactor:
func (z_ ZoomRange) ContainsZoomFactor(zoomFactor float64) bool {
	rv := objc.Send[bool](z_.ID, objc.Sel("containsZoomFactor:"), zoomFactor)
	return rv
}







// The range’s maximum zoom factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVZoomRange/maxZoomFactor
func (z_ ZoomRange) MaxZoomFactor() float64 {
	rv := objc.Send[float64](z_.ID, objc.Sel("maxZoomFactor"))
	return rv
}


// The range’s minimum zoom factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVZoomRange/minZoomFactor
func (z_ ZoomRange) MinZoomFactor() float64 {
	rv := objc.Send[float64](z_.ID, objc.Sel("minZoomFactor"))
	return rv
}


// A maximum zoom factor the format allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videomaxzoomfactor
func (z_ ZoomRange) VideoMaxZoomFactor() float64 {
	rv := objc.Send[float64](z_.ID, objc.Sel("videoMaxZoomFactor"))
	return rv
}


// A maximum zoom factor the format allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videomaxzoomfactor
func (z_ ZoomRange) SetVideoMaxZoomFactor(value float64) {
	objc.Send[objc.ID](z_.ID, objc.Sel("setVideoMaxZoomFactor:"), value)
}


// A threshold at which the system upscales pixel data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videozoomfactorupscalethreshold
func (z_ ZoomRange) VideoZoomFactorUpscaleThreshold() float64 {
	rv := objc.Send[float64](z_.ID, objc.Sel("videoZoomFactorUpscaleThreshold"))
	return rv
}


// A threshold at which the system upscales pixel data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videozoomfactorupscalethreshold
func (z_ ZoomRange) SetVideoZoomFactorUpscaleThreshold(value float64) {
	objc.Send[objc.ID](z_.ID, objc.Sel("setVideoZoomFactorUpscaleThreshold:"), value)
}


// A Boolean value that indicates whether the format supports zoom factors outside the range supported for depth delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/zoomfactorsoutsideofvideozoomrangesfordepthdeliverysupported
func (z_ ZoomRange) ZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported() bool {
	rv := objc.Send[bool](z_.ID, objc.Sel("zoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported"))
	return rv
}


// A Boolean value that indicates whether the format supports zoom factors outside the range supported for depth delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/zoomfactorsoutsideofvideozoomrangesfordepthdeliverysupported
func (z_ ZoomRange) SetZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported(value bool) {
	objc.Send[objc.ID](z_.ID, objc.Sel("setZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported:"), value)
}








