// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DDMatchPostalAddress */


/* debug [class_header]: Header for DDMatchPostalAddress */
// The class instance for the [DDMatchPostalAddress] class.
var (
	DDMatchPostalAddressClass     _DDMatchPostalAddressClass
	DDMatchPostalAddressClassOnce sync.Once
)

func getDDMatchPostalAddressClass() _DDMatchPostalAddressClass {
	DDMatchPostalAddressClassOnce.Do(func() {
		DDMatchPostalAddressClass = _DDMatchPostalAddressClass{objc.GetClass("DDMatchPostalAddress")}
	})
	return DDMatchPostalAddressClass
}

type _DDMatchPostalAddressClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DDMatchPostalAddress */
// An interface definition for the [DDMatchPostalAddress] class.
type IDDMatchPostalAddress interface {
	IDDMatch
	
/* debug [class_interface_properties]: Properties for DDMatchPostalAddress */
	// properties:
	City() objc.IObject /* cross-framework: NSString */
	Country() objc.IObject /* cross-framework: NSString */
	PostalCode() objc.IObject /* cross-framework: NSString */
	State() objc.IObject /* cross-framework: NSString */
	Street() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DDMatchPostalAddress */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DDMatchPostalAddress */
// Alloc allocates a new instance without initialization.
func (dc _DDMatchPostalAddressClass) Alloc() DDMatchPostalAddress {
	rv := objc.Send[DDMatchPostalAddress](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DDMatchPostalAddressClass) New() DDMatchPostalAddress {
	rv := objc.Send[DDMatchPostalAddress](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDMatchPostalAddress) Init() DDMatchPostalAddress {
	rv := objc.Send[DDMatchPostalAddress](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDMatchPostalAddress) Autorelease() DDMatchPostalAddress {
	rv := objc.Send[DDMatchPostalAddress](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDMatchPostalAddress creates a new DDMatchPostalAddress instance.
func NewDDMatchPostalAddress() DDMatchPostalAddress {
	return getDDMatchPostalAddressClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DDMatchPostalAddress */
// An object that contains a postal address that the data detection system matches.
//
// The DataDetection framework returns a postal address match in a object, which optionally contains the matching parts of a postal address: street, city, state, postal code, and country.


// An object that contains a postal address that the data detection system matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPostalAddress
type DDMatchPostalAddress struct {
	DDMatch
}

// DDMatchPostalAddressFrom constructs a [DDMatchPostalAddress] from an unsafe.Pointer.
//
// An object that contains a postal address that the data detection system matches.
func DDMatchPostalAddressFrom(ptr unsafe.Pointer) DDMatchPostalAddress {
	return DDMatchPostalAddress{
		DDMatch: DDMatchFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DDMatchPostalAddress *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DDMatchPostalAddress */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DDMatchPostalAddress */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DDMatchPostalAddress */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DDMatchPostalAddress */

// The city name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPostalAddress/city
func (d_ DDMatchPostalAddress) City() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("city"))
	return rv
}/* debug [instance_properties/getter]: city */


// The country or region name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPostalAddress/country
func (d_ DDMatchPostalAddress) Country() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("country"))
	return rv
}/* debug [instance_properties/getter]: country */


// The postal code in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPostalAddress/postalCode
func (d_ DDMatchPostalAddress) PostalCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("postalCode"))
	return rv
}/* debug [instance_properties/getter]: postalCode */


// The state name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPostalAddress/state
func (d_ DDMatchPostalAddress) State() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The street name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPostalAddress/street
func (d_ DDMatchPostalAddress) Street() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("street"))
	return rv
}/* debug [instance_properties/getter]: street */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DDMatchPostalAddress */



