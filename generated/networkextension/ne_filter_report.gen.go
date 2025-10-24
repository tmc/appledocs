// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEFilterReport */


/* debug [class_header]: Header for NEFilterReport */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFilterReport */
// An interface definition for the [NEFilterReport] class.
type INEFilterReport interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEFilterReport */
	// properties:
	Action() NEFilterAction
	BytesInboundCount() uint
	BytesOutboundCount() uint
	Event() NEFilterReportEvent
	Flow() INEFilterFlow
	ShouldReport() bool
	SetShouldReport(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFilterReport */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFilterReport */
// Alloc allocates a new instance without initialization.
func (nc _NEFilterReportClass) Alloc() NEFilterReport {
	rv := objc.Send[NEFilterReport](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFilterReport */
// The report of the data provider’s action on a flow.
//
// The system issues a report by calling your control provider’s method with a report instance when the data provider issues a verdict whose property is set to .


// The report of the data provider’s action on a flow.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFilterReport *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFilterReport */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFilterReport */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFilterReport */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFilterReport */

// The action taken on the reported flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/action
func (n_ NEFilterReport) Action() NEFilterAction {
	rv := objc.Send[NEFilterAction](n_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// The number of inbound bytes received from the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/bytesInboundCount
func (n_ NEFilterReport) BytesInboundCount() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("bytesInboundCount"))
	return rv
}/* debug [instance_properties/getter]: bytesInboundCount */


// The number of outbound bytes sent on the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/bytesOutboundCount
func (n_ NEFilterReport) BytesOutboundCount() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("bytesOutboundCount"))
	return rv
}/* debug [instance_properties/getter]: bytesOutboundCount */


// The type of event indicated by this report.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/event-swift.property
func (n_ NEFilterReport) Event() NEFilterReportEvent {
	rv := objc.Send[NEFilterReportEvent](n_.ID, objc.Sel("event"))
	return rv
}/* debug [instance_properties/getter]: event */


// The flow on which the associated action was taken.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/flow
func (n_ NEFilterReport) Flow() INEFilterFlow {
	rv := objc.Send[NEFilterFlow](n_.ID, objc.Sel("flow"))
	return rv
}/* debug [instance_properties/getter]: flow */


// A Boolean value that indicates whether to send a report to the control provider when processing this verdict.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterverdict/shouldreport
func (n_ NEFilterReport) ShouldReport() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("shouldReport"))
	return rv
}/* debug [instance_properties/getter]: shouldReport */


// A Boolean value that indicates whether to send a report to the control provider when processing this verdict.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterverdict/shouldreport
func (n_ NEFilterReport) SetShouldReport(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setShouldReport:"), value)
}/* debug [instance_properties/setter]: shouldReport */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFilterReport */



