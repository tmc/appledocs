// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEHotspotHelper */


/* debug [class_header]: Header for NEHotspotHelper */
// The class instance for the [NEHotspotHelper] class.
var (
	NEHotspotHelperClass     _NEHotspotHelperClass
	NEHotspotHelperClassOnce sync.Once
)

func getNEHotspotHelperClass() _NEHotspotHelperClass {
	NEHotspotHelperClassOnce.Do(func() {
		NEHotspotHelperClass = _NEHotspotHelperClass{objc.GetClass("NEHotspotHelper")}
	})
	return NEHotspotHelperClass
}

type _NEHotspotHelperClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEHotspotHelper */
// An interface definition for the [NEHotspotHelper] class.
type INEHotspotHelper interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEHotspotHelper */
	// properties:
	KNEHotspotHelperOptionDisplayName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEHotspotHelper */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEHotspotHelper */
// Alloc allocates a new instance without initialization.
func (nc _NEHotspotHelperClass) Alloc() NEHotspotHelper {
	rv := objc.Send[NEHotspotHelper](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEHotspotHelperClass) New() NEHotspotHelper {
	rv := objc.Send[NEHotspotHelper](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEHotspotHelper) Init() NEHotspotHelper {
	rv := objc.Send[NEHotspotHelper](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEHotspotHelper) Autorelease() NEHotspotHelper {
	rv := objc.Send[NEHotspotHelper](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEHotspotHelper creates a new NEHotspotHelper instance.
func NewNEHotspotHelper() NEHotspotHelper {
	return getNEHotspotHelperClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEHotspotHelper */
// A class to register a hotspot helper.
//
// The API gives your app the ability to perform custom authentication for Wi-Fi Hotspots. It gives users a way to seamlessly connect to a large aggregated network of Wi-Fi Hotspots. The API lets your app configure those hotspots.


// A class to register a hotspot helper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelper
type NEHotspotHelper struct {
	objectivec.Object
}

// NEHotspotHelperFrom constructs a [NEHotspotHelper] from an unsafe.Pointer.
//
// A class to register a hotspot helper.
func NEHotspotHelperFrom(ptr unsafe.Pointer) NEHotspotHelper {
	return NEHotspotHelper{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEHotspotHelper *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEHotspotHelper */

// Terminate the authentication session for a Hotspot network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelper/logoff(_:)
func (nc _NEHotspotHelperClass) Logoff(network INEHotspotNetwork) bool {
	rv := objc.Send[bool](objc.ID(nc.class), objc.Sel("logoff:"), network)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Logoff) */


// Register the application as a Hotspot Helper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelper/register(options:queue:handler:)
func (nc _NEHotspotHelperClass) RegisterWithOptionsQueueHandler(options foundation.IDictionary, queue objectivec.IObject, handler objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(nc.class), objc.Sel("registerWithOptions:queue:handler:"), options, queue, handler)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterWithOptionsQueueHandler) */


// Return the list of network interfaces managed by the Hotspot Helper infrastructure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelper/supportedNetworkInterfaces()
func (nc _NEHotspotHelperClass) SupportedNetworkInterfaces() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(nc.class), objc.Sel("supportedNetworkInterfaces"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedNetworkInterfaces) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEHotspotHelper */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEHotspotHelper */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEHotspotHelper */

// The string displayed in Wi-Fi Settings for a network handled by the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/knehotspothelperoptiondisplayname
func (n_ NEHotspotHelper) KNEHotspotHelperOptionDisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("kNEHotspotHelperOptionDisplayName"))
	return rv
}/* debug [instance_properties/getter]: kNEHotspotHelperOptionDisplayName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEHotspotHelper */


