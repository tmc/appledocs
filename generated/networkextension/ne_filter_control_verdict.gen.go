// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [NEFilterControlVerdict] class.
var (
	NEFilterControlVerdictClass     _NEFilterControlVerdictClass
	NEFilterControlVerdictClassOnce sync.Once
)

func getNEFilterControlVerdictClass() _NEFilterControlVerdictClass {
	NEFilterControlVerdictClassOnce.Do(func() {
		NEFilterControlVerdictClass = _NEFilterControlVerdictClass{objc.GetClass("NEFilterControlVerdict")}
	})
	return NEFilterControlVerdictClass
}

type _NEFilterControlVerdictClass struct {
	class objc.Class
}





// An interface definition for the [NEFilterControlVerdict] class.
type INEFilterControlVerdict interface {
	INEFilterNewFlowVerdict
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEFilterControlVerdictClass) Alloc() NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterControlVerdictClass) New() NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterControlVerdict) Init() NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterControlVerdict) Autorelease() NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterControlVerdict creates a new NEFilterControlVerdict instance.
func NewNEFilterControlVerdict() NEFilterControlVerdict {
	return getNEFilterControlVerdictClass().New()
}





// The result from a filter control provider.


// The result from a filter control provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlVerdict
type NEFilterControlVerdict struct {
	NEFilterNewFlowVerdict
}

// NEFilterControlVerdictFrom constructs a [NEFilterControlVerdict] from an unsafe.Pointer.
//
// The result from a filter control provider.
func NEFilterControlVerdictFrom(ptr unsafe.Pointer) NEFilterControlVerdict {
	return NEFilterControlVerdict{
		NEFilterNewFlowVerdict: NEFilterNewFlowVerdictFrom(ptr),
	}
}










// Create a verdict that indicates to the system that all of the flow’s data should be allowed to pass to its final destination, and that the filtering rules have been updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlVerdict/allow(withUpdateRules:)
func (nc _NEFilterControlVerdictClass) AllowVerdictWithUpdateRules(updateRules bool) NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](objc.ID(nc.class), objc.Sel("allowVerdictWithUpdateRules:"), updateRules)
	return rv
}


// Create a verdict that indicates to the system that all of the flow’s data should be dropped, and that the filtering rules have been updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlVerdict/drop(withUpdateRules:)
func (nc _NEFilterControlVerdictClass) DropVerdictWithUpdateRules(updateRules bool) NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](objc.ID(nc.class), objc.Sel("dropVerdictWithUpdateRules:"), updateRules)
	return rv
}


// Create a verdict that indicates to the system that the filtering rules have been updated, and that the Filter Data Provider needs to make a decision about the flow’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlVerdict/updateRules()
func (nc _NEFilterControlVerdictClass) UpdateRules() NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](objc.ID(nc.class), objc.Sel("updateRules"))
	return rv
}






















