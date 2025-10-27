// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [NETransparentProxyManager] class.
var (
	NETransparentProxyManagerClass     _NETransparentProxyManagerClass
	NETransparentProxyManagerClassOnce sync.Once
)

func getNETransparentProxyManagerClass() _NETransparentProxyManagerClass {
	NETransparentProxyManagerClassOnce.Do(func() {
		NETransparentProxyManagerClass = _NETransparentProxyManagerClass{objc.GetClass("NETransparentProxyManager")}
	})
	return NETransparentProxyManagerClass
}

type _NETransparentProxyManagerClass struct {
	class objc.Class
}





// An interface definition for the [NETransparentProxyManager] class.
type INETransparentProxyManager interface {
	INEVPNManager
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NETransparentProxyManagerClass) Alloc() NETransparentProxyManager {
	rv := objc.Send[NETransparentProxyManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NETransparentProxyManagerClass) New() NETransparentProxyManager {
	rv := objc.Send[NETransparentProxyManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NETransparentProxyManager) Init() NETransparentProxyManager {
	rv := objc.Send[NETransparentProxyManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NETransparentProxyManager) Autorelease() NETransparentProxyManager {
	rv := objc.Send[NETransparentProxyManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETransparentProxyManager creates a new NETransparentProxyManager instance.
func NewNETransparentProxyManager() NETransparentProxyManager {
	return getNETransparentProxyManagerClass().New()
}





// An object that configures and controls transparent proxies.


// An object that configures and controls transparent proxies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyManager
type NETransparentProxyManager struct {
	NEVPNManager
}

// NETransparentProxyManagerFrom constructs a [NETransparentProxyManager] from an unsafe.Pointer.
//
// An object that configures and controls transparent proxies.
func NETransparentProxyManagerFrom(ptr unsafe.Pointer) NETransparentProxyManager {
	return NETransparentProxyManager{
		NEVPNManager: NEVPNManagerFrom(ptr),
	}
}










// Loads all previously-saved transparent proxy configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyManager/loadAllFromPreferences(completionHandler:)
func (nc _NETransparentProxyManagerClass) LoadAllFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("loadAllFromPreferencesWithCompletionHandler:"), completionHandler)
}























