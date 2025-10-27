// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [NEAppPushManager] class.
type INEAppPushManager interface {
	objectivec.IObject
	

	// properties:
	NEAppPushErrorDomain() foundation.foundation.INSString
	IsActive() bool
	SetIsActive(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)


	

	// methods:


}





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










// Loads all saved manager configurations asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/loadAllFromPreferences(completionHandler:)
func (nc _NEAppPushManagerClass) LoadAllFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("loadAllFromPreferencesWithCompletionHandler:"), completionHandler)
}

















// The error domain string for local push errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppusherrordomain
func (n_ NEAppPushManager) NEAppPushErrorDomain() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEAppPushErrorDomain"))
	return rv
}


// A Boolean value that indicates whether a configuration is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/isactive
func (n_ NEAppPushManager) IsActive() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isActive"))
	return rv
}


// A Boolean value that indicates whether a configuration is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/isactive
func (n_ NEAppPushManager) SetIsActive(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsActive:"), value)
}


// A property you use to toggle enabling the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/isenabled
func (n_ NEAppPushManager) IsEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isEnabled"))
	return rv
}


// A property you use to toggle enabling the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/isenabled
func (n_ NEAppPushManager) SetIsEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsEnabled:"), value)
}







