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
	AudioChannels() []CaptureAudioChannel
	AutomaticallyAdjustsVideoMirroring() bool
	SetAutomaticallyAdjustsVideoMirroring(value bool)
	InputPorts() []CaptureInputPort
	Active() bool
	Enabled() bool
	SetEnabled(value bool)
	SupportsVideoFieldMode() bool
	SupportsVideoMaxFrameDuration() bool
	SupportsVideoMinFrameDuration() bool
	VideoMirrored() bool
	SetVideoMirrored(value bool)
	SupportsVideoMirroring() bool
	SupportsVideoOrientation() bool
	Output() IAVCaptureOutput
	VideoFieldMode() VideoFieldMode
	SetVideoFieldMode(value VideoFieldMode)
	VideoMaxFrameDuration() objc.IObject /* cross-framework: Time */
	SetVideoMaxFrameDuration(value objc.IObject /* cross-framework: Time */)
	VideoMinFrameDuration() objc.IObject /* cross-framework: Time */
	SetVideoMinFrameDuration(value objc.IObject /* cross-framework: Time */)
	VideoOrientation() CaptureVideoOrientation
	SetVideoOrientation(value CaptureVideoOrientation)
	VideoPreviewLayer() IAVCaptureVideoPreviewLayer
	VideoRotationAngle() float64
	SetVideoRotationAngle(value float64)
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


	

	// methods:
	IsVideoRotationAngleSupported(videoRotationAngle float64) bool


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureConnectionClass) Alloc() CaptureConnection {
	rv := objc.Send[CaptureConnection](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates a capture connection that represents a connection between an input port and a video preview layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/init(inputPort:videoPreviewLayer:)
func NewCaptureConnectionWithInputPortVideoPreviewLayer(port IAVCaptureInputPort, layer IAVCaptureVideoPreviewLayer) CaptureConnection {
	instance := getCaptureConnectionClass().Alloc()
	rv := objc.Send[CaptureConnection](instance.ID, objc.Sel("initWithInputPort:videoPreviewLayer:"), port, layer)
	rv.Autorelease()
	return rv
}


// Creates a capture connection that represents a connection between multiple input ports and an output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/init(inputPorts:output:)
func NewCaptureConnectionWithInputPortsOutput(ports []CaptureInputPort, output IAVCaptureOutput) CaptureConnection {
	instance := getCaptureConnectionClass().Alloc()
	rv := objc.Send[CaptureConnection](instance.ID, objc.Sel("initWithInputPorts:output:"), ports, output)
	rv.Autorelease()
	return rv
}







// Returns a capture connection that represents a connection between an input port and a video preview layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/connectionWithInputPort:videoPreviewLayer:
func (cc _CaptureConnectionClass) ConnectionWithInputPortVideoPreviewLayer(port IAVCaptureInputPort, layer IAVCaptureVideoPreviewLayer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("connectionWithInputPort:videoPreviewLayer:"), port, layer)
	return rv
}


// Returns a capture connection that represents a connection between multiple input ports and an output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/connectionWithInputPorts:output:
func (cc _CaptureConnectionClass) ConnectionWithInputPortsOutput(ports []CaptureInputPort, output IAVCaptureOutput) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("connectionWithInputPorts:output:"), ports, output)
	return rv
}












// Returns a Boolean value that indicates whether the connection supports a rotation angle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoRotationAngleSupported(_:)
func (c_ CaptureConnection) IsVideoRotationAngleSupported(videoRotationAngle float64) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoRotationAngleSupported:"), videoRotationAngle)
	return rv
}







// An array of audio channels that the connection provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/audioChannels
func (c_ CaptureConnection) AudioChannels() []CaptureAudioChannel {
	rv := objc.Send[[]CaptureAudioChannel](c_.ID, objc.Sel("audioChannels"))
	return rv
}


// A Boolean value that indicates whether you can enable mirroring based on a session’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/automaticallyAdjustsVideoMirroring
func (c_ CaptureConnection) AutomaticallyAdjustsVideoMirroring() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyAdjustsVideoMirroring"))
	return rv
}


// A Boolean value that indicates whether you can enable mirroring based on a session’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/automaticallyAdjustsVideoMirroring
func (c_ CaptureConnection) SetAutomaticallyAdjustsVideoMirroring(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallyAdjustsVideoMirroring:"), value)
}


