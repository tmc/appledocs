// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVFrameRateRange */


/* debug [class_header]: Header for AVFrameRateRange */
// The class instance for the [FrameRateRange] class.
var (
	FrameRateRangeClass     _FrameRateRangeClass
	FrameRateRangeClassOnce sync.Once
)

func getFrameRateRangeClass() _FrameRateRangeClass {
	FrameRateRangeClassOnce.Do(func() {
		FrameRateRangeClass = _FrameRateRangeClass{objc.GetClass("AVFrameRateRange")}
	})
	return FrameRateRangeClass
}

type _FrameRateRangeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FrameRateRange */
// An interface definition for the [FrameRateRange] class.
type IFrameRateRange interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FrameRateRange */
	// properties:
	MaxFrameDuration() objc.IObject /* cross-framework: Time */
	MaxFrameRate() objectivec.IObject
	MinFrameDuration() objc.IObject /* cross-framework: Time */
	MinFrameRate() objectivec.IObject
	IsAutoVideoFrameRateSupported() bool
	SetIsAutoVideoFrameRateSupported(value bool)
	IsMultiCamSupported() bool
	SetIsMultiCamSupported(value bool)
	IsVideoBinned() bool
	SetIsVideoBinned(value bool)
	IsVideoHDRSupported() bool
	SetIsVideoHDRSupported(value bool)
	VideoSupportedFrameRateRanges() IAVFrameRateRange
	SetVideoSupportedFrameRateRanges(value IAVFrameRateRange)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FrameRateRange */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FrameRateRange */
// Alloc allocates a new instance without initialization.
func (fc _FrameRateRangeClass) Alloc() FrameRateRange {
	rv := objc.Send[FrameRateRange](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FrameRateRangeClass) New() FrameRateRange {
	rv := objc.Send[FrameRateRange](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FrameRateRange) Init() FrameRateRange {
	rv := objc.Send[FrameRateRange](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FrameRateRange) Autorelease() FrameRateRange {
	rv := objc.Send[FrameRateRange](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFrameRateRange creates a new FrameRateRange instance.
func NewFrameRateRange() FrameRateRange {
	return getFrameRateRangeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FrameRateRange */
// An immutable type that represents a range of valid frame rates.
//
// An AVFrameRateRange object is immutable. An object wraps a CMFormatDescription and expresses a range of valid video frame rates as an array of objects. An object uses to describe the formats it supports and the currently-active format.


// An immutable type that represents a range of valid frame rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFrameRateRange
type FrameRateRange struct {
	objectivec.Object
}

// FrameRateRangeFrom constructs a [FrameRateRange] from an unsafe.Pointer.
//
// An immutable type that represents a range of valid frame rates.
func FrameRateRangeFrom(ptr unsafe.Pointer) FrameRateRange {
	return FrameRateRange{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FrameRateRange *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FrameRateRange */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FrameRateRange */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FrameRateRange */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FrameRateRange */

// The maximum frame duration supported by the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFrameRateRange/maxFrameDuration
func (f_ FrameRateRange) MaxFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](f_.ID, objc.Sel("maxFrameDuration"))
	return rv
}/* debug [instance_properties/getter]: maxFrameDuration */


// The maximum frame rate supported by the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFrameRateRange/maxFrameRate
func (f_ FrameRateRange) MaxFrameRate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("maxFrameRate"))
	return rv
}/* debug [instance_properties/getter]: maxFrameRate */


// The minimum frame duration supported by the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFrameRateRange/minFrameDuration
func (f_ FrameRateRange) MinFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](f_.ID, objc.Sel("minFrameDuration"))
	return rv
}/* debug [instance_properties/getter]: minFrameDuration */


// The minimum frame rate supported by the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFrameRateRange/minFrameRate
func (f_ FrameRateRange) MinFrameRate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("minFrameRate"))
	return rv
}/* debug [instance_properties/getter]: minFrameRate */


// A Boolean value that Indicates whether the format supports performing automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isautovideoframeratesupported
func (f_ FrameRateRange) IsAutoVideoFrameRateSupported() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isAutoVideoFrameRateSupported"))
	return rv
}/* debug [instance_properties/getter]: isAutoVideoFrameRateSupported */


// A Boolean value that Indicates whether the format supports performing automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isautovideoframeratesupported
func (f_ FrameRateRange) SetIsAutoVideoFrameRateSupported(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsAutoVideoFrameRateSupported:"), value)
}/* debug [instance_properties/setter]: isAutoVideoFrameRateSupported */


// A Boolean value that indicates whether a multi-camera capture session supports this format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/ismulticamsupported
func (f_ FrameRateRange) IsMultiCamSupported() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isMultiCamSupported"))
	return rv
}/* debug [instance_properties/getter]: isMultiCamSupported */


// A Boolean value that indicates whether a multi-camera capture session supports this format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/ismulticamsupported
func (f_ FrameRateRange) SetIsMultiCamSupported(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsMultiCamSupported:"), value)
}/* debug [instance_properties/setter]: isMultiCamSupported */


// A Boolean value that indicates whether the format produces video data in a binned format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isvideobinned
func (f_ FrameRateRange) IsVideoBinned() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isVideoBinned"))
	return rv
}/* debug [instance_properties/getter]: isVideoBinned */


// A Boolean value that indicates whether the format produces video data in a binned format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isvideobinned
func (f_ FrameRateRange) SetIsVideoBinned(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsVideoBinned:"), value)
}/* debug [instance_properties/setter]: isVideoBinned */


// A Boolean value that indicates whether the format supports high dynamic range streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isvideohdrsupported
func (f_ FrameRateRange) IsVideoHDRSupported() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isVideoHDRSupported"))
	return rv
}/* debug [instance_properties/getter]: isVideoHDRSupported */


// A Boolean value that indicates whether the format supports high dynamic range streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isvideohdrsupported
func (f_ FrameRateRange) SetIsVideoHDRSupported(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsVideoHDRSupported:"), value)
}/* debug [instance_properties/setter]: isVideoHDRSupported */


// A list of frame rate ranges that a format supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videosupportedframerateranges
func (f_ FrameRateRange) VideoSupportedFrameRateRanges() IAVFrameRateRange {
	rv := objc.Send[FrameRateRange](f_.ID, objc.Sel("videoSupportedFrameRateRanges"))
	return rv
}/* debug [instance_properties/getter]: videoSupportedFrameRateRanges */


// A list of frame rate ranges that a format supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videosupportedframerateranges
func (f_ FrameRateRange) SetVideoSupportedFrameRateRanges(value IAVFrameRateRange) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setVideoSupportedFrameRateRanges:"), value)
}/* debug [instance_properties/setter]: videoSupportedFrameRateRanges */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVFrameRateRange */



