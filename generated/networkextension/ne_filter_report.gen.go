// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEFilterReport] class.
var (
	NEFilterReportClass     _NEFilterReportClass
	NEFilterReportClassOnce sync.Once
)

func getNEFilterReportClass() _NEFilterReportClass {
	NEFilterReportClassOnce.Do(func() {
		NEFilterReportClass = _NEFilterReportClass{objc.GetClass("NEFilterReport")}
	})
	return NEFilterReportClass
}

type _NEFilterReportClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterReport] class.
type INEFilterReport interface {
	objectivec.IObject
}

// The report of the data provider’s action on a flow.
//
// The system issues a report by calling your control provider’s method with a report instance when the data provider issues a verdict whose property is set to .
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport
type NEFilterReport struct {
	objectivec.Object
}

// NEFilterReportFrom constructs a [NEFilterReport] from an unsafe.Pointer.
//
// The report of the data provider’s action on a flow.
func NEFilterReportFrom(ptr unsafe.Pointer) NEFilterReport {
	return NEFilterReport{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterReportClass) Alloc() NEFilterReport {
	rv := objc.Send[NEFilterReport](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterReportClass) New() NEFilterReport {
	rv := objc.Send[NEFilterReport](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterReport) Init() NEFilterReport {
	rv := objc.Send[NEFilterReport](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterReport) Autorelease() NEFilterReport {
	rv := objc.Send[NEFilterReport](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterReport creates a new NEFilterReport instance.
func NewNEFilterReport() NEFilterReport {
	return getNEFilterReportClass().New()
}


// The number of inbound bytes received from the flow.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterreport/bytesinboundcount
func (n_ NEFilterReport) BytesInboundCount() int {
	rv := objc.Send[int](n_.ID, objc.Sel("bytesInboundCount"))
	return rv
}


// SetBytesInboundCount sets the value of the bytesInboundCount property.
// The number of inbound bytes received from the flow.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterreport/bytesinboundcount
func (n_ NEFilterReport) SetBytesInboundCount(value int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setBytesInboundCount:"), value)
}

// The number of outbound bytes sent on the flow.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterreport/bytesoutboundcount
func (n_ NEFilterReport) BytesOutboundCount() int {
	rv := objc.Send[int](n_.ID, objc.Sel("bytesOutboundCount"))
	return rv
}


// SetBytesOutboundCount sets the value of the bytesOutboundCount property.
// The number of outbound bytes sent on the flow.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterreport/bytesoutboundcount
func (n_ NEFilterReport) SetBytesOutboundCount(value int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setBytesOutboundCount:"), value)
}

// The action taken on the reported flow.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterreport/action
func (n_ NEFilterReport) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// The action taken on the reported flow.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterreport/action
func (n_ NEFilterReport) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAction:"), value)
}

// The flow on which the associated action was taken.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterreport/flow
func (n_ NEFilterReport) Flow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("flow"))
	return rv
}


// SetFlow sets the value of the flow property.
// The flow on which the associated action was taken.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterreport/flow
func (n_ NEFilterReport) SetFlow(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFlow:"), value)
}

// The type of event indicated by this report.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterreport/event-swift.property
func (n_ NEFilterReport) Event() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("event"))
	return rv
}


// SetEvent sets the value of the event property.
// The type of event indicated by this report.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterreport/event-swift.property
func (n_ NEFilterReport) SetEvent(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEvent:"), value)
}

// A Boolean value that indicates whether to send a report to the control provider when processing this verdict.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterverdict/shouldreport
func (n_ NEFilterReport) ShouldReport() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("shouldReport"))
	return rv
}


// SetShouldReport sets the value of the shouldReport property.
// A Boolean value that indicates whether to send a report to the control provider when processing this verdict.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterverdict/shouldreport
func (n_ NEFilterReport) SetShouldReport(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setShouldReport:"), value)
}



