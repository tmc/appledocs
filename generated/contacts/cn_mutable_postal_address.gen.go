// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
func (c_ CNMutablePostalAddress) City() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("city"))
	return rv
}


// The city name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/city
func (c_ CNMutablePostalAddress) SetCity(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCity:"), objc.String(value))
}


// The country or region name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/country
func (c_ CNMutablePostalAddress) Country() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("country"))
	return rv
}


// The country or region name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/country
func (c_ CNMutablePostalAddress) SetCountry(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCountry:"), objc.String(value))
}


// The ISO country code, using the ISO 3166-1 alpha-2 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/isocountrycode
func (c_ CNMutablePostalAddress) IsoCountryCode() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("isoCountryCode"))
	return rv
}


// The ISO country code, using the ISO 3166-1 alpha-2 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/isocountrycode
func (c_ CNMutablePostalAddress) SetIsoCountryCode(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsoCountryCode:"), objc.String(value))
}


// The postal code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/postalcode
func (c_ CNMutablePostalAddress) PostalCode() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("postalCode"))
	return rv
}


// The postal code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/postalcode
func (c_ CNMutablePostalAddress) SetPostalCode(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPostalCode:"), objc.String(value))
}


// The state name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/state
func (c_ CNMutablePostalAddress) State() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("state"))
	return rv
}


// The state name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/state
func (c_ CNMutablePostalAddress) SetState(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setState:"), objc.String(value))
}


// The street name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/street
func (c_ CNMutablePostalAddress) Street() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("street"))
	return rv
}


// The street name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/street
func (c_ CNMutablePostalAddress) SetStreet(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStreet:"), objc.String(value))
}


// The subadministrative area (such as a county or other region) in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/subadministrativearea
func (c_ CNMutablePostalAddress) SubAdministrativeArea() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("subAdministrativeArea"))
	return rv
}


// The subadministrative area (such as a county or other region) in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/subadministrativearea
func (c_ CNMutablePostalAddress) SetSubAdministrativeArea(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubAdministrativeArea:"), objc.String(value))
}


// Additional information associated with the location, typically defined at the city or town level, in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/sublocality
func (c_ CNMutablePostalAddress) SubLocality() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("subLocality"))
	return rv
}


// Additional information associated with the location, typically defined at the city or town level, in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/sublocality
func (c_ CNMutablePostalAddress) SetSubLocality(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubLocality:"), objc.String(value))
}



