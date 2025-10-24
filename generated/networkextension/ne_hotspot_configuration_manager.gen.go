// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEHotspotConfigurationManager */


/* debug [class_header]: Header for NEHotspotConfigurationManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEHotspotConfigurationManager */
// An interface definition for the [NEHotspotConfigurationManager] class.
type INEHotspotConfigurationManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEHotspotConfigurationManager */
	// properties:
	NEHotspotConfigurationErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEHotspotConfigurationManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEHotspotConfigurationManager */
// Alloc allocates a new instance without initialization.
func (nc _NEHotspotConfigurationManagerClass) Alloc() NEHotspotConfigurationManager {
	rv := objc.Send[NEHotspotConfigurationManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEHotspotConfigurationManager */
// A manager that applies and removes hotspot configurations of Wi-Fi networks.
//
// When your app creates a new hotspot configuration using and applies it to a Wi-Fi network or attempts to update a previously configured network, the device prompts the user for approval. Without explicit user consent, your app can’t make configuration changes. Your app can use or to delete a configuration that it has added, but not a configuration added by another app or user. The user can also delete configured networks using Settings > Wi-Fi. When your app is uninstalled, iOS removes the configurations of all networks your app has configured, including their keychain entries. Hotspot Configuration Manager errors are listed in .


// A manager that applies and removes hotspot configurations of Wi-Fi networks.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEHotspotConfigurationManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEHotspotConfigurationManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEHotspotConfigurationManager */

// Instantiates as a singleton, so it can be shared.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationManager/shared
func (nc _NEHotspotConfigurationManagerClass) SharedManager() NEHotspotConfigurationManager {
	rv := objc.Send[NEHotspotConfigurationManager](objc.ID(nc.class), objc.Sel("sharedManager"))
	return rv
}/* debug [class_properties_class/property]: sharedManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEHotspotConfigurationManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEHotspotConfigurationManager */

// The domain string for errors involving hotspot configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nehotspotconfigurationerrordomain
func (n_ NEHotspotConfigurationManager) NEHotspotConfigurationErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEHotspotConfigurationErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: NEHotspotConfigurationErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEHotspotConfigurationManager */


