// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [NEFilterRemediationVerdict] class.
var (
	NEFilterRemediationVerdictClass     _NEFilterRemediationVerdictClass
	NEFilterRemediationVerdictClassOnce sync.Once
)

func getNEFilterRemediationVerdictClass() _NEFilterRemediationVerdictClass {
	NEFilterRemediationVerdictClassOnce.Do(func() {
		NEFilterRemediationVerdictClass = _NEFilterRemediationVerdictClass{objc.GetClass("NEFilterRemediationVerdict")}
	})
	return NEFilterRemediationVerdictClass
}

type _NEFilterRemediationVerdictClass struct {
	class objc.Class
}





// An interface definition for the [NEFilterRemediationVerdict] class.
type INEFilterRemediationVerdict interface {
	INEFilterVerdict
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEFilterRemediationVerdictClass) Alloc() NEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterRemediationVerdictClass) New() NEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterRemediationVerdict) Init() NEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterRemediationVerdict) Autorelease() NEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterRemediationVerdict creates a new NEFilterRemediationVerdict instance.
func NewNEFilterRemediationVerdict() NEFilterRemediationVerdict {
	return getNEFilterRemediationVerdictClass().New()
}





// The result from a filter data provider after the user requests remediation for a blocked flow.


// The result from a filter data provider after the user requests remediation for a blocked flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRemediationVerdict
type NEFilterRemediationVerdict struct {
	NEFilterVerdict
}

// NEFilterRemediationVerdictFrom constructs a [NEFilterRemediationVerdict] from an unsafe.Pointer.
//
// The result from a filter data provider after the user requests remediation for a blocked flow.
func NEFilterRemediationVerdictFrom(ptr unsafe.Pointer) NEFilterRemediationVerdict {
	return NEFilterRemediationVerdict{
		NEFilterVerdict: NEFilterVerdictFrom(ptr),
	}
}










// Create a verdict that indicates to the system that the Filter Data Provider will allow the flow to pass to its final destination when/if the flow is requested again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRemediationVerdict/allow()
func (nc _NEFilterRemediationVerdictClass) AllowVerdict() NEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](objc.ID(nc.class), objc.Sel("allowVerdict"))
	return rv
}


// Create a verdict that indicates to the system that the Filter Data Provider will continue to block the flow of network data if/when the flow is requested again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRemediationVerdict/drop()
func (nc _NEFilterRemediationVerdictClass) DropVerdict() NEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](objc.ID(nc.class), objc.Sel("dropVerdict"))
	return rv
}


// Create a verdict that indicates to the system that the Filter Data Provider needs the filtering rules to be updated before it can make a remediation decision about the current flow of network data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRemediationVerdict/needRules()
func (nc _NEFilterRemediationVerdictClass) NeedRulesVerdict() NEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](objc.ID(nc.class), objc.Sel("needRulesVerdict"))
	return rv
}






















