// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterMoveToClosestFrequencyParams] class.
var (
	MTRLevelControlClusterMoveToClosestFrequencyParamsClass     _MTRLevelControlClusterMoveToClosestFrequencyParamsClass
	MTRLevelControlClusterMoveToClosestFrequencyParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveToClosestFrequencyParamsClass() _MTRLevelControlClusterMoveToClosestFrequencyParamsClass {
	MTRLevelControlClusterMoveToClosestFrequencyParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveToClosestFrequencyParamsClass = _MTRLevelControlClusterMoveToClosestFrequencyParamsClass{objc.GetClass("MTRLevelControlClusterMoveToClosestFrequencyParams")}
	})
	return MTRLevelControlClusterMoveToClosestFrequencyParamsClass
}

type _MTRLevelControlClusterMoveToClosestFrequencyParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterMoveToClosestFrequencyParams] class.
type IMTRLevelControlClusterMoveToClosestFrequencyParams interface {
	objectivec.IObject
	// properties:
	Frequency() objc.IObject /* cross-framework: NSNumber */
	SetFrequency(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToClosestFrequencyParams
type MTRLevelControlClusterMoveToClosestFrequencyParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveToClosestFrequencyParamsFrom constructs a [MTRLevelControlClusterMoveToClosestFrequencyParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveToClosestFrequencyParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveToClosestFrequencyParams {
	return MTRLevelControlClusterMoveToClosestFrequencyParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveToClosestFrequencyParamsClass) Alloc() MTRLevelControlClusterMoveToClosestFrequencyParams {
	rv := objc.Send[MTRLevelControlClusterMoveToClosestFrequencyParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterMoveToClosestFrequencyParamsClass) New() MTRLevelControlClusterMoveToClosestFrequencyParams {
	rv := objc.Send[MTRLevelControlClusterMoveToClosestFrequencyParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) Init() MTRLevelControlClusterMoveToClosestFrequencyParams {
	rv := objc.Send[MTRLevelControlClusterMoveToClosestFrequencyParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) Autorelease() MTRLevelControlClusterMoveToClosestFrequencyParams {
	rv := objc.Send[MTRLevelControlClusterMoveToClosestFrequencyParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveToClosestFrequencyParams creates a new MTRLevelControlClusterMoveToClosestFrequencyParams instance.
func NewMTRLevelControlClusterMoveToClosestFrequencyParams() MTRLevelControlClusterMoveToClosestFrequencyParams {
	return getMTRLevelControlClusterMoveToClosestFrequencyParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetoclosestfrequencyparams/frequency
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) Frequency() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("frequency"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetoclosestfrequencyparams/frequency
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) SetFrequency(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFrequency:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetoclosestfrequencyparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetoclosestfrequencyparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetoclosestfrequencyparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetoclosestfrequencyparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



