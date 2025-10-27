// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [NEFilterDataVerdict] class.
var (
	NEFilterDataVerdictClass     _NEFilterDataVerdictClass
	NEFilterDataVerdictClassOnce sync.Once
)

func getNEFilterDataVerdictClass() _NEFilterDataVerdictClass {
	NEFilterDataVerdictClassOnce.Do(func() {
		NEFilterDataVerdictClass = _NEFilterDataVerdictClass{objc.GetClass("NEFilterDataVerdict")}
	})
	return NEFilterDataVerdictClass
}

type _NEFilterDataVerdictClass struct {
	class objc.Class
}





// An interface definition for the [NEFilterDataVerdict] class.
type INEFilterDataVerdict interface {
	INEFilterVerdict
	

	// properties:
	StatisticsReportFrequency() NEFilterReportFrequency
	SetStatisticsReportFrequency(value NEFilterReportFrequency)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEFilterDataVerdictClass) Alloc() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterDataVerdictClass) New() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterDataVerdict) Init() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterDataVerdict) Autorelease() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterDataVerdict creates a new NEFilterDataVerdict instance.
func NewNEFilterDataVerdict() NEFilterDataVerdict {
	return getNEFilterDataVerdictClass().New()
}





// The result from a filter data provder for subsequent chunks of data on a flow.
//
// Return this verdict type from the various methods of .


// The result from a filter data provder for subsequent chunks of data on a flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict
type NEFilterDataVerdict struct {
	NEFilterVerdict
}

// NEFilterDataVerdictFrom constructs a [NEFilterDataVerdict] from an unsafe.Pointer.
//
// The result from a filter data provder for subsequent chunks of data on a flow.
func NEFilterDataVerdictFrom(ptr unsafe.Pointer) NEFilterDataVerdict {
	return NEFilterDataVerdict{
		NEFilterVerdict: NEFilterVerdictFrom(ptr),
	}
}






// Creates a verdict that tells the system to pass a chunk of network data to its final destination, and specifies the next chunk of data to provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/init(passBytes:peekBytes:)
func NewNEFilterDataVerdictWithPassBytesPeekBytes(passBytes uint, peekBytes uint) NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(getNEFilterDataVerdictClass().class), objc.Sel("dataVerdictWithPassBytes:peekBytes:"), passBytes, peekBytes)
	return rv
}







// Creates a verdict that tells the system to pass the current chunk of network data and all subsequent data for the current flow to its final destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/allow()
func (nc _NEFilterDataVerdictClass) AllowVerdict() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("allowVerdict"))
	return rv
}


// Creates a verdict that tells the system to drop the current chunk of network data and all subsequent data for the current flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/drop()
func (nc _NEFilterDataVerdictClass) DropVerdict() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("dropVerdict"))
	return rv
}


// Creates a verdict that tells the system to pass a chunk of network data to its final destination, and specifies the next chunk of data to provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/init(passBytes:peekBytes:)
func (nc _NEFilterDataVerdictClass) DataVerdictWithPassBytesPeekBytes(passBytes uint, peekBytes uint) NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("dataVerdictWithPassBytes:peekBytes:"), passBytes, peekBytes)
	return rv
}


// Creates a verdict that tells the system that the Filter Control Provider needs to update the rules before making a decision about the flow’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/needRules()
func (nc _NEFilterDataVerdictClass) NeedRulesVerdict() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("needRulesVerdict"))
	return rv
}


// Creates a verdict that tells the system to pause the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/pause()
func (nc _NEFilterDataVerdictClass) PauseVerdict() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("pauseVerdict"))
	return rv
}


// Creates a verdict to drop the current chunk of network data and all subsequent data for the current flow, and provides a remediation URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/remediateVerdict(withRemediationURLMapKey:remediationButtonTextMapKey:)
func (nc _NEFilterDataVerdictClass) RemediateVerdictWithRemediationURLMapKeyRemediationButtonTextMapKey(remediationURLMapKey foundation.foundation.INSString, remediationButtonTextMapKey foundation.foundation.INSString) NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("remediateVerdictWithRemediationURLMapKey:remediationButtonTextMapKey:"), remediationURLMapKey, remediationButtonTextMapKey)
	return rv
}

















// The frequencty at which to provide flow statistics to the data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/statisticsReportFrequency
func (n_ NEFilterDataVerdict) StatisticsReportFrequency() NEFilterReportFrequency {
	rv := objc.Send[NEFilterReportFrequency](n_.ID, objc.Sel("statisticsReportFrequency"))
	return rv
}


// The frequencty at which to provide flow statistics to the data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/statisticsReportFrequency
func (n_ NEFilterDataVerdict) SetStatisticsReportFrequency(value NEFilterReportFrequency) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setStatisticsReportFrequency:"), value)
}







