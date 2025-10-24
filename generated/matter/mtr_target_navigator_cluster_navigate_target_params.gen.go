// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRTargetNavigatorClusterNavigateTargetParams] class.
var (
	MTRTargetNavigatorClusterNavigateTargetParamsClass     _MTRTargetNavigatorClusterNavigateTargetParamsClass
	MTRTargetNavigatorClusterNavigateTargetParamsClassOnce sync.Once
)

func getMTRTargetNavigatorClusterNavigateTargetParamsClass() _MTRTargetNavigatorClusterNavigateTargetParamsClass {
	MTRTargetNavigatorClusterNavigateTargetParamsClassOnce.Do(func() {
		MTRTargetNavigatorClusterNavigateTargetParamsClass = _MTRTargetNavigatorClusterNavigateTargetParamsClass{objc.GetClass("MTRTargetNavigatorClusterNavigateTargetParams")}
	})
	return MTRTargetNavigatorClusterNavigateTargetParamsClass
}

type _MTRTargetNavigatorClusterNavigateTargetParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTargetNavigatorClusterNavigateTargetParams] class.
type IMTRTargetNavigatorClusterNavigateTargetParams interface {
	objectivec.IObject
	// properties:
	Data() objc.IObject /* cross-framework: NSString */
	SetData(value objc.IObject /* cross-framework: NSString */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	Target() objc.IObject /* cross-framework: NSNumber */
	SetTarget(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterNavigateTargetParams
type MTRTargetNavigatorClusterNavigateTargetParams struct {
	objectivec.Object
}

// MTRTargetNavigatorClusterNavigateTargetParamsFrom constructs a [MTRTargetNavigatorClusterNavigateTargetParams] from an unsafe.Pointer.
func MTRTargetNavigatorClusterNavigateTargetParamsFrom(ptr unsafe.Pointer) MTRTargetNavigatorClusterNavigateTargetParams {
	return MTRTargetNavigatorClusterNavigateTargetParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTargetNavigatorClusterNavigateTargetParamsClass) Alloc() MTRTargetNavigatorClusterNavigateTargetParams {
	rv := objc.Send[MTRTargetNavigatorClusterNavigateTargetParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTargetNavigatorClusterNavigateTargetParamsClass) New() MTRTargetNavigatorClusterNavigateTargetParams {
	rv := objc.Send[MTRTargetNavigatorClusterNavigateTargetParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) Init() MTRTargetNavigatorClusterNavigateTargetParams {
	rv := objc.Send[MTRTargetNavigatorClusterNavigateTargetParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) Autorelease() MTRTargetNavigatorClusterNavigateTargetParams {
	rv := objc.Send[MTRTargetNavigatorClusterNavigateTargetParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTargetNavigatorClusterNavigateTargetParams creates a new MTRTargetNavigatorClusterNavigateTargetParams instance.
func NewMTRTargetNavigatorClusterNavigateTargetParams() MTRTargetNavigatorClusterNavigateTargetParams {
	return getMTRTargetNavigatorClusterNavigateTargetParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/data
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) Data() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/data
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) SetData(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/serversideprocessingtimeout
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/serversideprocessingtimeout
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/target
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) Target() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("target"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/target
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) SetTarget(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTarget:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/timedinvoketimeoutms
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/timedinvoketimeoutms
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



