// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/coremedia"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [StreamConfiguration] class.
type IStreamConfiguration interface {
	objectivec.IObject
	// properties:
	CaptureDynamicRange() CaptureDynamicRange /* not a class type */
	SetCaptureDynamicRange(value CaptureDynamicRange /* not a class type */)
	QueueDepth() int
	SetQueueDepth(value int)
	BackgroundColor() objc.IObject /* cross-framework: Color */
	SetBackgroundColor(value objc.IObject /* cross-framework: Color */)
	CaptureMicrophone() bool
	SetCaptureMicrophone(value bool)
	CaptureResolution() CaptureResolutionType /* not a class type */
	SetCaptureResolution(value CaptureResolutionType /* not a class type */)
	CapturesAudio() bool
	SetCapturesAudio(value bool)
	CapturesShadowsOnly() bool
	SetCapturesShadowsOnly(value bool)
	ChannelCount() int
	SetChannelCount(value int)
	ColorMatrix() objc.IObject /* cross-framework: String */
	SetColorMatrix(value objc.IObject /* cross-framework: String */)
	ColorSpaceName() objc.IObject /* cross-framework: String */
	SetColorSpaceName(value objc.IObject /* cross-framework: String */)
	DestinationRect() objc.IObject /* cross-framework: Rect */
	SetDestinationRect(value objc.IObject /* cross-framework: Rect */)
	ExcludesCurrentProcessAudio() bool
	SetExcludesCurrentProcessAudio(value bool)
	Height() int
	SetHeight(value int)
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
	PresenterOverlayPrivacyAlertSetting() PresenterOverlayAlertSetting /* not a class type */
	SetPresenterOverlayPrivacyAlertSetting(value PresenterOverlayAlertSetting /* not a class type */)
	PreservesAspectRatio() bool
	SetPreservesAspectRatio(value bool)
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
	SourceRect() objc.IObject /* cross-framework: Rect */
	SetSourceRect(value objc.IObject /* cross-framework: Rect */)
	StreamName() objc.IObject /* cross-framework: NSString */
	SetStreamName(value objc.IObject /* cross-framework: NSString */)
	Width() int
	SetWidth(value int)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (sc _StreamConfigurationClass) Alloc() StreamConfiguration {
	rv := objc.Send[StreamConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureDynamicRange
func (s_ StreamConfiguration) CaptureDynamicRange() CaptureDynamicRange /* not a class type */ {
	rv := objc.Send[CaptureDynamicRange](s_.ID, objc.Sel("captureDynamicRange"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureDynamicRange
func (s_ StreamConfiguration) SetCaptureDynamicRange(value CaptureDynamicRange /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCaptureDynamicRange:"), value)
}


// The maximum number of frames for the queue to store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/queueDepth
func (s_ StreamConfiguration) QueueDepth() int {
	rv := objc.Send[int](s_.ID, objc.Sel("queueDepth"))
	return rv
}


// The maximum number of frames for the queue to store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/queueDepth
func (s_ StreamConfiguration) SetQueueDepth(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setQueueDepth:"), value)
}


// A background color for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/backgroundcolor
func (s_ StreamConfiguration) BackgroundColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[appkit.Color](s_.ID, objc.Sel("backgroundColor"))
	return rv
}


// A background color for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/backgroundcolor
func (s_ StreamConfiguration) SetBackgroundColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBackgroundColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/capturemicrophone
func (s_ StreamConfiguration) CaptureMicrophone() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("captureMicrophone"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/capturemicrophone
func (s_ StreamConfiguration) SetCaptureMicrophone(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCaptureMicrophone:"), value)
}


// The resolution at which to capture source content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/captureresolution
func (s_ StreamConfiguration) CaptureResolution() CaptureResolutionType /* not a class type */ {
	rv := objc.Send[CaptureResolutionType](s_.ID, objc.Sel("captureResolution"))
	return rv
}


// The resolution at which to capture source content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/captureresolution
func (s_ StreamConfiguration) SetCaptureResolution(value CaptureResolutionType /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCaptureResolution:"), value)
}


// A Boolean value that indicates whether to capture audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/capturesaudio
func (s_ StreamConfiguration) CapturesAudio() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("capturesAudio"))
	return rv
}


// A Boolean value that indicates whether to capture audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/capturesaudio
func (s_ StreamConfiguration) SetCapturesAudio(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCapturesAudio:"), value)
}


// A Boolean value that indicates if the stream only captures shadows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/capturesshadowsonly
func (s_ StreamConfiguration) CapturesShadowsOnly() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("capturesShadowsOnly"))
	return rv
}


// A Boolean value that indicates if the stream only captures shadows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/capturesshadowsonly
func (s_ StreamConfiguration) SetCapturesShadowsOnly(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCapturesShadowsOnly:"), value)
}


// The number of audio channels to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/channelcount
func (s_ StreamConfiguration) ChannelCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("channelCount"))
	return rv
}


// The number of audio channels to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/channelcount
func (s_ StreamConfiguration) SetChannelCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setChannelCount:"), value)
}


// A color matrix to apply to the output surface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/colormatrix
func (s_ StreamConfiguration) ColorMatrix() objc.IObject /* cross-framework: String */ {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("colorMatrix"))
	return rv
}


// A color matrix to apply to the output surface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/colormatrix
func (s_ StreamConfiguration) SetColorMatrix(value objc.IObject /* cross-framework: String */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setColorMatrix:"), value)
}


// A color space to use for the output buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/colorspacename
func (s_ StreamConfiguration) ColorSpaceName() objc.IObject /* cross-framework: String */ {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("colorSpaceName"))
	return rv
}


// A color space to use for the output buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/colorspacename
func (s_ StreamConfiguration) SetColorSpaceName(value objc.IObject /* cross-framework: String */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setColorSpaceName:"), value)
}


// A rectangle that specifies a destination into which to write the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/destinationrect
func (s_ StreamConfiguration) DestinationRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](s_.ID, objc.Sel("destinationRect"))
	return rv
}


// A rectangle that specifies a destination into which to write the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/destinationrect
func (s_ StreamConfiguration) SetDestinationRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDestinationRect:"), value)
}


// A Boolean value that indicates whether to exclude audio from your app during capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/excludescurrentprocessaudio
func (s_ StreamConfiguration) ExcludesCurrentProcessAudio() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("excludesCurrentProcessAudio"))
	return rv
}


// A Boolean value that indicates whether to exclude audio from your app during capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/excludescurrentprocessaudio
func (s_ StreamConfiguration) SetExcludesCurrentProcessAudio(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setExcludesCurrentProcessAudio:"), value)
}


// The height of the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/height
func (s_ StreamConfiguration) Height() int {
	rv := objc.Send[int](s_.ID, objc.Sel("height"))
	return rv
}


// The height of the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/height
func (s_ StreamConfiguration) SetHeight(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHeight:"), value)
}


// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in display style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/ignoreglobalclipdisplay
func (s_ StreamConfiguration) IgnoreGlobalClipDisplay() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreGlobalClipDisplay"))
	return rv
}


// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in display style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/ignoreglobalclipdisplay
func (s_ StreamConfiguration) SetIgnoreGlobalClipDisplay(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreGlobalClipDisplay:"), value)
}


// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in window style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/ignoreglobalclipsinglewindow
func (s_ StreamConfiguration) IgnoreGlobalClipSingleWindow() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreGlobalClipSingleWindow"))
	return rv
}


// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in window style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/ignoreglobalclipsinglewindow
func (s_ StreamConfiguration) SetIgnoreGlobalClipSingleWindow(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreGlobalClipSingleWindow:"), value)
}


// A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in display style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/ignoreshadowsdisplay
func (s_ StreamConfiguration) IgnoreShadowsDisplay() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreShadowsDisplay"))
	return rv
}


// A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in display style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/ignoreshadowsdisplay
func (s_ StreamConfiguration) SetIgnoreShadowsDisplay(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreShadowsDisplay:"), value)
}


// A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in window style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/ignoreshadowssinglewindow
func (s_ StreamConfiguration) IgnoreShadowsSingleWindow() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreShadowsSingleWindow"))
	return rv
}


// A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in window style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/ignoreshadowssinglewindow
func (s_ StreamConfiguration) SetIgnoreShadowsSingleWindow(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreShadowsSingleWindow:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/includechildwindows
func (s_ StreamConfiguration) IncludeChildWindows() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("includeChildWindows"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/includechildwindows
func (s_ StreamConfiguration) SetIncludeChildWindows(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIncludeChildWindows:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/microphonecapturedeviceid
func (s_ StreamConfiguration) MicrophoneCaptureDeviceID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("microphoneCaptureDeviceID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/microphonecapturedeviceid
func (s_ StreamConfiguration) SetMicrophoneCaptureDeviceID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMicrophoneCaptureDeviceID:"), value)
}


// The desired minimum time between frame updates, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/minimumframeinterval
func (s_ StreamConfiguration) MinimumFrameInterval() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](s_.ID, objc.Sel("minimumFrameInterval"))
	return rv
}


// The desired minimum time between frame updates, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/minimumframeinterval
func (s_ StreamConfiguration) SetMinimumFrameInterval(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumFrameInterval:"), value)
}


// A pixel format for sample buffers that a stream outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/pixelformat
func (s_ StreamConfiguration) PixelFormat() uint32 /* not a class type */ {
	rv := objc.Send[uint32](s_.ID, objc.Sel("pixelFormat"))
	return rv
}


// A pixel format for sample buffers that a stream outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/pixelformat
func (s_ StreamConfiguration) SetPixelFormat(value uint32 /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPixelFormat:"), value)
}


// A value indicating if alerts appear to presenters while using Presenter Overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/presenteroverlayprivacyalertsetting
func (s_ StreamConfiguration) PresenterOverlayPrivacyAlertSetting() PresenterOverlayAlertSetting /* not a class type */ {
	rv := objc.Send[PresenterOverlayAlertSetting](s_.ID, objc.Sel("presenterOverlayPrivacyAlertSetting"))
	return rv
}


// A value indicating if alerts appear to presenters while using Presenter Overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/presenteroverlayprivacyalertsetting
func (s_ StreamConfiguration) SetPresenterOverlayPrivacyAlertSetting(value PresenterOverlayAlertSetting /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPresenterOverlayPrivacyAlertSetting:"), value)
}


// A Boolean value that determines if the stream preserves aspect ratio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/preservesaspectratio
func (s_ StreamConfiguration) PreservesAspectRatio() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("preservesAspectRatio"))
	return rv
}


// A Boolean value that determines if the stream preserves aspect ratio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/preservesaspectratio
func (s_ StreamConfiguration) SetPreservesAspectRatio(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreservesAspectRatio:"), value)
}


// The sample rate for audio capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/samplerate
func (s_ StreamConfiguration) SampleRate() int {
	rv := objc.Send[int](s_.ID, objc.Sel("sampleRate"))
	return rv
}


// The sample rate for audio capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/samplerate
func (s_ StreamConfiguration) SetSampleRate(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSampleRate:"), value)
}


// A Boolean value that indicates whether to scale the output to fit the configured width and height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/scalestofit
func (s_ StreamConfiguration) ScalesToFit() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("scalesToFit"))
	return rv
}


// A Boolean value that indicates whether to scale the output to fit the configured width and height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/scalestofit
func (s_ StreamConfiguration) SetScalesToFit(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScalesToFit:"), value)
}


// A Boolean value that indicates if semitransparent content presents as opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/shouldbeopaque
func (s_ StreamConfiguration) ShouldBeOpaque() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldBeOpaque"))
	return rv
}


// A Boolean value that indicates if semitransparent content presents as opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/shouldbeopaque
func (s_ StreamConfiguration) SetShouldBeOpaque(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShouldBeOpaque:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/showmouseclicks
func (s_ StreamConfiguration) ShowMouseClicks() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showMouseClicks"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/showmouseclicks
func (s_ StreamConfiguration) SetShowMouseClicks(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowMouseClicks:"), value)
}


// A Boolean value that determines whether the cursor is visible in the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/showscursor
func (s_ StreamConfiguration) ShowsCursor() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsCursor"))
	return rv
}


// A Boolean value that determines whether the cursor is visible in the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/showscursor
func (s_ StreamConfiguration) SetShowsCursor(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsCursor:"), value)
}


// A rectangle that specifies the source area to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/sourcerect
func (s_ StreamConfiguration) SourceRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](s_.ID, objc.Sel("sourceRect"))
	return rv
}


// A rectangle that specifies the source area to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/sourcerect
func (s_ StreamConfiguration) SetSourceRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSourceRect:"), value)
}


// A name that you provide for identifying the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/streamname
func (s_ StreamConfiguration) StreamName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("streamName"))
	return rv
}


// A name that you provide for identifying the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/streamname
func (s_ StreamConfiguration) SetStreamName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStreamName:"), value)
}


// The width of the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/width
func (s_ StreamConfiguration) Width() int {
	rv := objc.Send[int](s_.ID, objc.Sel("width"))
	return rv
}


// The width of the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/width
func (s_ StreamConfiguration) SetWidth(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setWidth:"), value)
}



