// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccountLoginClusterGetSetupPINParams] class.
var (
	MTRAccountLoginClusterGetSetupPINParamsClass     _MTRAccountLoginClusterGetSetupPINParamsClass
	MTRAccountLoginClusterGetSetupPINParamsClassOnce sync.Once
)

func getMTRAccountLoginClusterGetSetupPINParamsClass() _MTRAccountLoginClusterGetSetupPINParamsClass {
	MTRAccountLoginClusterGetSetupPINParamsClassOnce.Do(func() {
		MTRAccountLoginClusterGetSetupPINParamsClass = _MTRAccountLoginClusterGetSetupPINParamsClass{objc.GetClass("MTRAccountLoginClusterGetSetupPINParams")}
	})
	return MTRAccountLoginClusterGetSetupPINParamsClass
}

type _MTRAccountLoginClusterGetSetupPINParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccountLoginClusterGetSetupPINParams] class.
type IMTRAccountLoginClusterGetSetupPINParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TempAccountIdentifier() objc.IObject /* cross-framework: NSString */
	SetTempAccountIdentifier(value objc.IObject /* cross-framework: NSString */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccountLoginClusterGetSetupPINParams
type MTRAccountLoginClusterGetSetupPINParams struct {
	objectivec.Object
}

// MTRAccountLoginClusterGetSetupPINParamsFrom constructs a [MTRAccountLoginClusterGetSetupPINParams] from an unsafe.Pointer.
func MTRAccountLoginClusterGetSetupPINParamsFrom(ptr unsafe.Pointer) MTRAccountLoginClusterGetSetupPINParams {
	return MTRAccountLoginClusterGetSetupPINParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccountLoginClusterGetSetupPINParamsClass) Alloc() MTRAccountLoginClusterGetSetupPINParams {
	rv := objc.Send[MTRAccountLoginClusterGetSetupPINParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccountLoginClusterGetSetupPINParamsClass) New() MTRAccountLoginClusterGetSetupPINParams {
	rv := objc.Send[MTRAccountLoginClusterGetSetupPINParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccountLoginClusterGetSetupPINParams) Init() MTRAccountLoginClusterGetSetupPINParams {
	rv := objc.Send[MTRAccountLoginClusterGetSetupPINParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccountLoginClusterGetSetupPINParams) Autorelease() MTRAccountLoginClusterGetSetupPINParams {
	rv := objc.Send[MTRAccountLoginClusterGetSetupPINParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccountLoginClusterGetSetupPINParams creates a new MTRAccountLoginClusterGetSetupPINParams instance.
func NewMTRAccountLoginClusterGetSetupPINParams() MTRAccountLoginClusterGetSetupPINParams {
	return getMTRAccountLoginClusterGetSetupPINParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclustergetsetuppinparams/serversideprocessingtimeout
func (m_ MTRAccountLoginClusterGetSetupPINParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclustergetsetuppinparams/serversideprocessingtimeout
func (m_ MTRAccountLoginClusterGetSetupPINParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclustergetsetuppinparams/tempaccountidentifier
func (m_ MTRAccountLoginClusterGetSetupPINParams) TempAccountIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("tempAccountIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclustergetsetuppinparams/tempaccountidentifier
func (m_ MTRAccountLoginClusterGetSetupPINParams) SetTempAccountIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTempAccountIdentifier:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclustergetsetuppinparams/timedinvoketimeoutms
func (m_ MTRAccountLoginClusterGetSetupPINParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclustergetsetuppinparams/timedinvoketimeoutms
func (m_ MTRAccountLoginClusterGetSetupPINParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



