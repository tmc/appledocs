// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaInputClusterSelectInputParams] class.
var (
	MTRMediaInputClusterSelectInputParamsClass     _MTRMediaInputClusterSelectInputParamsClass
	MTRMediaInputClusterSelectInputParamsClassOnce sync.Once
)

func getMTRMediaInputClusterSelectInputParamsClass() _MTRMediaInputClusterSelectInputParamsClass {
	MTRMediaInputClusterSelectInputParamsClassOnce.Do(func() {
		MTRMediaInputClusterSelectInputParamsClass = _MTRMediaInputClusterSelectInputParamsClass{objc.GetClass("MTRMediaInputClusterSelectInputParams")}
	})
	return MTRMediaInputClusterSelectInputParamsClass
}

type _MTRMediaInputClusterSelectInputParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaInputClusterSelectInputParams] class.
type IMTRMediaInputClusterSelectInputParams interface {
	objectivec.IObject
	// properties:
	Index() objc.IObject /* cross-framework: NSNumber */
	SetIndex(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaInputClusterSelectInputParams
type MTRMediaInputClusterSelectInputParams struct {
	objectivec.Object
}

// MTRMediaInputClusterSelectInputParamsFrom constructs a [MTRMediaInputClusterSelectInputParams] from an unsafe.Pointer.
func MTRMediaInputClusterSelectInputParamsFrom(ptr unsafe.Pointer) MTRMediaInputClusterSelectInputParams {
	return MTRMediaInputClusterSelectInputParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaInputClusterSelectInputParamsClass) Alloc() MTRMediaInputClusterSelectInputParams {
	rv := objc.Send[MTRMediaInputClusterSelectInputParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaInputClusterSelectInputParamsClass) New() MTRMediaInputClusterSelectInputParams {
	rv := objc.Send[MTRMediaInputClusterSelectInputParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaInputClusterSelectInputParams) Init() MTRMediaInputClusterSelectInputParams {
	rv := objc.Send[MTRMediaInputClusterSelectInputParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaInputClusterSelectInputParams) Autorelease() MTRMediaInputClusterSelectInputParams {
	rv := objc.Send[MTRMediaInputClusterSelectInputParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaInputClusterSelectInputParams creates a new MTRMediaInputClusterSelectInputParams instance.
func NewMTRMediaInputClusterSelectInputParams() MTRMediaInputClusterSelectInputParams {
	return getMTRMediaInputClusterSelectInputParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterselectinputparams/index
func (m_ MTRMediaInputClusterSelectInputParams) Index() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("index"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterselectinputparams/index
func (m_ MTRMediaInputClusterSelectInputParams) SetIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterselectinputparams/serversideprocessingtimeout
func (m_ MTRMediaInputClusterSelectInputParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterselectinputparams/serversideprocessingtimeout
func (m_ MTRMediaInputClusterSelectInputParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterselectinputparams/timedinvoketimeoutms
func (m_ MTRMediaInputClusterSelectInputParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterselectinputparams/timedinvoketimeoutms
func (m_ MTRMediaInputClusterSelectInputParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



