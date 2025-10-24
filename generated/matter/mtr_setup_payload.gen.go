// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	CommissioningFlow() MTRCommissioningFlow
	SetCommissioningFlow(value MTRCommissioningFlow)
	DiscoveryCapabilities() MTRDiscoveryCapabilities
	SetDiscoveryCapabilities(value MTRDiscoveryCapabilities)
	Discriminator() objc.IObject /* cross-framework: NSNumber */
	SetDiscriminator(value objc.IObject /* cross-framework: NSNumber */)
	HasShortDiscriminator() bool
	SetHasShortDiscriminator(value bool)
	ProductID() objc.IObject /* cross-framework: NSNumber */
	SetProductID(value objc.IObject /* cross-framework: NSNumber */)
	RendezvousInformation() objc.IObject /* cross-framework: NSNumber */
	SetRendezvousInformation(value objc.IObject /* cross-framework: NSNumber */)
	SerialNumber() objc.IObject /* cross-framework: NSString */
	SetSerialNumber(value objc.IObject /* cross-framework: NSString */)
	SetUpPINCode() objc.IObject /* cross-framework: NSNumber */
	SetSetUpPINCode(value objc.IObject /* cross-framework: NSNumber */)
	SetupPasscode() objc.IObject /* cross-framework: NSNumber */
	SetSetupPasscode(value objc.IObject /* cross-framework: NSNumber */)
	VendorElements() IMTROptionalQRCodeInfo
	SetVendorElements(value IMTROptionalQRCodeInfo)
	VendorID() objc.IObject /* cross-framework: NSNumber */
	SetVendorID(value objc.IObject /* cross-framework: NSNumber */)
	Version() objc.IObject /* cross-framework: NSNumber */
	SetVersion(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSetupPayload/init(onboardingPayload:)
func NewMTRSetupPayloadWithOnboardingPayloadError(onboardingPayload objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) MTRSetupPayload {
	rv := objc.Send[MTRSetupPayload](objc.ID(getMTRSetupPayloadClass().class), objc.Sel("setupPayloadWithOnboardingPayload:error:"), onboardingPayload, error_)
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSetupPayload/init(onboardingPayload:)
func (mc _MTRSetupPayloadClass) SetupPayloadWithOnboardingPayloadError(onboardingPayload objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) MTRSetupPayload {
	rv := objc.Send[MTRSetupPayload](objc.ID(mc.class), objc.Sel("setupPayloadWithOnboardingPayload:error:"), onboardingPayload, error_)
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/commissioningflow
func (m_ MTRSetupPayload) CommissioningFlow() MTRCommissioningFlow {
	rv := objc.Send[MTRCommissioningFlow](m_.ID, objc.Sel("commissioningFlow"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/commissioningflow
func (m_ MTRSetupPayload) SetCommissioningFlow(value MTRCommissioningFlow) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommissioningFlow:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/discoverycapabilities
func (m_ MTRSetupPayload) DiscoveryCapabilities() MTRDiscoveryCapabilities {
	rv := objc.Send[MTRDiscoveryCapabilities](m_.ID, objc.Sel("discoveryCapabilities"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/discoverycapabilities
func (m_ MTRSetupPayload) SetDiscoveryCapabilities(value MTRDiscoveryCapabilities) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDiscoveryCapabilities:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/discriminator
func (m_ MTRSetupPayload) Discriminator() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("discriminator"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/discriminator
func (m_ MTRSetupPayload) SetDiscriminator(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDiscriminator:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/hasshortdiscriminator
func (m_ MTRSetupPayload) HasShortDiscriminator() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasShortDiscriminator"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/hasshortdiscriminator
func (m_ MTRSetupPayload) SetHasShortDiscriminator(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasShortDiscriminator:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/productid
func (m_ MTRSetupPayload) ProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productID"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/productid
func (m_ MTRSetupPayload) SetProductID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/rendezvousinformation
func (m_ MTRSetupPayload) RendezvousInformation() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rendezvousInformation"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/rendezvousinformation
func (m_ MTRSetupPayload) SetRendezvousInformation(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRendezvousInformation:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/serialnumber
func (m_ MTRSetupPayload) SerialNumber() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("serialNumber"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/serialnumber
func (m_ MTRSetupPayload) SetSerialNumber(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSerialNumber:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/setuppincode
func (m_ MTRSetupPayload) SetUpPINCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("setUpPINCode"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/setuppincode
func (m_ MTRSetupPayload) SetSetUpPINCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSetUpPINCode:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/setuppasscode
func (m_ MTRSetupPayload) SetupPasscode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("setupPasscode"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/setuppasscode
func (m_ MTRSetupPayload) SetSetupPasscode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSetupPasscode:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/vendorelements
func (m_ MTRSetupPayload) VendorElements() IMTROptionalQRCodeInfo {
	rv := objc.Send[MTROptionalQRCodeInfo](m_.ID, objc.Sel("vendorElements"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/vendorelements
func (m_ MTRSetupPayload) SetVendorElements(value IMTROptionalQRCodeInfo) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorElements:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/vendorid
func (m_ MTRSetupPayload) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/vendorid
func (m_ MTRSetupPayload) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/version
func (m_ MTRSetupPayload) Version() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("version"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsetuppayload/version
func (m_ MTRSetupPayload) SetVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVersion:"), value)
}
