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
	ExternalStorageDevices() IAVExternalStorageDevice
	SetExternalStorageDevices(value IAVExternalStorageDevice)
	// methods:
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

// Alloc allocates a new instance without initialization.
func (ec _ExternalStorageDeviceDiscoverySessionClass) Alloc() ExternalStorageDeviceDiscoverySession {
	rv := objc.Send[ExternalStorageDeviceDiscoverySession](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// An array of external storage devices the session updates as individual devices connect or disconnect from the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevicediscoverysession/externalstoragedevices
func (e_ ExternalStorageDeviceDiscoverySession) ExternalStorageDevices() IAVExternalStorageDevice {
	rv := objc.Send[ExternalStorageDevice](e_.ID, objc.Sel("externalStorageDevices"))
	return rv
}


// An array of external storage devices the session updates as individual devices connect or disconnect from the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevicediscoverysession/externalstoragedevices
func (e_ ExternalStorageDeviceDiscoverySession) SetExternalStorageDevices(value IAVExternalStorageDevice) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExternalStorageDevices:"), value)
}



