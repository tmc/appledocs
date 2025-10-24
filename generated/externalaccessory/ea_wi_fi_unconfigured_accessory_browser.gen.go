// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class EAWiFiUnconfiguredAccessoryBrowser */


/* debug [class_header]: Header for EAWiFiUnconfiguredAccessoryBrowser */
// The class instance for the [EAWiFiUnconfiguredAccessoryBrowser] class.
var (
	EAWiFiUnconfiguredAccessoryBrowserClass     _EAWiFiUnconfiguredAccessoryBrowserClass
	EAWiFiUnconfiguredAccessoryBrowserClassOnce sync.Once
)

func getEAWiFiUnconfiguredAccessoryBrowserClass() _EAWiFiUnconfiguredAccessoryBrowserClass {
	EAWiFiUnconfiguredAccessoryBrowserClassOnce.Do(func() {
		EAWiFiUnconfiguredAccessoryBrowserClass = _EAWiFiUnconfiguredAccessoryBrowserClass{objc.GetClass("EAWiFiUnconfiguredAccessoryBrowser")}
	})
	return EAWiFiUnconfiguredAccessoryBrowserClass
}

type _EAWiFiUnconfiguredAccessoryBrowserClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EAWiFiUnconfiguredAccessoryBrowser */
// An interface definition for the [EAWiFiUnconfiguredAccessoryBrowser] class.
type IEAWiFiUnconfiguredAccessoryBrowser interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for EAWiFiUnconfiguredAccessoryBrowser */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EAWiFiUnconfiguredAccessoryBrowser */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EAWiFiUnconfiguredAccessoryBrowser */
// Alloc allocates a new instance without initialization.
func (ec _EAWiFiUnconfiguredAccessoryBrowserClass) Alloc() EAWiFiUnconfiguredAccessoryBrowser {
	rv := objc.Send[EAWiFiUnconfiguredAccessoryBrowser](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EAWiFiUnconfiguredAccessoryBrowserClass) New() EAWiFiUnconfiguredAccessoryBrowser {
	rv := objc.Send[EAWiFiUnconfiguredAccessoryBrowser](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EAWiFiUnconfiguredAccessoryBrowser) Init() EAWiFiUnconfiguredAccessoryBrowser {
	rv := objc.Send[EAWiFiUnconfiguredAccessoryBrowser](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EAWiFiUnconfiguredAccessoryBrowser) Autorelease() EAWiFiUnconfiguredAccessoryBrowser {
	rv := objc.Send[EAWiFiUnconfiguredAccessoryBrowser](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEAWiFiUnconfiguredAccessoryBrowser creates a new EAWiFiUnconfiguredAccessoryBrowser instance.
func NewEAWiFiUnconfiguredAccessoryBrowser() EAWiFiUnconfiguredAccessoryBrowser {
	return getEAWiFiUnconfiguredAccessoryBrowserClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EAWiFiUnconfiguredAccessoryBrowser */
// An object you use to scan for wireless accessories and configure them for use with the user’s app.
//
// The class gives your app access to the MFi Wireless Accessory Configuration process. You use a browser object to scan for unconfigured accessories, connect them to the user’s Wi-Fi infrastructure, and configure attributes of the accessories. An accessory is represented by an instance of .


// An object you use to scan for wireless accessories and configure them for use with the user’s app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser
type EAWiFiUnconfiguredAccessoryBrowser struct {
	objectivec.Object
}

// EAWiFiUnconfiguredAccessoryBrowserFrom constructs a [EAWiFiUnconfiguredAccessoryBrowser] from an unsafe.Pointer.
//
// An object you use to scan for wireless accessories and configure them for use with the user’s app.
func EAWiFiUnconfiguredAccessoryBrowserFrom(ptr unsafe.Pointer) EAWiFiUnconfiguredAccessoryBrowser {
	return EAWiFiUnconfiguredAccessoryBrowser{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EAWiFiUnconfiguredAccessoryBrowser */

// Creates a browser object that scans for unconfigured accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/init(delegate:queue:)
func NewEAWiFiUnconfiguredAccessoryBrowserWithDelegateQueue(delegate unsafe.Pointer, queue unsafe.Pointer) EAWiFiUnconfiguredAccessoryBrowser {
	instance := getEAWiFiUnconfiguredAccessoryBrowserClass().Alloc()
	rv := objc.Send[EAWiFiUnconfiguredAccessoryBrowser](instance.ID, objc.Sel("initWithDelegate:queue:"), delegate, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewEAWiFiUnconfiguredAccessoryBrowserWithDelegateQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EAWiFiUnconfiguredAccessoryBrowser */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EAWiFiUnconfiguredAccessoryBrowser */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EAWiFiUnconfiguredAccessoryBrowser */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EAWiFiUnconfiguredAccessoryBrowser */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EAWiFiUnconfiguredAccessoryBrowser */


