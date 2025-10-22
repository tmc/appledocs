// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEFilterNewFlowVerdict] class.
var (
	NEFilterNewFlowVerdictClass     _NEFilterNewFlowVerdictClass
	NEFilterNewFlowVerdictClassOnce sync.Once
)

func getNEFilterNewFlowVerdictClass() _NEFilterNewFlowVerdictClass {
	NEFilterNewFlowVerdictClassOnce.Do(func() {
		NEFilterNewFlowVerdictClass = _NEFilterNewFlowVerdictClass{objc.GetClass("NEFilterNewFlowVerdict")}
	})
	return NEFilterNewFlowVerdictClass
}

type _NEFilterNewFlowVerdictClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterNewFlowVerdict] class.
type INEFilterNewFlowVerdict interface {
	INEFilterVerdict
	StatisticsReportFrequency() unsafe.Pointer
	SetStatisticsReportFrequency(value unsafe.Pointer)
}

// The result from a filter data provder after the initial examination of a flow.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict
type NEFilterNewFlowVerdict struct {
	NEFilterVerdict
}

// NEFilterNewFlowVerdictFrom constructs a [NEFilterNewFlowVerdict] from an unsafe.Pointer.
//
// The result from a filter data provder after the initial examination of a flow.
func NEFilterNewFlowVerdictFrom(ptr unsafe.Pointer) NEFilterNewFlowVerdict {
	return NEFilterNewFlowVerdict{
		NEFilterVerdict: NEFilterVerdictFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterNewFlowVerdictClass) Alloc() NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterNewFlowVerdictClass) New() NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterNewFlowVerdict) Init() NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterNewFlowVerdict) Autorelease() NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterNewFlowVerdict creates a new NEFilterNewFlowVerdict instance.
func NewNEFilterNewFlowVerdict() NEFilterNewFlowVerdict {
	return getNEFilterNewFlowVerdictClass().New()
}


// The frequency at which the data provider receives reports.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilternewflowverdict/statisticsreportfrequency
func (n_ NEFilterNewFlowVerdict) StatisticsReportFrequency() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("statisticsReportFrequency"))
	return rv
}


// SetStatisticsReportFrequency sets the value of the statisticsReportFrequency property.
// The frequency at which the data provider receives reports.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilternewflowverdict/statisticsreportfrequency
func (n_ NEFilterNewFlowVerdict) SetStatisticsReportFrequency(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setStatisticsReportFrequency:"), value)
}



