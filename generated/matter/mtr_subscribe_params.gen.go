// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRSubscribeParams */


/* debug [class_header]: Header for MTRSubscribeParams */
// The class instance for the [MTRSubscribeParams] class.
var (
	MTRSubscribeParamsClass     _MTRSubscribeParamsClass
	MTRSubscribeParamsClassOnce sync.Once
)

func getMTRSubscribeParamsClass() _MTRSubscribeParamsClass {
	MTRSubscribeParamsClassOnce.Do(func() {
		MTRSubscribeParamsClass = _MTRSubscribeParamsClass{objc.GetClass("MTRSubscribeParams")}
	})
	return MTRSubscribeParamsClass
}

type _MTRSubscribeParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRSubscribeParams */
// An interface definition for the [MTRSubscribeParams] class.
type IMTRSubscribeParams interface {
	IMTRReadParams
	
/* debug [class_interface_properties]: Properties for MTRSubscribeParams */
	// properties:
	AutoResubscribe() objc.IObject /* cross-framework: NSNumber */
	SetAutoResubscribe(value objc.IObject /* cross-framework: NSNumber */)
	KeepPreviousSubscriptions() objc.IObject /* cross-framework: NSNumber */
	SetKeepPreviousSubscriptions(value objc.IObject /* cross-framework: NSNumber */)
	MaxInterval() objc.IObject /* cross-framework: NSNumber */
	SetMaxInterval(value objc.IObject /* cross-framework: NSNumber */)
	MinInterval() objc.IObject /* cross-framework: NSNumber */
	SetMinInterval(value objc.IObject /* cross-framework: NSNumber */)
	ReplaceExistingSubscriptions() bool
	SetReplaceExistingSubscriptions(value bool)
	ReportEventsUrgently() bool
	SetReportEventsUrgently(value bool)
	ResubscribeAutomatically() bool
	SetResubscribeAutomatically(value bool)
	ShouldReplaceExistingSubscriptions() bool
	SetShouldReplaceExistingSubscriptions(value bool)
	ShouldReportEventsUrgently() bool
	SetShouldReportEventsUrgently(value bool)
	ShouldResubscribeAutomatically() bool
	SetShouldResubscribeAutomatically(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRSubscribeParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRSubscribeParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRSubscribeParamsClass) Alloc() MTRSubscribeParams {
	rv := objc.Send[MTRSubscribeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRSubscribeParamsClass) New() MTRSubscribeParams {
	rv := objc.Send[MTRSubscribeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSubscribeParams) Init() MTRSubscribeParams {
	rv := objc.Send[MTRSubscribeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSubscribeParams) Autorelease() MTRSubscribeParams {
	rv := objc.Send[MTRSubscribeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSubscribeParams creates a new MTRSubscribeParams instance.
func NewMTRSubscribeParams() MTRSubscribeParams {
	return getMTRSubscribeParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRSubscribeParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams
type MTRSubscribeParams struct {
	MTRReadParams
}

// MTRSubscribeParamsFrom constructs a [MTRSubscribeParams] from an unsafe.Pointer.
func MTRSubscribeParamsFrom(ptr unsafe.Pointer) MTRSubscribeParams {
	return MTRSubscribeParams{
		MTRReadParams: MTRReadParamsFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRSubscribeParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/init(minInterval:maxInterval:)
func NewMTRSubscribeParamsWithMinIntervalMaxInterval(minInterval objc.IObject /* cross-framework: NSNumber */, maxInterval objc.IObject /* cross-framework: NSNumber */) MTRSubscribeParams {
	instance := getMTRSubscribeParamsClass().Alloc()
	rv := objc.Send[MTRSubscribeParams](instance.ID, objc.Sel("initWithMinInterval:maxInterval:"), minInterval, maxInterval)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRSubscribeParamsWithMinIntervalMaxInterval */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRSubscribeParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRSubscribeParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRSubscribeParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRSubscribeParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/autoResubscribe
func (m_ MTRSubscribeParams) AutoResubscribe() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("autoResubscribe"))
	return rv
}/* debug [instance_properties/getter]: autoResubscribe */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/autoResubscribe
func (m_ MTRSubscribeParams) SetAutoResubscribe(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutoResubscribe:"), value)
}/* debug [instance_properties/setter]: autoResubscribe */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/keepPreviousSubscriptions
func (m_ MTRSubscribeParams) KeepPreviousSubscriptions() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("keepPreviousSubscriptions"))
	return rv
}/* debug [instance_properties/getter]: keepPreviousSubscriptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/keepPreviousSubscriptions
func (m_ MTRSubscribeParams) SetKeepPreviousSubscriptions(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeepPreviousSubscriptions:"), value)
}/* debug [instance_properties/setter]: keepPreviousSubscriptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/maxInterval
func (m_ MTRSubscribeParams) MaxInterval() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxInterval"))
	return rv
}/* debug [instance_properties/getter]: maxInterval */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/maxInterval
func (m_ MTRSubscribeParams) SetMaxInterval(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxInterval:"), value)
}/* debug [instance_properties/setter]: maxInterval */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/minInterval
func (m_ MTRSubscribeParams) MinInterval() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minInterval"))
	return rv
}/* debug [instance_properties/getter]: minInterval */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/minInterval
func (m_ MTRSubscribeParams) SetMinInterval(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinInterval:"), value)
}/* debug [instance_properties/setter]: minInterval */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/shouldReplaceExistingSubscriptions
func (m_ MTRSubscribeParams) ReplaceExistingSubscriptions() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("replaceExistingSubscriptions"))
	return rv
}/* debug [instance_properties/getter]: replaceExistingSubscriptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/shouldReplaceExistingSubscriptions
func (m_ MTRSubscribeParams) SetReplaceExistingSubscriptions(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReplaceExistingSubscriptions:"), value)
}/* debug [instance_properties/setter]: replaceExistingSubscriptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/shouldReportEventsUrgently
func (m_ MTRSubscribeParams) ReportEventsUrgently() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("reportEventsUrgently"))
	return rv
}/* debug [instance_properties/getter]: reportEventsUrgently */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/shouldReportEventsUrgently
func (m_ MTRSubscribeParams) SetReportEventsUrgently(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReportEventsUrgently:"), value)
}/* debug [instance_properties/setter]: reportEventsUrgently */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/shouldResubscribeAutomatically
func (m_ MTRSubscribeParams) ResubscribeAutomatically() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("resubscribeAutomatically"))
	return rv
}/* debug [instance_properties/getter]: resubscribeAutomatically */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSubscribeParams/shouldResubscribeAutomatically
func (m_ MTRSubscribeParams) SetResubscribeAutomatically(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResubscribeAutomatically:"), value)
}/* debug [instance_properties/setter]: resubscribeAutomatically */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldreplaceexistingsubscriptions
func (m_ MTRSubscribeParams) ShouldReplaceExistingSubscriptions() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldReplaceExistingSubscriptions"))
	return rv
}/* debug [instance_properties/getter]: shouldReplaceExistingSubscriptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldreplaceexistingsubscriptions
func (m_ MTRSubscribeParams) SetShouldReplaceExistingSubscriptions(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldReplaceExistingSubscriptions:"), value)
}/* debug [instance_properties/setter]: shouldReplaceExistingSubscriptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldreporteventsurgently
func (m_ MTRSubscribeParams) ShouldReportEventsUrgently() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldReportEventsUrgently"))
	return rv
}/* debug [instance_properties/getter]: shouldReportEventsUrgently */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldreporteventsurgently
func (m_ MTRSubscribeParams) SetShouldReportEventsUrgently(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldReportEventsUrgently:"), value)
}/* debug [instance_properties/setter]: shouldReportEventsUrgently */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldresubscribeautomatically
func (m_ MTRSubscribeParams) ShouldResubscribeAutomatically() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldResubscribeAutomatically"))
	return rv
}/* debug [instance_properties/getter]: shouldResubscribeAutomatically */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldresubscribeautomatically
func (m_ MTRSubscribeParams) SetShouldResubscribeAutomatically(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldResubscribeAutomatically:"), value)
}/* debug [instance_properties/setter]: shouldResubscribeAutomatically */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRSubscribeParams */


