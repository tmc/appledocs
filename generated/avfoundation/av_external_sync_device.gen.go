// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ExternalSyncDevice] class.
var (
	ExternalSyncDeviceClass     _ExternalSyncDeviceClass
	ExternalSyncDeviceClassOnce sync.Once
)

func getExternalSyncDeviceClass() _ExternalSyncDeviceClass {
	ExternalSyncDeviceClassOnce.Do(func() {
		ExternalSyncDeviceClass = _ExternalSyncDeviceClass{objc.GetClass("AVExternalSyncDevice")}
	})
	return ExternalSyncDeviceClass
}

type _ExternalSyncDeviceClass struct {
	class objc.Class
}





// An interface definition for the [ExternalSyncDevice] class.
type IExternalSyncDevice interface {
	objectivec.IObject
	

	// properties:
	Clock() ClockRef /* not a class type */
	ProductID() objectivec.IObject
	SignalCompensationDelay() objc.IObject /* cross-framework: Time */
	SetSignalCompensationDelay(value objc.IObject /* cross-framework: Time */)
	Status() ExternalSyncDeviceStatus
	Uuid() foundation.UUID
	VendorID() objectivec.IObject


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ec _ExternalSyncDeviceClass) Alloc() ExternalSyncDevice {
	rv := objc.Send[ExternalSyncDevice](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _ExternalSyncDeviceClass) New() ExternalSyncDevice {
	rv := objc.Send[ExternalSyncDevice](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExternalSyncDevice) Init() ExternalSyncDevice {
	rv := objc.Send[ExternalSyncDevice](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExternalSyncDevice) Autorelease() ExternalSyncDevice {
	rv := objc.Send[ExternalSyncDevice](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExternalSyncDevice creates a new ExternalSyncDevice instance.
func NewExternalSyncDevice() ExternalSyncDevice {
	return getExternalSyncDeviceClass().New()
}





// An external sync device connected to a host device that can be used to drive the timing of an internal component, such as a camera sensor.
//
// Each instance of corresponds to a physical external device that can drive an internal component, like a camera readout. You cannot create instances of . Instead, you obtain an array of all currently available external sync devices using .


// An external sync device connected to a host device that can be used to drive the timing of an internal component, such as a camera sensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice
type ExternalSyncDevice struct {
	objectivec.Object
}

// ExternalSyncDeviceFrom constructs a [ExternalSyncDevice] from an unsafe.Pointer.
//
// An external sync device connected to a host device that can be used to drive the timing of an internal component, such as a camera sensor.
func ExternalSyncDeviceFrom(ptr unsafe.Pointer) ExternalSyncDevice {
	return ExternalSyncDevice{objectivec.Object{objc.ID(ptr)}}
}

























// A clock representing the source of time from the external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/clock
func (e_ ExternalSyncDevice) Clock() ClockRef /* not a class type */ {
	rv := objc.Send[ClockRef](e_.ID, objc.Sel("clock"))
	return rv
}


// The USB product identifier associated with the external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/productID
func (e_ ExternalSyncDevice) ProductID() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("productID"))
	return rv
}


// Delay to wait before starting the frame capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/signalCompensationDelay
func (e_ ExternalSyncDevice) SignalCompensationDelay() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](e_.ID, objc.Sel("signalCompensationDelay"))
	return rv
}


// Delay to wait before starting the frame capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/signalCompensationDelay
func (e_ ExternalSyncDevice) SetSignalCompensationDelay(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSignalCompensationDelay:"), value)
}


// The status of the externally connected device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/status
func (e_ ExternalSyncDevice) Status() ExternalSyncDeviceStatus {
	rv := objc.Send[ExternalSyncDeviceStatus](e_.ID, objc.Sel("status"))
	return rv
}


// A unique identifier for an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/uuid
func (e_ ExternalSyncDevice) Uuid() foundation.UUID {
	rv := objc.Send[foundation.UUID](e_.ID, objc.Sel("uuid"))
	return rv
}


// The USB vendor identifier associated with the external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/vendorID
func (e_ ExternalSyncDevice) VendorID() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("vendorID"))
	return rv
}








