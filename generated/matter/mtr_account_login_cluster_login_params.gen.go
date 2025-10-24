// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccountLoginClusterLoginParams] class.
var (
	MTRAccountLoginClusterLoginParamsClass     _MTRAccountLoginClusterLoginParamsClass
	MTRAccountLoginClusterLoginParamsClassOnce sync.Once
)

func getMTRAccountLoginClusterLoginParamsClass() _MTRAccountLoginClusterLoginParamsClass {
	MTRAccountLoginClusterLoginParamsClassOnce.Do(func() {
		MTRAccountLoginClusterLoginParamsClass = _MTRAccountLoginClusterLoginParamsClass{objc.GetClass("MTRAccountLoginClusterLoginParams")}
	})
	return MTRAccountLoginClusterLoginParamsClass
}

type _MTRAccountLoginClusterLoginParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccountLoginClusterLoginParams] class.
type IMTRAccountLoginClusterLoginParams interface {
	objectivec.IObject
	// properties:
	Node() objc.IObject /* cross-framework: NSNumber */
	SetNode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	SetupPIN() objc.IObject /* cross-framework: NSString */
	SetSetupPIN(value objc.IObject /* cross-framework: NSString */)
	TempAccountIdentifier() objc.IObject /* cross-framework: NSString */
	SetTempAccountIdentifier(value objc.IObject /* cross-framework: NSString */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccountLoginClusterLoginParams
type MTRAccountLoginClusterLoginParams struct {
	objectivec.Object
}

// MTRAccountLoginClusterLoginParamsFrom constructs a [MTRAccountLoginClusterLoginParams] from an unsafe.Pointer.
func MTRAccountLoginClusterLoginParamsFrom(ptr unsafe.Pointer) MTRAccountLoginClusterLoginParams {
	return MTRAccountLoginClusterLoginParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccountLoginClusterLoginParamsClass) Alloc() MTRAccountLoginClusterLoginParams {
	rv := objc.Send[MTRAccountLoginClusterLoginParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccountLoginClusterLoginParamsClass) New() MTRAccountLoginClusterLoginParams {
	rv := objc.Send[MTRAccountLoginClusterLoginParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccountLoginClusterLoginParams) Init() MTRAccountLoginClusterLoginParams {
	rv := objc.Send[MTRAccountLoginClusterLoginParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccountLoginClusterLoginParams) Autorelease() MTRAccountLoginClusterLoginParams {
	rv := objc.Send[MTRAccountLoginClusterLoginParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccountLoginClusterLoginParams creates a new MTRAccountLoginClusterLoginParams instance.
func NewMTRAccountLoginClusterLoginParams() MTRAccountLoginClusterLoginParams {
	return getMTRAccountLoginClusterLoginParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/node
func (m_ MTRAccountLoginClusterLoginParams) Node() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("node"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/node
func (m_ MTRAccountLoginClusterLoginParams) SetNode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNode:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/serversideprocessingtimeout
func (m_ MTRAccountLoginClusterLoginParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/serversideprocessingtimeout
func (m_ MTRAccountLoginClusterLoginParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/setuppin
func (m_ MTRAccountLoginClusterLoginParams) SetupPIN() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("setupPIN"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/setuppin
func (m_ MTRAccountLoginClusterLoginParams) SetSetupPIN(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSetupPIN:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/tempaccountidentifier
func (m_ MTRAccountLoginClusterLoginParams) TempAccountIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("tempAccountIdentifier"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/tempaccountidentifier
func (m_ MTRAccountLoginClusterLoginParams) SetTempAccountIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTempAccountIdentifier:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/timedinvoketimeoutms
func (m_ MTRAccountLoginClusterLoginParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/timedinvoketimeoutms
func (m_ MTRAccountLoginClusterLoginParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
