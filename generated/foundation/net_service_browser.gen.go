// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NetServiceBrowser] class.
var netServiceBrowserClass = _NetServiceBrowserClass{objc.GetClass("NSNetServiceBrowser")}

type _NetServiceBrowserClass struct {
	class objc.Class
}

// An interface definition for the [NetServiceBrowser] class.
type INetServiceBrowser interface {
	objectivec.IObject
	RemoveFromRunLoopForMode(aRunLoop unsafe.Pointer, mode unsafe.Pointer)
	ScheduleInRunLoopForMode(aRunLoop unsafe.Pointer, mode unsafe.Pointer)
	SearchForBrowsableDomains()
	SearchForRegistrationDomains()
	SearchForServicesOfTypeInDomain(type_ string, domainString string)
	Stop()
}

// A network service browser that finds published services on a network using multicast DNS. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return netServiceBrowserClass.New()
}




// Removes the receiver from the specified run loop. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/remove(from:forMode:)
func (n_ NetServiceBrowser) RemoveFromRunLoopForMode(aRunLoop unsafe.Pointer, mode unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeFromRunLoop:forMode:"), aRunLoop, mode)
}
// Adds the receiver to the specified run loop. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/schedule(in:forMode:)
func (n_ NetServiceBrowser) ScheduleInRunLoopForMode(aRunLoop unsafe.Pointer, mode unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("scheduleInRunLoop:forMode:"), aRunLoop, mode)
}
// Initiates a search for domains visible to the host. This method returns immediately. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/searchForBrowsableDomains()
func (n_ NetServiceBrowser) SearchForBrowsableDomains() {
	objc.Send[objc.ID](n_.ID, objc.Sel("searchForBrowsableDomains"))
}
// Initiates a search for domains in which the host may register services. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/searchForRegistrationDomains()
func (n_ NetServiceBrowser) SearchForRegistrationDomains() {
	objc.Send[objc.ID](n_.ID, objc.Sel("searchForRegistrationDomains"))
}
// Starts a search for services of a particular type within a specific domain. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/searchForServices(ofType:inDomain:)
func (n_ NetServiceBrowser) SearchForServicesOfTypeInDomain(type_ string, domainString string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("searchForServicesOfType:inDomain:"), type_, domainString)
}
// Halts a currently running search or resolution. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/stop()
func (n_ NetServiceBrowser) Stop() {
	objc.Send[objc.ID](n_.ID, objc.Sel("stop"))
}

