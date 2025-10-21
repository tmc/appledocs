// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDevice] class.
var (
	MTRDeviceClass     _MTRDeviceClass
	MTRDeviceClassOnce sync.Once
)

func getMTRDeviceClass() _MTRDeviceClass {
	MTRDeviceClassOnce.Do(func() {
		MTRDeviceClass = _MTRDeviceClass{objc.GetClass("MTRDevice")}
	})
	return MTRDeviceClass
}

type _MTRDeviceClass struct {
	class objc.Class
}

// An interface definition for the [MTRDevice] class.
type IMTRDevice interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice
type MTRDevice struct {
	objectivec.Object
}

// MTRDeviceFrom constructs a [MTRDevice] from an unsafe.Pointer.
func MTRDeviceFrom(ptr unsafe.Pointer) MTRDevice {
	return MTRDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceClass) Alloc() MTRDevice {
	rv := objc.Send[MTRDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceClass) New() MTRDevice {
	rv := objc.Send[MTRDevice](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDevice) Init() MTRDevice {
	rv := objc.Send[MTRDevice](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDevice) Autorelease() MTRDevice {
	rv := objc.Send[MTRDevice](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDevice creates a new MTRDevice instance.
func NewMTRDevice() MTRDevice {
	return getMTRDeviceClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/devicecacheprimed
func (m_ MTRDevice) DeviceCachePrimed() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("deviceCachePrimed"))
	return rv
}


// SetDeviceCachePrimed sets the value of the deviceCachePrimed property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/devicecacheprimed
func (m_ MTRDevice) SetDeviceCachePrimed(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceCachePrimed:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/devicecontroller
func (m_ MTRDevice) DeviceController() MTRDeviceController {
	rv := objc.Send[MTRDeviceController](m_.ID, objc.Sel("deviceController"))
	return rv
}


// SetDeviceController sets the value of the deviceController property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/devicecontroller
func (m_ MTRDevice) SetDeviceController(value IMTRDeviceController) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceController:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/estimatedstarttime
func (m_ MTRDevice) EstimatedStartTime() foundation.Date {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("estimatedStartTime"))
	return rv
}


// SetEstimatedStartTime sets the value of the estimatedStartTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/estimatedstarttime
func (m_ MTRDevice) SetEstimatedStartTime(value foundation.IDate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEstimatedStartTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/estimatedsubscriptionlatency
func (m_ MTRDevice) EstimatedSubscriptionLatency() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("estimatedSubscriptionLatency"))
	return rv
}


// SetEstimatedSubscriptionLatency sets the value of the estimatedSubscriptionLatency property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/estimatedsubscriptionlatency
func (m_ MTRDevice) SetEstimatedSubscriptionLatency(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEstimatedSubscriptionLatency:"), value)
}

// Network commissioning features supported by the device.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/networkcommissioningfeatures
func (m_ MTRDevice) NetworkCommissioningFeatures() MTRNetworkCommissioningFeature {
	rv := objc.Send[MTRNetworkCommissioningFeature](m_.ID, objc.Sel("networkCommissioningFeatures"))
	return rv
}


// SetNetworkCommissioningFeatures sets the value of the networkCommissioningFeatures property.
// Network commissioning features supported by the device.

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/networkcommissioningfeatures
func (m_ MTRDevice) SetNetworkCommissioningFeatures(value IMTRNetworkCommissioningFeature) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkCommissioningFeatures:"), value)
}

// The Product Identifier associated with the device.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/productid
func (m_ MTRDevice) ProductID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("productID"))
	return rv
}


// SetProductID sets the value of the productID property.
// The Product Identifier associated with the device.

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/productid
func (m_ MTRDevice) SetProductID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/state
func (m_ MTRDevice) State() MTRDeviceState {
	rv := objc.Send[MTRDeviceState](m_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/state
func (m_ MTRDevice) SetState(value MTRDeviceState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}

// The Vendor Identifier associated with the device.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/vendorid
func (m_ MTRDevice) VendorID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("vendorID"))
	return rv
}


// SetVendorID sets the value of the vendorID property.
// The Vendor Identifier associated with the device.

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/vendorid
func (m_ MTRDevice) SetVendorID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}



