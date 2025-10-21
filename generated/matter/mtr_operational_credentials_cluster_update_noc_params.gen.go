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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/icacvalue
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) IcacValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("icacValue"))
	return rv
}


// SetIcacValue sets the value of the icacValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/icacvalue
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) SetIcacValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIcacValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/nocvalue
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) NocValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nocValue"))
	return rv
}


// SetNocValue sets the value of the nocValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatenocparams/nocvalue
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) SetNocValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNocValue:"), value)
}



