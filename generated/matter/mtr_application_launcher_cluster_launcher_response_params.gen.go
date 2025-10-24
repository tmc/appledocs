// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRApplicationLauncherClusterLauncherResponseParams] class.
var (
	MTRApplicationLauncherClusterLauncherResponseParamsClass     _MTRApplicationLauncherClusterLauncherResponseParamsClass
	MTRApplicationLauncherClusterLauncherResponseParamsClassOnce sync.Once
)

func getMTRApplicationLauncherClusterLauncherResponseParamsClass() _MTRApplicationLauncherClusterLauncherResponseParamsClass {
	MTRApplicationLauncherClusterLauncherResponseParamsClassOnce.Do(func() {
		MTRApplicationLauncherClusterLauncherResponseParamsClass = _MTRApplicationLauncherClusterLauncherResponseParamsClass{objc.GetClass("MTRApplicationLauncherClusterLauncherResponseParams")}
	})
	return MTRApplicationLauncherClusterLauncherResponseParamsClass
}

type _MTRApplicationLauncherClusterLauncherResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRApplicationLauncherClusterLauncherResponseParams] class.
type IMTRApplicationLauncherClusterLauncherResponseParams interface {
	objectivec.IObject
	// properties:
	Data() objc.IObject /* cross-framework: Data */
	SetData(value objc.IObject /* cross-framework: Data */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLauncherResponseParams
type MTRApplicationLauncherClusterLauncherResponseParams struct {
	objectivec.Object
}

// MTRApplicationLauncherClusterLauncherResponseParamsFrom constructs a [MTRApplicationLauncherClusterLauncherResponseParams] from an unsafe.Pointer.
func MTRApplicationLauncherClusterLauncherResponseParamsFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterLauncherResponseParams {
	return MTRApplicationLauncherClusterLauncherResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterLauncherResponseParamsClass) Alloc() MTRApplicationLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRApplicationLauncherClusterLauncherResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRApplicationLauncherClusterLauncherResponseParamsClass) New() MTRApplicationLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRApplicationLauncherClusterLauncherResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) Init() MTRApplicationLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRApplicationLauncherClusterLauncherResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) Autorelease() MTRApplicationLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRApplicationLauncherClusterLauncherResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterLauncherResponseParams creates a new MTRApplicationLauncherClusterLauncherResponseParams instance.
func NewMTRApplicationLauncherClusterLauncherResponseParams() MTRApplicationLauncherClusterLauncherResponseParams {
	return getMTRApplicationLauncherClusterLauncherResponseParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterlauncherresponseparams/data
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) Data() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterlauncherresponseparams/data
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) SetData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterlauncherresponseparams/status
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterlauncherresponseparams/status
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterlauncherresponseparams/timedinvoketimeoutms
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterlauncherresponseparams/timedinvoketimeoutms
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



