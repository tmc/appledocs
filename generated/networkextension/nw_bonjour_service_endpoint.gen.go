// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NWBonjourServiceEndpoint] class.
var (
	NWBonjourServiceEndpointClass     _NWBonjourServiceEndpointClass
	NWBonjourServiceEndpointClassOnce sync.Once
)

func getNWBonjourServiceEndpointClass() _NWBonjourServiceEndpointClass {
	NWBonjourServiceEndpointClassOnce.Do(func() {
		NWBonjourServiceEndpointClass = _NWBonjourServiceEndpointClass{objc.GetClass("NWBonjourServiceEndpoint")}
	})
	return NWBonjourServiceEndpointClass
}

type _NWBonjourServiceEndpointClass struct {
	class objc.Class
}

// An interface definition for the [NWBonjourServiceEndpoint] class.
type INWBonjourServiceEndpoint interface {
	INWEndpoint
	Domain() string
	Type() string
	Name() string
	SetName(value string)
}

// A network endpoint specified as a Bonjour service name, type, and domain.
//
// For example, the Bonjour service has the name , the type , and the domain .
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint
type NWBonjourServiceEndpoint struct {
	NWEndpoint
}

// NWBonjourServiceEndpointFrom constructs a [NWBonjourServiceEndpoint] from an unsafe.Pointer.
//
// A network endpoint specified as a Bonjour service name, type, and domain.
func NWBonjourServiceEndpointFrom(ptr unsafe.Pointer) NWBonjourServiceEndpoint {
	return NWBonjourServiceEndpoint{
		NWEndpoint: NWEndpointFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NWBonjourServiceEndpointClass) Alloc() NWBonjourServiceEndpoint {
	rv := objc.Send[NWBonjourServiceEndpoint](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NWBonjourServiceEndpointClass) New() NWBonjourServiceEndpoint {
	rv := objc.Send[NWBonjourServiceEndpoint](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NWBonjourServiceEndpoint) Init() NWBonjourServiceEndpoint {
	rv := objc.Send[NWBonjourServiceEndpoint](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NWBonjourServiceEndpoint) Autorelease() NWBonjourServiceEndpoint {
	rv := objc.Send[NWBonjourServiceEndpoint](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWBonjourServiceEndpoint creates a new NWBonjourServiceEndpoint instance.
func NewNWBonjourServiceEndpoint() NWBonjourServiceEndpoint {
	return getNWBonjourServiceEndpointClass().New()
}




// Create an endpoint with a Bonjour service name, type, and domain. All fields must be specified.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint/init(name:type:domain:)
func NewNWBonjourServiceEndpointWithNameTypeDomain(name string, type_ string, domain string) NWBonjourServiceEndpoint {
	rv := objc.Send[NWBonjourServiceEndpoint](objc.ID(getNWBonjourServiceEndpointClass().class), objc.Sel("endpointWithName:type:domain:"), objc.String(name), objc.String(type_), objc.String(domain))
	return rv
}


// Create an endpoint with a Bonjour service name, type, and domain. All fields must be specified.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint/init(name:type:domain:)
func (nc _NWBonjourServiceEndpointClass) EndpointWithNameTypeDomain(name string, type_ string, domain string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("endpointWithName:type:domain:"), objc.String(name), objc.String(type_), objc.String(domain))
	return rv
}

// The endpoint’s Bonjour service domain, such as .
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint/domain
func (n_ NWBonjourServiceEndpoint) Domain() string {
	rv := objc.Send[string](n_.ID, objc.Sel("domain"))
	return rv
}

// The endpoint’s Bonjour service type.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint/type
func (n_ NWBonjourServiceEndpoint) Type() string {
	rv := objc.Send[string](n_.ID, objc.Sel("type"))
	return rv
}

// The endpoint’s Bonjour service name.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwbonjourserviceendpoint/name
func (n_ NWBonjourServiceEndpoint) Name() string {
	rv := objc.Send[string](n_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The endpoint’s Bonjour service name.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwbonjourserviceendpoint/name
func (n_ NWBonjourServiceEndpoint) SetName(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setName:"), objc.String(value))
}


