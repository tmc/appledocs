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
	// properties:
	ActiveVideoStabilizationMode() AVCaptureVideoStabilizationMode /* foo */
	SetActiveVideoStabilizationMode(value AVCaptureVideoStabilizationMode /* foo */)
	AudioChannels() IAVCaptureAudioChannel
	SetAudioChannels(value IAVCaptureAudioChannel)
	AutomaticallyAdjustsVideoMirroring() bool /* primitive/slice/pointer */
	SetAutomaticallyAdjustsVideoMirroring(value bool /* primitive/slice/pointer */)
	EnablesVideoStabilizationWhenAvailable() bool /* primitive/slice/pointer */
	SetEnablesVideoStabilizationWhenAvailable(value bool /* primitive/slice/pointer */)
	InputPorts() AVCaptureInputPort /* foo */
	SetInputPorts(value AVCaptureInputPort /* foo */)
	IsActive() bool /* primitive/slice/pointer */
	SetIsActive(value bool /* primitive/slice/pointer */)
	IsCameraIntrinsicMatrixDeliveryEnabled() bool /* primitive/slice/pointer */
	SetIsCameraIntrinsicMatrixDeliveryEnabled(value bool /* primitive/slice/pointer */)
	IsCameraIntrinsicMatrixDeliverySupported() bool /* primitive/slice/pointer */
	SetIsCameraIntrinsicMatrixDeliverySupported(value bool /* primitive/slice/pointer */)
	IsEnabled() bool /* primitive/slice/pointer */
	SetIsEnabled(value bool /* primitive/slice/pointer */)
	IsVideoFieldModeSupported() bool /* primitive/slice/pointer */
	SetIsVideoFieldModeSupported(value bool /* primitive/slice/pointer */)
	IsVideoMaxFrameDurationSupported() bool /* primitive/slice/pointer */
	SetIsVideoMaxFrameDurationSupported(value bool /* primitive/slice/pointer */)
	IsVideoMinFrameDurationSupported() bool /* primitive/slice/pointer */
	SetIsVideoMinFrameDurationSupported(value bool /* primitive/slice/pointer */)
	IsVideoMirrored() bool /* primitive/slice/pointer */
	SetIsVideoMirrored(value bool /* primitive/slice/pointer */)
	IsVideoMirroringSupported() bool /* primitive/slice/pointer */
	SetIsVideoMirroringSupported(value bool /* primitive/slice/pointer */)
	IsVideoOrientationSupported() bool /* primitive/slice/pointer */
	SetIsVideoOrientationSupported(value bool /* primitive/slice/pointer */)
	IsVideoStabilizationEnabled() bool /* primitive/slice/pointer */
	SetIsVideoStabilizationEnabled(value bool /* primitive/slice/pointer */)
	IsVideoStabilizationSupported() bool /* primitive/slice/pointer */
	SetIsVideoStabilizationSupported(value bool /* primitive/slice/pointer */)
	Output() IAVCaptureOutput
	SetOutput(value IAVCaptureOutput)
	PreferredVideoStabilizationMode() AVCaptureVideoStabilizationMode /* foo */
	SetPreferredVideoStabilizationMode(value AVCaptureVideoStabilizationMode /* foo */)
	VideoFieldMode() AVVideoFieldMode /* foo */
	SetVideoFieldMode(value AVVideoFieldMode /* foo */)
	VideoMaxFrameDuration() CMTime /* foo */
	SetVideoMaxFrameDuration(value CMTime /* foo */)
	VideoMaxScaleAndCropFactor() float64 /* primitive/slice/pointer */
	SetVideoMaxScaleAndCropFactor(value float64 /* primitive/slice/pointer */)
	VideoMinFrameDuration() CMTime /* foo */
	SetVideoMinFrameDuration(value CMTime /* foo */)
	VideoOrientation() AVCaptureVideoOrientation /* foo */
	SetVideoOrientation(value AVCaptureVideoOrientation /* foo */)
	VideoPreviewLayer() IAVCaptureVideoPreviewLayer
	SetVideoPreviewLayer(value IAVCaptureVideoPreviewLayer)
	VideoRotationAngle() float64 /* primitive/slice/pointer */
	SetVideoRotationAngle(value float64 /* primitive/slice/pointer */)
	VideoScaleAndCropFactor() float64 /* primitive/slice/pointer */
	SetVideoScaleAndCropFactor(value float64 /* primitive/slice/pointer */)
	// methods:
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



// The connection’s current stabilization mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/activevideostabilizationmode
func (c_ CaptureConnection) ActiveVideoStabilizationMode() AVCaptureVideoStabilizationMode /* foo */ {
	rv := objc.Send[CaptureVideoStabilizationMode](c_.ID, objc.Sel("activeVideoStabilizationMode"))
	return rv
}


// The connection’s current stabilization mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/activevideostabilizationmode
func (c_ CaptureConnection) SetActiveVideoStabilizationMode(value AVCaptureVideoStabilizationMode /* foo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveVideoStabilizationMode:"), value)
}


// An array of audio channels that the connection provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/audiochannels
func (c_ CaptureConnection) AudioChannels() IAVCaptureAudioChannel {
	rv := objc.Send[CaptureAudioChannel](c_.ID, objc.Sel("audioChannels"))
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
func (c_ CaptureConnection) AutomaticallyAdjustsVideoMirroring() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyAdjustsVideoMirroring"))
	return rv
}


// A Boolean value that indicates whether you can enable mirroring based on a session’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/automaticallyadjustsvideomirroring
func (c_ CaptureConnection) SetAutomaticallyAdjustsVideoMirroring(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallyAdjustsVideoMirroring:"), value)
}


// A Boolean value that indicates whether the system enables video stabilization when it’s available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/enablesvideostabilizationwhenavailable
func (c_ CaptureConnection) EnablesVideoStabilizationWhenAvailable() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("enablesVideoStabilizationWhenAvailable"))
	return rv
}


// A Boolean value that indicates whether the system enables video stabilization when it’s available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/enablesvideostabilizationwhenavailable
func (c_ CaptureConnection) SetEnablesVideoStabilizationWhenAvailable(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnablesVideoStabilizationWhenAvailable:"), value)
}


// An array of the connection’s input ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/inputports
func (c_ CaptureConnection) InputPorts() AVCaptureInputPort /* foo */ {
	rv := objc.Send[CaptureInputPort](c_.ID, objc.Sel("inputPorts"))
	return rv
}


// An array of the connection’s input ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/inputports
func (c_ CaptureConnection) SetInputPorts(value AVCaptureInputPort /* foo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInputPorts:"), value)
}


// Indicates whether the connection is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isactive
func (c_ CaptureConnection) IsActive() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isActive"))
	return rv
}


// Indicates whether the connection is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isactive
func (c_ CaptureConnection) SetIsActive(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsActive:"), value)
}


// A Boolean value that indicates whether the connection can configure the capture pipeline to deliver camera intrinsics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliveryenabled
func (c_ CaptureConnection) IsCameraIntrinsicMatrixDeliveryEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraIntrinsicMatrixDeliveryEnabled"))
	return rv
}


// A Boolean value that indicates whether the connection can configure the capture pipeline to deliver camera intrinsics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliveryenabled
func (c_ CaptureConnection) SetIsCameraIntrinsicMatrixDeliveryEnabled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraIntrinsicMatrixDeliveryEnabled:"), value)
}


// A Boolean value that indicates whether the capture connection currently supports delivering camera intrinsics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliverysupported
func (c_ CaptureConnection) IsCameraIntrinsicMatrixDeliverySupported() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraIntrinsicMatrixDeliverySupported"))
	return rv
}


// A Boolean value that indicates whether the capture connection currently supports delivering camera intrinsics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliverysupported
func (c_ CaptureConnection) SetIsCameraIntrinsicMatrixDeliverySupported(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraIntrinsicMatrixDeliverySupported:"), value)
}


// Turns the connection on and off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isenabled
func (c_ CaptureConnection) IsEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}


// Turns the connection on and off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isenabled
func (c_ CaptureConnection) SetIsEnabled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value that indicates whether the connection supports setting a video field mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideofieldmodesupported
func (c_ CaptureConnection) IsVideoFieldModeSupported() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoFieldModeSupported"))
	return rv
}


// A Boolean value that indicates whether the connection supports setting a video field mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideofieldmodesupported
func (c_ CaptureConnection) SetIsVideoFieldModeSupported(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoFieldModeSupported:"), value)
}


// A Boolean value that indicates whether the connection supports a maximum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomaxframedurationsupported
func (c_ CaptureConnection) IsVideoMaxFrameDurationSupported() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMaxFrameDurationSupported"))
	return rv
}


// A Boolean value that indicates whether the connection supports a maximum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomaxframedurationsupported
func (c_ CaptureConnection) SetIsVideoMaxFrameDurationSupported(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMaxFrameDurationSupported:"), value)
}


// A Boolean value that indicates whether the connection supports a minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideominframedurationsupported
func (c_ CaptureConnection) IsVideoMinFrameDurationSupported() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMinFrameDurationSupported"))
	return rv
}


// A Boolean value that indicates whether the connection supports a minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideominframedurationsupported
func (c_ CaptureConnection) SetIsVideoMinFrameDurationSupported(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMinFrameDurationSupported:"), value)
}


// A Boolean value that indicates whether the connection horizontally flips the video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirrored
func (c_ CaptureConnection) IsVideoMirrored() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMirrored"))
	return rv
}


// A Boolean value that indicates whether the connection horizontally flips the video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirrored
func (c_ CaptureConnection) SetIsVideoMirrored(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMirrored:"), value)
}


// A Boolean value that indicates whether the connection supports video mirroring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirroringsupported
func (c_ CaptureConnection) IsVideoMirroringSupported() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMirroringSupported"))
	return rv
}


// A Boolean value that indicates whether the connection supports video mirroring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirroringsupported
func (c_ CaptureConnection) SetIsVideoMirroringSupported(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMirroringSupported:"), value)
}


// A Boolean value that indicates whether the connection supports changing the orientation of the video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideoorientationsupported
func (c_ CaptureConnection) IsVideoOrientationSupported() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoOrientationSupported"))
	return rv
}


// A Boolean value that indicates whether the connection supports changing the orientation of the video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideoorientationsupported
func (c_ CaptureConnection) SetIsVideoOrientationSupported(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoOrientationSupported:"), value)
}


// A Boolean value that indicates whether video stabilization is active for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationenabled
func (c_ CaptureConnection) IsVideoStabilizationEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoStabilizationEnabled"))
	return rv
}


// A Boolean value that indicates whether video stabilization is active for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationenabled
func (c_ CaptureConnection) SetIsVideoStabilizationEnabled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoStabilizationEnabled:"), value)
}


// A Boolean value that indicates whether this connection supports video stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationsupported
func (c_ CaptureConnection) IsVideoStabilizationSupported() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoStabilizationSupported"))
	return rv
}


// A Boolean value that indicates whether this connection supports video stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationsupported
func (c_ CaptureConnection) SetIsVideoStabilizationSupported(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoStabilizationSupported:"), value)
}


// The connection’s output port, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/output
func (c_ CaptureConnection) Output() IAVCaptureOutput {
	rv := objc.Send[CaptureOutput](c_.ID, objc.Sel("output"))
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
func (c_ CaptureConnection) PreferredVideoStabilizationMode() AVCaptureVideoStabilizationMode /* foo */ {
	rv := objc.Send[CaptureVideoStabilizationMode](c_.ID, objc.Sel("preferredVideoStabilizationMode"))
	return rv
}


// The stabilization mode that’s the most appropriate for a video connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/preferredvideostabilizationmode
func (c_ CaptureConnection) SetPreferredVideoStabilizationMode(value AVCaptureVideoStabilizationMode /* foo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreferredVideoStabilizationMode:"), value)
}


// A setting that tells the connection how to interlace video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videofieldmode
func (c_ CaptureConnection) VideoFieldMode() AVVideoFieldMode /* foo */ {
	rv := objc.Send[VideoFieldMode](c_.ID, objc.Sel("videoFieldMode"))
	return rv
}


// A setting that tells the connection how to interlace video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videofieldmode
func (c_ CaptureConnection) SetVideoFieldMode(value AVVideoFieldMode /* foo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoFieldMode:"), value)
}


// The largest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videomaxframeduration
func (c_ CaptureConnection) VideoMaxFrameDuration() CMTime /* foo */ {
	rv := objc.Send[Time](c_.ID, objc.Sel("videoMaxFrameDuration"))
	return rv
}


// The largest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videomaxframeduration
func (c_ CaptureConnection) SetVideoMaxFrameDuration(value CMTime /* foo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMaxFrameDuration:"), value)
}


// The connection’s maximum video scale and crop factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videomaxscaleandcropfactor
func (c_ CaptureConnection) VideoMaxScaleAndCropFactor() float64 /* primitive/slice/pointer */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMaxScaleAndCropFactor"))
	return rv
}


// The connection’s maximum video scale and crop factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videomaxscaleandcropfactor
func (c_ CaptureConnection) SetVideoMaxScaleAndCropFactor(value float64 /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMaxScaleAndCropFactor:"), value)
}


// The smallest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videominframeduration
func (c_ CaptureConnection) VideoMinFrameDuration() CMTime /* foo */ {
	rv := objc.Send[Time](c_.ID, objc.Sel("videoMinFrameDuration"))
	return rv
}


// The smallest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videominframeduration
func (c_ CaptureConnection) SetVideoMinFrameDuration(value CMTime /* foo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMinFrameDuration:"), value)
}


// An orientation that tells the connection how to rotate a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videoorientation
func (c_ CaptureConnection) VideoOrientation() AVCaptureVideoOrientation /* foo */ {
	rv := objc.Send[CaptureVideoOrientation](c_.ID, objc.Sel("videoOrientation"))
	return rv
}


// An orientation that tells the connection how to rotate a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videoorientation
func (c_ CaptureConnection) SetVideoOrientation(value AVCaptureVideoOrientation /* foo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoOrientation:"), value)
}


// The video preview layer associated with the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videopreviewlayer
func (c_ CaptureConnection) VideoPreviewLayer() IAVCaptureVideoPreviewLayer {
	rv := objc.Send[CaptureVideoPreviewLayer](c_.ID, objc.Sel("videoPreviewLayer"))
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
func (c_ CaptureConnection) VideoRotationAngle() float64 /* primitive/slice/pointer */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoRotationAngle"))
	return rv
}


// A rotation angle the connection applies to a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videorotationangle
func (c_ CaptureConnection) SetVideoRotationAngle(value float64 /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoRotationAngle:"), value)
}


// The current scale and crop factor the video output uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videoscaleandcropfactor
func (c_ CaptureConnection) VideoScaleAndCropFactor() float64 /* primitive/slice/pointer */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoScaleAndCropFactor"))
	return rv
}


// The current scale and crop factor the video output uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/videoscaleandcropfactor
func (c_ CaptureConnection) SetVideoScaleAndCropFactor(value float64 /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoScaleAndCropFactor:"), value)
}



