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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/adminvendorid
func (m_ MTROperationalCredentialsClusterAddNOCParams) AdminVendorId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("adminVendorId"))
	return rv
}


// SetAdminVendorId sets the value of the adminVendorId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/adminvendorid
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetAdminVendorId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAdminVendorId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/caseadminsubject
func (m_ MTROperationalCredentialsClusterAddNOCParams) CaseAdminSubject() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("caseAdminSubject"))
	return rv
}


// SetCaseAdminSubject sets the value of the caseAdminSubject property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/caseadminsubject
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetCaseAdminSubject(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCaseAdminSubject:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/icacvalue
func (m_ MTROperationalCredentialsClusterAddNOCParams) IcacValue() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("icacValue"))
	return rv
}


// SetIcacValue sets the value of the icacValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/icacvalue
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetIcacValue(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIcacValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/ipkvalue
func (m_ MTROperationalCredentialsClusterAddNOCParams) IpkValue() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("ipkValue"))
	return rv
}


// SetIpkValue sets the value of the ipkValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/ipkvalue
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetIpkValue(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIpkValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/nocvalue
func (m_ MTROperationalCredentialsClusterAddNOCParams) NocValue() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("nocValue"))
	return rv
}


// SetNocValue sets the value of the nocValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/nocvalue
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetNocValue(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNocValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterAddNOCParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAddNOCParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusteraddnocparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterAddNOCParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



