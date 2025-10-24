// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterCertificateChainRequestParams */


/* debug [class_header]: Header for MTROperationalCredentialsClusterCertificateChainRequestParams */
// The class instance for the [MTROperationalCredentialsClusterCertificateChainRequestParams] class.
var (
	MTROperationalCredentialsClusterCertificateChainRequestParamsClass     _MTROperationalCredentialsClusterCertificateChainRequestParamsClass
	MTROperationalCredentialsClusterCertificateChainRequestParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterCertificateChainRequestParamsClass() _MTROperationalCredentialsClusterCertificateChainRequestParamsClass {
	MTROperationalCredentialsClusterCertificateChainRequestParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterCertificateChainRequestParamsClass = _MTROperationalCredentialsClusterCertificateChainRequestParamsClass{objc.GetClass("MTROperationalCredentialsClusterCertificateChainRequestParams")}
	})
	return MTROperationalCredentialsClusterCertificateChainRequestParamsClass
}

type _MTROperationalCredentialsClusterCertificateChainRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterCertificateChainRequestParams */
// An interface definition for the [MTROperationalCredentialsClusterCertificateChainRequestParams] class.
type IMTROperationalCredentialsClusterCertificateChainRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterCertificateChainRequestParams */
	// properties:
	CertificateType() objc.IObject /* cross-framework: NSNumber */
	SetCertificateType(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterCertificateChainRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterCertificateChainRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterCertificateChainRequestParamsClass) Alloc() MTROperationalCredentialsClusterCertificateChainRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterCertificateChainRequestParamsClass) New() MTROperationalCredentialsClusterCertificateChainRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) Init() MTROperationalCredentialsClusterCertificateChainRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) Autorelease() MTROperationalCredentialsClusterCertificateChainRequestParams {
	rv := objc.Send[MTROperationalCredentialsClusterCertificateChainRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterCertificateChainRequestParams creates a new MTROperationalCredentialsClusterCertificateChainRequestParams instance.
func NewMTROperationalCredentialsClusterCertificateChainRequestParams() MTROperationalCredentialsClusterCertificateChainRequestParams {
	return getMTROperationalCredentialsClusterCertificateChainRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterCertificateChainRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainRequestParams
type MTROperationalCredentialsClusterCertificateChainRequestParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterCertificateChainRequestParamsFrom constructs a [MTROperationalCredentialsClusterCertificateChainRequestParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterCertificateChainRequestParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterCertificateChainRequestParams {
	return MTROperationalCredentialsClusterCertificateChainRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterCertificateChainRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterCertificateChainRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterCertificateChainRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterCertificateChainRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterCertificateChainRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainRequestParams/certificateType
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) CertificateType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("certificateType"))
	return rv
}/* debug [instance_properties/getter]: certificateType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainRequestParams/certificateType
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) SetCertificateType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCertificateType:"), value)
}/* debug [instance_properties/setter]: certificateType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainRequestParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainRequestParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainRequestParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterCertificateChainRequestParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterCertificateChainRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterCertificateChainRequestParams */



