// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKNavigation */

/* debug [class_header]: Header for WKNavigation */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for Navigation */
// An interface definition for the [Navigation] class.
type INavigation interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for Navigation */
	// properties:
	EffectiveContentMode() ContentMode
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for Navigation */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for Navigation */
// Alloc allocates a new instance without initialization.
func (nc _NavigationClass) Alloc() Navigation {
	rv := objc.Send[Navigation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for Navigation */
// An object that tracks the loading progress of a webpage.
//
// A object uniquely identifies a load request for a webpage. When you ask a web view to load content or navigate to a page, the web view returns a object that identifies your request. As the load operation progresses, the web view reports progress of that operation to various methods of its navigation delegate, passing them the matching object.

// An object that tracks the loading progress of a webpage.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for Navigation */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for Navigation */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for Navigation */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for Navigation */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for Navigation */

// The content mode WebKit uses to load the webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigation/effectiveContentMode
func (n_ Navigation) EffectiveContentMode() ContentMode {
	rv := objc.Send[ContentMode](n_.ID, objc.Sel("effectiveContentMode"))
	return rv
} /* debug [instance_properties/getter]: effectiveContentMode */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKNavigation */
