// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// Represents a physical external storage device that stores media assets.
//
// Each storage device instance corresponds to a physical external storage device where the system can media assets. You can access all of the currently available external storage devices with the object’s property.
//
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


// Requests access to an external storage device on behalf of your app, which can present a dialog to a person on their device’s display.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/requestAccess(completionHandler:)
func (ec _ExternalStorageDeviceClass) RequestAccessWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("requestAccessWithCompletionHandler:"), handler)
}

// A Boolean value that indicates whether the system has a connection to the external storage device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/isConnected
func (e_ ExternalStorageDevice) Connected() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("connected"))
	return rv
}



