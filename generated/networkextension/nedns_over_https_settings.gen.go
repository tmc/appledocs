// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEDNSOverHTTPSSettings] class.
var (
	NEDNSOverHTTPSSettingsClass     _NEDNSOverHTTPSSettingsClass
	NEDNSOverHTTPSSettingsClassOnce sync.Once
)

func getNEDNSOverHTTPSSettingsClass() _NEDNSOverHTTPSSettingsClass {
	NEDNSOverHTTPSSettingsClassOnce.Do(func() {
		NEDNSOverHTTPSSettingsClass = _NEDNSOverHTTPSSettingsClass{objc.GetClass("NEDNSOverHTTPSSettings")}
	})
	return NEDNSOverHTTPSSettingsClass
}

type _NEDNSOverHTTPSSettingsClass struct {
	class objc.Class
}

// An interface definition for the [NEDNSOverHTTPSSettings] class.
type INEDNSOverHTTPSSettings interface {
	INEDNSSettings
}

// The DNS resolver settings for a DNS-over-HTTPS server.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverHTTPSSettings
type NEDNSOverHTTPSSettings struct {
	NEDNSSettings
}

// NEDNSOverHTTPSSettingsFrom constructs a [NEDNSOverHTTPSSettings] from an unsafe.Pointer.
//
// The DNS resolver settings for a DNS-over-HTTPS server.
func NEDNSOverHTTPSSettingsFrom(ptr unsafe.Pointer) NEDNSOverHTTPSSettings {
	return NEDNSOverHTTPSSettings{
		NEDNSSettings: NEDNSSettingsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEDNSOverHTTPSSettingsClass) Alloc() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEDNSOverHTTPSSettingsClass) New() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEDNSOverHTTPSSettings) Init() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEDNSOverHTTPSSettings) Autorelease() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSOverHTTPSSettings creates a new NEDNSOverHTTPSSettings instance.
func NewNEDNSOverHTTPSSettings() NEDNSOverHTTPSSettings {
	return getNEDNSOverHTTPSSettingsClass().New()
}




