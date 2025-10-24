// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralCommissioningClusterArmFailSafeParams] class.
var (
	MTRGeneralCommissioningClusterArmFailSafeParamsClass     _MTRGeneralCommissioningClusterArmFailSafeParamsClass
	MTRGeneralCommissioningClusterArmFailSafeParamsClassOnce sync.Once
)

func getMTRGeneralCommissioningClusterArmFailSafeParamsClass() _MTRGeneralCommissioningClusterArmFailSafeParamsClass {
	MTRGeneralCommissioningClusterArmFailSafeParamsClassOnce.Do(func() {
		MTRGeneralCommissioningClusterArmFailSafeParamsClass = _MTRGeneralCommissioningClusterArmFailSafeParamsClass{objc.GetClass("MTRGeneralCommissioningClusterArmFailSafeParams")}
	})
	return MTRGeneralCommissioningClusterArmFailSafeParamsClass
}

type _MTRGeneralCommissioningClusterArmFailSafeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralCommissioningClusterArmFailSafeParams] class.
type IMTRGeneralCommissioningClusterArmFailSafeParams interface {
	objectivec.IObject
	// properties:
	Breadcrumb() objc.IObject /* cross-framework: NSNumber */
	SetBreadcrumb(value objc.IObject /* cross-framework: NSNumber */)
	ExpiryLengthSeconds() objc.IObject /* cross-framework: NSNumber */
	SetExpiryLengthSeconds(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralCommissioningClusterArmFailSafeParams
type MTRGeneralCommissioningClusterArmFailSafeParams struct {
	objectivec.Object
}

// MTRGeneralCommissioningClusterArmFailSafeParamsFrom constructs a [MTRGeneralCommissioningClusterArmFailSafeParams] from an unsafe.Pointer.
func MTRGeneralCommissioningClusterArmFailSafeParamsFrom(ptr unsafe.Pointer) MTRGeneralCommissioningClusterArmFailSafeParams {
	return MTRGeneralCommissioningClusterArmFailSafeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralCommissioningClusterArmFailSafeParamsClass) Alloc() MTRGeneralCommissioningClusterArmFailSafeParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralCommissioningClusterArmFailSafeParamsClass) New() MTRGeneralCommissioningClusterArmFailSafeParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) Init() MTRGeneralCommissioningClusterArmFailSafeParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) Autorelease() MTRGeneralCommissioningClusterArmFailSafeParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralCommissioningClusterArmFailSafeParams creates a new MTRGeneralCommissioningClusterArmFailSafeParams instance.
func NewMTRGeneralCommissioningClusterArmFailSafeParams() MTRGeneralCommissioningClusterArmFailSafeParams {
	return getMTRGeneralCommissioningClusterArmFailSafeParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/breadcrumb
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) Breadcrumb() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("breadcrumb"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/breadcrumb
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) SetBreadcrumb(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBreadcrumb:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/expirylengthseconds
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) ExpiryLengthSeconds() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("expiryLengthSeconds"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/expirylengthseconds
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) SetExpiryLengthSeconds(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExpiryLengthSeconds:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/serversideprocessingtimeout
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/serversideprocessingtimeout
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/timedinvoketimeoutms
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclusterarmfailsafeparams/timedinvoketimeoutms
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
