// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CNMutablePostalAddress] class.
var (
	CNMutablePostalAddressClass     _CNMutablePostalAddressClass
	CNMutablePostalAddressClassOnce sync.Once
)

func getCNMutablePostalAddressClass() _CNMutablePostalAddressClass {
	CNMutablePostalAddressClassOnce.Do(func() {
		CNMutablePostalAddressClass = _CNMutablePostalAddressClass{objc.GetClass("CNMutablePostalAddress")}
	})
	return CNMutablePostalAddressClass
}

type _CNMutablePostalAddressClass struct {
	class objc.Class
}

// An interface definition for the [CNMutablePostalAddress] class.
type ICNMutablePostalAddress interface {
	ICNPostalAddress
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
	// methods:
}

// A mutable representation of the postal address for a contact.
//
// is not a thread-safe class. To remove properties when saving a mutable postal address, set string properties to empty values.


// A mutable representation of the postal address for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress
type CNMutablePostalAddress struct {
	CNPostalAddress
}

// CNMutablePostalAddressFrom constructs a [CNMutablePostalAddress] from an unsafe.Pointer.
//
// A mutable representation of the postal address for a contact.
func CNMutablePostalAddressFrom(ptr unsafe.Pointer) CNMutablePostalAddress {
	return CNMutablePostalAddress{
		CNPostalAddress: CNPostalAddressFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNMutablePostalAddressClass) Alloc() CNMutablePostalAddress {
	rv := objc.Send[CNMutablePostalAddress](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNMutablePostalAddressClass) New() CNMutablePostalAddress {
	rv := objc.Send[CNMutablePostalAddress](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNMutablePostalAddress) Init() CNMutablePostalAddress {
	rv := objc.Send[CNMutablePostalAddress](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNMutablePostalAddress) Autorelease() CNMutablePostalAddress {
	rv := objc.Send[CNMutablePostalAddress](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNMutablePostalAddress creates a new CNMutablePostalAddress instance.
func NewCNMutablePostalAddress() CNMutablePostalAddress {
	return getCNMutablePostalAddressClass().New()
}



// The city name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/city
func (c_ CNMutablePostalAddress) City() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("city"))
	return rv
}


// The city name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/city
func (c_ CNMutablePostalAddress) SetCity(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCity:"), value)
}


// The country or region name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/country
func (c_ CNMutablePostalAddress) Country() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("country"))
	return rv
}


// The country or region name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/country
func (c_ CNMutablePostalAddress) SetCountry(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCountry:"), value)
}


// The ISO country code, using the ISO 3166-1 alpha-2 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/isocountrycode
func (c_ CNMutablePostalAddress) IsoCountryCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("isoCountryCode"))
	return rv
}


// The ISO country code, using the ISO 3166-1 alpha-2 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/isocountrycode
func (c_ CNMutablePostalAddress) SetIsoCountryCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsoCountryCode:"), value)
}


// The postal code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/postalcode
func (c_ CNMutablePostalAddress) PostalCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("postalCode"))
	return rv
}


// The postal code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/postalcode
func (c_ CNMutablePostalAddress) SetPostalCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPostalCode:"), value)
}


// The state name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/state
func (c_ CNMutablePostalAddress) State() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("state"))
	return rv
}


// The state name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/state
func (c_ CNMutablePostalAddress) SetState(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setState:"), value)
}


// The street name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/street
func (c_ CNMutablePostalAddress) Street() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("street"))
	return rv
}


// The street name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/street
func (c_ CNMutablePostalAddress) SetStreet(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStreet:"), value)
}


// The subadministrative area (such as a county or other region) in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/subadministrativearea
func (c_ CNMutablePostalAddress) SubAdministrativeArea() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subAdministrativeArea"))
	return rv
}


// The subadministrative area (such as a county or other region) in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/subadministrativearea
func (c_ CNMutablePostalAddress) SetSubAdministrativeArea(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubAdministrativeArea:"), value)
}


// Additional information associated with the location, typically defined at the city or town level, in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/sublocality
func (c_ CNMutablePostalAddress) SubLocality() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subLocality"))
	return rv
}


// Additional information associated with the location, typically defined at the city or town level, in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/sublocality
func (c_ CNMutablePostalAddress) SetSubLocality(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubLocality:"), value)
}



