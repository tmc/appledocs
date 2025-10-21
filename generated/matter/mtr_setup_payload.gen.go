// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSetupPayload] class.
var (
	MTRSetupPayloadClass     _MTRSetupPayloadClass
	MTRSetupPayloadClassOnce sync.Once
)

func getMTRSetupPayloadClass() _MTRSetupPayloadClass {
	MTRSetupPayloadClassOnce.Do(func() {
		MTRSetupPayloadClass = _MTRSetupPayloadClass{objc.GetClass("MTRSetupPayload")}
	})
	return MTRSetupPayloadClass
}

type _MTRSetupPayloadClass struct {
	class objc.Class
}

// An interface definition for the [MTRSetupPayload] class.
type IMTRSetupPayload interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSetupPayload
type MTRSetupPayload struct {
	objectivec.Object
}

// MTRSetupPayloadFrom constructs a [MTRSetupPayload] from an unsafe.Pointer.
func MTRSetupPayloadFrom(ptr unsafe.Pointer) MTRSetupPayload {
	return MTRSetupPayload{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSetupPayloadClass) Alloc() MTRSetupPayload {
	rv := objc.Send[MTRSetupPayload](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSetupPayloadClass) New() MTRSetupPayload {
	rv := objc.Send[MTRSetupPayload](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSetupPayload) Init() MTRSetupPayload {
	rv := objc.Send[MTRSetupPayload](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSetupPayload) Autorelease() MTRSetupPayload {
	rv := objc.Send[MTRSetupPayload](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSetupPayload creates a new MTRSetupPayload instance.
func NewMTRSetupPayload() MTRSetupPayload {
	return getMTRSetupPayloadClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSetupPayload/init(onboardingPayload:)
func NewMTRSetupPayloadWithOnboardingPayloadError(onboardingPayload string, error_ unsafe.Pointer) MTRSetupPayload {
	rv := objc.Send[MTRSetupPayload](objc.ID(getMTRSetupPayloadClass().class), objc.Sel("setupPayloadWithOnboardingPayload:error:"), objc.String(onboardingPayload), error_)
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSetupPayload/init(onboardingPayload:)
func (mc _MTRSetupPayloadClass) SetupPayloadWithOnboardingPayloadError(onboardingPayload string, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("setupPayloadWithOnboardingPayload:error:"), objc.String(onboardingPayload), error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/setuppincode
func (m_ MTRSetupPayload) SetUpPINCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("setUpPINCode"))
	return rv
}


// SetSetUpPINCode sets the value of the setUpPINCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/setuppincode
func (m_ MTRSetupPayload) SetSetUpPINCode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSetUpPINCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/discriminator
func (m_ MTRSetupPayload) Discriminator() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("discriminator"))
	return rv
}


// SetDiscriminator sets the value of the discriminator property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/discriminator
func (m_ MTRSetupPayload) SetDiscriminator(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDiscriminator:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/setuppasscode
func (m_ MTRSetupPayload) SetupPasscode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("setupPasscode"))
	return rv
}


// SetSetupPasscode sets the value of the setupPasscode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/setuppasscode
func (m_ MTRSetupPayload) SetSetupPasscode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSetupPasscode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/vendorelements
func (m_ MTRSetupPayload) VendorElements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("vendorElements"))
	return rv
}


// SetVendorElements sets the value of the vendorElements property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/vendorelements
func (m_ MTRSetupPayload) SetVendorElements(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorElements:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/discoverycapabilities
func (m_ MTRSetupPayload) DiscoveryCapabilities() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("discoveryCapabilities"))
	return rv
}


// SetDiscoveryCapabilities sets the value of the discoveryCapabilities property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/discoverycapabilities
func (m_ MTRSetupPayload) SetDiscoveryCapabilities(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDiscoveryCapabilities:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/hasshortdiscriminator
func (m_ MTRSetupPayload) HasShortDiscriminator() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasShortDiscriminator"))
	return rv
}


// SetHasShortDiscriminator sets the value of the hasShortDiscriminator property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/hasshortdiscriminator
func (m_ MTRSetupPayload) SetHasShortDiscriminator(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasShortDiscriminator:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/vendorid
func (m_ MTRSetupPayload) VendorID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("vendorID"))
	return rv
}


// SetVendorID sets the value of the vendorID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/vendorid
func (m_ MTRSetupPayload) SetVendorID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/rendezvousinformation
func (m_ MTRSetupPayload) RendezvousInformation() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rendezvousInformation"))
	return rv
}


// SetRendezvousInformation sets the value of the rendezvousInformation property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/rendezvousinformation
func (m_ MTRSetupPayload) SetRendezvousInformation(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRendezvousInformation:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/serialnumber
func (m_ MTRSetupPayload) SerialNumber() string {
	rv := objc.Send[string](m_.ID, objc.Sel("serialNumber"))
	return rv
}


// SetSerialNumber sets the value of the serialNumber property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/serialnumber
func (m_ MTRSetupPayload) SetSerialNumber(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSerialNumber:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/productid
func (m_ MTRSetupPayload) ProductID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("productID"))
	return rv
}


// SetProductID sets the value of the productID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/productid
func (m_ MTRSetupPayload) SetProductID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/version
func (m_ MTRSetupPayload) Version() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("version"))
	return rv
}


// SetVersion sets the value of the version property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/version
func (m_ MTRSetupPayload) SetVersion(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVersion:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/commissioningflow
func (m_ MTRSetupPayload) CommissioningFlow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("commissioningFlow"))
	return rv
}


// SetCommissioningFlow sets the value of the commissioningFlow property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/commissioningflow
func (m_ MTRSetupPayload) SetCommissioningFlow(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommissioningFlow:"), value)
}


