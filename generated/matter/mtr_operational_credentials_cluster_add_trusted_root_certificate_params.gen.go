// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterAddTrustedRootCertificateParams */


/* debug [class_header]: Header for MTROperationalCredentialsClusterAddTrustedRootCertificateParams */
// The class instance for the [MTROperationalCredentialsClusterAddTrustedRootCertificateParams] class.
var (
	MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass     _MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass
	MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass() _MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass {
	MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass = _MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass{objc.GetClass("MTROperationalCredentialsClusterAddTrustedRootCertificateParams")}
	})
	return MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass
}

type _MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterAddTrustedRootCertificateParams */
// An interface definition for the [MTROperationalCredentialsClusterAddTrustedRootCertificateParams] class.
type IMTROperationalCredentialsClusterAddTrustedRootCertificateParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterAddTrustedRootCertificateParams */
	// properties:
	RootCACertificate() objc.IObject /* cross-framework: NSData */
	SetRootCACertificate(value objc.IObject /* cross-framework: NSData */)
	RootCertificate() objc.IObject /* cross-framework: NSData */
	SetRootCertificate(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterAddTrustedRootCertificateParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterAddTrustedRootCertificateParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass) Alloc() MTROperationalCredentialsClusterAddTrustedRootCertificateParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddTrustedRootCertificateParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass) New() MTROperationalCredentialsClusterAddTrustedRootCertificateParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddTrustedRootCertificateParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) Init() MTROperationalCredentialsClusterAddTrustedRootCertificateParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddTrustedRootCertificateParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) Autorelease() MTROperationalCredentialsClusterAddTrustedRootCertificateParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddTrustedRootCertificateParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterAddTrustedRootCertificateParams creates a new MTROperationalCredentialsClusterAddTrustedRootCertificateParams instance.
func NewMTROperationalCredentialsClusterAddTrustedRootCertificateParams() MTROperationalCredentialsClusterAddTrustedRootCertificateParams {
	return getMTROperationalCredentialsClusterAddTrustedRootCertificateParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterAddTrustedRootCertificateParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddTrustedRootCertificateParams
type MTROperationalCredentialsClusterAddTrustedRootCertificateParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterAddTrustedRootCertificateParamsFrom constructs a [MTROperationalCredentialsClusterAddTrustedRootCertificateParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterAddTrustedRootCertificateParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterAddTrustedRootCertificateParams {
	return MTROperationalCredentialsClusterAddTrustedRootCertificateParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterAddTrustedRootCertificateParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterAddTrustedRootCertificateParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterAddTrustedRootCertificateParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterAddTrustedRootCertificateParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterAddTrustedRootCertificateParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddTrustedRootCertificateParams/rootCACertificate
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) RootCACertificate() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("rootCACertificate"))
	return rv
}/* debug [instance_properties/getter]: rootCACertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddTrustedRootCertificateParams/rootCACertificate
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) SetRootCACertificate(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCACertificate:"), value)
}/* debug [instance_properties/setter]: rootCACertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddTrustedRootCertificateParams/rootCertificate
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) RootCertificate() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("rootCertificate"))
	return rv
}/* debug [instance_properties/getter]: rootCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddTrustedRootCertificateParams/rootCertificate
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) SetRootCertificate(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCertificate:"), value)
}/* debug [instance_properties/setter]: rootCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddTrustedRootCertificateParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddTrustedRootCertificateParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddTrustedRootCertificateParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddTrustedRootCertificateParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterAddTrustedRootCertificateParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterAddTrustedRootCertificateParams */



