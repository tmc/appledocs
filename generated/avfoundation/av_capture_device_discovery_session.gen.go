// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CaptureDeviceDiscoverySession] class.
type ICaptureDeviceDiscoverySession interface {
	objectivec.IObject
	

	// properties:
	Devices() []CaptureDevice


	

	// methods:


}





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






// Creates a discovery session that finds devices that match the specified criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession/init(deviceTypes:mediaType:position:)
func NewCaptureDeviceDiscoverySessionWithDeviceTypesMediaTypePosition(deviceTypes []string, mediaType MediaType, position CaptureDevicePosition) CaptureDeviceDiscoverySession {
	rv := objc.Send[CaptureDeviceDiscoverySession](objc.ID(getCaptureDeviceDiscoverySessionClass().class), objc.Sel("discoverySessionWithDeviceTypes:mediaType:position:"), deviceTypes, mediaType, position)
	return rv
}







// Creates a discovery session that finds devices that match the specified criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession/init(deviceTypes:mediaType:position:)
func (cc _CaptureDeviceDiscoverySessionClass) DiscoverySessionWithDeviceTypesMediaTypePosition(deviceTypes []string, mediaType MediaType, position CaptureDevicePosition) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("discoverySessionWithDeviceTypes:mediaType:position:"), deviceTypes, mediaType, position)
	return rv
}

















// A list of devices that match the search criteria of the discovery session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession/devices
func (c_ CaptureDeviceDiscoverySession) Devices() []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("devices"))
	return rv
}







