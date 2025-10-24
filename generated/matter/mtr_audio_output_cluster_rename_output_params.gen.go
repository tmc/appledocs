// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAudioOutputClusterRenameOutputParams] class.
var (
	MTRAudioOutputClusterRenameOutputParamsClass     _MTRAudioOutputClusterRenameOutputParamsClass
	MTRAudioOutputClusterRenameOutputParamsClassOnce sync.Once
)

func getMTRAudioOutputClusterRenameOutputParamsClass() _MTRAudioOutputClusterRenameOutputParamsClass {
	MTRAudioOutputClusterRenameOutputParamsClassOnce.Do(func() {
		MTRAudioOutputClusterRenameOutputParamsClass = _MTRAudioOutputClusterRenameOutputParamsClass{objc.GetClass("MTRAudioOutputClusterRenameOutputParams")}
	})
	return MTRAudioOutputClusterRenameOutputParamsClass
}

type _MTRAudioOutputClusterRenameOutputParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAudioOutputClusterRenameOutputParams] class.
type IMTRAudioOutputClusterRenameOutputParams interface {
	objectivec.IObject
	// properties:
	Index() objc.IObject /* cross-framework: NSNumber */
	SetIndex(value objc.IObject /* cross-framework: NSNumber */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAudioOutputClusterRenameOutputParams
type MTRAudioOutputClusterRenameOutputParams struct {
	objectivec.Object
}

// MTRAudioOutputClusterRenameOutputParamsFrom constructs a [MTRAudioOutputClusterRenameOutputParams] from an unsafe.Pointer.
func MTRAudioOutputClusterRenameOutputParamsFrom(ptr unsafe.Pointer) MTRAudioOutputClusterRenameOutputParams {
	return MTRAudioOutputClusterRenameOutputParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAudioOutputClusterRenameOutputParamsClass) Alloc() MTRAudioOutputClusterRenameOutputParams {
	rv := objc.Send[MTRAudioOutputClusterRenameOutputParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAudioOutputClusterRenameOutputParamsClass) New() MTRAudioOutputClusterRenameOutputParams {
	rv := objc.Send[MTRAudioOutputClusterRenameOutputParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAudioOutputClusterRenameOutputParams) Init() MTRAudioOutputClusterRenameOutputParams {
	rv := objc.Send[MTRAudioOutputClusterRenameOutputParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAudioOutputClusterRenameOutputParams) Autorelease() MTRAudioOutputClusterRenameOutputParams {
	rv := objc.Send[MTRAudioOutputClusterRenameOutputParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAudioOutputClusterRenameOutputParams creates a new MTRAudioOutputClusterRenameOutputParams instance.
func NewMTRAudioOutputClusterRenameOutputParams() MTRAudioOutputClusterRenameOutputParams {
	return getMTRAudioOutputClusterRenameOutputParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusterrenameoutputparams/index
func (m_ MTRAudioOutputClusterRenameOutputParams) Index() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("index"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusterrenameoutputparams/index
func (m_ MTRAudioOutputClusterRenameOutputParams) SetIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndex:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusterrenameoutputparams/name
func (m_ MTRAudioOutputClusterRenameOutputParams) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusterrenameoutputparams/name
func (m_ MTRAudioOutputClusterRenameOutputParams) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusterrenameoutputparams/serversideprocessingtimeout
func (m_ MTRAudioOutputClusterRenameOutputParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusterrenameoutputparams/serversideprocessingtimeout
func (m_ MTRAudioOutputClusterRenameOutputParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusterrenameoutputparams/timedinvoketimeoutms
func (m_ MTRAudioOutputClusterRenameOutputParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusterrenameoutputparams/timedinvoketimeoutms
func (m_ MTRAudioOutputClusterRenameOutputParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
