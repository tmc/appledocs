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
}

// An object that represents a connection from a capture input to a capture output.
//
// Capture inputs have one or more input ports (instances of ). Capture outputs can accept data from one or more sources (for example, an object accepts both video and audio data). You can add an instance to a session using the method only if the method returns . When using the or method, the session forms connections automatically between all compatible inputs and outputs. You only need to add connections manually when adding an input or output with no connections. You can also use connections to enable or disable the flow of data from a given input or to a given output.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMinFrameDuration
func (c_ CaptureConnection) VideoMinFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoMinFrameDuration"))
	return rv
}


// SetVideoMinFrameDuration sets the value of the videoMinFrameDuration property.
// The smallest time interval the connection can apply between consecutive video frames.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMinFrameDuration
func (c_ CaptureConnection) SetVideoMinFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMinFrameDuration:"), value)
}

// An orientation that tells the connection how to rotate a video flowing through it.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoOrientation
func (c_ CaptureConnection) VideoOrientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoOrientation"))
	return rv
}


// SetVideoOrientation sets the value of the videoOrientation property.
// An orientation that tells the connection how to rotate a video flowing through it.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoOrientation
func (c_ CaptureConnection) SetVideoOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoOrientation:"), value)
}

// The connection’s current stabilization mode.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/activevideostabilizationmode
func (c_ CaptureConnection) ActiveVideoStabilizationMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("activeVideoStabilizationMode"))
	return rv
}


// SetActiveVideoStabilizationMode sets the value of the activeVideoStabilizationMode property.
// The connection’s current stabilization mode.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/activevideostabilizationmode
func (c_ CaptureConnection) SetActiveVideoStabilizationMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveVideoStabilizationMode:"), value)
}

// An array of audio channels that the connection provides.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/audiochannels
func (c_ CaptureConnection) AudioChannels() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("audioChannels"))
	return rv
}


// SetAudioChannels sets the value of the audioChannels property.
// An array of audio channels that the connection provides.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/audiochannels
func (c_ CaptureConnection) SetAudioChannels(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioChannels:"), value)
}

// A Boolean value that indicates whether you can enable mirroring based on a session’s configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/automaticallyadjustsvideomirroring
func (c_ CaptureConnection) AutomaticallyAdjustsVideoMirroring() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyAdjustsVideoMirroring"))
	return rv
}


// SetAutomaticallyAdjustsVideoMirroring sets the value of the automaticallyAdjustsVideoMirroring property.
// A Boolean value that indicates whether you can enable mirroring based on a session’s configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/automaticallyadjustsvideomirroring
func (c_ CaptureConnection) SetAutomaticallyAdjustsVideoMirroring(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallyAdjustsVideoMirroring:"), value)
}

// A Boolean value that indicates whether the system enables video stabilization when it’s available.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/enablesvideostabilizationwhenavailable
func (c_ CaptureConnection) EnablesVideoStabilizationWhenAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enablesVideoStabilizationWhenAvailable"))
	return rv
}


// SetEnablesVideoStabilizationWhenAvailable sets the value of the enablesVideoStabilizationWhenAvailable property.
// A Boolean value that indicates whether the system enables video stabilization when it’s available.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/enablesvideostabilizationwhenavailable
func (c_ CaptureConnection) SetEnablesVideoStabilizationWhenAvailable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnablesVideoStabilizationWhenAvailable:"), value)
}

// An array of the connection’s input ports.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/inputports
func (c_ CaptureConnection) InputPorts() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("inputPorts"))
	return rv
}


// SetInputPorts sets the value of the inputPorts property.
// An array of the connection’s input ports.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/inputports
func (c_ CaptureConnection) SetInputPorts(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInputPorts:"), value)
}

// Indicates whether the connection is active.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isactive
func (c_ CaptureConnection) IsActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isActive"))
	return rv
}


// SetIsActive sets the value of the isActive property.
// Indicates whether the connection is active.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isactive
func (c_ CaptureConnection) SetIsActive(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsActive:"), value)
}

// A Boolean value that indicates whether the connection can configure the capture pipeline to deliver camera intrinsics information.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliveryenabled
func (c_ CaptureConnection) IsCameraIntrinsicMatrixDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraIntrinsicMatrixDeliveryEnabled"))
	return rv
}


