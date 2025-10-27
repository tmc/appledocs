// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	

	// properties:
	MaximumSSLProtocolVersion() uint
	SetMaximumSSLProtocolVersion(value uint)
	MinimumSSLProtocolVersion() uint
	SetMinimumSSLProtocolVersion(value uint)
	SSLCipherSuites() unsafe.Pointer
	SetSSLCipherSuites(value unsafe.Pointer)
	TLSSessionID() foundation.foundation.INSData
	SetTLSSessionID(value foundation.foundation.INSData)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NWTLSParametersClass) Alloc() NWTLSParameters {
	rv := objc.Send[NWTLSParameters](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// TLS properties for creating a connection.


// TLS properties for creating a connection.
//
// [Full Topic]
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

























// The maximum allowed value to use when negotiating TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/maximumSSLProtocolVersion
func (n_ NWTLSParameters) MaximumSSLProtocolVersion() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("maximumSSLProtocolVersion"))
	return rv
}


// The maximum allowed value to use when negotiating TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/maximumSSLProtocolVersion
func (n_ NWTLSParameters) SetMaximumSSLProtocolVersion(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumSSLProtocolVersion:"), value)
}


// The minimum allowed value to use when negotiating TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/minimumSSLProtocolVersion
func (n_ NWTLSParameters) MinimumSSLProtocolVersion() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("minimumSSLProtocolVersion"))
	return rv
}


// The minimum allowed value to use when negotiating TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/minimumSSLProtocolVersion
func (n_ NWTLSParameters) SetMinimumSSLProtocolVersion(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumSSLProtocolVersion:"), value)
}


// The set of allowed cipher suites when negotiating TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/sslCipherSuites
func (n_ NWTLSParameters) SSLCipherSuites() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("SSLCipherSuites"))
	return rv
}


// The set of allowed cipher suites when negotiating TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/sslCipherSuites
func (n_ NWTLSParameters) SetSSLCipherSuites(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSSLCipherSuites:"), value)
}


// The Session ID to use for the associated TCP connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/tlsSessionID
func (n_ NWTLSParameters) TLSSessionID() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("TLSSessionID"))
	return rv
}


// The Session ID to use for the associated TCP connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/tlsSessionID
func (n_ NWTLSParameters) SetTLSSessionID(value foundation.foundation.INSData) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTLSSessionID:"), value)
}








