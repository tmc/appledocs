// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestSpecificParams] class.
var (
	MTRTestClusterClusterTestSpecificParamsClass     _MTRTestClusterClusterTestSpecificParamsClass
	MTRTestClusterClusterTestSpecificParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestSpecificParamsClass() _MTRTestClusterClusterTestSpecificParamsClass {
	MTRTestClusterClusterTestSpecificParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestSpecificParamsClass = _MTRTestClusterClusterTestSpecificParamsClass{objc.GetClass("MTRTestClusterClusterTestSpecificParams")}
	})
	return MTRTestClusterClusterTestSpecificParamsClass
}

type _MTRTestClusterClusterTestSpecificParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestSpecificParams] class.
type IMTRTestClusterClusterTestSpecificParams interface {
	IMTRUnitTestingClusterTestSpecificParams
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestSpecificParams
type MTRTestClusterClusterTestSpecificParams struct {
	MTRUnitTestingClusterTestSpecificParams
}

// MTRTestClusterClusterTestSpecificParamsFrom constructs a [MTRTestClusterClusterTestSpecificParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestSpecificParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestSpecificParams {
	return MTRTestClusterClusterTestSpecificParams{
		MTRUnitTestingClusterTestSpecificParams: MTRUnitTestingClusterTestSpecificParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestSpecificParamsClass) Alloc() MTRTestClusterClusterTestSpecificParams {
	rv := objc.Send[MTRTestClusterClusterTestSpecificParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestSpecificParamsClass) New() MTRTestClusterClusterTestSpecificParams {
	rv := objc.Send[MTRTestClusterClusterTestSpecificParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestSpecificParams) Init() MTRTestClusterClusterTestSpecificParams {
	rv := objc.Send[MTRTestClusterClusterTestSpecificParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestSpecificParams) Autorelease() MTRTestClusterClusterTestSpecificParams {
	rv := objc.Send[MTRTestClusterClusterTestSpecificParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestSpecificParams creates a new MTRTestClusterClusterTestSpecificParams instance.
func NewMTRTestClusterClusterTestSpecificParams() MTRTestClusterClusterTestSpecificParams {
	return getMTRTestClusterClusterTestSpecificParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestspecificparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestSpecificParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestspecificparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestSpecificParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestspecificparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestSpecificParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestspecificparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestSpecificParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
