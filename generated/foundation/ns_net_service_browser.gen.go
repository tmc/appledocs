// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSNetServiceBrowser */


/* debug [class_header]: Header for NSNetServiceBrowser */
// The class instance for the [NetServiceBrowser] class.
var (
	NetServiceBrowserClass     _NetServiceBrowserClass
	NetServiceBrowserClassOnce sync.Once
)

func getNetServiceBrowserClass() _NetServiceBrowserClass {
	NetServiceBrowserClassOnce.Do(func() {
		NetServiceBrowserClass = _NetServiceBrowserClass{objc.GetClass("NSNetServiceBrowser")}
	})
	return NetServiceBrowserClass
}

type _NetServiceBrowserClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NetServiceBrowser */
// An interface definition for the [NetServiceBrowser] class.
type INetServiceBrowser interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NetServiceBrowser */
	// properties:
	IncludesPeerToPeer() bool
	SetIncludesPeerToPeer(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NetServiceBrowser */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NetServiceBrowser */
// Alloc allocates a new instance without initialization.
func (nc _NetServiceBrowserClass) Alloc() NetServiceBrowser {
	rv := objc.Send[NetServiceBrowser](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NetServiceBrowserClass) New() NetServiceBrowser {
	rv := objc.Send[NetServiceBrowser](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NetServiceBrowser) Init() NetServiceBrowser {
	rv := objc.Send[NetServiceBrowser](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NetServiceBrowser) Autorelease() NetServiceBrowser {
	rv := objc.Send[NetServiceBrowser](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNetServiceBrowser creates a new NetServiceBrowser instance.
func NewNetServiceBrowser() NetServiceBrowser {
	return getNetServiceBrowserClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NetServiceBrowser */
// A network service browser that finds published services on a network using multicast DNS.
//
// Services can range from standard services, such as HTTP and FTP, to custom services defined by other applications. You can use a network service browser in your code to obtain the list of accessible domains and then to obtain an object for each discovered service. Each network service browser performs one search at a time, so if you want to perform multiple simultaneous searches, use multiple network service browsers. A network service browser performs all searches asynchronously using the current run loop to execute the search in the background. Results from a search are returned through the associated delegate object, which your client application must provide. Searching proceeds in the background until the object receives a message. To use an object to search for services, allocate it, initialize it, and assign a delegate. (If you wish, you can also use the and methods to execute searches on a run loop other than the current one.) Once your object is ready, you begin by gathering the list of accessible domains using either the or methods. From the list of returned domains, you can pick one and use the method to search for services in that domain. The class provides two ways to search for domains. In most cases, your client should use the method to search only for local domains to which the host machine has registration authority. This is the preferred method for accessing domains as it guarantees that the host machine can connect to services in the returned domains. Access to domains outside this list may be more limited.


// A network service browser that finds published services on a network using multicast DNS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser
type NetServiceBrowser struct {
	objectivec.Object
}

// NetServiceBrowserFrom constructs a [NetServiceBrowser] from an unsafe.Pointer.
//
// A network service browser that finds published services on a network using multicast DNS.
func NetServiceBrowserFrom(ptr unsafe.Pointer) NetServiceBrowser {
	return NetServiceBrowser{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NetServiceBrowser */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NetServiceBrowser */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NetServiceBrowser */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NetServiceBrowser */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NetServiceBrowser */

// Whether to browse over peer-to-peer Bluetooth and Wi-Fi, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/includesPeerToPeer
func (n_ NetServiceBrowser) IncludesPeerToPeer() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("includesPeerToPeer"))
	return rv
}/* debug [instance_properties/getter]: includesPeerToPeer */


// Whether to browse over peer-to-peer Bluetooth and Wi-Fi, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/includesPeerToPeer
func (n_ NetServiceBrowser) SetIncludesPeerToPeer(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludesPeerToPeer:"), value)
}/* debug [instance_properties/setter]: includesPeerToPeer */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSNetServiceBrowser */


