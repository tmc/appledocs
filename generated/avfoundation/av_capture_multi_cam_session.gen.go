// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVCaptureMultiCamSession */


/* debug [class_header]: Header for AVCaptureMultiCamSession */
// The class instance for the [CaptureMultiCamSession] class.
var (
	CaptureMultiCamSessionClass     _CaptureMultiCamSessionClass
	CaptureMultiCamSessionClassOnce sync.Once
)

func getCaptureMultiCamSessionClass() _CaptureMultiCamSessionClass {
	CaptureMultiCamSessionClassOnce.Do(func() {
		CaptureMultiCamSessionClass = _CaptureMultiCamSessionClass{objc.GetClass("AVCaptureMultiCamSession")}
	})
	return CaptureMultiCamSessionClass
}

type _CaptureMultiCamSessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureMultiCamSession */
// An interface definition for the [CaptureMultiCamSession] class.
type ICaptureMultiCamSession interface {
	ICaptureSession
	
/* debug [class_interface_properties]: Properties for CaptureMultiCamSession */
	// properties:
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureMultiCamSession */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureMultiCamSession */
// Alloc allocates a new instance without initialization.
func (cc _CaptureMultiCamSessionClass) Alloc() CaptureMultiCamSession {
	rv := objc.Send[CaptureMultiCamSession](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureMultiCamSessionClass) New() CaptureMultiCamSession {
	rv := objc.Send[CaptureMultiCamSession](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureMultiCamSession) Init() CaptureMultiCamSession {
	rv := objc.Send[CaptureMultiCamSession](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureMultiCamSession) Autorelease() CaptureMultiCamSession {
	rv := objc.Send[CaptureMultiCamSession](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureMultiCamSession creates a new CaptureMultiCamSession instance.
func NewCaptureMultiCamSession() CaptureMultiCamSession {
	return getCaptureMultiCamSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureMultiCamSession */
// A capture session that supports simultaneous capture from multiple inputs of the same media type.
//
// The session preset for a multicamera session is always . Set each capture device’s value to the desired quality of service. You can dynamically enable and disable this session’s individual camera inputs without interrupting capture preview. To stop an individual camera, disable all of its connections or connected ports. The camera then stops streaming data to save power and bandwidth. Other inputs that are streaming data through the session are unaffected.


// A capture session that supports simultaneous capture from multiple inputs of the same media type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession
type CaptureMultiCamSession struct {
	CaptureSession
}

// CaptureMultiCamSessionFrom constructs a [CaptureMultiCamSession] from an unsafe.Pointer.
//
// A capture session that supports simultaneous capture from multiple inputs of the same media type.
func CaptureMultiCamSessionFrom(ptr unsafe.Pointer) CaptureMultiCamSession {
	return CaptureMultiCamSession{
		CaptureSession: CaptureSessionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureMultiCamSession *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureMultiCamSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureMultiCamSession */

// A Boolean value that indicates whether this device supports multi-camera sessions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession/isMultiCamSupported
func (cc _CaptureMultiCamSessionClass) MultiCamSupported() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("multiCamSupported"))
	return rv
}/* debug [class_properties_class/property]: multiCamSupported */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureMultiCamSession */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureMultiCamSession */

// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureMultiCamSession) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("activeFormat"))
	return rv
}/* debug [instance_properties/getter]: activeFormat */


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureMultiCamSession) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveFormat:"), value)
}/* debug [instance_properties/setter]: activeFormat */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureMultiCamSession */


