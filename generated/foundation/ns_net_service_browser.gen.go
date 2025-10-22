// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [NetServiceBrowser] class.
type INetServiceBrowser interface {
	objectivec.IObject
	RemoveFromRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode)
	ScheduleInRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode)
	SearchForBrowsableDomains()
	SearchForRegistrationDomains()
	SearchForServicesOfTypeInDomain(type_ string, domainString string)
	Stop()
}

// A network service browser that finds published services on a network using multicast DNS.
//
// Services can range from standard services, such as HTTP and FTP, to custom services defined by other applications. You can use a network service browser in your code to obtain the list of accessible domains and then to obtain an object for each discovered service. Each network service browser performs one search at a time, so if you want to perform multiple simultaneous searches, use multiple network service browsers. A network service browser performs all searches asynchronously using the current run loop to execute the search in the background. Results from a search are returned through the associated delegate object, which your client application must provide. Searching proceeds in the background until the object receives a message. To use an object to search for services, allocate it, initialize it, and assign a delegate. (If you wish, you can also use the and methods to execute searches on a run loop other than the current one.) Once your object is ready, you begin by gathering the list of accessible domains using either the or methods. From the list of returned domains, you can pick one and use the method to search for services in that domain. The class provides two ways to search for domains. In most cases, your client should use the method to search only for local domains to which the host machine has registration authority. This is the preferred method for accessing domains as it guarantees that the host machine can connect to services in the returned domains. Access to domains outside this list may be more limited.
//
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

// Alloc allocates a new instance without initialization.
func (nc _NetServiceBrowserClass) Alloc() NetServiceBrowser {
	rv := objc.Send[NetServiceBrowser](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Removes the receiver from the specified run loop.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/remove(from:forMode:)
func (n_ NetServiceBrowser) RemoveFromRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeFromRunLoop:forMode:"), aRunLoop, mode)
}

// Adds the receiver to the specified run loop.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/schedule(in:forMode:)
func (n_ NetServiceBrowser) ScheduleInRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.Send[objc.ID](n_.ID, objc.Sel("scheduleInRunLoop:forMode:"), aRunLoop, mode)
}

// Initiates a search for domains visible to the host. This method returns immediately.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/searchForBrowsableDomains()
func (n_ NetServiceBrowser) SearchForBrowsableDomains() {
	objc.Send[objc.ID](n_.ID, objc.Sel("searchForBrowsableDomains"))
}

// Initiates a search for domains in which the host may register services.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/searchForRegistrationDomains()
func (n_ NetServiceBrowser) SearchForRegistrationDomains() {
	objc.Send[objc.ID](n_.ID, objc.Sel("searchForRegistrationDomains"))
}

// Starts a search for services of a particular type within a specific domain.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/searchForServices(ofType:inDomain:)
func (n_ NetServiceBrowser) SearchForServicesOfTypeInDomain(type_ string, domainString string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("searchForServicesOfType:inDomain:"), objc.String(type_), objc.String(domainString))
}

// Halts a currently running search or resolution.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/stop()
func (n_ NetServiceBrowser) Stop() {
	objc.Send[objc.ID](n_.ID, objc.Sel("stop"))
}

// The delegate object for this instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/delegate
func (n_ NetServiceBrowser) Delegate() objc.ID {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object for this instance.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/delegate
func (n_ NetServiceBrowser) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDelegate:"), value)
}

// Whether to browse over peer-to-peer Bluetooth and Wi-Fi, if available.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/includesPeerToPeer
func (n_ NetServiceBrowser) IncludesPeerToPeer() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("includesPeerToPeer"))
	return rv
}


// SetIncludesPeerToPeer sets the value of the includesPeerToPeer property.
// Whether to browse over peer-to-peer Bluetooth and Wi-Fi, if available.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/includesPeerToPeer
func (n_ NetServiceBrowser) SetIncludesPeerToPeer(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludesPeerToPeer:"), value)
}


