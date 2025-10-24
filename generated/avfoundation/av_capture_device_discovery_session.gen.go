// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureDeviceDiscoverySession */


/* debug [class_header]: Header for AVCaptureDeviceDiscoverySession */
// The class instance for the [CaptureDeviceDiscoverySession] class.
var (
	CaptureDeviceDiscoverySessionClass     _CaptureDeviceDiscoverySessionClass
	CaptureDeviceDiscoverySessionClassOnce sync.Once
)

func getCaptureDeviceDiscoverySessionClass() _CaptureDeviceDiscoverySessionClass {
	CaptureDeviceDiscoverySessionClassOnce.Do(func() {
		CaptureDeviceDiscoverySessionClass = _CaptureDeviceDiscoverySessionClass{objc.GetClass("AVCaptureDeviceDiscoverySession")}
	})
	return CaptureDeviceDiscoverySessionClass
}

type _CaptureDeviceDiscoverySessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureDeviceDiscoverySession */
// An interface definition for the [CaptureDeviceDiscoverySession] class.
type ICaptureDeviceDiscoverySession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureDeviceDiscoverySession */
	// properties:
	Devices() []CaptureDevice
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureDeviceDiscoverySession */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureDeviceDiscoverySession */
// Alloc allocates a new instance without initialization.
func (cc _CaptureDeviceDiscoverySessionClass) Alloc() CaptureDeviceDiscoverySession {
	rv := objc.Send[CaptureDeviceDiscoverySession](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureDeviceDiscoverySessionClass) New() CaptureDeviceDiscoverySession {
	rv := objc.Send[CaptureDeviceDiscoverySession](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDeviceDiscoverySession) Init() CaptureDeviceDiscoverySession {
	rv := objc.Send[CaptureDeviceDiscoverySession](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDeviceDiscoverySession) Autorelease() CaptureDeviceDiscoverySession {
	rv := objc.Send[CaptureDeviceDiscoverySession](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDeviceDiscoverySession creates a new CaptureDeviceDiscoverySession instance.
func NewCaptureDeviceDiscoverySession() CaptureDeviceDiscoverySession {
	return getCaptureDeviceDiscoverySessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureDeviceDiscoverySession */
// An object that finds capture devices that match specific search criteria.
//
// After creating a device discovery session, query its property to find a device to use for capture. You can also key-value observe this property to monitor changes to the list of available devices.


// An object that finds capture devices that match specific search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession
type CaptureDeviceDiscoverySession struct {
	objectivec.Object
}

// CaptureDeviceDiscoverySessionFrom constructs a [CaptureDeviceDiscoverySession] from an unsafe.Pointer.
//
// An object that finds capture devices that match specific search criteria.
func CaptureDeviceDiscoverySessionFrom(ptr unsafe.Pointer) CaptureDeviceDiscoverySession {
	return CaptureDeviceDiscoverySession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureDeviceDiscoverySession */

// Creates a discovery session that finds devices that match the specified criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession/init(deviceTypes:mediaType:position:)
func NewCaptureDeviceDiscoverySessionWithDeviceTypesMediaTypePosition(deviceTypes []string, mediaType MediaType /* typedef */, position CaptureDevicePosition) CaptureDeviceDiscoverySession {
	rv := objc.Send[CaptureDeviceDiscoverySession](objc.ID(getCaptureDeviceDiscoverySessionClass().class), objc.Sel("discoverySessionWithDeviceTypes:mediaType:position:"), deviceTypes, mediaType, position)
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureDeviceDiscoverySessionWithDeviceTypesMediaTypePosition */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureDeviceDiscoverySession */

// Creates a discovery session that finds devices that match the specified criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession/init(deviceTypes:mediaType:position:)
func (cc _CaptureDeviceDiscoverySessionClass) DiscoverySessionWithDeviceTypesMediaTypePosition(deviceTypes []string, mediaType MediaType /* typedef */, position CaptureDevicePosition) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("discoverySessionWithDeviceTypes:mediaType:position:"), deviceTypes, mediaType, position)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DiscoverySessionWithDeviceTypesMediaTypePosition) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureDeviceDiscoverySession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureDeviceDiscoverySession */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureDeviceDiscoverySession */

// A list of devices that match the search criteria of the discovery session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession/devices
func (c_ CaptureDeviceDiscoverySession) Devices() []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("devices"))
	return rv
}/* debug [instance_properties/getter]: devices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureDeviceDiscoverySession */


