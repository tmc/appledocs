// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SCStreamConfiguration */


/* debug [class_header]: Header for SCStreamConfiguration */
// The class instance for the [StreamConfiguration] class.
var (
	StreamConfigurationClass     _StreamConfigurationClass
	StreamConfigurationClassOnce sync.Once
)

func getStreamConfigurationClass() _StreamConfigurationClass {
	StreamConfigurationClassOnce.Do(func() {
		StreamConfigurationClass = _StreamConfigurationClass{objc.GetClass("SCStreamConfiguration")}
	})
	return StreamConfigurationClass
}

type _StreamConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StreamConfiguration */
// An interface definition for the [StreamConfiguration] class.
type IStreamConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for StreamConfiguration */
	// properties:
	BackgroundColor() ColorRef /* not a class type */
	SetBackgroundColor(value ColorRef /* not a class type */)
	CaptureDynamicRange() CaptureDynamicRange
	SetCaptureDynamicRange(value CaptureDynamicRange)
	CaptureMicrophone() bool
	SetCaptureMicrophone(value bool)
	CaptureResolution() CaptureResolutionType
	SetCaptureResolution(value CaptureResolutionType)
	CapturesAudio() bool
	SetCapturesAudio(value bool)
	CapturesShadowsOnly() bool
	SetCapturesShadowsOnly(value bool)
	ChannelCount() int
	SetChannelCount(value int)
	ColorMatrix() StringRef /* not a class type */
	SetColorMatrix(value StringRef /* not a class type */)
	ColorSpaceName() StringRef /* not a class type */
	SetColorSpaceName(value StringRef /* not a class type */)
	DestinationRect() corefoundation.CGRect
	SetDestinationRect(value corefoundation.CGRect)
	ExcludesCurrentProcessAudio() bool
	SetExcludesCurrentProcessAudio(value bool)
	Height() uintptr /* not a class type */
	SetHeight(value uintptr /* not a class type */)
	IgnoreGlobalClipDisplay() bool
	SetIgnoreGlobalClipDisplay(value bool)
	IgnoreGlobalClipSingleWindow() bool
	SetIgnoreGlobalClipSingleWindow(value bool)
	IgnoreShadowsDisplay() bool
	SetIgnoreShadowsDisplay(value bool)
	IgnoreShadowsSingleWindow() bool
	SetIgnoreShadowsSingleWindow(value bool)
	IncludeChildWindows() bool
	SetIncludeChildWindows(value bool)
	MicrophoneCaptureDeviceID() objc.IObject /* cross-framework: NSString */
	SetMicrophoneCaptureDeviceID(value objc.IObject /* cross-framework: NSString */)
	MinimumFrameInterval() objc.IObject /* cross-framework: Time */
	SetMinimumFrameInterval(value objc.IObject /* cross-framework: Time */)
	PixelFormat() uint32 /* not a class type */
	SetPixelFormat(value uint32 /* not a class type */)
	PresenterOverlayPrivacyAlertSetting() PresenterOverlayAlertSetting
	SetPresenterOverlayPrivacyAlertSetting(value PresenterOverlayAlertSetting)
	PreservesAspectRatio() bool
	SetPreservesAspectRatio(value bool)
	QueueDepth() int
	SetQueueDepth(value int)
	SampleRate() int
	SetSampleRate(value int)
	ScalesToFit() bool
	SetScalesToFit(value bool)
	ShouldBeOpaque() bool
	SetShouldBeOpaque(value bool)
	ShowMouseClicks() bool
	SetShowMouseClicks(value bool)
	ShowsCursor() bool
	SetShowsCursor(value bool)
	SourceRect() corefoundation.CGRect
	SetSourceRect(value corefoundation.CGRect)
	StreamName() objc.IObject /* cross-framework: NSString */
	SetStreamName(value objc.IObject /* cross-framework: NSString */)
	Width() uintptr /* not a class type */
	SetWidth(value uintptr /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StreamConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StreamConfiguration */
// Alloc allocates a new instance without initialization.
func (sc _StreamConfigurationClass) Alloc() StreamConfiguration {
	rv := objc.Send[StreamConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StreamConfigurationClass) New() StreamConfiguration {
	rv := objc.Send[StreamConfiguration](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StreamConfiguration) Init() StreamConfiguration {
	rv := objc.Send[StreamConfiguration](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StreamConfiguration) Autorelease() StreamConfiguration {
	rv := objc.Send[StreamConfiguration](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStreamConfiguration creates a new StreamConfiguration instance.
func NewStreamConfiguration() StreamConfiguration {
	return getStreamConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StreamConfiguration */
// An instance that provides the output configuration for a stream.
//
// Creating an instance of this class provides a default configuration for a stream. Only configure its properties if you need to customize the output.


// An instance that provides the output configuration for a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration
type StreamConfiguration struct {
	objectivec.Object
}

// StreamConfigurationFrom constructs a [StreamConfiguration] from an unsafe.Pointer.
//
// An instance that provides the output configuration for a stream.
func StreamConfigurationFrom(ptr unsafe.Pointer) StreamConfiguration {
	return StreamConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StreamConfiguration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/init(preset:)
func NewStreamConfigurationWithPreset(preset StreamConfigurationPreset) StreamConfiguration {
	rv := objc.Send[StreamConfiguration](objc.ID(getStreamConfigurationClass().class), objc.Sel("streamConfigurationWithPreset:"), preset)
	return rv
}/* debug [class_init_methods/constructor]: NewStreamConfigurationWithPreset */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StreamConfiguration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/init(preset:)
func (sc _StreamConfigurationClass) StreamConfigurationWithPreset(preset StreamConfigurationPreset) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("streamConfigurationWithPreset:"), preset)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StreamConfigurationWithPreset) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StreamConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StreamConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StreamConfiguration */

// A background color for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/backgroundColor
func (s_ StreamConfiguration) BackgroundColor() ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](s_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// A background color for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/backgroundColor
func (s_ StreamConfiguration) SetBackgroundColor(value ColorRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureDynamicRange
func (s_ StreamConfiguration) CaptureDynamicRange() CaptureDynamicRange {
	rv := objc.Send[CaptureDynamicRange](s_.ID, objc.Sel("captureDynamicRange"))
	return rv
}/* debug [instance_properties/getter]: captureDynamicRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureDynamicRange
func (s_ StreamConfiguration) SetCaptureDynamicRange(value CaptureDynamicRange) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCaptureDynamicRange:"), value)
}/* debug [instance_properties/setter]: captureDynamicRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureMicrophone
func (s_ StreamConfiguration) CaptureMicrophone() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("captureMicrophone"))
	return rv
}/* debug [instance_properties/getter]: captureMicrophone */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureMicrophone
func (s_ StreamConfiguration) SetCaptureMicrophone(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCaptureMicrophone:"), value)
}/* debug [instance_properties/setter]: captureMicrophone */


// The resolution at which to capture source content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureResolution
func (s_ StreamConfiguration) CaptureResolution() CaptureResolutionType {
	rv := objc.Send[CaptureResolutionType](s_.ID, objc.Sel("captureResolution"))
	return rv
}/* debug [instance_properties/getter]: captureResolution */


// The resolution at which to capture source content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureResolution
func (s_ StreamConfiguration) SetCaptureResolution(value CaptureResolutionType) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCaptureResolution:"), value)
}/* debug [instance_properties/setter]: captureResolution */


// A Boolean value that indicates whether to capture audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/capturesAudio
func (s_ StreamConfiguration) CapturesAudio() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("capturesAudio"))
	return rv
}/* debug [instance_properties/getter]: capturesAudio */


// A Boolean value that indicates whether to capture audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/capturesAudio
func (s_ StreamConfiguration) SetCapturesAudio(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCapturesAudio:"), value)
}/* debug [instance_properties/setter]: capturesAudio */


// A Boolean value that indicates if the stream only captures shadows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/capturesShadowsOnly
func (s_ StreamConfiguration) CapturesShadowsOnly() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("capturesShadowsOnly"))
	return rv
}/* debug [instance_properties/getter]: capturesShadowsOnly */


// A Boolean value that indicates if the stream only captures shadows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/capturesShadowsOnly
func (s_ StreamConfiguration) SetCapturesShadowsOnly(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCapturesShadowsOnly:"), value)
}/* debug [instance_properties/setter]: capturesShadowsOnly */


// The number of audio channels to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/channelCount
func (s_ StreamConfiguration) ChannelCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("channelCount"))
	return rv
}/* debug [instance_properties/getter]: channelCount */


// The number of audio channels to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/channelCount
func (s_ StreamConfiguration) SetChannelCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setChannelCount:"), value)
}/* debug [instance_properties/setter]: channelCount */


// A color matrix to apply to the output surface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/colorMatrix
func (s_ StreamConfiguration) ColorMatrix() StringRef /* not a class type */ {
	rv := objc.Send[StringRef](s_.ID, objc.Sel("colorMatrix"))
	return rv
}/* debug [instance_properties/getter]: colorMatrix */


// A color matrix to apply to the output surface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/colorMatrix
func (s_ StreamConfiguration) SetColorMatrix(value StringRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setColorMatrix:"), value)
}/* debug [instance_properties/setter]: colorMatrix */


// A color space to use for the output buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/colorSpaceName
func (s_ StreamConfiguration) ColorSpaceName() StringRef /* not a class type */ {
	rv := objc.Send[StringRef](s_.ID, objc.Sel("colorSpaceName"))
	return rv
}/* debug [instance_properties/getter]: colorSpaceName */


// A color space to use for the output buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/colorSpaceName
func (s_ StreamConfiguration) SetColorSpaceName(value StringRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setColorSpaceName:"), value)
}/* debug [instance_properties/setter]: colorSpaceName */


// A rectangle that specifies a destination into which to write the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/destinationRect
func (s_ StreamConfiguration) DestinationRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s_.ID, objc.Sel("destinationRect"))
	return rv
}/* debug [instance_properties/getter]: destinationRect */


// A rectangle that specifies a destination into which to write the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/destinationRect
func (s_ StreamConfiguration) SetDestinationRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDestinationRect:"), value)
}/* debug [instance_properties/setter]: destinationRect */


// A Boolean value that indicates whether to exclude audio from your app during capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/excludesCurrentProcessAudio
func (s_ StreamConfiguration) ExcludesCurrentProcessAudio() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("excludesCurrentProcessAudio"))
	return rv
}/* debug [instance_properties/getter]: excludesCurrentProcessAudio */


// A Boolean value that indicates whether to exclude audio from your app during capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/excludesCurrentProcessAudio
func (s_ StreamConfiguration) SetExcludesCurrentProcessAudio(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setExcludesCurrentProcessAudio:"), value)
}/* debug [instance_properties/setter]: excludesCurrentProcessAudio */


// The height of the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/height
func (s_ StreamConfiguration) Height() uintptr /* not a class type */ {
	rv := objc.Send[uintptr](s_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// The height of the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/height
func (s_ StreamConfiguration) SetHeight(value uintptr /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHeight:"), value)
}/* debug [instance_properties/setter]: height */


// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in display style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreGlobalClipDisplay
func (s_ StreamConfiguration) IgnoreGlobalClipDisplay() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreGlobalClipDisplay"))
	return rv
}/* debug [instance_properties/getter]: ignoreGlobalClipDisplay */


// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in display style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreGlobalClipDisplay
func (s_ StreamConfiguration) SetIgnoreGlobalClipDisplay(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreGlobalClipDisplay:"), value)
}/* debug [instance_properties/setter]: ignoreGlobalClipDisplay */


// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in window style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreGlobalClipSingleWindow
func (s_ StreamConfiguration) IgnoreGlobalClipSingleWindow() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreGlobalClipSingleWindow"))
	return rv
}/* debug [instance_properties/getter]: ignoreGlobalClipSingleWindow */


// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in window style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreGlobalClipSingleWindow
func (s_ StreamConfiguration) SetIgnoreGlobalClipSingleWindow(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreGlobalClipSingleWindow:"), value)
}/* debug [instance_properties/setter]: ignoreGlobalClipSingleWindow */


// A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in display style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreShadowsDisplay
func (s_ StreamConfiguration) IgnoreShadowsDisplay() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreShadowsDisplay"))
	return rv
}/* debug [instance_properties/getter]: ignoreShadowsDisplay */


// A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in display style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreShadowsDisplay
func (s_ StreamConfiguration) SetIgnoreShadowsDisplay(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreShadowsDisplay:"), value)
}/* debug [instance_properties/setter]: ignoreShadowsDisplay */


// A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in window style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreShadowsSingleWindow
func (s_ StreamConfiguration) IgnoreShadowsSingleWindow() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreShadowsSingleWindow"))
	return rv
}/* debug [instance_properties/getter]: ignoreShadowsSingleWindow */


// A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in window style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreShadowsSingleWindow
func (s_ StreamConfiguration) SetIgnoreShadowsSingleWindow(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreShadowsSingleWindow:"), value)
}/* debug [instance_properties/setter]: ignoreShadowsSingleWindow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/includeChildWindows
func (s_ StreamConfiguration) IncludeChildWindows() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("includeChildWindows"))
	return rv
}/* debug [instance_properties/getter]: includeChildWindows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/includeChildWindows
func (s_ StreamConfiguration) SetIncludeChildWindows(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIncludeChildWindows:"), value)
}/* debug [instance_properties/setter]: includeChildWindows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/microphoneCaptureDeviceID
func (s_ StreamConfiguration) MicrophoneCaptureDeviceID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("microphoneCaptureDeviceID"))
	return rv
}/* debug [instance_properties/getter]: microphoneCaptureDeviceID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/microphoneCaptureDeviceID
func (s_ StreamConfiguration) SetMicrophoneCaptureDeviceID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMicrophoneCaptureDeviceID:"), value)
}/* debug [instance_properties/setter]: microphoneCaptureDeviceID */


// The desired minimum time between frame updates, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/minimumFrameInterval
func (s_ StreamConfiguration) MinimumFrameInterval() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](s_.ID, objc.Sel("minimumFrameInterval"))
	return rv
}/* debug [instance_properties/getter]: minimumFrameInterval */


// The desired minimum time between frame updates, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/minimumFrameInterval
func (s_ StreamConfiguration) SetMinimumFrameInterval(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumFrameInterval:"), value)
}/* debug [instance_properties/setter]: minimumFrameInterval */


// A pixel format for sample buffers that a stream outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/pixelFormat
func (s_ StreamConfiguration) PixelFormat() uint32 /* not a class type */ {
	rv := objc.Send[uint32](s_.ID, objc.Sel("pixelFormat"))
	return rv
}/* debug [instance_properties/getter]: pixelFormat */


// A pixel format for sample buffers that a stream outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/pixelFormat
func (s_ StreamConfiguration) SetPixelFormat(value uint32 /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPixelFormat:"), value)
}/* debug [instance_properties/setter]: pixelFormat */


// A value indicating if alerts appear to presenters while using Presenter Overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/presenterOverlayPrivacyAlertSetting
func (s_ StreamConfiguration) PresenterOverlayPrivacyAlertSetting() PresenterOverlayAlertSetting {
	rv := objc.Send[PresenterOverlayAlertSetting](s_.ID, objc.Sel("presenterOverlayPrivacyAlertSetting"))
	return rv
}/* debug [instance_properties/getter]: presenterOverlayPrivacyAlertSetting */


// A value indicating if alerts appear to presenters while using Presenter Overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/presenterOverlayPrivacyAlertSetting
func (s_ StreamConfiguration) SetPresenterOverlayPrivacyAlertSetting(value PresenterOverlayAlertSetting) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPresenterOverlayPrivacyAlertSetting:"), value)
}/* debug [instance_properties/setter]: presenterOverlayPrivacyAlertSetting */


// A Boolean value that determines if the stream preserves aspect ratio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/preservesAspectRatio
func (s_ StreamConfiguration) PreservesAspectRatio() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("preservesAspectRatio"))
	return rv
}/* debug [instance_properties/getter]: preservesAspectRatio */


// A Boolean value that determines if the stream preserves aspect ratio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/preservesAspectRatio
func (s_ StreamConfiguration) SetPreservesAspectRatio(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreservesAspectRatio:"), value)
}/* debug [instance_properties/setter]: preservesAspectRatio */


// The maximum number of frames for the queue to store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/queueDepth
func (s_ StreamConfiguration) QueueDepth() int {
	rv := objc.Send[int](s_.ID, objc.Sel("queueDepth"))
	return rv
}/* debug [instance_properties/getter]: queueDepth */


// The maximum number of frames for the queue to store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/queueDepth
func (s_ StreamConfiguration) SetQueueDepth(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setQueueDepth:"), value)
}/* debug [instance_properties/setter]: queueDepth */


// The sample rate for audio capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/sampleRate
func (s_ StreamConfiguration) SampleRate() int {
	rv := objc.Send[int](s_.ID, objc.Sel("sampleRate"))
	return rv
}/* debug [instance_properties/getter]: sampleRate */


// The sample rate for audio capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/sampleRate
func (s_ StreamConfiguration) SetSampleRate(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSampleRate:"), value)
}/* debug [instance_properties/setter]: sampleRate */


// A Boolean value that indicates whether to scale the output to fit the configured width and height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/scalesToFit
func (s_ StreamConfiguration) ScalesToFit() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("scalesToFit"))
	return rv
}/* debug [instance_properties/getter]: scalesToFit */


// A Boolean value that indicates whether to scale the output to fit the configured width and height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/scalesToFit
func (s_ StreamConfiguration) SetScalesToFit(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScalesToFit:"), value)
}/* debug [instance_properties/setter]: scalesToFit */


// A Boolean value that indicates if semitransparent content presents as opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/shouldBeOpaque
func (s_ StreamConfiguration) ShouldBeOpaque() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldBeOpaque"))
	return rv
}/* debug [instance_properties/getter]: shouldBeOpaque */


// A Boolean value that indicates if semitransparent content presents as opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/shouldBeOpaque
func (s_ StreamConfiguration) SetShouldBeOpaque(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShouldBeOpaque:"), value)
}/* debug [instance_properties/setter]: shouldBeOpaque */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/showMouseClicks
func (s_ StreamConfiguration) ShowMouseClicks() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showMouseClicks"))
	return rv
}/* debug [instance_properties/getter]: showMouseClicks */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/showMouseClicks
func (s_ StreamConfiguration) SetShowMouseClicks(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowMouseClicks:"), value)
}/* debug [instance_properties/setter]: showMouseClicks */


// A Boolean value that determines whether the cursor is visible in the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/showsCursor
func (s_ StreamConfiguration) ShowsCursor() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsCursor"))
	return rv
}/* debug [instance_properties/getter]: showsCursor */


// A Boolean value that determines whether the cursor is visible in the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/showsCursor
func (s_ StreamConfiguration) SetShowsCursor(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsCursor:"), value)
}/* debug [instance_properties/setter]: showsCursor */


// A rectangle that specifies the source area to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/sourceRect
func (s_ StreamConfiguration) SourceRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s_.ID, objc.Sel("sourceRect"))
	return rv
}/* debug [instance_properties/getter]: sourceRect */


// A rectangle that specifies the source area to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/sourceRect
func (s_ StreamConfiguration) SetSourceRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSourceRect:"), value)
}/* debug [instance_properties/setter]: sourceRect */


// A name that you provide for identifying the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/streamName
func (s_ StreamConfiguration) StreamName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("streamName"))
	return rv
}/* debug [instance_properties/getter]: streamName */


// A name that you provide for identifying the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/streamName
func (s_ StreamConfiguration) SetStreamName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStreamName:"), value)
}/* debug [instance_properties/setter]: streamName */


// The width of the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/width
func (s_ StreamConfiguration) Width() uintptr /* not a class type */ {
	rv := objc.Send[uintptr](s_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// The width of the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/width
func (s_ StreamConfiguration) SetWidth(value uintptr /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SCStreamConfiguration */


