// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAdministratorCommissioningClusterOpenCommissioningWindowParams] class.
var (
	MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass     _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass
	MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClassOnce sync.Once
)

func getMTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass() _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass {
	MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClassOnce.Do(func() {
		MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass = _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass{objc.GetClass("MTRAdministratorCommissioningClusterOpenCommissioningWindowParams")}
	})
	return MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass
}

type _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAdministratorCommissioningClusterOpenCommissioningWindowParams] class.
type IMTRAdministratorCommissioningClusterOpenCommissioningWindowParams interface {
	objectivec.IObject
	// properties:
	CommissioningTimeout() objc.IObject /* cross-framework: NSNumber */
	SetCommissioningTimeout(value objc.IObject /* cross-framework: NSNumber */)
	Discriminator() objc.IObject /* cross-framework: NSNumber */
	SetDiscriminator(value objc.IObject /* cross-framework: NSNumber */)
	Iterations() objc.IObject /* cross-framework: NSNumber */
	SetIterations(value objc.IObject /* cross-framework: NSNumber */)
	PakePasscodeVerifier() objc.IObject /* cross-framework: Data */
	SetPakePasscodeVerifier(value objc.IObject /* cross-framework: Data */)
	PakeVerifier() objc.IObject /* cross-framework: Data */
	SetPakeVerifier(value objc.IObject /* cross-framework: Data */)
	Salt() objc.IObject /* cross-framework: Data */
	SetSalt(value objc.IObject /* cross-framework: Data */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAdministratorCommissioningClusterOpenCommissioningWindowParams
type MTRAdministratorCommissioningClusterOpenCommissioningWindowParams struct {
	objectivec.Object
}

// MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsFrom constructs a [MTRAdministratorCommissioningClusterOpenCommissioningWindowParams] from an unsafe.Pointer.
func MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsFrom(ptr unsafe.Pointer) MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	return MTRAdministratorCommissioningClusterOpenCommissioningWindowParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass) Alloc() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenCommissioningWindowParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass) New() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenCommissioningWindowParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) Init() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenCommissioningWindowParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) Autorelease() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenCommissioningWindowParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAdministratorCommissioningClusterOpenCommissioningWindowParams creates a new MTRAdministratorCommissioningClusterOpenCommissioningWindowParams instance.
func NewMTRAdministratorCommissioningClusterOpenCommissioningWindowParams() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	return getMTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/commissioningtimeout
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) CommissioningTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("commissioningTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/commissioningtimeout
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetCommissioningTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommissioningTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/discriminator
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) Discriminator() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("discriminator"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/discriminator
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetDiscriminator(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDiscriminator:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/iterations
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) Iterations() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("iterations"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/iterations
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetIterations(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIterations:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/pakepasscodeverifier
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) PakePasscodeVerifier() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("pakePasscodeVerifier"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/pakepasscodeverifier
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetPakePasscodeVerifier(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPakePasscodeVerifier:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/pakeverifier
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) PakeVerifier() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("pakeVerifier"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/pakeverifier
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetPakeVerifier(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPakeVerifier:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/salt
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) Salt() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("salt"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/salt
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetSalt(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSalt:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/serversideprocessingtimeout
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/serversideprocessingtimeout
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/timedinvoketimeoutms
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtradministratorcommissioningclusteropencommissioningwindowparams/timedinvoketimeoutms
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
