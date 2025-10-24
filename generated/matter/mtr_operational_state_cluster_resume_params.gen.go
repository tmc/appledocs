// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalStateClusterResumeParams] class.
var (
	MTROperationalStateClusterResumeParamsClass     _MTROperationalStateClusterResumeParamsClass
	MTROperationalStateClusterResumeParamsClassOnce sync.Once
)

func getMTROperationalStateClusterResumeParamsClass() _MTROperationalStateClusterResumeParamsClass {
	MTROperationalStateClusterResumeParamsClassOnce.Do(func() {
		MTROperationalStateClusterResumeParamsClass = _MTROperationalStateClusterResumeParamsClass{objc.GetClass("MTROperationalStateClusterResumeParams")}
	})
	return MTROperationalStateClusterResumeParamsClass
}

type _MTROperationalStateClusterResumeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalStateClusterResumeParams] class.
type IMTROperationalStateClusterResumeParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalStateClusterResumeParams
type MTROperationalStateClusterResumeParams struct {
	objectivec.Object
}

// MTROperationalStateClusterResumeParamsFrom constructs a [MTROperationalStateClusterResumeParams] from an unsafe.Pointer.
func MTROperationalStateClusterResumeParamsFrom(ptr unsafe.Pointer) MTROperationalStateClusterResumeParams {
	return MTROperationalStateClusterResumeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalStateClusterResumeParamsClass) Alloc() MTROperationalStateClusterResumeParams {
	rv := objc.Send[MTROperationalStateClusterResumeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalStateClusterResumeParamsClass) New() MTROperationalStateClusterResumeParams {
	rv := objc.Send[MTROperationalStateClusterResumeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalStateClusterResumeParams) Init() MTROperationalStateClusterResumeParams {
	rv := objc.Send[MTROperationalStateClusterResumeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalStateClusterResumeParams) Autorelease() MTROperationalStateClusterResumeParams {
	rv := objc.Send[MTROperationalStateClusterResumeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalStateClusterResumeParams creates a new MTROperationalStateClusterResumeParams instance.
func NewMTROperationalStateClusterResumeParams() MTROperationalStateClusterResumeParams {
	return getMTROperationalStateClusterResumeParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusterresumeparams/serversideprocessingtimeout
func (m_ MTROperationalStateClusterResumeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusterresumeparams/serversideprocessingtimeout
func (m_ MTROperationalStateClusterResumeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusterresumeparams/timedinvoketimeoutms
func (m_ MTROperationalStateClusterResumeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusterresumeparams/timedinvoketimeoutms
func (m_ MTROperationalStateClusterResumeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