// An array of the connection’s input ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/inputPorts
func (c_ CaptureConnection) InputPorts() []CaptureInputPort {
	rv := objc.Send[[]CaptureInputPort](c_.ID, objc.Sel("inputPorts"))
	return rv
}


// Indicates whether the connection is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isActive
func (c_ CaptureConnection) Active() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("active"))
	return rv
}


// Turns the connection on and off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isEnabled
func (c_ CaptureConnection) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}


// Turns the connection on and off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isEnabled
func (c_ CaptureConnection) SetEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabled:"), value)
}


// A Boolean value that indicates whether the connection supports setting a video field mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoFieldModeSupported
func (c_ CaptureConnection) SupportsVideoFieldMode() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsVideoFieldMode"))
	return rv
}


// A Boolean value that indicates whether the connection supports a maximum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoMaxFrameDurationSupported
func (c_ CaptureConnection) SupportsVideoMaxFrameDuration() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsVideoMaxFrameDuration"))
	return rv
}


// A Boolean value that indicates whether the connection supports a minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoMinFrameDurationSupported
func (c_ CaptureConnection) SupportsVideoMinFrameDuration() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsVideoMinFrameDuration"))
	return rv
}


// A Boolean value that indicates whether the connection horizontally flips the video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoMirrored
func (c_ CaptureConnection) VideoMirrored() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("videoMirrored"))
	return rv
}


// A Boolean value that indicates whether the connection horizontally flips the video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoMirrored
func (c_ CaptureConnection) SetVideoMirrored(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMirrored:"), value)
}


// A Boolean value that indicates whether the connection supports video mirroring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoMirroringSupported
func (c_ CaptureConnection) SupportsVideoMirroring() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsVideoMirroring"))
	return rv
}


// A Boolean value that indicates whether the connection supports changing the orientation of the video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoOrientationSupported
func (c_ CaptureConnection) SupportsVideoOrientation() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsVideoOrientation"))
	return rv
}


// The connection’s output port, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/output
func (c_ CaptureConnection) Output() IAVCaptureOutput {
	rv := objc.Send[CaptureOutput](c_.ID, objc.Sel("output"))
	return rv
}


// A setting that tells the connection how to interlace video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoFieldMode
func (c_ CaptureConnection) VideoFieldMode() VideoFieldMode {
	rv := objc.Send[VideoFieldMode](c_.ID, objc.Sel("videoFieldMode"))
	return rv
}


// A setting that tells the connection how to interlace video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoFieldMode
func (c_ CaptureConnection) SetVideoFieldMode(value VideoFieldMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoFieldMode:"), value)
}


// The largest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMaxFrameDuration
func (c_ CaptureConnection) VideoMaxFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("videoMaxFrameDuration"))
	return rv
}


// The largest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMaxFrameDuration
func (c_ CaptureConnection) SetVideoMaxFrameDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMaxFrameDuration:"), value)
}


// The smallest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMinFrameDuration
func (c_ CaptureConnection) VideoMinFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("videoMinFrameDuration"))
	return rv
}


// The smallest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMinFrameDuration
func (c_ CaptureConnection) SetVideoMinFrameDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMinFrameDuration:"), value)
}


// An orientation that tells the connection how to rotate a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoOrientation
func (c_ CaptureConnection) VideoOrientation() CaptureVideoOrientation {
	rv := objc.Send[CaptureVideoOrientation](c_.ID, objc.Sel("videoOrientation"))
	return rv
}


// An orientation that tells the connection how to rotate a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoOrientation
func (c_ CaptureConnection) SetVideoOrientation(value CaptureVideoOrientation) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoOrientation:"), value)
}


// The video preview layer associated with the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoPreviewLayer
func (c_ CaptureConnection) VideoPreviewLayer() IAVCaptureVideoPreviewLayer {
	rv := objc.Send[CaptureVideoPreviewLayer](c_.ID, objc.Sel("videoPreviewLayer"))
	return rv
}


// A rotation angle the connection applies to a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoRotationAngle
func (c_ CaptureConnection) VideoRotationAngle() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoRotationAngle"))
	return rv
}


// A rotation angle the connection applies to a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoRotationAngle
func (c_ CaptureConnection) SetVideoRotationAngle(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoRotationAngle:"), value)
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







