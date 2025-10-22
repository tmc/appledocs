// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NWTLSParameters] class.
var (
	NWTLSParametersClass     _NWTLSParametersClass
	NWTLSParametersClassOnce sync.Once
)

func getNWTLSParametersClass() _NWTLSParametersClass {
	NWTLSParametersClassOnce.Do(func() {
		NWTLSParametersClass = _NWTLSParametersClass{objc.GetClass("NWTLSParameters")}
	})
	return NWTLSParametersClass
}

type _NWTLSParametersClass struct {
	class objc.Class
}

// An interface definition for the [NWTLSParameters] class.
type INWTLSParameters interface {
	objectivec.IObject
	MaximumSSLProtocolVersion() int
	SetMaximumSSLProtocolVersion(value int)
	MinimumSSLProtocolVersion() int
	SetMinimumSSLProtocolVersion(value int)
	SslCipherSuites() foundation.Number
	SetSslCipherSuites(value foundation.INumber)
	TlsSessionID() foundation.Data
	SetTlsSessionID(value foundation.IData)
}

// TLS properties for creating a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters
type NWTLSParameters struct {
	objectivec.Object
}

// NWTLSParametersFrom constructs a [NWTLSParameters] from an unsafe.Pointer.
//
// TLS properties for creating a connection.
func NWTLSParametersFrom(ptr unsafe.Pointer) NWTLSParameters {
	return NWTLSParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NWTLSParametersClass) Alloc() NWTLSParameters {
	rv := objc.Send[NWTLSParameters](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NWTLSParametersClass) New() NWTLSParameters {
	rv := objc.Send[NWTLSParameters](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NWTLSParameters) Init() NWTLSParameters {
	rv := objc.Send[NWTLSParameters](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NWTLSParameters) Autorelease() NWTLSParameters {
	rv := objc.Send[NWTLSParameters](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWTLSParameters creates a new NWTLSParameters instance.
func NewNWTLSParameters() NWTLSParameters {
	return getNWTLSParametersClass().New()
}


// The maximum allowed
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtlsparameters/maximumsslprotocolversion
func (n_ NWTLSParameters) MaximumSSLProtocolVersion() int {
	rv := objc.Send[int](n_.ID, objc.Sel("maximumSSLProtocolVersion"))
	return rv
}


// SetMaximumSSLProtocolVersion sets the value of the maximumSSLProtocolVersion property.
// The maximum allowed

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtlsparameters/maximumsslprotocolversion
func (n_ NWTLSParameters) SetMaximumSSLProtocolVersion(value int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumSSLProtocolVersion:"), value)
}

// The minimum allowed
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtlsparameters/minimumsslprotocolversion
func (n_ NWTLSParameters) MinimumSSLProtocolVersion() int {
	rv := objc.Send[int](n_.ID, objc.Sel("minimumSSLProtocolVersion"))
	return rv
}


// SetMinimumSSLProtocolVersion sets the value of the minimumSSLProtocolVersion property.
// The minimum allowed

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtlsparameters/minimumsslprotocolversion
func (n_ NWTLSParameters) SetMinimumSSLProtocolVersion(value int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumSSLProtocolVersion:"), value)
}

// The set of allowed cipher suites when negotiating TLS.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtlsparameters/sslciphersuites
func (n_ NWTLSParameters) SslCipherSuites() foundation.Number {
	rv := objc.Send[foundation.Number](n_.ID, objc.Sel("sslCipherSuites"))
	return rv
}


// SetSslCipherSuites sets the value of the sslCipherSuites property.
// The set of allowed cipher suites when negotiating TLS.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtlsparameters/sslciphersuites
func (n_ NWTLSParameters) SetSslCipherSuites(value foundation.INumber) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSslCipherSuites:"), value)
}

// The Session ID to use for the associated TCP connection.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtlsparameters/tlssessionid
func (n_ NWTLSParameters) TlsSessionID() foundation.Data {
	rv := objc.Send[foundation.Data](n_.ID, objc.Sel("tlsSessionID"))
	return rv
}


// SetTlsSessionID sets the value of the tlsSessionID property.
// The Session ID to use for the associated TCP connection.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtlsparameters/tlssessionid
func (n_ NWTLSParameters) SetTlsSessionID(value foundation.IData) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTlsSessionID:"), value)
}



