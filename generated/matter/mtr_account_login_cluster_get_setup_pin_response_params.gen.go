// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccountLoginClusterGetSetupPINResponseParams] class.
var (
	MTRAccountLoginClusterGetSetupPINResponseParamsClass     _MTRAccountLoginClusterGetSetupPINResponseParamsClass
	MTRAccountLoginClusterGetSetupPINResponseParamsClassOnce sync.Once
)

func getMTRAccountLoginClusterGetSetupPINResponseParamsClass() _MTRAccountLoginClusterGetSetupPINResponseParamsClass {
	MTRAccountLoginClusterGetSetupPINResponseParamsClassOnce.Do(func() {
		MTRAccountLoginClusterGetSetupPINResponseParamsClass = _MTRAccountLoginClusterGetSetupPINResponseParamsClass{objc.GetClass("MTRAccountLoginClusterGetSetupPINResponseParams")}
	})
	return MTRAccountLoginClusterGetSetupPINResponseParamsClass
}

type _MTRAccountLoginClusterGetSetupPINResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccountLoginClusterGetSetupPINResponseParams] class.
type IMTRAccountLoginClusterGetSetupPINResponseParams interface {
	objectivec.IObject
	// properties:
	SetupPIN() objc.IObject /* cross-framework: NSString */
	SetSetupPIN(value objc.IObject /* cross-framework: NSString */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccountLoginClusterGetSetupPINResponseParams
type MTRAccountLoginClusterGetSetupPINResponseParams struct {
	objectivec.Object
}

// MTRAccountLoginClusterGetSetupPINResponseParamsFrom constructs a [MTRAccountLoginClusterGetSetupPINResponseParams] from an unsafe.Pointer.
func MTRAccountLoginClusterGetSetupPINResponseParamsFrom(ptr unsafe.Pointer) MTRAccountLoginClusterGetSetupPINResponseParams {
	return MTRAccountLoginClusterGetSetupPINResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccountLoginClusterGetSetupPINResponseParamsClass) Alloc() MTRAccountLoginClusterGetSetupPINResponseParams {
	rv := objc.Send[MTRAccountLoginClusterGetSetupPINResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccountLoginClusterGetSetupPINResponseParamsClass) New() MTRAccountLoginClusterGetSetupPINResponseParams {
	rv := objc.Send[MTRAccountLoginClusterGetSetupPINResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccountLoginClusterGetSetupPINResponseParams) Init() MTRAccountLoginClusterGetSetupPINResponseParams {
	rv := objc.Send[MTRAccountLoginClusterGetSetupPINResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccountLoginClusterGetSetupPINResponseParams) Autorelease() MTRAccountLoginClusterGetSetupPINResponseParams {
	rv := objc.Send[MTRAccountLoginClusterGetSetupPINResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccountLoginClusterGetSetupPINResponseParams creates a new MTRAccountLoginClusterGetSetupPINResponseParams instance.
func NewMTRAccountLoginClusterGetSetupPINResponseParams() MTRAccountLoginClusterGetSetupPINResponseParams {
	return getMTRAccountLoginClusterGetSetupPINResponseParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclustergetsetuppinresponseparams/setuppin
func (m_ MTRAccountLoginClusterGetSetupPINResponseParams) SetupPIN() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("setupPIN"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclustergetsetuppinresponseparams/setuppin
func (m_ MTRAccountLoginClusterGetSetupPINResponseParams) SetSetupPIN(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSetupPIN:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclustergetsetuppinresponseparams/timedinvoketimeoutms
func (m_ MTRAccountLoginClusterGetSetupPINResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclustergetsetuppinresponseparams/timedinvoketimeoutms
func (m_ MTRAccountLoginClusterGetSetupPINResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