// SetIsCameraIntrinsicMatrixDeliveryEnabled sets the value of the isCameraIntrinsicMatrixDeliveryEnabled property.
// A Boolean value that indicates whether the connection can configure the capture pipeline to deliver camera intrinsics information.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliveryenabled
func (c_ CaptureConnection) SetIsCameraIntrinsicMatrixDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraIntrinsicMatrixDeliveryEnabled:"), value)
}

// A Boolean value that indicates whether the capture connection currently supports delivering camera intrinsics information.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliverysupported
func (c_ CaptureConnection) IsCameraIntrinsicMatrixDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraIntrinsicMatrixDeliverySupported"))
	return rv
}


// SetIsCameraIntrinsicMatrixDeliverySupported sets the value of the isCameraIntrinsicMatrixDeliverySupported property.
// A Boolean value that indicates whether the capture connection currently supports delivering camera intrinsics information.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliverysupported
func (c_ CaptureConnection) SetIsCameraIntrinsicMatrixDeliverySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraIntrinsicMatrixDeliverySupported:"), value)
}

// Turns the connection on and off.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isenabled
func (c_ CaptureConnection) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
// Turns the connection on and off.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isenabled
func (c_ CaptureConnection) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}

// A Boolean value that indicates whether the connection supports setting a video field mode.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideofieldmodesupported
func (c_ CaptureConnection) IsVideoFieldModeSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoFieldModeSupported"))
	return rv
}


// SetIsVideoFieldModeSupported sets the value of the isVideoFieldModeSupported property.
// A Boolean value that indicates whether the connection supports setting a video field mode.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideofieldmodesupported
func (c_ CaptureConnection) SetIsVideoFieldModeSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoFieldModeSupported:"), value)
}

// A Boolean value that indicates whether the connection supports a maximum frame duration.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomaxframedurationsupported
func (c_ CaptureConnection) IsVideoMaxFrameDurationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMaxFrameDurationSupported"))
	return rv
}


// SetIsVideoMaxFrameDurationSupported sets the value of the isVideoMaxFrameDurationSupported property.
// A Boolean value that indicates whether the connection supports a maximum frame duration.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomaxframedurationsupported
func (c_ CaptureConnection) SetIsVideoMaxFrameDurationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMaxFrameDurationSupported:"), value)
}

// A Boolean value that indicates whether the connection supports a minimum frame duration.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideominframedurationsupported
func (c_ CaptureConnection) IsVideoMinFrameDurationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMinFrameDurationSupported"))
	return rv
}


// SetIsVideoMinFrameDurationSupported sets the value of the isVideoMinFrameDurationSupported property.
// A Boolean value that indicates whether the connection supports a minimum frame duration.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideominframedurationsupported
func (c_ CaptureConnection) SetIsVideoMinFrameDurationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMinFrameDurationSupported:"), value)
}

// A Boolean value that indicates whether the connection horizontally flips the video flowing through it.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirrored
func (c_ CaptureConnection) IsVideoMirrored() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMirrored"))
	return rv
}


// SetIsVideoMirrored sets the value of the isVideoMirrored property.
// A Boolean value that indicates whether the connection horizontally flips the video flowing through it.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirrored
func (c_ CaptureConnection) SetIsVideoMirrored(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMirrored:"), value)
}

// A Boolean value that indicates whether the connection supports video mirroring.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirroringsupported
func (c_ CaptureConnection) IsVideoMirroringSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMirroringSupported"))
	return rv
}


// SetIsVideoMirroringSupported sets the value of the isVideoMirroringSupported property.
// A Boolean value that indicates whether the connection supports video mirroring.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirroringsupported
func (c_ CaptureConnection) SetIsVideoMirroringSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMirroringSupported:"), value)
}

// A Boolean value that indicates whether the connection supports changing the orientation of the video.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideoorientationsupported
func (c_ CaptureConnection) IsVideoOrientationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoOrientationSupported"))
	return rv
}


// SetIsVideoOrientationSupported sets the value of the isVideoOrientationSupported property.
// A Boolean value that indicates whether the connection supports changing the orientation of the video.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideoorientationsupported
func (c_ CaptureConnection) SetIsVideoOrientationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoOrientationSupported:"), value)
}

// A Boolean value that indicates whether video stabilization is active for the connection.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationenabled
func (c_ CaptureConnection) IsVideoStabilizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoStabilizationEnabled"))
	return rv
}


