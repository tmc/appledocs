// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTROperationalCredentialsClusterAddNOCParams] class.
type IMTROperationalCredentialsClusterAddNOCParams interface {
	objectivec.IObject
	// properties:
	AdminVendorId() objc.IObject /* cross-framework: NSNumber */
	SetAdminVendorId(value objc.IObject /* cross-framework: NSNumber */)
	CaseAdminSubject() objc.IObject /* cross-framework: NSNumber */
	SetCaseAdminSubject(value objc.IObject /* cross-framework: NSNumber */)
	IcacValue() objc.IObject /* cross-framework: Data */
	SetIcacValue(value objc.IObject /* cross-framework: Data */)
	IpkValue() objc.IObject /* cross-framework: Data */
	SetIpkValue(value objc.IObject /* cross-framework: Data */)
	NocValue() objc.IObject /* cross-framework: Data */
	SetNocValue(value objc.IObject /* cross-framework: Data */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams
type MTROperationalCredentialsClusterAddNOCParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterAddNOCParamsFrom constructs a [MTROperationalCredentialsClusterAddNOCParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterAddNOCParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterAddNOCParams {
	return MTROperationalCredentialsClusterAddNOCParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterAddNOCParamsClass) Alloc() MTROperationalCredentialsClusterAddNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddNOCParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/adminvendorid
func (m_ MTROperationalCredentialsClusterAddNOCParams) AdminVendorId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("adminVendorId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/adminvendorid
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetAdminVendorId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAdminVendorId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/caseadminsubject
func (m_ MTROperationalCredentialsClusterAddNOCParams) CaseAdminSubject() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("caseAdminSubject"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/caseadminsubject
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetCaseAdminSubject(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCaseAdminSubject:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/icacvalue
func (m_ MTROperationalCredentialsClusterAddNOCParams) IcacValue() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("icacValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/icacvalue
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetIcacValue(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIcacValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/ipkvalue
func (m_ MTROperationalCredentialsClusterAddNOCParams) IpkValue() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("ipkValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/ipkvalue
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetIpkValue(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIpkValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/nocvalue
func (m_ MTROperationalCredentialsClusterAddNOCParams) NocValue() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("nocValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/nocvalue
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetNocValue(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNocValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterAddNOCParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAddNOCParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



