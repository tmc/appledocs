// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
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
}

// An instance that provides the output configuration for a stream.
//
// Creating an instance of this class provides a default configuration for a stream. Only configure its properties if you need to customize the output.
//
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


//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureDynamicRange
func (s_ StreamConfiguration) CaptureDynamicRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("captureDynamicRange"))
	return rv
}


// SetCaptureDynamicRange sets the value of the captureDynamicRange property.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureDynamicRange
func (s_ StreamConfiguration) SetCaptureDynamicRange(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCaptureDynamicRange:"), value)
}

// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in window style.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreGlobalClipSingleWindow
func (s_ StreamConfiguration) IgnoreGlobalClipSingleWindow() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreGlobalClipSingleWindow"))
	return rv
}


// SetIgnoreGlobalClipSingleWindow sets the value of the ignoreGlobalClipSingleWindow property.
// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in window style.

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreGlobalClipSingleWindow
func (s_ StreamConfiguration) SetIgnoreGlobalClipSingleWindow(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreGlobalClipSingleWindow:"), value)
}

// The maximum number of frames for the queue to store.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/queueDepth
func (s_ StreamConfiguration) QueueDepth() int {
	rv := objc.Send[int](s_.ID, objc.Sel("queueDepth"))
	return rv
}


// SetQueueDepth sets the value of the queueDepth property.
// The maximum number of frames for the queue to store.

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/queueDepth
func (s_ StreamConfiguration) SetQueueDepth(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setQueueDepth:"), value)
}

// A background color for the output.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/backgroundcolor
func (s_ StreamConfiguration) BackgroundColor() appkit.Color {
	rv := objc.Send[appkit.Color](s_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// A background color for the output.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/backgroundcolor
func (s_ StreamConfiguration) SetBackgroundColor(value appkit.IColor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBackgroundColor:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/capturemicrophone
func (s_ StreamConfiguration) CaptureMicrophone() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("captureMicrophone"))
	return rv
}


// SetCaptureMicrophone sets the value of the captureMicrophone property.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/capturemicrophone
func (s_ StreamConfiguration) SetCaptureMicrophone(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCaptureMicrophone:"), value)
}

// The resolution at which to capture source content.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/captureresolution
func (s_ StreamConfiguration) CaptureResolution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("captureResolution"))
	return rv
}


// SetCaptureResolution sets the value of the captureResolution property.
// The resolution at which to capture source content.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/captureresolution
func (s_ StreamConfiguration) SetCaptureResolution(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCaptureResolution:"), value)
}

// A Boolean value that indicates whether to capture audio.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/capturesaudio
func (s_ StreamConfiguration) CapturesAudio() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("capturesAudio"))
	return rv
}


// SetCapturesAudio sets the value of the capturesAudio property.
// A Boolean value that indicates whether to capture audio.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/capturesaudio
func (s_ StreamConfiguration) SetCapturesAudio(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCapturesAudio:"), value)
}

// A Boolean value that indicates if the stream only captures shadows.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/capturesshadowsonly
func (s_ StreamConfiguration) CapturesShadowsOnly() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("capturesShadowsOnly"))
	return rv
}


// SetCapturesShadowsOnly sets the value of the capturesShadowsOnly property.
// A Boolean value that indicates if the stream only captures shadows.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/capturesshadowsonly
func (s_ StreamConfiguration) SetCapturesShadowsOnly(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCapturesShadowsOnly:"), value)
}

// The number of audio channels to capture.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/channelcount
func (s_ StreamConfiguration) ChannelCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("channelCount"))
	return rv
}


// SetChannelCount sets the value of the channelCount property.
// The number of audio channels to capture.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/channelcount
func (s_ StreamConfiguration) SetChannelCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setChannelCount:"), value)
}

// A color matrix to apply to the output surface.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/colormatrix
func (s_ StreamConfiguration) ColorMatrix() foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("colorMatrix"))
	return rv
}


// SetColorMatrix sets the value of the colorMatrix property.
// A color matrix to apply to the output surface.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/colormatrix
func (s_ StreamConfiguration) SetColorMatrix(value foundation.IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setColorMatrix:"), value)
}

// A color space to use for the output buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/colorspacename
func (s_ StreamConfiguration) ColorSpaceName() foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("colorSpaceName"))
	return rv
}


