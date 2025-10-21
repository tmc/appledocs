// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEIPv6Settings] class.
var (
	NEIPv6SettingsClass     _NEIPv6SettingsClass
	NEIPv6SettingsClassOnce sync.Once
)

func getNEIPv6SettingsClass() _NEIPv6SettingsClass {
	NEIPv6SettingsClassOnce.Do(func() {
		NEIPv6SettingsClass = _NEIPv6SettingsClass{objc.GetClass("NEIPv6Settings")}
	})
	return NEIPv6SettingsClass
}

type _NEIPv6SettingsClass struct {
	class objc.Class
}

// An interface definition for the [NEIPv6Settings] class.
type INEIPv6Settings interface {
	objectivec.IObject
}

// The IPv6 settings of an IP layer network tunnel.
//
// To specify the IPv6 settings of a packet tunnel, set its . property to an instance of this class.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings
type NEIPv6Settings struct {
	objectivec.Object
}

// NEIPv6SettingsFrom constructs a [NEIPv6Settings] from an unsafe.Pointer.
//
// The IPv6 settings of an IP layer network tunnel.
func NEIPv6SettingsFrom(ptr unsafe.Pointer) NEIPv6Settings {
	return NEIPv6Settings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEIPv6SettingsClass) Alloc() NEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEIPv6SettingsClass) New() NEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEIPv6Settings) Init() NEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEIPv6Settings) Autorelease() NEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEIPv6Settings creates a new NEIPv6Settings instance.
func NewNEIPv6Settings() NEIPv6Settings {
	return getNEIPv6SettingsClass().New()
}


// The IPv6 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/excludedRoutes
func (n_ NEIPv6Settings) ExcludedRoutes() []NEIPv6Route {
	rv := objc.Send[[]NEIPv6Route](n_.ID, objc.Sel("excludedRoutes"))
	return rv
}


// SetExcludedRoutes sets the value of the excludedRoutes property.
// The IPv6 network traffic that the system routes to the primary physical interface, not the TUN interface.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/excludedRoutes
func (n_ NEIPv6Settings) SetExcludedRoutes(value []NEIPv6Route) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedRoutes:"), nsArray)
}
// The IPv6 network traffic that the system routes to the TUN interface.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/includedRoutes
func (n_ NEIPv6Settings) IncludedRoutes() []NEIPv6Route {
	rv := objc.Send[[]NEIPv6Route](n_.ID, objc.Sel("includedRoutes"))
	return rv
}


// SetIncludedRoutes sets the value of the includedRoutes property.
// The IPv6 network traffic that the system routes to the TUN interface.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/includedRoutes
func (n_ NEIPv6Settings) SetIncludedRoutes(value []NEIPv6Route) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludedRoutes:"), nsArray)
}


