// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKAddress */


/* debug [class_header]: Header for MKAddress */
// The class instance for the [MKAddress] class.
var (
	MKAddressClass     _MKAddressClass
	MKAddressClassOnce sync.Once
)

func getMKAddressClass() _MKAddressClass {
	MKAddressClassOnce.Do(func() {
		MKAddressClass = _MKAddressClass{objc.GetClass("MKAddress")}
	})
	return MKAddressClass
}

type _MKAddressClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKAddress */
// An interface definition for the [MKAddress] class.
type IMKAddress interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKAddress */
	// properties:
	FullAddress() objc.IObject /* cross-framework: NSString */
	ShortAddress() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKAddress */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKAddress */
// Alloc allocates a new instance without initialization.
func (mc _MKAddressClass) Alloc() MKAddress {
	rv := objc.Send[MKAddress](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKAddressClass) New() MKAddress {
	rv := objc.Send[MKAddress](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKAddress) Init() MKAddress {
	rv := objc.Send[MKAddress](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKAddress) Autorelease() MKAddress {
	rv := objc.Send[MKAddress](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKAddress creates a new MKAddress instance.
func NewMKAddress() MKAddress {
	return getMKAddressClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKAddress */
// A class that contains a full address, and, optionally, a short address.
//
// MapKit capabilities, such as Search and Reverse geocoding, populate the of a with a full address, and a short address, if the framework has one. When presenting a Place Card using an or a selection accessory on an annotation you created using an , MapKit uses the full address provided if you create the using .


// A class that contains a full address, and, optionally, a short address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddress
type MKAddress struct {
	objectivec.Object
}

// MKAddressFrom constructs a [MKAddress] from an unsafe.Pointer.
//
// A class that contains a full address, and, optionally, a short address.
func MKAddressFrom(ptr unsafe.Pointer) MKAddress {
	return MKAddress{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKAddress */

// Initializes a new address with a location’s full address using a string and a short address that provides an abbreviated form of the address such as a street address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddress/init(fullAddress:shortAddress:)
func NewMKAddressWithFullAddressShortAddress(fullAddress objc.IObject /* cross-framework: NSString */, shortAddress objc.IObject /* cross-framework: NSString */) MKAddress {
	instance := getMKAddressClass().Alloc()
	rv := objc.Send[MKAddress](instance.ID, objc.Sel("initWithFullAddress:shortAddress:"), fullAddress, shortAddress)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKAddressWithFullAddressShortAddress */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKAddress */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKAddress */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKAddress */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKAddress */

// A string that represents a place’s full address
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddress/fullAddress
func (m_ MKAddress) FullAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("fullAddress"))
	return rv
}/* debug [instance_properties/getter]: fullAddress */


// A string that represents the short address of a location, such as it’s street address and city.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddress/shortAddress
func (m_ MKAddress) ShortAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("shortAddress"))
	return rv
}/* debug [instance_properties/getter]: shortAddress */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKAddress */


