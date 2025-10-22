// Code generated from Apple documentation for DeviceCheck. DO NOT EDIT.

package devicecheck

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DCDevice] class.
var (
	DCDeviceClass     _DCDeviceClass
	DCDeviceClassOnce sync.Once
)

func getDCDeviceClass() _DCDeviceClass {
	DCDeviceClassOnce.Do(func() {
		DCDeviceClass = _DCDeviceClass{objc.GetClass("DCDevice")}
	})
	return DCDeviceClass
}

type _DCDeviceClass struct {
	class objc.Class
}

// An interface definition for the [DCDevice] class.
type IDCDevice interface {
	objectivec.IObject
	GenerateTokenWithCompletionHandler(completion unsafe.Pointer)
	Supported() bool
	IsSupported() bool
	SetIsSupported(value bool)
}

// A representation of a device that provides a unique, authenticated token.
//
// Use the shared instance of the class to generate a token that identifies a device. Call the method to get the token, and then send it to your server: On your server, combine the token with an authentication key that you obtain from Apple, and use the result to request access to two per-device binary digits (bits). After authenticating the device, Apple passes the current values of the bits, along with the date they were last modified, to your server. Your server applies its business logic to this information and communicates the results to your app. For more information about server-side procedures, see . Apple records the bits for you, and reports the bits back to you, but you’re responsible for keeping track of what the bits mean. You’re also responsible for determining when to reset the bits for a given device; for example, when a user sells the device to someone else.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCDevice
type DCDevice struct {
	objectivec.Object
}

// DCDeviceFrom constructs a [DCDevice] from an unsafe.Pointer.
//
// A representation of a device that provides a unique, authenticated token.
func DCDeviceFrom(ptr unsafe.Pointer) DCDevice {
	return DCDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DCDeviceClass) Alloc() DCDevice {
	rv := objc.Send[DCDevice](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DCDeviceClass) New() DCDevice {
	rv := objc.Send[DCDevice](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DCDevice) Init() DCDevice {
	rv := objc.Send[DCDevice](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DCDevice) Autorelease() DCDevice {
	rv := objc.Send[DCDevice](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDCDevice creates a new DCDevice instance.
func NewDCDevice() DCDevice {
	return getDCDeviceClass().New()
}


// A representation of the device for which you want to query the two bits of data.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCDevice/current
func (dc _DCDeviceClass) CurrentDevice() DCDevice {
	rv := objc.Send[DCDevice](objc.ID(dc.class), objc.Sel("currentDevice"))
	return rv
}
// Generates a token that identifies the current device.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCDevice/generateToken(completionHandler:)
func (d_ DCDevice) GenerateTokenWithCompletionHandler(completion unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("generateTokenWithCompletionHandler:"), completion)
}

// A representation of the device for which you want to query the two bits of data.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCDevice/current
func (d_ DCDevice) CurrentDevice() DCDevice {
	rv := objc.Send[DCDevice](d_.ID, objc.Sel("currentDevice"))
	return rv
}

// A Boolean value that indicates whether the device supports the DeviceCheck API.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCDevice/isSupported
func (d_ DCDevice) Supported() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("supported"))
	return rv
}

// A Boolean value that indicates whether the device supports the DeviceCheck
//
// [Full Topic]: https://developer.apple.com/documentation/devicecheck/dcdevice/issupported
func (d_ DCDevice) IsSupported() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isSupported"))
	return rv
}


// SetIsSupported sets the value of the isSupported property.
// A Boolean value that indicates whether the device supports the DeviceCheck

//
// [Full Topic]: https://developer.apple.com/documentation/devicecheck/dcdevice/issupported
func (d_ DCDevice) SetIsSupported(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsSupported:"), value)
}




