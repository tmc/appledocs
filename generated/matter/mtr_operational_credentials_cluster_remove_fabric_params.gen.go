// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterRemoveFabricParams] class.
var (
	MTROperationalCredentialsClusterRemoveFabricParamsClass     _MTROperationalCredentialsClusterRemoveFabricParamsClass
	MTROperationalCredentialsClusterRemoveFabricParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterRemoveFabricParamsClass() _MTROperationalCredentialsClusterRemoveFabricParamsClass {
	MTROperationalCredentialsClusterRemoveFabricParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterRemoveFabricParamsClass = _MTROperationalCredentialsClusterRemoveFabricParamsClass{objc.GetClass("MTROperationalCredentialsClusterRemoveFabricParams")}
	})
	return MTROperationalCredentialsClusterRemoveFabricParamsClass
}

type _MTROperationalCredentialsClusterRemoveFabricParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterRemoveFabricParams] class.
type IMTROperationalCredentialsClusterRemoveFabricParams interface {
	objectivec.IObject
	// properties:
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterRemoveFabricParams
type MTROperationalCredentialsClusterRemoveFabricParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterRemoveFabricParamsFrom constructs a [MTROperationalCredentialsClusterRemoveFabricParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterRemoveFabricParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterRemoveFabricParams {
	return MTROperationalCredentialsClusterRemoveFabricParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterRemoveFabricParamsClass) Alloc() MTROperationalCredentialsClusterRemoveFabricParams {
	rv := objc.Send[MTROperationalCredentialsClusterRemoveFabricParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterRemoveFabricParamsClass) New() MTROperationalCredentialsClusterRemoveFabricParams {
	rv := objc.Send[MTROperationalCredentialsClusterRemoveFabricParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) Init() MTROperationalCredentialsClusterRemoveFabricParams {
	rv := objc.Send[MTROperationalCredentialsClusterRemoveFabricParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) Autorelease() MTROperationalCredentialsClusterRemoveFabricParams {
	rv := objc.Send[MTROperationalCredentialsClusterRemoveFabricParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterRemoveFabricParams creates a new MTROperationalCredentialsClusterRemoveFabricParams instance.
func NewMTROperationalCredentialsClusterRemoveFabricParams() MTROperationalCredentialsClusterRemoveFabricParams {
	return getMTROperationalCredentialsClusterRemoveFabricParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterremovefabricparams/fabricindex
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterremovefabricparams/fabricindex
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterremovefabricparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterremovefabricparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterremovefabricparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterremovefabricparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterRemoveFabricParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



