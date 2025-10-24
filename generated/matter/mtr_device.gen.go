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
	// properties:
	DeviceCachePrimed() bool
	SetDeviceCachePrimed(value bool)
	DeviceController() IMTRDeviceController
	SetDeviceController(value IMTRDeviceController)
	EstimatedStartTime() objc.IObject /* cross-framework: Date */
	SetEstimatedStartTime(value objc.IObject /* cross-framework: Date */)
	EstimatedSubscriptionLatency() objc.IObject /* cross-framework: NSNumber */
	SetEstimatedSubscriptionLatency(value objc.IObject /* cross-framework: NSNumber */)
	NetworkCommissioningFeatures() MTRNetworkCommissioningFeature
	SetNetworkCommissioningFeatures(value MTRNetworkCommissioningFeature)
	ProductID() objc.IObject /* cross-framework: NSNumber */
	SetProductID(value objc.IObject /* cross-framework: NSNumber */)
	State() MTRDeviceState
	SetState(value MTRDeviceState)
	VendorID() objc.IObject /* cross-framework: NSNumber */
	SetVendorID(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/devicecacheprimed
func (m_ MTRDevice) DeviceCachePrimed() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("deviceCachePrimed"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/devicecacheprimed
func (m_ MTRDevice) SetDeviceCachePrimed(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceCachePrimed:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/devicecontroller
func (m_ MTRDevice) DeviceController() IMTRDeviceController {
	rv := objc.Send[MTRDeviceController](m_.ID, objc.Sel("deviceController"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/devicecontroller
func (m_ MTRDevice) SetDeviceController(value IMTRDeviceController) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceController:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/estimatedstarttime
func (m_ MTRDevice) EstimatedStartTime() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("estimatedStartTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/estimatedstarttime
func (m_ MTRDevice) SetEstimatedStartTime(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEstimatedStartTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/estimatedsubscriptionlatency
func (m_ MTRDevice) EstimatedSubscriptionLatency() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("estimatedSubscriptionLatency"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/estimatedsubscriptionlatency
func (m_ MTRDevice) SetEstimatedSubscriptionLatency(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEstimatedSubscriptionLatency:"), value)
}


// Network commissioning features supported by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/networkcommissioningfeatures
func (m_ MTRDevice) NetworkCommissioningFeatures() MTRNetworkCommissioningFeature {
	rv := objc.Send[MTRNetworkCommissioningFeature](m_.ID, objc.Sel("networkCommissioningFeatures"))
	return rv
}


// Network commissioning features supported by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/networkcommissioningfeatures
func (m_ MTRDevice) SetNetworkCommissioningFeatures(value MTRNetworkCommissioningFeature) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkCommissioningFeatures:"), value)
}


// The Product Identifier associated with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/productid
func (m_ MTRDevice) ProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productID"))
	return rv
}


// The Product Identifier associated with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/productid
func (m_ MTRDevice) SetProductID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/state
func (m_ MTRDevice) State() MTRDeviceState {
	rv := objc.Send[MTRDeviceState](m_.ID, objc.Sel("state"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/state
func (m_ MTRDevice) SetState(value MTRDeviceState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}


// The Vendor Identifier associated with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/vendorid
func (m_ MTRDevice) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}


// The Vendor Identifier associated with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevice/vendorid
func (m_ MTRDevice) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}



