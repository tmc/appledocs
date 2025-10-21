// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [NWHostEndpoint] class.
var (
	NWHostEndpointClass     _NWHostEndpointClass
	NWHostEndpointClassOnce sync.Once
)

func getNWHostEndpointClass() _NWHostEndpointClass {
	NWHostEndpointClassOnce.Do(func() {
		NWHostEndpointClass = _NWHostEndpointClass{objc.GetClass("NWHostEndpoint")}
	})
	return NWHostEndpointClass
}

type _NWHostEndpointClass struct {
	class objc.Class
}

// An interface definition for the [NWHostEndpoint] class.
type INWHostEndpoint interface {
	INWEndpoint
}

// A network endpoint specified by DNS name (or IP address) and port.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint
type NWHostEndpoint struct {
	NWEndpoint
}

// NWHostEndpointFrom constructs a [NWHostEndpoint] from an unsafe.Pointer.
//
// A network endpoint specified by DNS name (or IP address) and port.
func NWHostEndpointFrom(ptr unsafe.Pointer) NWHostEndpoint {
	return NWHostEndpoint{
		NWEndpoint: NWEndpointFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NWHostEndpointClass) Alloc() NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NWHostEndpointClass) New() NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NWHostEndpoint) Init() NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NWHostEndpoint) Autorelease() NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWHostEndpoint creates a new NWHostEndpoint instance.
func NewNWHostEndpoint() NWHostEndpoint {
	return getNWHostEndpointClass().New()
}




// Create a host endpoint with a hostname and port.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint/init(hostname:port:)
func NewNWHostEndpointWithHostnamePort(hostname appkit.string, port appkit.string) NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](objc.ID(getNWHostEndpointClass().class), objc.Sel("endpointWithHostname:port:"), hostname, port)
	return rv
}


// Create a host endpoint with a hostname and port.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint/init(hostname:port:)
func (nc _NWHostEndpointClass) EndpointWithHostnamePort(hostname appkit.string, port appkit.string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("endpointWithHostname:port:"), hostname, port)
	return rv
}

// The endpoint’s hostname.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint/hostname
func (n_ NWHostEndpoint) Hostname() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("hostname"))
	return rv
}

// The endpoint’s port, represented as a string.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint/port
func (n_ NWHostEndpoint) Port() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("port"))
	return rv
}


