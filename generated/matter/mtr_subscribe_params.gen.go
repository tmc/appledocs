// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [MTRSubscribeParams] class.
type IMTRSubscribeParams interface {
	IMTRReadParams
	// properties:
	AutoResubscribe() objc.IObject /* cross-framework: NSNumber */
	SetAutoResubscribe(value objc.IObject /* cross-framework: NSNumber */)
	KeepPreviousSubscriptions() objc.IObject /* cross-framework: NSNumber */
	SetKeepPreviousSubscriptions(value objc.IObject /* cross-framework: NSNumber */)
	MaxInterval() objc.IObject /* cross-framework: NSNumber */
	SetMaxInterval(value objc.IObject /* cross-framework: NSNumber */)
	MinInterval() objc.IObject /* cross-framework: NSNumber */
	SetMinInterval(value objc.IObject /* cross-framework: NSNumber */)
	ShouldReplaceExistingSubscriptions() bool
	SetShouldReplaceExistingSubscriptions(value bool)
	ShouldReportEventsUrgently() bool
	SetShouldReportEventsUrgently(value bool)
	ShouldResubscribeAutomatically() bool
	SetShouldResubscribeAutomatically(value bool)
	// methods:
}



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

// Alloc allocates a new instance without initialization.
func (mc _MTRSubscribeParamsClass) Alloc() MTRSubscribeParams {
	rv := objc.Send[MTRSubscribeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/autoresubscribe
func (m_ MTRSubscribeParams) AutoResubscribe() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("autoResubscribe"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/autoresubscribe
func (m_ MTRSubscribeParams) SetAutoResubscribe(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutoResubscribe:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/keepprevioussubscriptions
func (m_ MTRSubscribeParams) KeepPreviousSubscriptions() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("keepPreviousSubscriptions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/keepprevioussubscriptions
func (m_ MTRSubscribeParams) SetKeepPreviousSubscriptions(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeepPreviousSubscriptions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/maxinterval
func (m_ MTRSubscribeParams) MaxInterval() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxInterval"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/maxinterval
func (m_ MTRSubscribeParams) SetMaxInterval(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxInterval:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/mininterval
func (m_ MTRSubscribeParams) MinInterval() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minInterval"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/mininterval
func (m_ MTRSubscribeParams) SetMinInterval(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinInterval:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldreplaceexistingsubscriptions
func (m_ MTRSubscribeParams) ShouldReplaceExistingSubscriptions() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldReplaceExistingSubscriptions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldreplaceexistingsubscriptions
func (m_ MTRSubscribeParams) SetShouldReplaceExistingSubscriptions(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldReplaceExistingSubscriptions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldreporteventsurgently
func (m_ MTRSubscribeParams) ShouldReportEventsUrgently() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldReportEventsUrgently"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldreporteventsurgently
func (m_ MTRSubscribeParams) SetShouldReportEventsUrgently(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldReportEventsUrgently:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldresubscribeautomatically
func (m_ MTRSubscribeParams) ShouldResubscribeAutomatically() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldResubscribeAutomatically"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldresubscribeautomatically
func (m_ MTRSubscribeParams) SetShouldResubscribeAutomatically(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldResubscribeAutomatically:"), value)
}