// SetIsVideoStabilizationEnabled sets the value of the isVideoStabilizationEnabled property.
// A Boolean value that indicates whether video stabilization is active for the connection.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationenabled
func (c_ CaptureConnection) SetIsVideoStabilizationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoStabilizationEnabled:"), value)
}

// A Boolean value that indicates whether this connection supports video stabilization.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationsupported
func (c_ CaptureConnection) IsVideoStabilizationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoStabilizationSupported"))
	return rv
}


// SetIsVideoStabilizationSupported sets the value of the isVideoStabilizationSupported property.
// A Boolean value that indicates whether this connection supports video stabilization.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationsupported
func (c_ CaptureConnection) SetIsVideoStabilizationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoStabilizationSupported:"), value)
}

// The connection’s output port, if applicable.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/output
func (c_ CaptureConnection) Output() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("output"))
	return rv
}


// SetOutput sets the value of the output property.
// The connection’s output port, if applicable.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/output
func (c_ CaptureConnection) SetOutput(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutput:"), value)
}

// The stabilization mode that’s the most appropriate for a video connection.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/preferredvideostabilizationmode
func (c_ CaptureConnection) PreferredVideoStabilizationMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("preferredVideoStabilizationMode"))
	return rv
}


// SetPreferredVideoStabilizationMode sets the value of the preferredVideoStabilizationMode property.
// The stabilization mode that’s the most appropriate for a video connection.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/preferredvideostabilizationmode
func (c_ CaptureConnection) SetPreferredVideoStabilizationMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredVideoStabilizationMode:"), value)
}

// A setting that tells the connection how to interlace video flowing through it.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videofieldmode
func (c_ CaptureConnection) VideoFieldMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoFieldMode"))
	return rv
}


// SetVideoFieldMode sets the value of the videoFieldMode property.
// A setting that tells the connection how to interlace video flowing through it.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videofieldmode
func (c_ CaptureConnection) SetVideoFieldMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoFieldMode:"), value)
}

// The largest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videomaxframeduration
func (c_ CaptureConnection) VideoMaxFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoMaxFrameDuration"))
	return rv
}


// SetVideoMaxFrameDuration sets the value of the videoMaxFrameDuration property.
// The largest time interval the connection can apply between consecutive video frames.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videomaxframeduration
func (c_ CaptureConnection) SetVideoMaxFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMaxFrameDuration:"), value)
}

// The connection’s maximum video scale and crop factor.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videomaxscaleandcropfactor
func (c_ CaptureConnection) VideoMaxScaleAndCropFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMaxScaleAndCropFactor"))
	return rv
}


// SetVideoMaxScaleAndCropFactor sets the value of the videoMaxScaleAndCropFactor property.
// The connection’s maximum video scale and crop factor.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videomaxscaleandcropfactor
func (c_ CaptureConnection) SetVideoMaxScaleAndCropFactor(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMaxScaleAndCropFactor:"), value)
}

// The video preview layer associated with the connection.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videopreviewlayer
func (c_ CaptureConnection) VideoPreviewLayer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoPreviewLayer"))
	return rv
}


// SetVideoPreviewLayer sets the value of the videoPreviewLayer property.
// The video preview layer associated with the connection.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videopreviewlayer
func (c_ CaptureConnection) SetVideoPreviewLayer(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoPreviewLayer:"), value)
}

// A rotation angle the connection applies to a video flowing through it.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videorotationangle
func (c_ CaptureConnection) VideoRotationAngle() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoRotationAngle"))
	return rv
}


// SetVideoRotationAngle sets the value of the videoRotationAngle property.
// A rotation angle the connection applies to a video flowing through it.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videorotationangle
func (c_ CaptureConnection) SetVideoRotationAngle(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoRotationAngle:"), value)
}

// The current scale and crop factor the video output uses.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videoscaleandcropfactor
func (c_ CaptureConnection) VideoScaleAndCropFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoScaleAndCropFactor"))
	return rv
}


// SetVideoScaleAndCropFactor sets the value of the videoScaleAndCropFactor property.
// The current scale and crop factor the video output uses.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videoscaleandcropfactor
func (c_ CaptureConnection) SetVideoScaleAndCropFactor(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoScaleAndCropFactor:"), value)
}



