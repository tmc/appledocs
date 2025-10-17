// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [NetServiceBrowser] class.
var NetServiceBrowserClass objc.Class

func init() {
	NetServiceBrowserClass = objc.GetClass("NSNetServiceBrowser")
}

type NetServiceBrowser struct {
	objc.ID
}

func NetServiceBrowserFrom(ptr unsafe.Pointer) NetServiceBrowser {
	return NetServiceBrowser{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc NetServiceBrowser) Alloc() NetServiceBrowser {
	ret := objc.ID(NetServiceBrowserClass).Send(objc.RegisterName("alloc"))
	return NetServiceBrowser{ret}
}

// Init initializes the instance.
func (n_ NetServiceBrowser) Init() NetServiceBrowser {
	ret := n_.ID.Send(objc.RegisterName("init"))
	return NetServiceBrowser{ret}
}
// Initializes an allocated   object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetServiceBrowser/init()
func NewNetServiceBrowser() NetServiceBrowser {
	instance := NetServiceBrowser{}.Alloc()
	instance = instance.Init()
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Removes the receiver from the specified run loop. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetServiceBrowser/remove(from:forMode:)
func (n_ NetServiceBrowser) RemoveFromRunLoopForMode(aRunLoop unsafe.Pointer, mode unsafe.Pointer) {
	sel := objc.RegisterName("removeFromRunLoop:forMode:")
	n_.ID.Send(sel, aRunLoop, mode)
}
// Adds the receiver to the specified run loop. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetServiceBrowser/schedule(in:forMode:)
func (n_ NetServiceBrowser) ScheduleInRunLoopForMode(aRunLoop unsafe.Pointer, mode unsafe.Pointer) {
	sel := objc.RegisterName("scheduleInRunLoop:forMode:")
	n_.ID.Send(sel, aRunLoop, mode)
}
// Initiates a search for domains visible to the host. This method returns immediately. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetServiceBrowser/searchForBrowsableDomains()
func (n_ NetServiceBrowser) SearchForBrowsableDomains() {
	sel := objc.RegisterName("searchForBrowsableDomains")
	n_.ID.Send(sel)
}
// Initiates a search for domains in which the host may register services. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetServiceBrowser/searchForRegistrationDomains()
func (n_ NetServiceBrowser) SearchForRegistrationDomains() {
	sel := objc.RegisterName("searchForRegistrationDomains")
	n_.ID.Send(sel)
}
// Starts a search for services of a particular type within a specific domain. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetServiceBrowser/searchForServices(ofType:inDomain:)
func (n_ NetServiceBrowser) SearchForServicesOfTypeInDomain(type_ string, domainString string) {
	sel := objc.RegisterName("searchForServicesOfType:inDomain:")
	n_.ID.Send(sel, type_, domainString)
}
// Halts a currently running search or resolution. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetServiceBrowser/stop()
func (n_ NetServiceBrowser) Stop() {
	sel := objc.RegisterName("stop")
	n_.ID.Send(sel)
}

