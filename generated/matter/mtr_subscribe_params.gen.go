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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldresubscribeautomatically
func (m_ MTRSubscribeParams) ShouldResubscribeAutomatically() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldResubscribeAutomatically"))
	return rv
}


// SetShouldResubscribeAutomatically sets the value of the shouldResubscribeAutomatically property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldresubscribeautomatically
func (m_ MTRSubscribeParams) SetShouldResubscribeAutomatically(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldResubscribeAutomatically:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/maxinterval
func (m_ MTRSubscribeParams) MaxInterval() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("maxInterval"))
	return rv
}


// SetMaxInterval sets the value of the maxInterval property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/maxinterval
func (m_ MTRSubscribeParams) SetMaxInterval(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxInterval:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/keepprevioussubscriptions
func (m_ MTRSubscribeParams) KeepPreviousSubscriptions() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("keepPreviousSubscriptions"))
	return rv
}


// SetKeepPreviousSubscriptions sets the value of the keepPreviousSubscriptions property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/keepprevioussubscriptions
func (m_ MTRSubscribeParams) SetKeepPreviousSubscriptions(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeepPreviousSubscriptions:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldreporteventsurgently
func (m_ MTRSubscribeParams) ShouldReportEventsUrgently() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldReportEventsUrgently"))
	return rv
}


// SetShouldReportEventsUrgently sets the value of the shouldReportEventsUrgently property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldreporteventsurgently
func (m_ MTRSubscribeParams) SetShouldReportEventsUrgently(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldReportEventsUrgently:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/mininterval
func (m_ MTRSubscribeParams) MinInterval() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("minInterval"))
	return rv
}


// SetMinInterval sets the value of the minInterval property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/mininterval
func (m_ MTRSubscribeParams) SetMinInterval(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinInterval:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/autoresubscribe
func (m_ MTRSubscribeParams) AutoResubscribe() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("autoResubscribe"))
	return rv
}


// SetAutoResubscribe sets the value of the autoResubscribe property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/autoresubscribe
func (m_ MTRSubscribeParams) SetAutoResubscribe(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutoResubscribe:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldreplaceexistingsubscriptions
func (m_ MTRSubscribeParams) ShouldReplaceExistingSubscriptions() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldReplaceExistingSubscriptions"))
	return rv
}


// SetShouldReplaceExistingSubscriptions sets the value of the shouldReplaceExistingSubscriptions property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsubscribeparams/shouldreplaceexistingsubscriptions
func (m_ MTRSubscribeParams) SetShouldReplaceExistingSubscriptions(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldReplaceExistingSubscriptions:"), value)
}



