// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralCommissioningClusterArmFailSafeResponseParams] class.
var (
	MTRGeneralCommissioningClusterArmFailSafeResponseParamsClass     _MTRGeneralCommissioningClusterArmFailSafeResponseParamsClass
	MTRGeneralCommissioningClusterArmFailSafeResponseParamsClassOnce sync.Once
)

func getMTRGeneralCommissioningClusterArmFailSafeResponseParamsClass() _MTRGeneralCommissioningClusterArmFailSafeResponseParamsClass {
	MTRGeneralCommissioningClusterArmFailSafeResponseParamsClassOnce.Do(func() {
		MTRGeneralCommissioningClusterArmFailSafeResponseParamsClass = _MTRGeneralCommissioningClusterArmFailSafeResponseParamsClass{objc.GetClass("MTRGeneralCommissioningClusterArmFailSafeResponseParams")}
	})
	return MTRGeneralCommissioningClusterArmFailSafeResponseParamsClass
}

type _MTRGeneralCommissioningClusterArmFailSafeResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralCommissioningClusterArmFailSafeResponseParams] class.
type IMTRGeneralCommissioningClusterArmFailSafeResponseParams interface {
	objectivec.IObject
	// properties:
	DebugText() objc.IObject /* cross-framework: NSString */
	SetDebugText(value objc.IObject /* cross-framework: NSString */)
	ErrorCode() objc.IObject /* cross-framework: NSNumber */
	SetErrorCode(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralCommissioningClusterArmFailSafeResponseParams
type MTRGeneralCommissioningClusterArmFailSafeResponseParams struct {
	objectivec.Object
}

// MTRGeneralCommissioningClusterArmFailSafeResponseParamsFrom constructs a [MTRGeneralCommissioningClusterArmFailSafeResponseParams] from an unsafe.Pointer.
func MTRGeneralCommissioningClusterArmFailSafeResponseParamsFrom(ptr unsafe.Pointer) MTRGeneralCommissioningClusterArmFailSafeResponseParams {
	return MTRGeneralCommissioningClusterArmFailSafeResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralCommissioningClusterArmFailSafeResponseParamsClass) Alloc() MTRGeneralCommissioningClusterArmFailSafeResponseParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralCommissioningClusterArmFailSafeResponseParamsClass) New() MTRGeneralCommissioningClusterArmFailSafeResponseParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralCommissioningClusterArmFailSafeResponseParams) Init() MTRGeneralCommissioningClusterArmFailSafeResponseParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralCommissioningClusterArmFailSafeResponseParams) Autorelease() MTRGeneralCommissioningClusterArmFailSafeResponseParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralCommissioningClusterArmFailSafeResponseParams creates a new MTRGeneralCommissioningClusterArmFailSafeResponseParams instance.
func NewMTRGeneralCommissioningClusterArmFailSafeResponseParams() MTRGeneralCommissioningClusterArmFailSafeResponseParams {
	return getMTRGeneralCommissioningClusterArmFailSafeResponseParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsaferesponseparams/debugtext
func (m_ MTRGeneralCommissioningClusterArmFailSafeResponseParams) DebugText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("debugText"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsaferesponseparams/debugtext
func (m_ MTRGeneralCommissioningClusterArmFailSafeResponseParams) SetDebugText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDebugText:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsaferesponseparams/errorcode
func (m_ MTRGeneralCommissioningClusterArmFailSafeResponseParams) ErrorCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("errorCode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsaferesponseparams/errorcode
func (m_ MTRGeneralCommissioningClusterArmFailSafeResponseParams) SetErrorCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorCode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsaferesponseparams/timedinvoketimeoutms
func (m_ MTRGeneralCommissioningClusterArmFailSafeResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsaferesponseparams/timedinvoketimeoutms
func (m_ MTRGeneralCommissioningClusterArmFailSafeResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



