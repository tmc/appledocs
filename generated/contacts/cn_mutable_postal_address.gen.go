// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CNMutablePostalAddress */


/* debug [class_header]: Header for CNMutablePostalAddress */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNMutablePostalAddress */
// An interface definition for the [CNMutablePostalAddress] class.
type ICNMutablePostalAddress interface {
	ICNPostalAddress
	
/* debug [class_interface_properties]: Properties for CNMutablePostalAddress */
	// properties:
	City() objc.IObject /* cross-framework: NSString */
	SetCity(value objc.IObject /* cross-framework: NSString */)
	Country() objc.IObject /* cross-framework: NSString */
	SetCountry(value objc.IObject /* cross-framework: NSString */)
	ISOCountryCode() objc.IObject /* cross-framework: NSString */
	SetISOCountryCode(value objc.IObject /* cross-framework: NSString */)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNMutablePostalAddress */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNMutablePostalAddress */
// Alloc allocates a new instance without initialization.
func (cc _CNMutablePostalAddressClass) Alloc() CNMutablePostalAddress {
	rv := objc.Send[CNMutablePostalAddress](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNMutablePostalAddress */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNMutablePostalAddress *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNMutablePostalAddress */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNMutablePostalAddress */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNMutablePostalAddress */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNMutablePostalAddress */

// The city name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/city
func (c_ CNMutablePostalAddress) City() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("city"))
	return rv
}/* debug [instance_properties/getter]: city */


// The city name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/city
func (c_ CNMutablePostalAddress) SetCity(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCity:"), value)
}/* debug [instance_properties/setter]: city */


// The country or region name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/country
func (c_ CNMutablePostalAddress) Country() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("country"))
	return rv
}/* debug [instance_properties/getter]: country */


// The country or region name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/country
func (c_ CNMutablePostalAddress) SetCountry(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCountry:"), value)
}/* debug [instance_properties/setter]: country */


// The ISO country code, using the ISO 3166-1 alpha-2 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/isoCountryCode
func (c_ CNMutablePostalAddress) ISOCountryCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("ISOCountryCode"))
	return rv
}/* debug [instance_properties/getter]: ISOCountryCode */


// The ISO country code, using the ISO 3166-1 alpha-2 standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/isoCountryCode
func (c_ CNMutablePostalAddress) SetISOCountryCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setISOCountryCode:"), value)
}/* debug [instance_properties/setter]: ISOCountryCode */


// The postal code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/postalCode
func (c_ CNMutablePostalAddress) PostalCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("postalCode"))
	return rv
}/* debug [instance_properties/getter]: postalCode */


// The postal code of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/postalCode
func (c_ CNMutablePostalAddress) SetPostalCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPostalCode:"), value)
}/* debug [instance_properties/setter]: postalCode */


// The state name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/state
func (c_ CNMutablePostalAddress) State() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The state name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/state
func (c_ CNMutablePostalAddress) SetState(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */


// The street name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/street
func (c_ CNMutablePostalAddress) Street() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("street"))
	return rv
}/* debug [instance_properties/getter]: street */


// The street name of the address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/street
func (c_ CNMutablePostalAddress) SetStreet(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStreet:"), value)
}/* debug [instance_properties/setter]: street */


// The subadministrative area (such as a county or other region) in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/subAdministrativeArea
func (c_ CNMutablePostalAddress) SubAdministrativeArea() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subAdministrativeArea"))
	return rv
}/* debug [instance_properties/getter]: subAdministrativeArea */


// The subadministrative area (such as a county or other region) in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/subAdministrativeArea
func (c_ CNMutablePostalAddress) SetSubAdministrativeArea(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubAdministrativeArea:"), value)
}/* debug [instance_properties/setter]: subAdministrativeArea */


// Additional information associated with the location, typically defined at the city or town level, in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/subLocality
func (c_ CNMutablePostalAddress) SubLocality() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subLocality"))
	return rv
}/* debug [instance_properties/getter]: subLocality */


// Additional information associated with the location, typically defined at the city or town level, in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutablePostalAddress/subLocality
func (c_ CNMutablePostalAddress) SetSubLocality(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubLocality:"), value)
}/* debug [instance_properties/setter]: subLocality */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNMutablePostalAddress */



