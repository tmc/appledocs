// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterAddNOCParams */


/* debug [class_header]: Header for MTROperationalCredentialsClusterAddNOCParams */
// The class instance for the [MTROperationalCredentialsClusterAddNOCParams] class.
var (
	MTROperationalCredentialsClusterAddNOCParamsClass     _MTROperationalCredentialsClusterAddNOCParamsClass
	MTROperationalCredentialsClusterAddNOCParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterAddNOCParamsClass() _MTROperationalCredentialsClusterAddNOCParamsClass {
	MTROperationalCredentialsClusterAddNOCParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterAddNOCParamsClass = _MTROperationalCredentialsClusterAddNOCParamsClass{objc.GetClass("MTROperationalCredentialsClusterAddNOCParams")}
	})
	return MTROperationalCredentialsClusterAddNOCParamsClass
}

type _MTROperationalCredentialsClusterAddNOCParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterAddNOCParams */
// An interface definition for the [MTROperationalCredentialsClusterAddNOCParams] class.
type IMTROperationalCredentialsClusterAddNOCParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterAddNOCParams */
	// properties:
	AdminVendorId() objc.IObject /* cross-framework: NSNumber */
	SetAdminVendorId(value objc.IObject /* cross-framework: NSNumber */)
	CaseAdminSubject() objc.IObject /* cross-framework: NSNumber */
	SetCaseAdminSubject(value objc.IObject /* cross-framework: NSNumber */)
	IcacValue() objc.IObject /* cross-framework: NSData */
	SetIcacValue(value objc.IObject /* cross-framework: NSData */)
	IpkValue() objc.IObject /* cross-framework: NSData */
	SetIpkValue(value objc.IObject /* cross-framework: NSData */)
	NocValue() objc.IObject /* cross-framework: NSData */
	SetNocValue(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterAddNOCParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterAddNOCParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterAddNOCParamsClass) Alloc() MTROperationalCredentialsClusterAddNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddNOCParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterAddNOCParamsClass) New() MTROperationalCredentialsClusterAddNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddNOCParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterAddNOCParams) Init() MTROperationalCredentialsClusterAddNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddNOCParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterAddNOCParams) Autorelease() MTROperationalCredentialsClusterAddNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddNOCParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterAddNOCParams creates a new MTROperationalCredentialsClusterAddNOCParams instance.
func NewMTROperationalCredentialsClusterAddNOCParams() MTROperationalCredentialsClusterAddNOCParams {
	return getMTROperationalCredentialsClusterAddNOCParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterAddNOCParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams
type MTROperationalCredentialsClusterAddNOCParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterAddNOCParamsFrom constructs a [MTROperationalCredentialsClusterAddNOCParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterAddNOCParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterAddNOCParams {
	return MTROperationalCredentialsClusterAddNOCParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterAddNOCParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterAddNOCParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterAddNOCParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterAddNOCParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterAddNOCParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams/adminVendorId
func (m_ MTROperationalCredentialsClusterAddNOCParams) AdminVendorId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("adminVendorId"))
	return rv
}/* debug [instance_properties/getter]: adminVendorId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams/adminVendorId
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetAdminVendorId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAdminVendorId:"), value)
}/* debug [instance_properties/setter]: adminVendorId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams/caseAdminSubject
func (m_ MTROperationalCredentialsClusterAddNOCParams) CaseAdminSubject() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("caseAdminSubject"))
	return rv
}/* debug [instance_properties/getter]: caseAdminSubject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams/caseAdminSubject
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetCaseAdminSubject(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCaseAdminSubject:"), value)
}/* debug [instance_properties/setter]: caseAdminSubject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams/icacValue
func (m_ MTROperationalCredentialsClusterAddNOCParams) IcacValue() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("icacValue"))
	return rv
}/* debug [instance_properties/getter]: icacValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams/icacValue
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetIcacValue(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIcacValue:"), value)
}/* debug [instance_properties/setter]: icacValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams/ipkValue
func (m_ MTROperationalCredentialsClusterAddNOCParams) IpkValue() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("ipkValue"))
	return rv
}/* debug [instance_properties/getter]: ipkValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams/ipkValue
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetIpkValue(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIpkValue:"), value)
}/* debug [instance_properties/setter]: ipkValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams/nocValue
func (m_ MTROperationalCredentialsClusterAddNOCParams) NocValue() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("nocValue"))
	return rv
}/* debug [instance_properties/getter]: nocValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams/nocValue
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetNocValue(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNocValue:"), value)
}/* debug [instance_properties/setter]: nocValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterAddNOCParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterAddNOCParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterAddNOCParams */



