// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDiagnosticLogsClusterRetrieveLogsResponseParams] class.
var (
	MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass     _MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass
	MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClassOnce sync.Once
)

func getMTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass() _MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass {
	MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClassOnce.Do(func() {
		MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass = _MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass{objc.GetClass("MTRDiagnosticLogsClusterRetrieveLogsResponseParams")}
	})
	return MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass
}

type _MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDiagnosticLogsClusterRetrieveLogsResponseParams] class.
type IMTRDiagnosticLogsClusterRetrieveLogsResponseParams interface {
	objectivec.IObject
	// properties:
	Content() objc.IObject /* cross-framework: Data */
	SetContent(value objc.IObject /* cross-framework: Data */)
	LogContent() objc.IObject /* cross-framework: Data */
	SetLogContent(value objc.IObject /* cross-framework: Data */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimeSinceBoot() objc.IObject /* cross-framework: NSNumber */
	SetTimeSinceBoot(value objc.IObject /* cross-framework: NSNumber */)
	TimeStamp() objc.IObject /* cross-framework: NSNumber */
	SetTimeStamp(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UtcTimeStamp() objc.IObject /* cross-framework: NSNumber */
	SetUtcTimeStamp(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDiagnosticLogsClusterRetrieveLogsResponseParams
type MTRDiagnosticLogsClusterRetrieveLogsResponseParams struct {
	objectivec.Object
}

// MTRDiagnosticLogsClusterRetrieveLogsResponseParamsFrom constructs a [MTRDiagnosticLogsClusterRetrieveLogsResponseParams] from an unsafe.Pointer.
func MTRDiagnosticLogsClusterRetrieveLogsResponseParamsFrom(ptr unsafe.Pointer) MTRDiagnosticLogsClusterRetrieveLogsResponseParams {
	return MTRDiagnosticLogsClusterRetrieveLogsResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass) Alloc() MTRDiagnosticLogsClusterRetrieveLogsResponseParams {
	rv := objc.Send[MTRDiagnosticLogsClusterRetrieveLogsResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass) New() MTRDiagnosticLogsClusterRetrieveLogsResponseParams {
	rv := objc.Send[MTRDiagnosticLogsClusterRetrieveLogsResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) Init() MTRDiagnosticLogsClusterRetrieveLogsResponseParams {
	rv := objc.Send[MTRDiagnosticLogsClusterRetrieveLogsResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) Autorelease() MTRDiagnosticLogsClusterRetrieveLogsResponseParams {
	rv := objc.Send[MTRDiagnosticLogsClusterRetrieveLogsResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDiagnosticLogsClusterRetrieveLogsResponseParams creates a new MTRDiagnosticLogsClusterRetrieveLogsResponseParams instance.
func NewMTRDiagnosticLogsClusterRetrieveLogsResponseParams() MTRDiagnosticLogsClusterRetrieveLogsResponseParams {
	return getMTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/content
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) Content() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("content"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/content
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) SetContent(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setContent:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/logcontent
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) LogContent() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("logContent"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/logcontent
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) SetLogContent(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLogContent:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/status
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/status
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/timesinceboot
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) TimeSinceBoot() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timeSinceBoot"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/timesinceboot
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) SetTimeSinceBoot(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeSinceBoot:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/timestamp
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) TimeStamp() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timeStamp"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/timestamp
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) SetTimeStamp(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeStamp:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/timedinvoketimeoutms
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/timedinvoketimeoutms
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/utctimestamp
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) UtcTimeStamp() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("utcTimeStamp"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/utctimestamp
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) SetUtcTimeStamp(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUtcTimeStamp:"), value)
}
