// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [DDMatchPostalAddress] class.
type IDDMatchPostalAddress interface {
	IDDMatch
	City() string
	Country() string
	PostalCode() string
	State() string
	Street() string
}

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

// Alloc allocates a new instance without initialization.
func (dc _DDMatchPostalAddressClass) Alloc() DDMatchPostalAddress {
	rv := objc.Send[DDMatchPostalAddress](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The city name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPostalAddress/city
func (d_ DDMatchPostalAddress) City() string {
	rv := objc.Send[string](d_.ID, objc.Sel("city"))
	return rv
}


// The country or region name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPostalAddress/country
func (d_ DDMatchPostalAddress) Country() string {
	rv := objc.Send[string](d_.ID, objc.Sel("country"))
	return rv
}


// The postal code in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPostalAddress/postalCode
func (d_ DDMatchPostalAddress) PostalCode() string {
	rv := objc.Send[string](d_.ID, objc.Sel("postalCode"))
	return rv
}


// The state name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPostalAddress/state
func (d_ DDMatchPostalAddress) State() string {
	rv := objc.Send[string](d_.ID, objc.Sel("state"))
	return rv
}


// The street name in a postal address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchPostalAddress/street
func (d_ DDMatchPostalAddress) Street() string {
	rv := objc.Send[string](d_.ID, objc.Sel("street"))
	return rv
}



