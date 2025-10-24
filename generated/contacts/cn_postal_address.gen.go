// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNPostalAddress */


/* debug [class_header]: Header for CNPostalAddress */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNPostalAddress */
// An interface definition for the [CNPostalAddress] class.
type ICNPostalAddress interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNPostalAddress */
	// properties:
	City() objc.IObject /* cross-framework: NSString */
	Country() objc.IObject /* cross-framework: NSString */
	ISOCountryCode() objc.IObject /* cross-framework: NSString */
	PostalCode() objc.IObject /* cross-framework: NSString */
	State() objc.IObject /* cross-framework: NSString */
	Street() objc.IObject /* cross-framework: NSString */
	SubAdministrativeArea() objc.IObject /* cross-framework: NSString */
	SubLocality() objc.IObject /* cross-framework: NSString */
	CNPostalAddressCityKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressCountryKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressISOCountryCodeKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressPostalCodeKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressStateKey() objc.IObject /* cross-framework: NSString */
	CNPostalAddressStreetKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNPostalAddress */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNPostalAddress */
// Alloc allocates a new instance without initialization.
func (cc _CNPostalAddressClass) Alloc() CNPostalAddress {
	rv := objc.Send[CNPostalAddress](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNPostalAddress */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNPostalAddress *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNPostalAddress */

// Returns the localized name for the property associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/localizedString(forKey:)
func (cc _CNPostalAddressClass) LocalizedStringForKey(key objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("localizedStringForKey:"), key)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocalizedStringForKey) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNPostalAddress */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNPostalAddress */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNPostalAddress */

// The city name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/city
func (c_ CNPostalAddress) City() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("city"))
	return rv
}/* debug [instance_properties/getter]: city */


// The country or region name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/country
func (c_ CNPostalAddress) Country() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("country"))
	return rv
}/* debug [instance_properties/getter]: country */


// The ISO country code for the country or region in a postal address, using the ISO 3166-1 alpha-2 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/isoCountryCode
func (c_ CNPostalAddress) ISOCountryCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("ISOCountryCode"))
	return rv
}/* debug [instance_properties/getter]: ISOCountryCode */


// The postal code in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/postalCode
func (c_ CNPostalAddress) PostalCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("postalCode"))
	return rv
}/* debug [instance_properties/getter]: postalCode */


// The state name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/state
func (c_ CNPostalAddress) State() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The street name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/street
func (c_ CNPostalAddress) Street() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("street"))
	return rv
}/* debug [instance_properties/getter]: street */


// The subadministrative area (such as a county or other region) in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/subAdministrativeArea
func (c_ CNPostalAddress) SubAdministrativeArea() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subAdministrativeArea"))
	return rv
}/* debug [instance_properties/getter]: subAdministrativeArea */


// Additional information associated with the location, typically defined at the city or town level, in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPostalAddress/subLocality
func (c_ CNPostalAddress) SubLocality() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subLocality"))
	return rv
}/* debug [instance_properties/getter]: subLocality */


// The city of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresscitykey
func (c_ CNPostalAddress) CNPostalAddressCityKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressCityKey"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressCityKey */


// The country or region name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresscountrykey
func (c_ CNPostalAddress) CNPostalAddressCountryKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressCountryKey"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressCountryKey */


// The ISO country code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressisocountrycodekey
func (c_ CNPostalAddress) CNPostalAddressISOCountryCodeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressISOCountryCodeKey"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressISOCountryCodeKey */


// The postal code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdresspostalcodekey
func (c_ CNPostalAddress) CNPostalAddressPostalCodeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressPostalCodeKey"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressPostalCodeKey */


// The state name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressstatekey
func (c_ CNPostalAddress) CNPostalAddressStateKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressStateKey"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressStateKey */


// The street name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdressstreetkey
func (c_ CNPostalAddress) CNPostalAddressStreetKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNPostalAddressStreetKey"))
	return rv
}/* debug [instance_properties/getter]: CNPostalAddressStreetKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNPostalAddress */





