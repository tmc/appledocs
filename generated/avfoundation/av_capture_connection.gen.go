// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureConnection */


/* debug [class_header]: Header for AVCaptureConnection */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureConnection */
// An interface definition for the [CaptureConnection] class.
type ICaptureConnection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureConnection */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureConnection */
	// methods:
	IsVideoRotationAngleSupported(videoRotationAngle float64) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureConnection */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureConnection */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureConnection */

// Creates a capture connection that represents a connection between an input port and a video preview layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/init(inputPort:videoPreviewLayer:)
func NewCaptureConnectionWithInputPortVideoPreviewLayer(port IAVCaptureInputPort, layer IAVCaptureVideoPreviewLayer) CaptureConnection {
	instance := getCaptureConnectionClass().Alloc()
	rv := objc.Send[CaptureConnection](instance.ID, objc.Sel("initWithInputPort:videoPreviewLayer:"), port, layer)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureConnectionWithInputPortVideoPreviewLayer */


// Creates a capture connection that represents a connection between multiple input ports and an output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/init(inputPorts:output:)
func NewCaptureConnectionWithInputPortsOutput(ports []CaptureInputPort, output IAVCaptureOutput) CaptureConnection {
	instance := getCaptureConnectionClass().Alloc()
	rv := objc.Send[CaptureConnection](instance.ID, objc.Sel("initWithInputPorts:output:"), ports, output)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureConnectionWithInputPortsOutput */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureConnection */

// Returns a capture connection that represents a connection between an input port and a video preview layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/connectionWithInputPort:videoPreviewLayer:
func (cc _CaptureConnectionClass) ConnectionWithInputPortVideoPreviewLayer(port IAVCaptureInputPort, layer IAVCaptureVideoPreviewLayer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("connectionWithInputPort:videoPreviewLayer:"), port, layer)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConnectionWithInputPortVideoPreviewLayer) */


// Returns a capture connection that represents a connection between multiple input ports and an output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/connectionWithInputPorts:output:
func (cc _CaptureConnectionClass) ConnectionWithInputPortsOutput(ports []CaptureInputPort, output IAVCaptureOutput) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("connectionWithInputPorts:output:"), ports, output)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConnectionWithInputPortsOutput) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureConnection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureConnection */

// Returns a Boolean value that indicates whether the connection supports a rotation angle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoRotationAngleSupported(_:)
func (c_ CaptureConnection) IsVideoRotationAngleSupported(videoRotationAngle float64) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoRotationAngleSupported:"), videoRotationAngle)
	return rv
}/* debug [instance_methods/method]: IsVideoRotationAngleSupported */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureConnection */

// An array of audio channels that the connection provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/audioChannels
func (c_ CaptureConnection) AudioChannels() []CaptureAudioChannel {
	rv := objc.Send[[]CaptureAudioChannel](c_.ID, objc.Sel("audioChannels"))
	return rv
}/* debug [instance_properties/getter]: audioChannels */


// A Boolean value that indicates whether you can enable mirroring based on a session’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/automaticallyAdjustsVideoMirroring
func (c_ CaptureConnection) AutomaticallyAdjustsVideoMirroring() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyAdjustsVideoMirroring"))
	return rv
}/* debug [instance_properties/getter]: automaticallyAdjustsVideoMirroring */


// A Boolean value that indicates whether you can enable mirroring based on a session’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/automaticallyAdjustsVideoMirroring
func (c_ CaptureConnection) SetAutomaticallyAdjustsVideoMirroring(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallyAdjustsVideoMirroring:"), value)
}/* debug [instance_properties/setter]: automaticallyAdjustsVideoMirroring */


// An array of the connection’s input ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/inputPorts
func (c_ CaptureConnection) InputPorts() []CaptureInputPort {
	rv := objc.Send[[]CaptureInputPort](c_.ID, objc.Sel("inputPorts"))
	return rv
}/* debug [instance_properties/getter]: inputPorts */


// Indicates whether the connection is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isActive
func (c_ CaptureConnection) Active() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("active"))
	return rv
}/* debug [instance_properties/getter]: active */


// Turns the connection on and off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isEnabled
func (c_ CaptureConnection) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// Turns the connection on and off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isEnabled
func (c_ CaptureConnection) SetEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// A Boolean value that indicates whether the connection supports setting a video field mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoFieldModeSupported
func (c_ CaptureConnection) SupportsVideoFieldMode() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsVideoFieldMode"))
	return rv
}/* debug [instance_properties/getter]: supportsVideoFieldMode */


// A Boolean value that indicates whether the connection supports a maximum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoMaxFrameDurationSupported
func (c_ CaptureConnection) SupportsVideoMaxFrameDuration() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsVideoMaxFrameDuration"))
	return rv
}/* debug [instance_properties/getter]: supportsVideoMaxFrameDuration */


// A Boolean value that indicates whether the connection supports a minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoMinFrameDurationSupported
func (c_ CaptureConnection) SupportsVideoMinFrameDuration() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsVideoMinFrameDuration"))
	return rv
}/* debug [instance_properties/getter]: supportsVideoMinFrameDuration */


// A Boolean value that indicates whether the connection horizontally flips the video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoMirrored
func (c_ CaptureConnection) VideoMirrored() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("videoMirrored"))
	return rv
}/* debug [instance_properties/getter]: videoMirrored */


// A Boolean value that indicates whether the connection horizontally flips the video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoMirrored
func (c_ CaptureConnection) SetVideoMirrored(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMirrored:"), value)
}/* debug [instance_properties/setter]: videoMirrored */


// A Boolean value that indicates whether the connection supports video mirroring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoMirroringSupported
func (c_ CaptureConnection) SupportsVideoMirroring() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsVideoMirroring"))
	return rv
}/* debug [instance_properties/getter]: supportsVideoMirroring */


// A Boolean value that indicates whether the connection supports changing the orientation of the video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoOrientationSupported
func (c_ CaptureConnection) SupportsVideoOrientation() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsVideoOrientation"))
	return rv
}/* debug [instance_properties/getter]: supportsVideoOrientation */


// The connection’s output port, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/output
func (c_ CaptureConnection) Output() IAVCaptureOutput {
	rv := objc.Send[CaptureOutput](c_.ID, objc.Sel("output"))
	return rv
}/* debug [instance_properties/getter]: output */


// A setting that tells the connection how to interlace video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoFieldMode
func (c_ CaptureConnection) VideoFieldMode() VideoFieldMode {
	rv := objc.Send[VideoFieldMode](c_.ID, objc.Sel("videoFieldMode"))
	return rv
}/* debug [instance_properties/getter]: videoFieldMode */


// A setting that tells the connection how to interlace video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoFieldMode
func (c_ CaptureConnection) SetVideoFieldMode(value VideoFieldMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoFieldMode:"), value)
}/* debug [instance_properties/setter]: videoFieldMode */


// The largest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMaxFrameDuration
func (c_ CaptureConnection) VideoMaxFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("videoMaxFrameDuration"))
	return rv
}/* debug [instance_properties/getter]: videoMaxFrameDuration */


// The largest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMaxFrameDuration
func (c_ CaptureConnection) SetVideoMaxFrameDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMaxFrameDuration:"), value)
}/* debug [instance_properties/setter]: videoMaxFrameDuration */


// The smallest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMinFrameDuration
func (c_ CaptureConnection) VideoMinFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("videoMinFrameDuration"))
	return rv
}/* debug [instance_properties/getter]: videoMinFrameDuration */


// The smallest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMinFrameDuration
func (c_ CaptureConnection) SetVideoMinFrameDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMinFrameDuration:"), value)
}/* debug [instance_properties/setter]: videoMinFrameDuration */


// An orientation that tells the connection how to rotate a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoOrientation
func (c_ CaptureConnection) VideoOrientation() CaptureVideoOrientation {
	rv := objc.Send[CaptureVideoOrientation](c_.ID, objc.Sel("videoOrientation"))
	return rv
}/* debug [instance_properties/getter]: videoOrientation */


// An orientation that tells the connection how to rotate a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoOrientation
func (c_ CaptureConnection) SetVideoOrientation(value CaptureVideoOrientation) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoOrientation:"), value)
}/* debug [instance_properties/setter]: videoOrientation */


// The video preview layer associated with the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoPreviewLayer
func (c_ CaptureConnection) VideoPreviewLayer() IAVCaptureVideoPreviewLayer {
	rv := objc.Send[CaptureVideoPreviewLayer](c_.ID, objc.Sel("videoPreviewLayer"))
	return rv
}/* debug [instance_properties/getter]: videoPreviewLayer */


// A rotation angle the connection applies to a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoRotationAngle
func (c_ CaptureConnection) VideoRotationAngle() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoRotationAngle"))
	return rv
}/* debug [instance_properties/getter]: videoRotationAngle */


// A rotation angle the connection applies to a video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoRotationAngle
func (c_ CaptureConnection) SetVideoRotationAngle(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoRotationAngle:"), value)
}/* debug [instance_properties/setter]: videoRotationAngle */


// Indicates whether the connection is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isactive
func (c_ CaptureConnection) IsActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isActive"))
	return rv
}/* debug [instance_properties/getter]: isActive */


// Indicates whether the connection is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isactive
func (c_ CaptureConnection) SetIsActive(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsActive:"), value)
}/* debug [instance_properties/setter]: isActive */


// A Boolean value that indicates whether the connection can configure the capture pipeline to deliver camera intrinsics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliveryenabled
func (c_ CaptureConnection) IsCameraIntrinsicMatrixDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraIntrinsicMatrixDeliveryEnabled"))
	return rv
}/* debug [instance_properties/getter]: isCameraIntrinsicMatrixDeliveryEnabled */


// A Boolean value that indicates whether the connection can configure the capture pipeline to deliver camera intrinsics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliveryenabled
func (c_ CaptureConnection) SetIsCameraIntrinsicMatrixDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraIntrinsicMatrixDeliveryEnabled:"), value)
}/* debug [instance_properties/setter]: isCameraIntrinsicMatrixDeliveryEnabled */


