// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// An object that configures a push provider and manages its life cycle.
//
// Your app can create as many instances as you need. Load your managers from the persistent store and set up their delegates immediately after the app launches, so they’re ready to handle incoming calls.
//
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

// Alloc allocates a new instance without initialization.
func (nc _NEAppPushManagerClass) Alloc() NEAppPushManager {
	rv := objc.Send[NEAppPushManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A delegate that receives incoming call information from the provider.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/delegate
func (n_ NEAppPushManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate that receives incoming call information from the provider.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/delegate
func (n_ NEAppPushManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDelegate:"), value)
}

// An array of Wi-Fi SSID strings that the system matches for local push activation.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/matchSSIDs
func (n_ NEAppPushManager) MatchSSIDs() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("matchSSIDs"))
	return rv
}


// SetMatchSSIDs sets the value of the matchSSIDs property.
// An array of Wi-Fi SSID strings that the system matches for local push activation.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/matchSSIDs
func (n_ NEAppPushManager) SetMatchSSIDs(value []string) {
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
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchSSIDs:"), nsArray)
}

// A string that contains the bundle identifier of the push provider.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/providerBundleIdentifier
func (n_ NEAppPushManager) ProviderBundleIdentifier() string {
	rv := objc.Send[string](n_.ID, objc.Sel("providerBundleIdentifier"))
	return rv
}


// SetProviderBundleIdentifier sets the value of the providerBundleIdentifier property.
// A string that contains the bundle identifier of the push provider.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/providerBundleIdentifier
func (n_ NEAppPushManager) SetProviderBundleIdentifier(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderBundleIdentifier:"), objc.String(value))
}



