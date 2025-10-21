// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEDNSOverTLSSettings] class.
var (
	NEDNSOverTLSSettingsClass     _NEDNSOverTLSSettingsClass
	NEDNSOverTLSSettingsClassOnce sync.Once
)

func getNEDNSOverTLSSettingsClass() _NEDNSOverTLSSettingsClass {
	NEDNSOverTLSSettingsClassOnce.Do(func() {
		NEDNSOverTLSSettingsClass = _NEDNSOverTLSSettingsClass{objc.GetClass("NEDNSOverTLSSettings")}
	})
	return NEDNSOverTLSSettingsClass
}

type _NEDNSOverTLSSettingsClass struct {
	class objc.Class
}

// An interface definition for the [NEDNSOverTLSSettings] class.
type INEDNSOverTLSSettings interface {
	INEDNSSettings
}

// The DNS resolver settings for a DNS-over-TLS server.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings
type NEDNSOverTLSSettings struct {
	NEDNSSettings
}

// NEDNSOverTLSSettingsFrom constructs a [NEDNSOverTLSSettings] from an unsafe.Pointer.
//
// The DNS resolver settings for a DNS-over-TLS server.
func NEDNSOverTLSSettingsFrom(ptr unsafe.Pointer) NEDNSOverTLSSettings {
	return NEDNSOverTLSSettings{
		NEDNSSettings: NEDNSSettingsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEDNSOverTLSSettingsClass) Alloc() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEDNSOverTLSSettingsClass) New() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEDNSOverTLSSettings) Init() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEDNSOverTLSSettings) Autorelease() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSOverTLSSettings creates a new NEDNSOverTLSSettings instance.
func NewNEDNSOverTLSSettings() NEDNSOverTLSSettings {
	return getNEDNSOverTLSSettingsClass().New()
}




