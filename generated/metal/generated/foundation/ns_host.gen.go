// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Host] class.
var (
	HostClass     _HostClass
	HostClassOnce sync.Once
)

func getHostClass() _HostClass {
	HostClassOnce.Do(func() {
		HostClass = _HostClass{objc.GetClass("NSHost")}
	})
	return HostClass
}

type _HostClass struct {
	class objc.Class
}

// An interface definition for the [Host] class.
type IHost interface {
	objectivec.IObject
	// properties:
	Address() IString
	SetAddress(value IString)
	Addresses() IString
	SetAddresses(value IString)
	LocalizedName() IString
	SetLocalizedName(value IString)
	Name() IString
	SetName(value IString)
	Names() IString
	SetNames(value IString)
	// methods:
}

// A representation of an individual host on the network.
//
// The class provides methods to access the network name and address information for a host. Instances of the class represent individual on a network. Use objects to get the current host’s names and addresses and to look up other hosts by name or by address. To create an object, use the , , or class methods (don’t use and ). These methods use available network administration services to discover all names and addresses for the host requested. They don’t attempt to contact the host itself, however. This approach avoids untimely delays due to a host being unavailable, but it may result in incomplete information about the host. An object contains all of the network addresses and names discovered for a given host by the network administration services. Each object may contain several addresses and have more than one name. If an object has more than one name, the additional names are variations on the same name, typically the basic host name plus the fully qualified domain name. For example, with a host name in the domain , an object can hold both the names and . methods are thread-safe.


// A representation of an individual host on the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Host
type Host struct {
	objectivec.Object
}

// HostFrom constructs a [Host] from an unsafe.Pointer.
//
// A representation of an individual host on the network.
func HostFrom(ptr unsafe.Pointer) Host {
	return Host{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HostClass) Alloc() Host {
	rv := objc.Send[Host](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HostClass) New() Host {
	rv := objc.Send[Host](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ Host) Init() Host {
	rv := objc.Send[Host](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ Host) Autorelease() Host {
	rv := objc.Send[Host](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHost creates a new Host instance.
func NewHost() Host {
	return getHostClass().New()
}



// Returns one of the network addresses of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/host/address
func (h_ Host) Address() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("address"))
	return rv
}


// Returns one of the network addresses of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/host/address
func (h_ Host) SetAddress(value IString) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAddress:"), value)
}


// Returns all the network addresses of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/host/addresses
func (h_ Host) Addresses() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("addresses"))
	return rv
}


// Returns all the network addresses of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/host/addresses
func (h_ Host) SetAddresses(value IString) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAddresses:"), value)
}


// Returns the name used as by default when publishing
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/host/localizedname
func (h_ Host) LocalizedName() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("localizedName"))
	return rv
}


// Returns the name used as by default when publishing
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/host/localizedname
func (h_ Host) SetLocalizedName(value IString) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setLocalizedName:"), value)
}


// Returns one of the hostnames of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/host/name
func (h_ Host) Name() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("name"))
	return rv
}


// Returns one of the hostnames of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/host/name
func (h_ Host) SetName(value IString) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setName:"), value)
}


// Returns all the hostnames of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/host/names
func (h_ Host) Names() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("names"))
	return rv
}


// Returns all the hostnames of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/host/names
func (h_ Host) SetNames(value IString) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setNames:"), value)
}



