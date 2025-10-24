// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterCertificateChainResponseParams */


/* debug [class_header]: Header for MTROperationalCredentialsClusterCertificateChainResponseParams */
// The class instance for the [MTROperationalCredentialsClusterCertificateChainResponseParams] class.
var (
	MTROperationalCredentialsClusterCertificateChainResponseParamsClass     _MTROperationalCredentialsClusterCertificateChainResponseParamsClass
	MTROperationalCredentialsClusterCertificateChainResponseParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterCertificateChainResponseParamsClass() _MTROperationalCredentialsClusterCertificateChainResponseParamsClass {
	MTROperationalCredentialsClusterCertificateChainResponseParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterCertificateChainResponseParamsClass = _MTROperationalCredentialsClusterCertificateChainResponseParamsClass{objc.GetClass("MTROperationalCredentialsClusterCertificateChainResponseParams")}
	})
	return MTROperationalCredentialsClusterCertificateChainResponseParamsClass
}

type _MTROperationalCredentialsClusterCertificateChainResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterCertificateChainResponseParams */
// An interface definition for the [MTROperationalCredentialsClusterCertificateChainResponseParams] class.
type IMTROperationalCredentialsClusterCertificateChainResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterCertificateChainResponseParams */
	// properties:
	Certificate() objc.IObject /* cross-framework: NSData */
	SetCertificate(value objc.IObject /* cross-framework: NSData */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterCertificateChainResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterCertificateChainResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterCertificateChainResponseParamsClass) Alloc() MTROperationalCredentialsClusterCertificateChainResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterCertificateChainResponseParamsClass) New() MTROperationalCredentialsClusterCertificateChainResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterCertificateChainResponseParams) Init() MTROperationalCredentialsClusterCertificateChainResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterCertificateChainResponseParams) Autorelease() MTROperationalCredentialsClusterCertificateChainResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterCertificateChainResponseParams creates a new MTROperationalCredentialsClusterCertificateChainResponseParams instance.
func NewMTROperationalCredentialsClusterCertificateChainResponseParams() MTROperationalCredentialsClusterCertificateChainResponseParams {
	return getMTROperationalCredentialsClusterCertificateChainResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterCertificateChainResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainResponseParams
type MTROperationalCredentialsClusterCertificateChainResponseParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterCertificateChainResponseParamsFrom constructs a [MTROperationalCredentialsClusterCertificateChainResponseParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterCertificateChainResponseParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterCertificateChainResponseParams {
	return MTROperationalCredentialsClusterCertificateChainResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterCertificateChainResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainResponseParams/init(responseValue:)
func NewMTROperationalCredentialsClusterCertificateChainResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTROperationalCredentialsClusterCertificateChainResponseParams {
	instance := getMTROperationalCredentialsClusterCertificateChainResponseParamsClass().Alloc()
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROperationalCredentialsClusterCertificateChainResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterCertificateChainResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterCertificateChainResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterCertificateChainResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterCertificateChainResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainResponseParams/certificate
func (m_ MTROperationalCredentialsClusterCertificateChainResponseParams) Certificate() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("certificate"))
	return rv
}/* debug [instance_properties/getter]: certificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainResponseParams/certificate
func (m_ MTROperationalCredentialsClusterCertificateChainResponseParams) SetCertificate(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCertificate:"), value)
}/* debug [instance_properties/setter]: certificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainResponseParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterCertificateChainResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainResponseParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterCertificateChainResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterCertificateChainResponseParams */


