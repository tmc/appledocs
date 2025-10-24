// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NWTLSParameters */


/* debug [class_header]: Header for NWTLSParameters */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NWTLSParameters */
// An interface definition for the [NWTLSParameters] class.
type INWTLSParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NWTLSParameters */
	// properties:
	MaximumSSLProtocolVersion() uint
	SetMaximumSSLProtocolVersion(value uint)
	MinimumSSLProtocolVersion() uint
	SetMinimumSSLProtocolVersion(value uint)
	SSLCipherSuites() unsafe.Pointer
	SetSSLCipherSuites(value unsafe.Pointer)
	TLSSessionID() objc.IObject /* cross-framework: NSData */
	SetTLSSessionID(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NWTLSParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NWTLSParameters */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NWTLSParameters */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NWTLSParameters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NWTLSParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NWTLSParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NWTLSParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NWTLSParameters */

// The maximum allowed value to use when negotiating TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/maximumSSLProtocolVersion
func (n_ NWTLSParameters) MaximumSSLProtocolVersion() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("maximumSSLProtocolVersion"))
	return rv
}/* debug [instance_properties/getter]: maximumSSLProtocolVersion */


// The maximum allowed value to use when negotiating TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/maximumSSLProtocolVersion
func (n_ NWTLSParameters) SetMaximumSSLProtocolVersion(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumSSLProtocolVersion:"), value)
}/* debug [instance_properties/setter]: maximumSSLProtocolVersion */


// The minimum allowed value to use when negotiating TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/minimumSSLProtocolVersion
func (n_ NWTLSParameters) MinimumSSLProtocolVersion() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("minimumSSLProtocolVersion"))
	return rv
}/* debug [instance_properties/getter]: minimumSSLProtocolVersion */


// The minimum allowed value to use when negotiating TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/minimumSSLProtocolVersion
func (n_ NWTLSParameters) SetMinimumSSLProtocolVersion(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumSSLProtocolVersion:"), value)
}/* debug [instance_properties/setter]: minimumSSLProtocolVersion */


// The set of allowed cipher suites when negotiating TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/sslCipherSuites
func (n_ NWTLSParameters) SSLCipherSuites() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("SSLCipherSuites"))
	return rv
}/* debug [instance_properties/getter]: SSLCipherSuites */


// The set of allowed cipher suites when negotiating TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/sslCipherSuites
func (n_ NWTLSParameters) SetSSLCipherSuites(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSSLCipherSuites:"), value)
}/* debug [instance_properties/setter]: SSLCipherSuites */


// The Session ID to use for the associated TCP connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/tlsSessionID
func (n_ NWTLSParameters) TLSSessionID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("TLSSessionID"))
	return rv
}/* debug [instance_properties/getter]: TLSSessionID */


// The Session ID to use for the associated TCP connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/tlsSessionID
func (n_ NWTLSParameters) SetTLSSessionID(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTLSSessionID:"), value)
}/* debug [instance_properties/setter]: TLSSessionID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NWTLSParameters */



