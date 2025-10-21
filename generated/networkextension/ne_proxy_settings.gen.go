// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEProxySettings] class.
var (
	NEProxySettingsClass     _NEProxySettingsClass
	NEProxySettingsClassOnce sync.Once
)

func getNEProxySettingsClass() _NEProxySettingsClass {
	NEProxySettingsClassOnce.Do(func() {
		NEProxySettingsClass = _NEProxySettingsClass{objc.GetClass("NEProxySettings")}
	})
	return NEProxySettingsClass
}

type _NEProxySettingsClass struct {
	class objc.Class
}

// An interface definition for the [NEProxySettings] class.
type INEProxySettings interface {
	objectivec.IObject
}

// contains HTTP proxy settings.
//
// is used in the context of a VPN configuration to specify the proxy that should be used for network traffic when the VPN is active. Instances of this class are thread safe.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings
type NEProxySettings struct {
	objectivec.Object
}

// NEProxySettingsFrom constructs a [NEProxySettings] from an unsafe.Pointer.
//
// contains HTTP proxy settings.
func NEProxySettingsFrom(ptr unsafe.Pointer) NEProxySettings {
	return NEProxySettings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEProxySettingsClass) Alloc() NEProxySettings {
	rv := objc.Send[NEProxySettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEProxySettingsClass) New() NEProxySettings {
	rv := objc.Send[NEProxySettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEProxySettings) Init() NEProxySettings {
	rv := objc.Send[NEProxySettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEProxySettings) Autorelease() NEProxySettings {
	rv := objc.Send[NEProxySettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEProxySettings creates a new NEProxySettings instance.
func NewNEProxySettings() NEProxySettings {
	return getNEProxySettingsClass().New()
}




