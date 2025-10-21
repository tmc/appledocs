// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Navigation] class.
var (
	NavigationClass     _NavigationClass
	NavigationClassOnce sync.Once
)

func getNavigationClass() _NavigationClass {
	NavigationClassOnce.Do(func() {
		NavigationClass = _NavigationClass{objc.GetClass("WKNavigation")}
	})
	return NavigationClass
}

type _NavigationClass struct {
	class objc.Class
}

// An interface definition for the [Navigation] class.
type INavigation interface {
	objectivec.IObject
}

// An object that tracks the loading progress of a webpage.
//
// A object uniquely identifies a load request for a webpage. When you ask a web view to load content or navigate to a page, the web view returns a object that identifies your request. As the load operation progresses, the web view reports progress of that operation to various methods of its navigation delegate, passing them the matching object.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigation
type Navigation struct {
	objectivec.Object
}

// NavigationFrom constructs a [Navigation] from an unsafe.Pointer.
//
// An object that tracks the loading progress of a webpage.
func NavigationFrom(ptr unsafe.Pointer) Navigation {
	return Navigation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NavigationClass) Alloc() Navigation {
	rv := objc.Send[Navigation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NavigationClass) New() Navigation {
	rv := objc.Send[Navigation](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ Navigation) Init() Navigation {
	rv := objc.Send[Navigation](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ Navigation) Autorelease() Navigation {
	rv := objc.Send[Navigation](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNavigation creates a new Navigation instance.
func NewNavigation() Navigation {
	return getNavigationClass().New()
}


// The content mode WebKit uses to load the webpage.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigation/effectiveContentMode
func (n_ Navigation) EffectiveContentMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("effectiveContentMode"))
	return rv
}



