// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterUpdateNOCParams] class.
var (
	MTROperationalCredentialsClusterUpdateNOCParamsClass     _MTROperationalCredentialsClusterUpdateNOCParamsClass
	MTROperationalCredentialsClusterUpdateNOCParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterUpdateNOCParamsClass() _MTROperationalCredentialsClusterUpdateNOCParamsClass {
	MTROperationalCredentialsClusterUpdateNOCParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterUpdateNOCParamsClass = _MTROperationalCredentialsClusterUpdateNOCParamsClass{objc.GetClass("MTROperationalCredentialsClusterUpdateNOCParams")}
	})
	return MTROperationalCredentialsClusterUpdateNOCParamsClass
}

type _MTROperationalCredentialsClusterUpdateNOCParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterUpdateNOCParams] class.
type IMTROperationalCredentialsClusterUpdateNOCParams interface {
	objectivec.IObject
	// properties:
	IcacValue() objc.IObject /* cross-framework: Data */
	SetIcacValue(value objc.IObject /* cross-framework: Data */)
	NocValue() objc.IObject /* cross-framework: Data */
	SetNocValue(value objc.IObject /* cross-framework: Data */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateNOCParams
type MTROperationalCredentialsClusterUpdateNOCParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterUpdateNOCParamsFrom constructs a [MTROperationalCredentialsClusterUpdateNOCParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterUpdateNOCParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterUpdateNOCParams {
	return MTROperationalCredentialsClusterUpdateNOCParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterUpdateNOCParamsClass) Alloc() MTROperationalCredentialsClusterUpdateNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateNOCParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterUpdateNOCParamsClass) New() MTROperationalCredentialsClusterUpdateNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateNOCParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) Init() MTROperationalCredentialsClusterUpdateNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateNOCParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) Autorelease() MTROperationalCredentialsClusterUpdateNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateNOCParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterUpdateNOCParams creates a new MTROperationalCredentialsClusterUpdateNOCParams instance.
func NewMTROperationalCredentialsClusterUpdateNOCParams() MTROperationalCredentialsClusterUpdateNOCParams {
	return getMTROperationalCredentialsClusterUpdateNOCParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/icacvalue
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) IcacValue() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("icacValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/icacvalue
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) SetIcacValue(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIcacValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/nocvalue
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) NocValue() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("nocValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/nocvalue
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) SetNocValue(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNocValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