// SetColorSpaceName sets the value of the colorSpaceName property.
// A color space to use for the output buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/colorspacename
func (s_ StreamConfiguration) SetColorSpaceName(value foundation.IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setColorSpaceName:"), value)
}

// A rectangle that specifies a destination into which to write the output.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/destinationrect
func (s_ StreamConfiguration) DestinationRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("destinationRect"))
	return rv
}


// SetDestinationRect sets the value of the destinationRect property.
// A rectangle that specifies a destination into which to write the output.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/destinationrect
func (s_ StreamConfiguration) SetDestinationRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDestinationRect:"), value)
}

// A Boolean value that indicates whether to exclude audio from your app during capture.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/excludescurrentprocessaudio
func (s_ StreamConfiguration) ExcludesCurrentProcessAudio() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("excludesCurrentProcessAudio"))
	return rv
}


// SetExcludesCurrentProcessAudio sets the value of the excludesCurrentProcessAudio property.
// A Boolean value that indicates whether to exclude audio from your app during capture.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/excludescurrentprocessaudio
func (s_ StreamConfiguration) SetExcludesCurrentProcessAudio(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setExcludesCurrentProcessAudio:"), value)
}

// The height of the output.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/height
func (s_ StreamConfiguration) Height() int {
	rv := objc.Send[int](s_.ID, objc.Sel("height"))
	return rv
}


// SetHeight sets the value of the height property.
// The height of the output.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/height
func (s_ StreamConfiguration) SetHeight(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHeight:"), value)
}

// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in display style.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/ignoreglobalclipdisplay
func (s_ StreamConfiguration) IgnoreGlobalClipDisplay() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreGlobalClipDisplay"))
	return rv
}


// SetIgnoreGlobalClipDisplay sets the value of the ignoreGlobalClipDisplay property.
// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in display style.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/ignoreglobalclipdisplay
func (s_ StreamConfiguration) SetIgnoreGlobalClipDisplay(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreGlobalClipDisplay:"), value)
}

// A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in display style.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/ignoreshadowsdisplay
func (s_ StreamConfiguration) IgnoreShadowsDisplay() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreShadowsDisplay"))
	return rv
}


// SetIgnoreShadowsDisplay sets the value of the ignoreShadowsDisplay property.
// A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in display style.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/ignoreshadowsdisplay
func (s_ StreamConfiguration) SetIgnoreShadowsDisplay(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreShadowsDisplay:"), value)
}

// A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in window style.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/ignoreshadowssinglewindow
func (s_ StreamConfiguration) IgnoreShadowsSingleWindow() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ignoreShadowsSingleWindow"))
	return rv
}


// SetIgnoreShadowsSingleWindow sets the value of the ignoreShadowsSingleWindow property.
// A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in window style.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/ignoreshadowssinglewindow
func (s_ StreamConfiguration) SetIgnoreShadowsSingleWindow(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIgnoreShadowsSingleWindow:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/includechildwindows
func (s_ StreamConfiguration) IncludeChildWindows() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("includeChildWindows"))
	return rv
}


// SetIncludeChildWindows sets the value of the includeChildWindows property.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/includechildwindows
func (s_ StreamConfiguration) SetIncludeChildWindows(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIncludeChildWindows:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/microphonecapturedeviceid
func (s_ StreamConfiguration) MicrophoneCaptureDeviceID() appkit.string {
	rv := objc.Send[appkit.string](s_.ID, objc.Sel("microphoneCaptureDeviceID"))
	return rv
}


// SetMicrophoneCaptureDeviceID sets the value of the microphoneCaptureDeviceID property.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/microphonecapturedeviceid
func (s_ StreamConfiguration) SetMicrophoneCaptureDeviceID(value appkit.string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMicrophoneCaptureDeviceID:"), value)
}

// The desired minimum time between frame updates, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/minimumframeinterval
func (s_ StreamConfiguration) MinimumFrameInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("minimumFrameInterval"))
	return rv
}


// SetMinimumFrameInterval sets the value of the minimumFrameInterval property.
// The desired minimum time between frame updates, in seconds.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/minimumframeinterval
func (s_ StreamConfiguration) SetMinimumFrameInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumFrameInterval:"), value)
}

// A pixel format for sample buffers that a stream outputs.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/pixelformat
func (s_ StreamConfiguration) PixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("pixelFormat"))
	return rv
}


// SetPixelFormat sets the value of the pixelFormat property.
// A pixel format for sample buffers that a stream outputs.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/pixelformat
func (s_ StreamConfiguration) SetPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPixelFormat:"), value)
}

// A value indicating if alerts appear to presenters while using Presenter Overlay.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/presenteroverlayprivacyalertsetting
func (s_ StreamConfiguration) PresenterOverlayPrivacyAlertSetting() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("presenterOverlayPrivacyAlertSetting"))
	return rv
}


// SetPresenterOverlayPrivacyAlertSetting sets the value of the presenterOverlayPrivacyAlertSetting property.
// A value indicating if alerts appear to presenters while using Presenter Overlay.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/presenteroverlayprivacyalertsetting
func (s_ StreamConfiguration) SetPresenterOverlayPrivacyAlertSetting(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPresenterOverlayPrivacyAlertSetting:"), value)
}

// A Boolean value that determines if the stream preserves aspect ratio.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/preservesaspectratio
func (s_ StreamConfiguration) PreservesAspectRatio() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("preservesAspectRatio"))
	return rv
}


// SetPreservesAspectRatio sets the value of the preservesAspectRatio property.
// A Boolean value that determines if the stream preserves aspect ratio.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/preservesaspectratio
func (s_ StreamConfiguration) SetPreservesAspectRatio(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreservesAspectRatio:"), value)
}

// The sample rate for audio capture.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/samplerate
func (s_ StreamConfiguration) SampleRate() int {
	rv := objc.Send[int](s_.ID, objc.Sel("sampleRate"))
	return rv
}


// SetSampleRate sets the value of the sampleRate property.
// The sample rate for audio capture.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/samplerate
func (s_ StreamConfiguration) SetSampleRate(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSampleRate:"), value)
}

// A Boolean value that indicates whether to scale the output to fit the configured width and height.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/scalestofit
func (s_ StreamConfiguration) ScalesToFit() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("scalesToFit"))
	return rv
}


// SetScalesToFit sets the value of the scalesToFit property.
// A Boolean value that indicates whether to scale the output to fit the configured width and height.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/scalestofit
func (s_ StreamConfiguration) SetScalesToFit(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScalesToFit:"), value)
}

// A Boolean value that indicates if semitransparent content presents as opaque.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/shouldbeopaque
func (s_ StreamConfiguration) ShouldBeOpaque() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldBeOpaque"))
	return rv
}


// SetShouldBeOpaque sets the value of the shouldBeOpaque property.
// A Boolean value that indicates if semitransparent content presents as opaque.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/shouldbeopaque
func (s_ StreamConfiguration) SetShouldBeOpaque(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShouldBeOpaque:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/showmouseclicks
func (s_ StreamConfiguration) ShowMouseClicks() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showMouseClicks"))
	return rv
}


// SetShowMouseClicks sets the value of the showMouseClicks property.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/showmouseclicks
func (s_ StreamConfiguration) SetShowMouseClicks(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowMouseClicks:"), value)
}

// A Boolean value that determines whether the cursor is visible in the stream.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/showscursor
func (s_ StreamConfiguration) ShowsCursor() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsCursor"))
	return rv
}


// SetShowsCursor sets the value of the showsCursor property.
// A Boolean value that determines whether the cursor is visible in the stream.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/showscursor
func (s_ StreamConfiguration) SetShowsCursor(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsCursor:"), value)
}

// A rectangle that specifies the source area to capture.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/sourcerect
func (s_ StreamConfiguration) SourceRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("sourceRect"))
	return rv
}


// SetSourceRect sets the value of the sourceRect property.
// A rectangle that specifies the source area to capture.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/sourcerect
func (s_ StreamConfiguration) SetSourceRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSourceRect:"), value)
}

// A name that you provide for identifying the stream.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/streamname
func (s_ StreamConfiguration) StreamName() appkit.string {
	rv := objc.Send[appkit.string](s_.ID, objc.Sel("streamName"))
	return rv
}


// SetStreamName sets the value of the streamName property.
// A name that you provide for identifying the stream.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/streamname
func (s_ StreamConfiguration) SetStreamName(value appkit.string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStreamName:"), value)
}

// The width of the output.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/width
func (s_ StreamConfiguration) Width() int {
	rv := objc.Send[int](s_.ID, objc.Sel("width"))
	return rv
}


// SetWidth sets the value of the width property.
// The width of the output.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scstreamconfiguration/width
func (s_ StreamConfiguration) SetWidth(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setWidth:"), value)
}



