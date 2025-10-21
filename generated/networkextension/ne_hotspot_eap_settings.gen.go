// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEHotspotEAPSettings] class.
var (
	NEHotspotEAPSettingsClass     _NEHotspotEAPSettingsClass
	NEHotspotEAPSettingsClassOnce sync.Once
)

func getNEHotspotEAPSettingsClass() _NEHotspotEAPSettingsClass {
	NEHotspotEAPSettingsClassOnce.Do(func() {
		NEHotspotEAPSettingsClass = _NEHotspotEAPSettingsClass{objc.GetClass("NEHotspotEAPSettings")}
	})
	return NEHotspotEAPSettingsClass
}

type _NEHotspotEAPSettingsClass struct {
	class objc.Class
}

// An interface definition for the [NEHotspotEAPSettings] class.
type INEHotspotEAPSettings interface {
	objectivec.IObject
}

// Extensible Authentication Protocol settings for configuring WPA and WPA2 enterprise Wi-Fi networks.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings
type NEHotspotEAPSettings struct {
	objectivec.Object
}

// NEHotspotEAPSettingsFrom constructs a [NEHotspotEAPSettings] from an unsafe.Pointer.
//
// Extensible Authentication Protocol settings for configuring WPA and WPA2 enterprise Wi-Fi networks.
func NEHotspotEAPSettingsFrom(ptr unsafe.Pointer) NEHotspotEAPSettings {
	return NEHotspotEAPSettings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEHotspotEAPSettingsClass) Alloc() NEHotspotEAPSettings {
	rv := objc.Send[NEHotspotEAPSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEHotspotEAPSettingsClass) New() NEHotspotEAPSettings {
	rv := objc.Send[NEHotspotEAPSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEHotspotEAPSettings) Init() NEHotspotEAPSettings {
	rv := objc.Send[NEHotspotEAPSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEHotspotEAPSettings) Autorelease() NEHotspotEAPSettings {
	rv := objc.Send[NEHotspotEAPSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEHotspotEAPSettings creates a new NEHotspotEAPSettings instance.
func NewNEHotspotEAPSettings() NEHotspotEAPSettings {
	return getNEHotspotEAPSettingsClass().New()
}