// A Boolean value that indicates whether the capture connection currently supports delivering camera intrinsics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliverysupported
func (c_ CaptureConnection) IsCameraIntrinsicMatrixDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraIntrinsicMatrixDeliverySupported"))
	return rv
}/* debug [instance_properties/getter]: isCameraIntrinsicMatrixDeliverySupported */


// A Boolean value that indicates whether the capture connection currently supports delivering camera intrinsics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/iscameraintrinsicmatrixdeliverysupported
func (c_ CaptureConnection) SetIsCameraIntrinsicMatrixDeliverySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraIntrinsicMatrixDeliverySupported:"), value)
}/* debug [instance_properties/setter]: isCameraIntrinsicMatrixDeliverySupported */


// Turns the connection on and off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isenabled
func (c_ CaptureConnection) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// Turns the connection on and off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isenabled
func (c_ CaptureConnection) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */


// A Boolean value that indicates whether the connection supports setting a video field mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideofieldmodesupported
func (c_ CaptureConnection) IsVideoFieldModeSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoFieldModeSupported"))
	return rv
}/* debug [instance_properties/getter]: isVideoFieldModeSupported */


// A Boolean value that indicates whether the connection supports setting a video field mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideofieldmodesupported
func (c_ CaptureConnection) SetIsVideoFieldModeSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoFieldModeSupported:"), value)
}/* debug [instance_properties/setter]: isVideoFieldModeSupported */


// A Boolean value that indicates whether the connection supports a maximum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomaxframedurationsupported
func (c_ CaptureConnection) IsVideoMaxFrameDurationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMaxFrameDurationSupported"))
	return rv
}/* debug [instance_properties/getter]: isVideoMaxFrameDurationSupported */


// A Boolean value that indicates whether the connection supports a maximum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomaxframedurationsupported
func (c_ CaptureConnection) SetIsVideoMaxFrameDurationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMaxFrameDurationSupported:"), value)
}/* debug [instance_properties/setter]: isVideoMaxFrameDurationSupported */


// A Boolean value that indicates whether the connection supports a minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideominframedurationsupported
func (c_ CaptureConnection) IsVideoMinFrameDurationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMinFrameDurationSupported"))
	return rv
}/* debug [instance_properties/getter]: isVideoMinFrameDurationSupported */


// A Boolean value that indicates whether the connection supports a minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideominframedurationsupported
func (c_ CaptureConnection) SetIsVideoMinFrameDurationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMinFrameDurationSupported:"), value)
}/* debug [instance_properties/setter]: isVideoMinFrameDurationSupported */


// A Boolean value that indicates whether the connection horizontally flips the video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirrored
func (c_ CaptureConnection) IsVideoMirrored() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMirrored"))
	return rv
}/* debug [instance_properties/getter]: isVideoMirrored */


// A Boolean value that indicates whether the connection horizontally flips the video flowing through it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirrored
func (c_ CaptureConnection) SetIsVideoMirrored(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMirrored:"), value)
}/* debug [instance_properties/setter]: isVideoMirrored */


// A Boolean value that indicates whether the connection supports video mirroring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirroringsupported
func (c_ CaptureConnection) IsVideoMirroringSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoMirroringSupported"))
	return rv
}/* debug [instance_properties/getter]: isVideoMirroringSupported */


// A Boolean value that indicates whether the connection supports video mirroring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideomirroringsupported
func (c_ CaptureConnection) SetIsVideoMirroringSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoMirroringSupported:"), value)
}/* debug [instance_properties/setter]: isVideoMirroringSupported */


// A Boolean value that indicates whether the connection supports changing the orientation of the video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideoorientationsupported
func (c_ CaptureConnection) IsVideoOrientationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoOrientationSupported"))
	return rv
}/* debug [instance_properties/getter]: isVideoOrientationSupported */


// A Boolean value that indicates whether the connection supports changing the orientation of the video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideoorientationsupported
func (c_ CaptureConnection) SetIsVideoOrientationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoOrientationSupported:"), value)
}/* debug [instance_properties/setter]: isVideoOrientationSupported */


// A Boolean value that indicates whether video stabilization is active for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationenabled
func (c_ CaptureConnection) IsVideoStabilizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoStabilizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: isVideoStabilizationEnabled */


// A Boolean value that indicates whether video stabilization is active for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationenabled
func (c_ CaptureConnection) SetIsVideoStabilizationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoStabilizationEnabled:"), value)
}/* debug [instance_properties/setter]: isVideoStabilizationEnabled */


// A Boolean value that indicates whether this connection supports video stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationsupported
func (c_ CaptureConnection) IsVideoStabilizationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoStabilizationSupported"))
	return rv
}/* debug [instance_properties/getter]: isVideoStabilizationSupported */


// A Boolean value that indicates whether this connection supports video stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/isvideostabilizationsupported
func (c_ CaptureConnection) SetIsVideoStabilizationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoStabilizationSupported:"), value)
}/* debug [instance_properties/setter]: isVideoStabilizationSupported */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureConnection */


