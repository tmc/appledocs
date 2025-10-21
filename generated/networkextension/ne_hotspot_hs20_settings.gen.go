// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [NEHotspotHS20Settings] class.
var (
	NEHotspotHS20SettingsClass     _NEHotspotHS20SettingsClass
	NEHotspotHS20SettingsClassOnce sync.Once
)

func getNEHotspotHS20SettingsClass() _NEHotspotHS20SettingsClass {
	NEHotspotHS20SettingsClassOnce.Do(func() {
		NEHotspotHS20SettingsClass = _NEHotspotHS20SettingsClass{objc.GetClass("NEHotspotHS20Settings")}
	})
	return NEHotspotHS20SettingsClass
}

type _NEHotspotHS20SettingsClass struct {
	class objc.Class
}

// An interface definition for the [NEHotspotHS20Settings] class.
type INEHotspotHS20Settings interface {
	objectivec.IObject
}

// Settings for configuring Hotspot 2.0 Wi-Fi networks.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHS20Settings
type NEHotspotHS20Settings struct {
	objectivec.Object
}

// NEHotspotHS20SettingsFrom constructs a [NEHotspotHS20Settings] from an unsafe.Pointer.
//
// Settings for configuring Hotspot 2.0 Wi-Fi networks.
func NEHotspotHS20SettingsFrom(ptr unsafe.Pointer) NEHotspotHS20Settings {
	return NEHotspotHS20Settings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEHotspotHS20SettingsClass) Alloc() NEHotspotHS20Settings {
	rv := objc.Send[NEHotspotHS20Settings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEHotspotHS20SettingsClass) New() NEHotspotHS20Settings {
	rv := objc.Send[NEHotspotHS20Settings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEHotspotHS20Settings) Init() NEHotspotHS20Settings {
	rv := objc.Send[NEHotspotHS20Settings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEHotspotHS20Settings) Autorelease() NEHotspotHS20Settings {
	rv := objc.Send[NEHotspotHS20Settings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEHotspotHS20Settings creates a new NEHotspotHS20Settings instance.
func NewNEHotspotHS20Settings() NEHotspotHS20Settings {
	return getNEHotspotHS20SettingsClass().New()
}




