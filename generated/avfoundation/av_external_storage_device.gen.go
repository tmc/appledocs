// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ExternalStorageDevice] class.
var (
	ExternalStorageDeviceClass     _ExternalStorageDeviceClass
	ExternalStorageDeviceClassOnce sync.Once
)

func getExternalStorageDeviceClass() _ExternalStorageDeviceClass {
	ExternalStorageDeviceClassOnce.Do(func() {
		ExternalStorageDeviceClass = _ExternalStorageDeviceClass{objc.GetClass("AVExternalStorageDevice")}
	})
	return ExternalStorageDeviceClass
}

type _ExternalStorageDeviceClass struct {
	class objc.Class
}

// An interface definition for the [ExternalStorageDevice] class.
type IExternalStorageDevice interface {
	objectivec.IObject
	DisplayName() string
	SetDisplayName(value string)
	FreeSize() int
	SetFreeSize(value int)
	IsConnected() bool
	SetIsConnected(value bool)
	IsNotRecommendedForCaptureUse() bool
	SetIsNotRecommendedForCaptureUse(value bool)
	TotalSize() int
	SetTotalSize(value int)
	Uuid() foundation.UUID
	SetUuid(value foundation.UUID)
	ExternalStorageDevices() IAVExternalStorageDevice
	SetExternalStorageDevices(value IAVExternalStorageDevice)
}

// Represents a physical external storage device that stores media assets.
//
// Each storage device instance corresponds to a physical external storage device where the system can media assets. You can access all of the currently available external storage devices with the object’s property.


// Represents a physical external storage device that stores media assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice
type ExternalStorageDevice struct {
	objectivec.Object
}

// ExternalStorageDeviceFrom constructs a [ExternalStorageDevice] from an unsafe.Pointer.
//
// Represents a physical external storage device that stores media assets.
func ExternalStorageDeviceFrom(ptr unsafe.Pointer) ExternalStorageDevice {
	return ExternalStorageDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _ExternalStorageDeviceClass) Alloc() ExternalStorageDevice {
	rv := objc.Send[ExternalStorageDevice](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _ExternalStorageDeviceClass) New() ExternalStorageDevice {
	rv := objc.Send[ExternalStorageDevice](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExternalStorageDevice) Init() ExternalStorageDevice {
	rv := objc.Send[ExternalStorageDevice](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExternalStorageDevice) Autorelease() ExternalStorageDevice {
	rv := objc.Send[ExternalStorageDevice](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExternalStorageDevice creates a new ExternalStorageDevice instance.
func NewExternalStorageDevice() ExternalStorageDevice {
	return getExternalStorageDeviceClass().New()
}



// The name of an external storage device that’s appropriate for a user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/displayname
func (e_ ExternalStorageDevice) DisplayName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("displayName"))
	return rv
}


// The name of an external storage device that’s appropriate for a user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/displayname
func (e_ ExternalStorageDevice) SetDisplayName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDisplayName:"), objc.String(value))
}


// The amount of free storage space, in bytes, that’s available on the external storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/freesize
func (e_ ExternalStorageDevice) FreeSize() int {
	rv := objc.Send[int](e_.ID, objc.Sel("freeSize"))
	return rv
}


// The amount of free storage space, in bytes, that’s available on the external storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/freesize
func (e_ ExternalStorageDevice) SetFreeSize(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setFreeSize:"), value)
}


// A Boolean value that indicates whether the system has a connection to the external storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/isconnected
func (e_ ExternalStorageDevice) IsConnected() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isConnected"))
	return rv
}


// A Boolean value that indicates whether the system has a connection to the external storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/isconnected
func (e_ ExternalStorageDevice) SetIsConnected(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsConnected:"), value)
}


// A Boolean value that indicates whether the external storage device is suitable for camera capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/isnotrecommendedforcaptureuse
func (e_ ExternalStorageDevice) IsNotRecommendedForCaptureUse() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isNotRecommendedForCaptureUse"))
	return rv
}


// A Boolean value that indicates whether the external storage device is suitable for camera capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/isnotrecommendedforcaptureuse
func (e_ ExternalStorageDevice) SetIsNotRecommendedForCaptureUse(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsNotRecommendedForCaptureUse:"), value)
}


// The total amount of storage space, in bytes, that’s available on the external storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/totalsize
func (e_ ExternalStorageDevice) TotalSize() int {
	rv := objc.Send[int](e_.ID, objc.Sel("totalSize"))
	return rv
}


// The total amount of storage space, in bytes, that’s available on the external storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/totalsize
func (e_ ExternalStorageDevice) SetTotalSize(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTotalSize:"), value)
}


// The external storage device’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/uuid
func (e_ ExternalStorageDevice) Uuid() foundation.UUID {
	rv := objc.Send[foundation.UUID](e_.ID, objc.Sel("uuid"))
	return rv
}


// The external storage device’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevice/uuid
func (e_ ExternalStorageDevice) SetUuid(value foundation.UUID) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUuid:"), value)
}


// An array of external storage devices the session updates as individual devices connect or disconnect from the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevicediscoverysession/externalstoragedevices
func (e_ ExternalStorageDevice) ExternalStorageDevices() IAVExternalStorageDevice {
	rv := objc.Send[AVExternalStorageDevice](e_.ID, objc.Sel("externalStorageDevices"))
	return rv
}


// An array of external storage devices the session updates as individual devices connect or disconnect from the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevicediscoverysession/externalstoragedevices
func (e_ ExternalStorageDevice) SetExternalStorageDevices(value IAVExternalStorageDevice) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExternalStorageDevices:"), value)
}



