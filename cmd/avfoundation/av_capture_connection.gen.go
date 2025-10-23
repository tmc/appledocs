// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CaptureConnection] class.
var (
	CaptureConnectionClass     _CaptureConnectionClass
	CaptureConnectionClassOnce sync.Once
)

func getCaptureConnectionClass() _CaptureConnectionClass {
	CaptureConnectionClassOnce.Do(func() {
		CaptureConnectionClass = _CaptureConnectionClass{objc.GetClass("AVCaptureConnection")}
	})
	return CaptureConnectionClass
}

type _CaptureConnectionClass struct {
	class objc.Class
}

// An interface definition for the [CaptureConnection] class.
type ICaptureConnection interface {
	objectivec.IObject
	VideoMinFrameDuration() unsafe.Pointer
	SetVideoMinFrameDuration(value unsafe.Pointer)
	VideoOrientation() unsafe.Pointer
	SetVideoOrientation(value unsafe.Pointer)
	ActiveVideoStabilizationMode() unsafe.Pointer
	SetActiveVideoStabilizationMode(value unsafe.Pointer)
	AudioChannels() AVCaptureAudioChannel
	SetAudioChannels(value IAVCaptureAudioChannel)
	AutomaticallyAdjustsVideoMirroring() bool
	SetAutomaticallyAdjustsVideoMirroring(value bool)
	EnablesVideoStabilizationWhenAvailable() bool
	SetEnablesVideoStabilizationWhenAvailable(value bool)
	InputPorts() AVCaptureInputPort
	SetInputPorts(value IAVCaptureInputPort)
	IsActive() bool
	SetIsActive(value bool)
	IsCameraIntrinsicMatrixDeliveryEnabled() bool
	SetIsCameraIntrinsicMatrixDeliveryEnabled(value bool)
	IsCameraIntrinsicMatrixDeliverySupported() bool
	SetIsCameraIntrinsicMatrixDeliverySupported(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsVideoFieldModeSupported() bool
	SetIsVideoFieldModeSupported(value bool)
	IsVideoMaxFrameDurationSupported() bool
	SetIsVideoMaxFrameDurationSupported(value bool)
	IsVideoMinFrameDurationSupported() bool
	SetIsVideoMinFrameDurationSupported(value bool)
	IsVideoMirrored() bool
	SetIsVideoMirrored(value bool)
	IsVideoMirroringSupported() bool
	SetIsVideoMirroringSupported(value bool)
	IsVideoOrientationSupported() bool
	SetIsVideoOrientationSupported(value bool)
	IsVideoStabilizationEnabled() bool
	SetIsVideoStabilizationEnabled(value bool)
	IsVideoStabilizationSupported() bool
	SetIsVideoStabilizationSupported(value bool)
	Output() AVCaptureOutput
	SetOutput(value IAVCaptureOutput)
	PreferredVideoStabilizationMode() unsafe.Pointer
	SetPreferredVideoStabilizationMode(value unsafe.Pointer)
	VideoFieldMode() unsafe.Pointer
	SetVideoFieldMode(value unsafe.Pointer)
	VideoMaxFrameDuration() unsafe.Pointer
	SetVideoMaxFrameDuration(value unsafe.Pointer)
	VideoMaxScaleAndCropFactor() float64
	SetVideoMaxScaleAndCropFactor(value float64)
	VideoPreviewLayer() AVCaptureVideoPreviewLayer
	SetVideoPreviewLayer(value IAVCaptureVideoPreviewLayer)
	VideoRotationAngle() float64
	SetVideoRotationAngle(value float64)
	VideoScaleAndCropFactor() float64
	SetVideoScaleAndCropFactor(value float64)
}

// An object that represents a connection from a capture input to a capture output.
//
// Capture inputs have one or more input ports (instances of ). Capture outputs can accept data from one or more sources (for example, an object accepts both video and audio data). You can add an instance to a session using the method only if the method returns . When using the or method, the session forms connections automatically between all compatible inputs and outputs. You only need to add connections manually when adding an input or output with no connections. You can also use connections to enable or disable the flow of data from a given input or to a given output.


// An object that represents a connection from a capture input to a capture output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection
type CaptureConnection struct {
	objectivec.Object
}

// CaptureConnectionFrom constructs a [CaptureConnection] from an unsafe.Pointer.
//
// An object that represents a connection from a capture input to a capture output.
func CaptureConnectionFrom(ptr unsafe.Pointer) CaptureConnection {
	return CaptureConnection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureConnectionClass) Alloc() CaptureConnection {
	rv := objc.Send[CaptureConnection](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureConnectionClass) New() CaptureConnection {
	rv := objc.Send[CaptureConnection](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureConnection) Init() CaptureConnection {
	rv := objc.Send[CaptureConnection](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureConnection) Autorelease() CaptureConnection {
	rv := objc.Send[CaptureConnection](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureConnection creates a new CaptureConnection instance.
func NewCaptureConnection() CaptureConnection {
	return getCaptureConnectionClass().New()
}



// The smallest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMinFrameDuration
func (c_ CaptureConnection) VideoMinFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoMinFrameDuration"))
	return rv
}


// The smallest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMinFrameDuration
func (c_ CaptureConnection) SetVideoMinFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMinFrameDuration:"), value)
}


// An orientation that tells the connection how to rotate a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoOrientation
func (c_ CaptureConnection) VideoOrientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoOrientation"))
	return rv
}


// An orientation that tells the connection how to rotate a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoOrientation
func (c_ CaptureConnection) SetVideoOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoOrientation:"), value)
}


// The connection’s current stabilization mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/activevideostabilizationmode
func (c_ CaptureConnection) ActiveVideoStabilizationMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("activeVideoStabilizationMode"))
	return rv
}


// The connection’s current stabilization mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/activevideostabilizationmode
func (c_ CaptureConnection) SetActiveVideoStabilizationMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveVideoStabilizationMode:"), value)
}


// An array of audio channels that the connection provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/audiochannels
func (c_ CaptureConnection) AudioChannels() AVCaptureAudioChannel {
	rv := objc.Send[AVCaptureAudioChannel](c_.ID, objc.Sel("audioChannels"))
	return rv
}


// An array of audio channels that the connection provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/audiochannels
func (c_ CaptureConnection) SetAudioChannels(value IAVCaptureAudioChannel) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioChannels:"), value)
}


// A Boolean value that indicates whether you can enable mirroring based on a session’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/automaticallyadjustsvideomirroring
func (c_ CaptureConnection) AutomaticallyAdjustsVideoMirroring() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyAdjustsVideoMirroring"))
	return rv
}


// A Boolean value that indicates whether you can enable mirroring based on a session’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/automaticallyadjustsvideomirroring
func (c_ CaptureConnection) SetAutomaticallyAdjustsVideoMirroring(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallyAdjustsVideoMirroring:"), value)
}


// A Boolean value that indicates whether the system enables video stabilization when it’s available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/enablesvideostabilizationwhenavailable
func (c_ CaptureConnection) EnablesVideoStabilizationWhenAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enablesVideoStabilizationWhenAvailable"))
	return rv
}


// A Boolean value that indicates whether the system enables video stabilization when it’s available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/enablesvideostabilizationwhenavailable
func (c_ CaptureConnection) SetEnablesVideoStabilizationWhenAvailable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnablesVideoStabilizationWhenAvailable:"), value)
}


// An array of the connection’s input ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/inputports
func (c_ CaptureConnection) InputPorts() AVCaptureInputPort {
	rv := objc.Send[AVCaptureInputPort](c_.ID, objc.Sel("inputPorts"))
	return rv
}


// An array of the connection’s input ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/inputports
func (c_ CaptureConnection) SetInputPorts(value IAVCaptureInputPort) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInputPorts:"), value)
}


// Indicates whether the connection is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isactive
func (c_ CaptureConnection) IsActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isActive"))
	return rv
}


// Indicates whether the connection is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isactive
func (c_ CaptureConnection) SetIsActive(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsActive:"), value)
}


// A Boolean value that indicates whether the connection can configure the capture pipeline to deliver camera intrinsics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliveryenabled
func (c_ CaptureConnection) IsCameraIntrinsicMatrixDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraIntrinsicMatrixDeliveryEnabled"))
	return rv
}


// A Boolean value that indicates whether the connection can configure the capture pipeline to deliver camera intrinsics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliveryenabled
func (c_ CaptureConnection) SetIsCameraIntrinsicMatrixDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraIntrinsicMatrixDeliveryEnabled:"), value)
}


// A Boolean value that indicates whether the capture connection currently supports delivering camera intrinsics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliverysupported
func (c_ CaptureConnection) IsCameraIntrinsicMatrixDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraIntrinsicMatrixDeliverySupported"))
	return rv
}


// A Boolean value that indicates whether the capture connection currently supports delivering camera intrinsics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliverysupported
func (c_ CaptureConnection) SetIsCameraIntrinsicMatrixDeliverySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraIntrinsicMatrixDeliverySupported:"), value)
}


// Turns the connection on and off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isenabled
func (c_ CaptureConnection) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}


// Turns the connection on and off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isenabled
func (c_ CaptureConnection) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value that indicates whether the connection supports setting a video field mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideofieldmodesupported
func (c_ CaptureConnection) IsVideoFieldModeSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoFieldModeSupported"))
	return rv
}


// A Boolean value that indicates whether the connection supports setting a video field mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideofieldmodesupported
func (c_ CaptureConnection) SetIsVideoFieldModeSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoFieldModeSupported:"), value)
}


// A Boolean value that indicates whether the connection supports a maximum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomaxframedurationsupported
func (c_ CaptureConnection) IsVideoMaxFrameDurationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMaxFrameDurationSupported"))
	return rv
}


// A Boolean value that indicates whether the connection supports a maximum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomaxframedurationsupported
func (c_ CaptureConnection) SetIsVideoMaxFrameDurationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMaxFrameDurationSupported:"), value)
}


// A Boolean value that indicates whether the connection supports a minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideominframedurationsupported
func (c_ CaptureConnection) IsVideoMinFrameDurationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMinFrameDurationSupported"))
	return rv
}


// A Boolean value that indicates whether the connection supports a minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideominframedurationsupported
func (c_ CaptureConnection) SetIsVideoMinFrameDurationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMinFrameDurationSupported:"), value)
}


// A Boolean value that indicates whether the connection horizontally flips the video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirrored
func (c_ CaptureConnection) IsVideoMirrored() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMirrored"))
	return rv
}


// A Boolean value that indicates whether the connection horizontally flips the video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirrored
func (c_ CaptureConnection) SetIsVideoMirrored(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMirrored:"), value)
}


// A Boolean value that indicates whether the connection supports video mirroring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirroringsupported
func (c_ CaptureConnection) IsVideoMirroringSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMirroringSupported"))
	return rv
}


// A Boolean value that indicates whether the connection supports video mirroring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirroringsupported
func (c_ CaptureConnection) SetIsVideoMirroringSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMirroringSupported:"), value)
}


// A Boolean value that indicates whether the connection supports changing the orientation of the video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideoorientationsupported
func (c_ CaptureConnection) IsVideoOrientationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoOrientationSupported"))
	return rv
}


// A Boolean value that indicates whether the connection supports changing the orientation of the video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideoorientationsupported
func (c_ CaptureConnection) SetIsVideoOrientationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoOrientationSupported:"), value)
}


// A Boolean value that indicates whether video stabilization is active for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationenabled
func (c_ CaptureConnection) IsVideoStabilizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoStabilizationEnabled"))
	return rv
}


// A Boolean value that indicates whether video stabilization is active for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationenabled
func (c_ CaptureConnection) SetIsVideoStabilizationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoStabilizationEnabled:"), value)
}


// A Boolean value that indicates whether this connection supports video stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationsupported
func (c_ CaptureConnection) IsVideoStabilizationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoStabilizationSupported"))
	return rv
}


// A Boolean value that indicates whether this connection supports video stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationsupported
func (c_ CaptureConnection) SetIsVideoStabilizationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoStabilizationSupported:"), value)
}


// The connection’s output port, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/output
func (c_ CaptureConnection) Output() AVCaptureOutput {
	rv := objc.Send[AVCaptureOutput](c_.ID, objc.Sel("output"))
	return rv
}


// The connection’s output port, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/output
func (c_ CaptureConnection) SetOutput(value IAVCaptureOutput) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutput:"), value)
}


// The stabilization mode that’s the most appropriate for a video connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/preferredvideostabilizationmode
func (c_ CaptureConnection) PreferredVideoStabilizationMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("preferredVideoStabilizationMode"))
	return rv
}


// The stabilization mode that’s the most appropriate for a video connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/preferredvideostabilizationmode
func (c_ CaptureConnection) SetPreferredVideoStabilizationMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredVideoStabilizationMode:"), value)
}


// A setting that tells the connection how to interlace video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videofieldmode
func (c_ CaptureConnection) VideoFieldMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoFieldMode"))
	return rv
}


// A setting that tells the connection how to interlace video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videofieldmode
func (c_ CaptureConnection) SetVideoFieldMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoFieldMode:"), value)
}


// The largest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videomaxframeduration
func (c_ CaptureConnection) VideoMaxFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoMaxFrameDuration"))
	return rv
}


// The largest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videomaxframeduration
func (c_ CaptureConnection) SetVideoMaxFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMaxFrameDuration:"), value)
}


// The connection’s maximum video scale and crop factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videomaxscaleandcropfactor
func (c_ CaptureConnection) VideoMaxScaleAndCropFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMaxScaleAndCropFactor"))
	return rv
}


// The connection’s maximum video scale and crop factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videomaxscaleandcropfactor
func (c_ CaptureConnection) SetVideoMaxScaleAndCropFactor(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMaxScaleAndCropFactor:"), value)
}


// The video preview layer associated with the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videopreviewlayer
func (c_ CaptureConnection) VideoPreviewLayer() AVCaptureVideoPreviewLayer {
	rv := objc.Send[AVCaptureVideoPreviewLayer](c_.ID, objc.Sel("videoPreviewLayer"))
	return rv
}


// The video preview layer associated with the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videopreviewlayer
func (c_ CaptureConnection) SetVideoPreviewLayer(value IAVCaptureVideoPreviewLayer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoPreviewLayer:"), value)
}


// A rotation angle the connection applies to a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videorotationangle
func (c_ CaptureConnection) VideoRotationAngle() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoRotationAngle"))
	return rv
}


// A rotation angle the connection applies to a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videorotationangle
func (c_ CaptureConnection) SetVideoRotationAngle(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoRotationAngle:"), value)
}


// The current scale and crop factor the video output uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videoscaleandcropfactor
func (c_ CaptureConnection) VideoScaleAndCropFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoScaleAndCropFactor"))
	return rv
}


// The current scale and crop factor the video output uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videoscaleandcropfactor
func (c_ CaptureConnection) SetVideoScaleAndCropFactor(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoScaleAndCropFactor:"), value)
}



