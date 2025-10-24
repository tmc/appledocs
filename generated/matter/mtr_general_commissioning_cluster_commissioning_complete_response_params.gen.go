// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralCommissioningClusterCommissioningCompleteResponseParams] class.
var (
	MTRGeneralCommissioningClusterCommissioningCompleteResponseParamsClass     _MTRGeneralCommissioningClusterCommissioningCompleteResponseParamsClass
	MTRGeneralCommissioningClusterCommissioningCompleteResponseParamsClassOnce sync.Once
)

func getMTRGeneralCommissioningClusterCommissioningCompleteResponseParamsClass() _MTRGeneralCommissioningClusterCommissioningCompleteResponseParamsClass {
	MTRGeneralCommissioningClusterCommissioningCompleteResponseParamsClassOnce.Do(func() {
		MTRGeneralCommissioningClusterCommissioningCompleteResponseParamsClass = _MTRGeneralCommissioningClusterCommissioningCompleteResponseParamsClass{objc.GetClass("MTRGeneralCommissioningClusterCommissioningCompleteResponseParams")}
	})
	return MTRGeneralCommissioningClusterCommissioningCompleteResponseParamsClass
}

type _MTRGeneralCommissioningClusterCommissioningCompleteResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralCommissioningClusterCommissioningCompleteResponseParams] class.
type IMTRGeneralCommissioningClusterCommissioningCompleteResponseParams interface {
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralCommissioningClusterCommissioningCompleteResponseParams
type MTRGeneralCommissioningClusterCommissioningCompleteResponseParams struct {
	objectivec.Object
}

// MTRGeneralCommissioningClusterCommissioningCompleteResponseParamsFrom constructs a [MTRGeneralCommissioningClusterCommissioningCompleteResponseParams] from an unsafe.Pointer.
func MTRGeneralCommissioningClusterCommissioningCompleteResponseParamsFrom(ptr unsafe.Pointer) MTRGeneralCommissioningClusterCommissioningCompleteResponseParams {
	return MTRGeneralCommissioningClusterCommissioningCompleteResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralCommissioningClusterCommissioningCompleteResponseParamsClass) Alloc() MTRGeneralCommissioningClusterCommissioningCompleteResponseParams {
	rv := objc.Send[MTRGeneralCommissioningClusterCommissioningCompleteResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralCommissioningClusterCommissioningCompleteResponseParamsClass) New() MTRGeneralCommissioningClusterCommissioningCompleteResponseParams {
	rv := objc.Send[MTRGeneralCommissioningClusterCommissioningCompleteResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralCommissioningClusterCommissioningCompleteResponseParams) Init() MTRGeneralCommissioningClusterCommissioningCompleteResponseParams {
	rv := objc.Send[MTRGeneralCommissioningClusterCommissioningCompleteResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralCommissioningClusterCommissioningCompleteResponseParams) Autorelease() MTRGeneralCommissioningClusterCommissioningCompleteResponseParams {
	rv := objc.Send[MTRGeneralCommissioningClusterCommissioningCompleteResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralCommissioningClusterCommissioningCompleteResponseParams creates a new MTRGeneralCommissioningClusterCommissioningCompleteResponseParams instance.
func NewMTRGeneralCommissioningClusterCommissioningCompleteResponseParams() MTRGeneralCommissioningClusterCommissioningCompleteResponseParams {
	return getMTRGeneralCommissioningClusterCommissioningCompleteResponseParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustercommissioningcompleteresponseparams/debugtext
func (m_ MTRGeneralCommissioningClusterCommissioningCompleteResponseParams) DebugText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("debugText"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustercommissioningcompleteresponseparams/debugtext
func (m_ MTRGeneralCommissioningClusterCommissioningCompleteResponseParams) SetDebugText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDebugText:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustercommissioningcompleteresponseparams/errorcode
func (m_ MTRGeneralCommissioningClusterCommissioningCompleteResponseParams) ErrorCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("errorCode"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustercommissioningcompleteresponseparams/errorcode
func (m_ MTRGeneralCommissioningClusterCommissioningCompleteResponseParams) SetErrorCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorCode:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustercommissioningcompleteresponseparams/timedinvoketimeoutms
func (m_ MTRGeneralCommissioningClusterCommissioningCompleteResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustercommissioningcompleteresponseparams/timedinvoketimeoutms
func (m_ MTRGeneralCommissioningClusterCommissioningCompleteResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
