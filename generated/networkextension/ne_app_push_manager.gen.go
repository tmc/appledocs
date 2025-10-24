// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEAppPushManager */


/* debug [class_header]: Header for NEAppPushManager */
// The class instance for the [NEAppPushManager] class.
var (
	NEAppPushManagerClass     _NEAppPushManagerClass
	NEAppPushManagerClassOnce sync.Once
)

func getNEAppPushManagerClass() _NEAppPushManagerClass {
	NEAppPushManagerClassOnce.Do(func() {
		NEAppPushManagerClass = _NEAppPushManagerClass{objc.GetClass("NEAppPushManager")}
	})
	return NEAppPushManagerClass
}

type _NEAppPushManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEAppPushManager */
// An interface definition for the [NEAppPushManager] class.
type INEAppPushManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEAppPushManager */
	// properties:
	NEAppPushErrorDomain() objc.IObject /* cross-framework: NSString */
	IsActive() bool
	SetIsActive(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEAppPushManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEAppPushManager */
// Alloc allocates a new instance without initialization.
func (nc _NEAppPushManagerClass) Alloc() NEAppPushManager {
	rv := objc.Send[NEAppPushManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEAppPushManagerClass) New() NEAppPushManager {
	rv := objc.Send[NEAppPushManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppPushManager) Init() NEAppPushManager {
	rv := objc.Send[NEAppPushManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppPushManager) Autorelease() NEAppPushManager {
	rv := objc.Send[NEAppPushManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppPushManager creates a new NEAppPushManager instance.
func NewNEAppPushManager() NEAppPushManager {
	return getNEAppPushManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEAppPushManager */
// An object that configures a push provider and manages its life cycle.
//
// Your app can create as many instances as you need. Load your managers from the persistent store and set up their delegates immediately after the app launches, so they’re ready to handle incoming calls.


// An object that configures a push provider and manages its life cycle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager
type NEAppPushManager struct {
	objectivec.Object
}

// NEAppPushManagerFrom constructs a [NEAppPushManager] from an unsafe.Pointer.
//
// An object that configures a push provider and manages its life cycle.
func NEAppPushManagerFrom(ptr unsafe.Pointer) NEAppPushManager {
	return NEAppPushManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEAppPushManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEAppPushManager */

// Loads all saved manager configurations asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/loadAllFromPreferences(completionHandler:)
func (nc _NEAppPushManagerClass) LoadAllFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("loadAllFromPreferencesWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadAllFromPreferencesWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEAppPushManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEAppPushManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEAppPushManager */

// The error domain string for local push errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppusherrordomain
func (n_ NEAppPushManager) NEAppPushErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEAppPushErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: NEAppPushErrorDomain */


// A Boolean value that indicates whether a configuration is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/isactive
func (n_ NEAppPushManager) IsActive() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isActive"))
	return rv
}/* debug [instance_properties/getter]: isActive */


// A Boolean value that indicates whether a configuration is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/isactive
func (n_ NEAppPushManager) SetIsActive(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsActive:"), value)
}/* debug [instance_properties/setter]: isActive */


// A property you use to toggle enabling the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/isenabled
func (n_ NEAppPushManager) IsEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A property you use to toggle enabling the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/isenabled
func (n_ NEAppPushManager) SetIsEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEAppPushManager */


