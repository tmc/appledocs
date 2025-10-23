// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	City() string /* primitive/slice/pointer. */
	SetCity(value string /* primitive/slice/pointer. */)
	Country() string /* primitive/slice/pointer. */
	SetCountry(value string /* primitive/slice/pointer. */)
	IsoCountryCode() string /* primitive/slice/pointer. */
	SetIsoCountryCode(value string /* primitive/slice/pointer. */)
	PostalCode() string /* primitive/slice/pointer. */
	SetPostalCode(value string /* primitive/slice/pointer. */)
	State() string /* primitive/slice/pointer. */
	SetState(value string /* primitive/slice/pointer. */)
	Street() string /* primitive/slice/pointer. */
	SetStreet(value string /* primitive/slice/pointer. */)
	SubAdministrativeArea() string /* primitive/slice/pointer. */
	SetSubAdministrativeArea(value string /* primitive/slice/pointer. */)
	SubLocality() string /* primitive/slice/pointer. */
	SetSubLocality(value string /* primitive/slice/pointer. */)
	CNPostalAddressCityKey() string /* primitive/slice/pointer. */
	CNPostalAddressCountryKey() string /* primitive/slice/pointer. */
	CNPostalAddressISOCountryCodeKey() string /* primitive/slice/pointer. */
	CNPostalAddressPostalCodeKey() string /* primitive/slice/pointer. */
	CNPostalAddressStateKey() string /* primitive/slice/pointer. */
	CNPostalAddressStreetKey() string /* primitive/slice/pointer. */
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
func (c_ CNPostalAddress) City() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("city"))
	return rv
}


// The city name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/city
func (c_ CNPostalAddress) SetCity(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCity:"), objc.String(value))
}


// The country or region name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/country
func (c_ CNPostalAddress) Country() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("country"))
	return rv
}


// The country or region name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/country
func (c_ CNPostalAddress) SetCountry(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCountry:"), objc.String(value))
}


// The ISO country code for the country or region in a postal address, using the ISO 3166-1 alpha-2 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/isocountrycode
func (c_ CNPostalAddress) IsoCountryCode() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("isoCountryCode"))
	return rv
}


// The ISO country code for the country or region in a postal address, using the ISO 3166-1 alpha-2 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/isocountrycode
func (c_ CNPostalAddress) SetIsoCountryCode(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsoCountryCode:"), objc.String(value))
}


// The postal code in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/postalcode
func (c_ CNPostalAddress) PostalCode() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("postalCode"))
	return rv
}


// The postal code in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/postalcode
func (c_ CNPostalAddress) SetPostalCode(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPostalCode:"), objc.String(value))
}


// The state name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/state
func (c_ CNPostalAddress) State() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("state"))
	return rv
}


// The state name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/state
func (c_ CNPostalAddress) SetState(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setState:"), objc.String(value))
}


// The street name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/street
func (c_ CNPostalAddress) Street() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("street"))
	return rv
}


// The street name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/street
func (c_ CNPostalAddress) SetStreet(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStreet:"), objc.String(value))
}


// The subadministrative area (such as a county or other region) in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/subadministrativearea
func (c_ CNPostalAddress) SubAdministrativeArea() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("subAdministrativeArea"))
	return rv
}


// The subadministrative area (such as a county or other region) in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/subadministrativearea
func (c_ CNPostalAddress) SetSubAdministrativeArea(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubAdministrativeArea:"), objc.String(value))
}


// Additional information associated with the location, typically defined at the city or town level, in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/sublocality
func (c_ CNPostalAddress) SubLocality() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("subLocality"))
	return rv
}


// Additional information associated with the location, typically defined at the city or town level, in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/sublocality
func (c_ CNPostalAddress) SetSubLocality(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubLocality:"), objc.String(value))
}


// The city of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresscitykey
func (c_ CNPostalAddress) CNPostalAddressCityKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressCityKey"))
	return rv
}


// The country or region name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresscountrykey
func (c_ CNPostalAddress) CNPostalAddressCountryKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressCountryKey"))
	return rv
}


// The ISO country code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressisocountrycodekey
func (c_ CNPostalAddress) CNPostalAddressISOCountryCodeKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressISOCountryCodeKey"))
	return rv
}


// The postal code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresspostalcodekey
func (c_ CNPostalAddress) CNPostalAddressPostalCodeKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressPostalCodeKey"))
	return rv
}


// The state name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressstatekey
func (c_ CNPostalAddress) CNPostalAddressStateKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressStateKey"))
	return rv
}


// The street name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressstreetkey
func (c_ CNPostalAddress) CNPostalAddressStreetKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNPostalAddressStreetKey"))
	return rv
}



