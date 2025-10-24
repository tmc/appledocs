// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Carrier] class.
var (
	CarrierClass     _CarrierClass
	CarrierClassOnce sync.Once
)

func getCarrierClass() _CarrierClass {
	CarrierClassOnce.Do(func() {
		CarrierClass = _CarrierClass{objc.GetClass("CTCarrier")}
	})
	return CarrierClass
}

type _CarrierClass struct {
	class objc.Class
}

// An interface definition for the [Carrier] class.
type ICarrier interface {
	objectivec.IObject
	// properties:
	AllowsVOIP() bool
	SetAllowsVOIP(value bool)
	IsoCountryCode() objc.IObject /* cross-framework: NSString */
	SetIsoCountryCode(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// Information about the user’s cellular service provider, such as its unique identifier and whether it allows VoIP calls on its network.


// Information about the user’s cellular service provider, such as its unique identifier and whether it allows VoIP calls on its network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCarrier
type Carrier struct {
	objectivec.Object
}

// CarrierFrom constructs a [Carrier] from an unsafe.Pointer.
//
// Information about the user’s cellular service provider, such as its unique identifier and whether it allows VoIP calls on its network.
func CarrierFrom(ptr unsafe.Pointer) Carrier {
	return Carrier{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CarrierClass) Alloc() Carrier {
	rv := objc.Send[Carrier](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CarrierClass) New() Carrier {
	rv := objc.Send[Carrier](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Carrier) Init() Carrier {
	rv := objc.Send[Carrier](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Carrier) Autorelease() Carrier {
	rv := objc.Send[Carrier](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCarrier creates a new Carrier instance.
func NewCarrier() Carrier {
	return getCarrierClass().New()
}



// Indicates if the carrier allows making VoIP calls on its network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcarrier/allowsvoip
func (c_ Carrier) AllowsVOIP() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsVOIP"))
	return rv
}


// Indicates if the carrier allows making VoIP calls on its network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcarrier/allowsvoip
func (c_ Carrier) SetAllowsVOIP(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsVOIP:"), value)
}


// The ISO country code for the user’s cellular service provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcarrier/isocountrycode
func (c_ Carrier) IsoCountryCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("isoCountryCode"))
	return rv
}


// The ISO country code for the user’s cellular service provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcarrier/isocountrycode
func (c_ Carrier) SetIsoCountryCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsoCountryCode:"), value)
}


