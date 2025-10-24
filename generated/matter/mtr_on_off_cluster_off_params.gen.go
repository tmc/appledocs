// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROnOffClusterOffParams] class.
var (
	MTROnOffClusterOffParamsClass     _MTROnOffClusterOffParamsClass
	MTROnOffClusterOffParamsClassOnce sync.Once
)

func getMTROnOffClusterOffParamsClass() _MTROnOffClusterOffParamsClass {
	MTROnOffClusterOffParamsClassOnce.Do(func() {
		MTROnOffClusterOffParamsClass = _MTROnOffClusterOffParamsClass{objc.GetClass("MTROnOffClusterOffParams")}
	})
	return MTROnOffClusterOffParamsClass
}

type _MTROnOffClusterOffParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROnOffClusterOffParams] class.
type IMTROnOffClusterOffParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOffParams
type MTROnOffClusterOffParams struct {
	objectivec.Object
}

// MTROnOffClusterOffParamsFrom constructs a [MTROnOffClusterOffParams] from an unsafe.Pointer.
func MTROnOffClusterOffParamsFrom(ptr unsafe.Pointer) MTROnOffClusterOffParams {
	return MTROnOffClusterOffParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROnOffClusterOffParamsClass) Alloc() MTROnOffClusterOffParams {
	rv := objc.Send[MTROnOffClusterOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROnOffClusterOffParamsClass) New() MTROnOffClusterOffParams {
	rv := objc.Send[MTROnOffClusterOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROnOffClusterOffParams) Init() MTROnOffClusterOffParams {
	rv := objc.Send[MTROnOffClusterOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROnOffClusterOffParams) Autorelease() MTROnOffClusterOffParams {
	rv := objc.Send[MTROnOffClusterOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROnOffClusterOffParams creates a new MTROnOffClusterOffParams instance.
func NewMTROnOffClusterOffParams() MTROnOffClusterOffParams {
	return getMTROnOffClusterOffParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteroffparams/serversideprocessingtimeout
func (m_ MTROnOffClusterOffParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteroffparams/serversideprocessingtimeout
func (m_ MTROnOffClusterOffParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteroffparams/timedinvoketimeoutms
func (m_ MTROnOffClusterOffParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteroffparams/timedinvoketimeoutms
func (m_ MTROnOffClusterOffParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



