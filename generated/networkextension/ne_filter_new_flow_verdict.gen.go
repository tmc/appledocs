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
	

	// properties:
	StatisticsReportFrequency() NEFilterReportFrequency
	SetStatisticsReportFrequency(value NEFilterReportFrequency)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEFilterNewFlowVerdictClass) Alloc() NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// The result from a filter data provder after the initial examination of a flow.


// The result from a filter data provder after the initial examination of a flow.
//
// [Full Topic]
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










// Create a verdict that indicates to the system that the all of the new flow’s data should be allowed to pass to its final destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/allow()
func (nc _NEFilterNewFlowVerdictClass) AllowVerdict() NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](objc.ID(nc.class), objc.Sel("allowVerdict"))
	return rv
}


// Create a verdict that indicates to the system that all of the new flow’s data should dropped, and the user should not be given the opportunity to request access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/drop()
func (nc _NEFilterNewFlowVerdictClass) DropVerdict() NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](objc.ID(nc.class), objc.Sel("dropVerdict"))
	return rv
}


// Create a verdict that indicates to the system that the filter needs to make a decision about a new flow after seeing a portion of the flow’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/filterDataVerdict(withFilterInbound:peekInboundBytes:filterOutbound:peekOutboundBytes:)
func (nc _NEFilterNewFlowVerdictClass) FilterDataVerdictWithFilterInboundPeekInboundBytesFilterOutboundPeekOutboundBytes(filterInbound bool, peekInboundBytes uint, filterOutbound bool, peekOutboundBytes uint) NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](objc.ID(nc.class), objc.Sel("filterDataVerdictWithFilterInbound:peekInboundBytes:filterOutbound:peekOutboundBytes:"), filterInbound, peekInboundBytes, filterOutbound, peekOutboundBytes)
	return rv
}


// Create a verdict that indicates to the system that the Filter Data Provider needs more information before it can make a decision about a new flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/needRules()
func (nc _NEFilterNewFlowVerdictClass) NeedRulesVerdict() NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](objc.ID(nc.class), objc.Sel("needRulesVerdict"))
	return rv
}


// Creates a verdict that tells the system to pause the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/pause()
func (nc _NEFilterNewFlowVerdictClass) PauseVerdict() NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](objc.ID(nc.class), objc.Sel("pauseVerdict"))
	return rv
}


// Create a verdict that indicates to the system that all of the new flow’s data should be dropped, but allow the user to request access by tapping or clicking on a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/remediateVerdict(withRemediationURLMapKey:remediationButtonTextMapKey:)
func (nc _NEFilterNewFlowVerdictClass) RemediateVerdictWithRemediationURLMapKeyRemediationButtonTextMapKey(remediationURLMapKey foundation.foundation.INSString, remediationButtonTextMapKey foundation.foundation.INSString) NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](objc.ID(nc.class), objc.Sel("remediateVerdictWithRemediationURLMapKey:remediationButtonTextMapKey:"), remediationURLMapKey, remediationButtonTextMapKey)
	return rv
}


// Create a verdict that indicates to the system that all of the new flow’s data should be allowed to pass to its final destination, but a string should first be appended to the new flow’s request URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/urlAppendStringVerdict(withMapKey:)
func (nc _NEFilterNewFlowVerdictClass) URLAppendStringVerdictWithMapKey(urlAppendMapKey foundation.foundation.INSString) NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](objc.ID(nc.class), objc.Sel("URLAppendStringVerdictWithMapKey:"), urlAppendMapKey)
	return rv
}

















// The frequency at which the data provider receives reports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/statisticsReportFrequency
func (n_ NEFilterNewFlowVerdict) StatisticsReportFrequency() NEFilterReportFrequency {
	rv := objc.Send[NEFilterReportFrequency](n_.ID, objc.Sel("statisticsReportFrequency"))
	return rv
}


// The frequency at which the data provider receives reports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/statisticsReportFrequency
func (n_ NEFilterNewFlowVerdict) SetStatisticsReportFrequency(value NEFilterReportFrequency) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setStatisticsReportFrequency:"), value)
}







