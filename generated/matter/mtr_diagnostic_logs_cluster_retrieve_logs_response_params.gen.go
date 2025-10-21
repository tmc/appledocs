// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/logcontent
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) LogContent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("logContent"))
	return rv
}


// SetLogContent sets the value of the logContent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/logcontent
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) SetLogContent(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLogContent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/utctimestamp
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) UtcTimeStamp() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("utcTimeStamp"))
	return rv
}


// SetUtcTimeStamp sets the value of the utcTimeStamp property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/utctimestamp
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) SetUtcTimeStamp(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUtcTimeStamp:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/content
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) Content() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("content"))
	return rv
}


// SetContent sets the value of the content property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/content
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) SetContent(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setContent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/timesinceboot
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) TimeSinceBoot() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timeSinceBoot"))
	return rv
}


// SetTimeSinceBoot sets the value of the timeSinceBoot property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/timesinceboot
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) SetTimeSinceBoot(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeSinceBoot:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/timedinvoketimeoutms
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/timedinvoketimeoutms
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/status
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/status
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) SetStatus(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/timestamp
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) TimeStamp() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timeStamp"))
	return rv
}


// SetTimeStamp sets the value of the timeStamp property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsresponseparams/timestamp
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) SetTimeStamp(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeStamp:"), value)
}



