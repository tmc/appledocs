// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NEFilterDataVerdict */


/* debug [class_header]: Header for NEFilterDataVerdict */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFilterDataVerdict */
// An interface definition for the [NEFilterDataVerdict] class.
type INEFilterDataVerdict interface {
	INEFilterVerdict
	
/* debug [class_interface_properties]: Properties for NEFilterDataVerdict */
	// properties:
	StatisticsReportFrequency() NEFilterReportFrequency
	SetStatisticsReportFrequency(value NEFilterReportFrequency)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFilterDataVerdict */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFilterDataVerdict */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFilterDataVerdict */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFilterDataVerdict */

// Creates a verdict that tells the system to pass a chunk of network data to its final destination, and specifies the next chunk of data to provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/init(passBytes:peekBytes:)
func NewNEFilterDataVerdictWithPassBytesPeekBytes(passBytes uint, peekBytes uint) NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(getNEFilterDataVerdictClass().class), objc.Sel("dataVerdictWithPassBytes:peekBytes:"), passBytes, peekBytes)
	return rv
}/* debug [class_init_methods/constructor]: NewNEFilterDataVerdictWithPassBytesPeekBytes */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFilterDataVerdict */

// Creates a verdict that tells the system to pass the current chunk of network data and all subsequent data for the current flow to its final destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/allow()
func (nc _NEFilterDataVerdictClass) AllowVerdict() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("allowVerdict"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AllowVerdict) */


// Creates a verdict that tells the system to drop the current chunk of network data and all subsequent data for the current flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/drop()
func (nc _NEFilterDataVerdictClass) DropVerdict() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("dropVerdict"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DropVerdict) */


// Creates a verdict that tells the system to pass a chunk of network data to its final destination, and specifies the next chunk of data to provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/init(passBytes:peekBytes:)
func (nc _NEFilterDataVerdictClass) DataVerdictWithPassBytesPeekBytes(passBytes uint, peekBytes uint) NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("dataVerdictWithPassBytes:peekBytes:"), passBytes, peekBytes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataVerdictWithPassBytesPeekBytes) */


// Creates a verdict that tells the system that the Filter Control Provider needs to update the rules before making a decision about the flow’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/needRules()
func (nc _NEFilterDataVerdictClass) NeedRulesVerdict() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("needRulesVerdict"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NeedRulesVerdict) */


// Creates a verdict that tells the system to pause the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/pause()
func (nc _NEFilterDataVerdictClass) PauseVerdict() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("pauseVerdict"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PauseVerdict) */


// Creates a verdict to drop the current chunk of network data and all subsequent data for the current flow, and provides a remediation URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/remediateVerdict(withRemediationURLMapKey:remediationButtonTextMapKey:)
func (nc _NEFilterDataVerdictClass) RemediateVerdictWithRemediationURLMapKeyRemediationButtonTextMapKey(remediationURLMapKey objc.IObject /* cross-framework: NSString */, remediationButtonTextMapKey objc.IObject /* cross-framework: NSString */) NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("remediateVerdictWithRemediationURLMapKey:remediationButtonTextMapKey:"), remediationURLMapKey, remediationButtonTextMapKey)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RemediateVerdictWithRemediationURLMapKeyRemediationButtonTextMapKey) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFilterDataVerdict */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFilterDataVerdict */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFilterDataVerdict */

// The frequencty at which to provide flow statistics to the data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/statisticsReportFrequency
func (n_ NEFilterDataVerdict) StatisticsReportFrequency() NEFilterReportFrequency {
	rv := objc.Send[NEFilterReportFrequency](n_.ID, objc.Sel("statisticsReportFrequency"))
	return rv
}/* debug [instance_properties/getter]: statisticsReportFrequency */


// The frequencty at which to provide flow statistics to the data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/statisticsReportFrequency
func (n_ NEFilterDataVerdict) SetStatisticsReportFrequency(value NEFilterReportFrequency) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setStatisticsReportFrequency:"), value)
}/* debug [instance_properties/setter]: statisticsReportFrequency */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFilterDataVerdict */


