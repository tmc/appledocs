// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	FullAddress() objc.IObject /* cross-framework: NSString */
	SetFullAddress(value objc.IObject /* cross-framework: NSString */)
	ShortAddress() objc.IObject /* cross-framework: NSString */
	SetShortAddress(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

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



// A string that represents a place’s full address
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddress/fulladdress
func (m_ MKAddress) FullAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("fullAddress"))
	return rv
}


// A string that represents a place’s full address
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddress/fulladdress
func (m_ MKAddress) SetFullAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFullAddress:"), value)
}


// A string that represents the short address of a location, such as it’s street address and city.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddress/shortaddress
func (m_ MKAddress) ShortAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("shortAddress"))
	return rv
}


// A string that represents the short address of a location, such as it’s street address and city.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddress/shortaddress
func (m_ MKAddress) SetShortAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShortAddress:"), value)
}



