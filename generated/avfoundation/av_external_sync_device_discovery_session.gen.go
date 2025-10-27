// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ExternalSyncDeviceDiscoverySession] class.
var (
	ExternalSyncDeviceDiscoverySessionClass     _ExternalSyncDeviceDiscoverySessionClass
	ExternalSyncDeviceDiscoverySessionClassOnce sync.Once
)

func getExternalSyncDeviceDiscoverySessionClass() _ExternalSyncDeviceDiscoverySessionClass {
	ExternalSyncDeviceDiscoverySessionClassOnce.Do(func() {
		ExternalSyncDeviceDiscoverySessionClass = _ExternalSyncDeviceDiscoverySessionClass{objc.GetClass("AVExternalSyncDeviceDiscoverySession")}
	})
	return ExternalSyncDeviceDiscoverySessionClass
}

type _ExternalSyncDeviceDiscoverySessionClass struct {
	class objc.Class
}





// An interface definition for the [ExternalSyncDeviceDiscoverySession] class.
type IExternalSyncDeviceDiscoverySession interface {
	objectivec.IObject
	

	// properties:
	Devices() []ExternalSyncDevice


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ec _ExternalSyncDeviceDiscoverySessionClass) Alloc() ExternalSyncDeviceDiscoverySession {
	rv := objc.Send[ExternalSyncDeviceDiscoverySession](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _ExternalSyncDeviceDiscoverySessionClass) New() ExternalSyncDeviceDiscoverySession {
	rv := objc.Send[ExternalSyncDeviceDiscoverySession](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExternalSyncDeviceDiscoverySession) Init() ExternalSyncDeviceDiscoverySession {
	rv := objc.Send[ExternalSyncDeviceDiscoverySession](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExternalSyncDeviceDiscoverySession) Autorelease() ExternalSyncDeviceDiscoverySession {
	rv := objc.Send[ExternalSyncDeviceDiscoverySession](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExternalSyncDeviceDiscoverySession creates a new ExternalSyncDeviceDiscoverySession instance.
func NewExternalSyncDeviceDiscoverySession() ExternalSyncDeviceDiscoverySession {
	return getExternalSyncDeviceDiscoverySessionClass().New()
}





// A means of discovering and monitoring connection / disconnection of external sync devices to the host.
//
// is a singleton that lists the external sync devices connected to the host. The client is expected to key-value observe the property for changes to the external sync devices list.


// A means of discovering and monitoring connection / disconnection of external sync devices to the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/DiscoverySession
type ExternalSyncDeviceDiscoverySession struct {
	objectivec.Object
}

// ExternalSyncDeviceDiscoverySessionFrom constructs a [ExternalSyncDeviceDiscoverySession] from an unsafe.Pointer.
//
// A means of discovering and monitoring connection / disconnection of external sync devices to the host.
func ExternalSyncDeviceDiscoverySessionFrom(ptr unsafe.Pointer) ExternalSyncDeviceDiscoverySession {
	return ExternalSyncDeviceDiscoverySession{objectivec.Object{objc.ID(ptr)}}
}















// Whether external sync devices are supported by this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/DiscoverySession/isSupported
func (ec _ExternalSyncDeviceDiscoverySessionClass) Supported() bool {
	rv := objc.Send[bool](objc.ID(ec.class), objc.Sel("supported"))
	return rv
}

// The singleton instance of the external sync source device discovery session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/DiscoverySession/shared
func (ec _ExternalSyncDeviceDiscoverySessionClass) SharedSession() ExternalSyncDeviceDiscoverySession {
	rv := objc.Send[ExternalSyncDeviceDiscoverySession](objc.ID(ec.class), objc.Sel("sharedSession"))
	return rv
}











// An array of external sync devices connected to this host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/DiscoverySession/devices
func (e_ ExternalSyncDeviceDiscoverySession) Devices() []ExternalSyncDevice {
	rv := objc.Send[[]ExternalSyncDevice](e_.ID, objc.Sel("devices"))
	return rv
}


// Whether external sync devices are supported by this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/DiscoverySession/isSupported
func (e_ ExternalSyncDeviceDiscoverySession) Supported() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("supported"))
	return rv
}


// The singleton instance of the external sync source device discovery session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/DiscoverySession/shared
func (e_ ExternalSyncDeviceDiscoverySession) SharedSession() IAVExternalSyncDeviceDiscoverySession {
	rv := objc.Send[ExternalSyncDeviceDiscoverySession](e_.ID, objc.Sel("sharedSession"))
	return rv
}








