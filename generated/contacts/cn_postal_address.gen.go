// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNPostalAddress] class.
var (
	CNPostalAddressClass     _CNPostalAddressClass
	CNPostalAddressClassOnce sync.Once
)

func getCNPostalAddressClass() _CNPostalAddressClass {
	CNPostalAddressClassOnce.Do(func() {
		CNPostalAddressClass = _CNPostalAddressClass{objc.GetClass("CNPostalAddress")}
	})
	return CNPostalAddressClass
}

type _CNPostalAddressClass struct {
	class objc.Class
}

// An interface definition for the [CNPostalAddress] class.
type ICNPostalAddress interface {
	objectivec.IObject
	// properties:
	City() objc.IObject /* cross-framework: NSString */
	SetCity(value objc.IObject /* cross-framework: NSString */)
	Country() objc.IObject /* cross-framework: NSString */
	SetCountry(value objc.IObject /* cross-framework: NSString */)
	IsoCountryCode() objc.IObject /* cross-framework: NSString */
	SetIsoCountryCode(value objc.IObject /* cross-framework: NSString */)
	PostalCode() objc.IObject /* cross-framework: NSString */
	SetPostalCode(value objc.IObject /* cross-framework: NSString */)
	State() objc.IObject /* cross-framework: NSString */
	SetState(value objc.IObject /* cross-framework: NSString */)
	Street() objc.IObject /* cross-framework: NSString */
	SetStreet(value objc.IObject /* cross-framework: NSString */)
	SubAdministrativeArea() objc.IObject /* cross-framework: NSString */
	SetSubAdministrativeArea(value objc.IObject /* cross-framework: NSString */)
	SubLocality() objc.IObject /* cross-framework: NSString */
	SetSubLocality(value objc.IObject /* cross-framework: NSString */)
	CNPostalAddressCityKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressCountryKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressISOCountryCodeKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressPostalCodeKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressStateKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressStreetKey() objc.IObject /* cross-framework: NSString */
	// methods:
}

// An immutable representation of the postal address for a contact.
//
// is a thread-safe class.


// An immutable representation of the postal address for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress
type CNPostalAddress struct {
	objectivec.Object
}

// CNPostalAddressFrom constructs a [CNPostalAddress] from an unsafe.Pointer.
//
// An immutable representation of the postal address for a contact.
func CNPostalAddressFrom(ptr unsafe.Pointer) CNPostalAddress {
	return CNPostalAddress{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNPostalAddressClass) Alloc() CNPostalAddress {
	rv := objc.Send[CNPostalAddress](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNPostalAddressClass) New() CNPostalAddress {
	rv := objc.Send[CNPostalAddress](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNPostalAddress) Init() CNPostalAddress {
	rv := objc.Send[CNPostalAddress](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNPostalAddress) Autorelease() CNPostalAddress {
	rv := objc.Send[CNPostalAddress](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNPostalAddress creates a new CNPostalAddress instance.
func NewCNPostalAddress() CNPostalAddress {
	return getCNPostalAddressClass().New()
}



// The city name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/city
func (c_ CNPostalAddress) City() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("city"))
	return rv
}


// The city name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/city
func (c_ CNPostalAddress) SetCity(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCity:"), value)
}


// The country or region name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/country
func (c_ CNPostalAddress) Country() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("country"))
	return rv
}


// The country or region name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/country
func (c_ CNPostalAddress) SetCountry(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCountry:"), value)
}


// The ISO country code for the country or region in a postal address, using the ISO 3166-1 alpha-2 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/isocountrycode
func (c_ CNPostalAddress) IsoCountryCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("isoCountryCode"))
	return rv
}


// The ISO country code for the country or region in a postal address, using the ISO 3166-1 alpha-2 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/isocountrycode
func (c_ CNPostalAddress) SetIsoCountryCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsoCountryCode:"), value)
}


// The postal code in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/postalcode
func (c_ CNPostalAddress) PostalCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("postalCode"))
	return rv
}


// The postal code in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/postalcode
func (c_ CNPostalAddress) SetPostalCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPostalCode:"), value)
}


// The state name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/state
func (c_ CNPostalAddress) State() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("state"))
	return rv
}


// The state name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/state
func (c_ CNPostalAddress) SetState(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setState:"), value)
}


// The street name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/street
func (c_ CNPostalAddress) Street() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("street"))
	return rv
}


// The street name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/street
func (c_ CNPostalAddress) SetStreet(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStreet:"), value)
}


// The subadministrative area (such as a county or other region) in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/subadministrativearea
func (c_ CNPostalAddress) SubAdministrativeArea() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subAdministrativeArea"))
	return rv
}


// The subadministrative area (such as a county or other region) in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/subadministrativearea
func (c_ CNPostalAddress) SetSubAdministrativeArea(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubAdministrativeArea:"), value)
}


// Additional information associated with the location, typically defined at the city or town level, in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/sublocality
func (c_ CNPostalAddress) SubLocality() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subLocality"))
	return rv
}


// Additional information associated with the location, typically defined at the city or town level, in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/sublocality
func (c_ CNPostalAddress) SetSubLocality(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubLocality:"), value)
}


// The city of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresscitykey
func (c_ CNPostalAddress) CNPostalAddressCityKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressCityKey"))
	return rv
}


// The country or region name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresscountrykey
func (c_ CNPostalAddress) CNPostalAddressCountryKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressCountryKey"))
	return rv
}


// The ISO country code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressisocountrycodekey
func (c_ CNPostalAddress) CNPostalAddressISOCountryCodeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressISOCountryCodeKey"))
	return rv
}


// The postal code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresspostalcodekey
func (c_ CNPostalAddress) CNPostalAddressPostalCodeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressPostalCodeKey"))
	return rv
}


// The state name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressstatekey
func (c_ CNPostalAddress) CNPostalAddressStateKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressStateKey"))
	return rv
}


// The street name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressstreetkey
func (c_ CNPostalAddress) CNPostalAddressStreetKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressStreetKey"))
	return rv
}



