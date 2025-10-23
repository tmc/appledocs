// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MKAddress] class.
type IMKAddress interface {
	objectivec.IObject
	FullAddress() string
	SetFullAddress(value string)
	ShortAddress() string
	SetShortAddress(value string)
}

// A class that contains a full address, and, optionally, a short address.


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

// Alloc allocates a new instance without initialization.
func (mc _MKAddressClass) Alloc() MKAddress {
	rv := objc.Send[MKAddress](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Initializes a new address with a location’s full address using a string and a short address that provides an abbreviated form of the address such as a street address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddress/init(fullAddress:shortAddress:)
func NewMKAddressWithFullAddressShortAddress(fullAddress string, shortAddress string) MKAddress {
	instance := getMKAddressClass().Alloc()
	rv := objc.Send[MKAddress](instance.ID, objc.Sel("initWithFullAddress:shortAddress:"), objc.String(fullAddress), objc.String(shortAddress))
	rv.Autorelease()
	return rv
}



// A string that represents a place’s full address
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddress/fulladdress
func (m_ MKAddress) FullAddress() string {
	rv := objc.Send[string](m_.ID, objc.Sel("fullAddress"))
	return rv
}


// A string that represents a place’s full address
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddress/fulladdress
func (m_ MKAddress) SetFullAddress(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFullAddress:"), objc.String(value))
}


// A string that represents the short address of a location, such as it’s street address and city.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddress/shortaddress
func (m_ MKAddress) ShortAddress() string {
	rv := objc.Send[string](m_.ID, objc.Sel("shortAddress"))
	return rv
}


// A string that represents the short address of a location, such as it’s street address and city.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddress/shortaddress
func (m_ MKAddress) SetShortAddress(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShortAddress:"), objc.String(value))
}


