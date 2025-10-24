// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalStateClusterStopParams] class.
var (
	MTROperationalStateClusterStopParamsClass     _MTROperationalStateClusterStopParamsClass
	MTROperationalStateClusterStopParamsClassOnce sync.Once
)

func getMTROperationalStateClusterStopParamsClass() _MTROperationalStateClusterStopParamsClass {
	MTROperationalStateClusterStopParamsClassOnce.Do(func() {
		MTROperationalStateClusterStopParamsClass = _MTROperationalStateClusterStopParamsClass{objc.GetClass("MTROperationalStateClusterStopParams")}
	})
	return MTROperationalStateClusterStopParamsClass
}

type _MTROperationalStateClusterStopParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalStateClusterStopParams] class.
type IMTROperationalStateClusterStopParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalStateClusterStopParams
type MTROperationalStateClusterStopParams struct {
	objectivec.Object
}

// MTROperationalStateClusterStopParamsFrom constructs a [MTROperationalStateClusterStopParams] from an unsafe.Pointer.
func MTROperationalStateClusterStopParamsFrom(ptr unsafe.Pointer) MTROperationalStateClusterStopParams {
	return MTROperationalStateClusterStopParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalStateClusterStopParamsClass) Alloc() MTROperationalStateClusterStopParams {
	rv := objc.Send[MTROperationalStateClusterStopParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalStateClusterStopParamsClass) New() MTROperationalStateClusterStopParams {
	rv := objc.Send[MTROperationalStateClusterStopParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalStateClusterStopParams) Init() MTROperationalStateClusterStopParams {
	rv := objc.Send[MTROperationalStateClusterStopParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalStateClusterStopParams) Autorelease() MTROperationalStateClusterStopParams {
	rv := objc.Send[MTROperationalStateClusterStopParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalStateClusterStopParams creates a new MTROperationalStateClusterStopParams instance.
func NewMTROperationalStateClusterStopParams() MTROperationalStateClusterStopParams {
	return getMTROperationalStateClusterStopParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusterstopparams/serversideprocessingtimeout
func (m_ MTROperationalStateClusterStopParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusterstopparams/serversideprocessingtimeout
func (m_ MTROperationalStateClusterStopParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusterstopparams/timedinvoketimeoutms
func (m_ MTROperationalStateClusterStopParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusterstopparams/timedinvoketimeoutms
func (m_ MTROperationalStateClusterStopParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



