// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// An immutable representation of the postal address for a contact.
//
// is a thread-safe class.
//
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


// Returns the localized name for the property associated with the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/localizedString(forKey:)
func (cc _CNPostalAddressClass) LocalizedStringForKey(key string) string {
	rv := objc.Send[string](objc.ID(cc.class), objc.Sel("localizedStringForKey:"), objc.String(key))
	return rv
}

// The city name in a postal address.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/city
func (c_ CNPostalAddress) City() string {
	rv := objc.Send[string](c_.ID, objc.Sel("city"))
	return rv
}

// The country or region name in a postal address.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/country
func (c_ CNPostalAddress) Country() string {
	rv := objc.Send[string](c_.ID, objc.Sel("country"))
	return rv
}

// The ISO country code for the country or region in a postal address, using the ISO 3166-1 alpha-2 standard.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/isoCountryCode
func (c_ CNPostalAddress) ISOCountryCode() string {
	rv := objc.Send[string](c_.ID, objc.Sel("ISOCountryCode"))
	return rv
}

// The postal code in a postal address.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/postalCode
func (c_ CNPostalAddress) PostalCode() string {
	rv := objc.Send[string](c_.ID, objc.Sel("postalCode"))
	return rv
}

// The state name in a postal address.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/state
func (c_ CNPostalAddress) State() string {
	rv := objc.Send[string](c_.ID, objc.Sel("state"))
	return rv
}

// The street name in a postal address.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/street
func (c_ CNPostalAddress) Street() string {
	rv := objc.Send[string](c_.ID, objc.Sel("street"))
	return rv
}

// The subadministrative area (such as a county or other region) in a postal address.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/subAdministrativeArea
func (c_ CNPostalAddress) SubAdministrativeArea() string {
	rv := objc.Send[string](c_.ID, objc.Sel("subAdministrativeArea"))
	return rv
}

// Additional information associated with the location, typically defined at the city or town level, in a postal address.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/subLocality
func (c_ CNPostalAddress) SubLocality() string {
	rv := objc.Send[string](c_.ID, objc.Sel("subLocality"))
	return rv
}



