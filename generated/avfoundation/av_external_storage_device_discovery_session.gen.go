// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ExternalStorageDeviceDiscoverySession] class.
var (
	ExternalStorageDeviceDiscoverySessionClass     _ExternalStorageDeviceDiscoverySessionClass
	ExternalStorageDeviceDiscoverySessionClassOnce sync.Once
)

func getExternalStorageDeviceDiscoverySessionClass() _ExternalStorageDeviceDiscoverySessionClass {
	ExternalStorageDeviceDiscoverySessionClassOnce.Do(func() {
		ExternalStorageDeviceDiscoverySessionClass = _ExternalStorageDeviceDiscoverySessionClass{objc.GetClass("AVExternalStorageDeviceDiscoverySession")}
	})
	return ExternalStorageDeviceDiscoverySessionClass
}

type _ExternalStorageDeviceDiscoverySessionClass struct {
	class objc.Class
}





// An interface definition for the [ExternalStorageDeviceDiscoverySession] class.
type IExternalStorageDeviceDiscoverySession interface {
	objectivec.IObject
	

	// properties:
	ExternalStorageDevices() []ExternalStorageDevice


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ec _ExternalStorageDeviceDiscoverySessionClass) Alloc() ExternalStorageDeviceDiscoverySession {
	rv := objc.Send[ExternalStorageDeviceDiscoverySession](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _ExternalStorageDeviceDiscoverySessionClass) New() ExternalStorageDeviceDiscoverySession {
	rv := objc.Send[ExternalStorageDeviceDiscoverySession](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExternalStorageDeviceDiscoverySession) Init() ExternalStorageDeviceDiscoverySession {
	rv := objc.Send[ExternalStorageDeviceDiscoverySession](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExternalStorageDeviceDiscoverySession) Autorelease() ExternalStorageDeviceDiscoverySession {
	rv := objc.Send[ExternalStorageDeviceDiscoverySession](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExternalStorageDeviceDiscoverySession creates a new ExternalStorageDeviceDiscoverySession instance.
func NewExternalStorageDeviceDiscoverySession() ExternalStorageDeviceDiscoverySession {
	return getExternalStorageDeviceDiscoverySessionClass().New()
}





// Informs your app when the external storage devices connect to and disconnect from the system.


// Informs your app when the external storage devices connect to and disconnect from the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession
type ExternalStorageDeviceDiscoverySession struct {
	objectivec.Object
}

// ExternalStorageDeviceDiscoverySessionFrom constructs a [ExternalStorageDeviceDiscoverySession] from an unsafe.Pointer.
//
// Informs your app when the external storage devices connect to and disconnect from the system.
func ExternalStorageDeviceDiscoverySessionFrom(ptr unsafe.Pointer) ExternalStorageDeviceDiscoverySession {
	return ExternalStorageDeviceDiscoverySession{objectivec.Object{objc.ID(ptr)}}
}















// A Boolean value that indicates whether the system supports external storage devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession/isSupported
func (ec _ExternalStorageDeviceDiscoverySessionClass) Supported() bool {
	rv := objc.Send[bool](objc.ID(ec.class), objc.Sel("supported"))
	return rv
}

// The system’s singleton device discovery session instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession/shared
func (ec _ExternalStorageDeviceDiscoverySessionClass) SharedSession() ExternalStorageDeviceDiscoverySession {
	rv := objc.Send[ExternalStorageDeviceDiscoverySession](objc.ID(ec.class), objc.Sel("sharedSession"))
	return rv
}











// An array of external storage devices the session updates as individual devices connect or disconnect from the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession/externalStorageDevices
func (e_ ExternalStorageDeviceDiscoverySession) ExternalStorageDevices() []ExternalStorageDevice {
	rv := objc.Send[[]ExternalStorageDevice](e_.ID, objc.Sel("externalStorageDevices"))
	return rv
}


// A Boolean value that indicates whether the system supports external storage devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession/isSupported
func (e_ ExternalStorageDeviceDiscoverySession) Supported() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("supported"))
	return rv
}


// The system’s singleton device discovery session instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession/shared
func (e_ ExternalStorageDeviceDiscoverySession) SharedSession() IAVExternalStorageDeviceDiscoverySession {
	rv := objc.Send[ExternalStorageDeviceDiscoverySession](e_.ID, objc.Sel("sharedSession"))
	return rv
}








