// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEHotspotConfigurationManager] class.
var (
	NEHotspotConfigurationManagerClass     _NEHotspotConfigurationManagerClass
	NEHotspotConfigurationManagerClassOnce sync.Once
)

func getNEHotspotConfigurationManagerClass() _NEHotspotConfigurationManagerClass {
	NEHotspotConfigurationManagerClassOnce.Do(func() {
		NEHotspotConfigurationManagerClass = _NEHotspotConfigurationManagerClass{objc.GetClass("NEHotspotConfigurationManager")}
	})
	return NEHotspotConfigurationManagerClass
}

type _NEHotspotConfigurationManagerClass struct {
	class objc.Class
}

// An interface definition for the [NEHotspotConfigurationManager] class.
type INEHotspotConfigurationManager interface {
	objectivec.IObject
}

// A manager that applies and removes hotspot configurations of Wi-Fi networks.
//
// When your app creates a new hotspot configuration using and applies it to a Wi-Fi network or attempts to update a previously configured network, the device prompts the user for approval. Without explicit user consent, your app can’t make configuration changes. Your app can use or to delete a configuration that it has added, but not a configuration added by another app or user. The user can also delete configured networks using Settings > Wi-Fi. When your app is uninstalled, iOS removes the configurations of all networks your app has configured, including their keychain entries. Hotspot Configuration Manager errors are listed in .
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationManager
type NEHotspotConfigurationManager struct {
	objectivec.Object
}

// NEHotspotConfigurationManagerFrom constructs a [NEHotspotConfigurationManager] from an unsafe.Pointer.
//
// A manager that applies and removes hotspot configurations of Wi-Fi networks.
func NEHotspotConfigurationManagerFrom(ptr unsafe.Pointer) NEHotspotConfigurationManager {
	return NEHotspotConfigurationManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEHotspotConfigurationManagerClass) Alloc() NEHotspotConfigurationManager {
	rv := objc.Send[NEHotspotConfigurationManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEHotspotConfigurationManagerClass) New() NEHotspotConfigurationManager {
	rv := objc.Send[NEHotspotConfigurationManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEHotspotConfigurationManager) Init() NEHotspotConfigurationManager {
	rv := objc.Send[NEHotspotConfigurationManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEHotspotConfigurationManager) Autorelease() NEHotspotConfigurationManager {
	rv := objc.Send[NEHotspotConfigurationManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEHotspotConfigurationManager creates a new NEHotspotConfigurationManager instance.
func NewNEHotspotConfigurationManager() NEHotspotConfigurationManager {
	return getNEHotspotConfigurationManagerClass().New()
}


// The domain string for errors involving hotspot configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotconfigurationerrordomain
func (n_ NEHotspotConfigurationManager) NEHotspotConfigurationErrorDomain() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("NEHotspotConfigurationErrorDomain"))
	return rv
}



