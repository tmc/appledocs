// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [NEFilterBrowserFlow] class.
var (
	NEFilterBrowserFlowClass     _NEFilterBrowserFlowClass
	NEFilterBrowserFlowClassOnce sync.Once
)

func getNEFilterBrowserFlowClass() _NEFilterBrowserFlowClass {
	NEFilterBrowserFlowClassOnce.Do(func() {
		NEFilterBrowserFlowClass = _NEFilterBrowserFlowClass{objc.GetClass("NEFilterBrowserFlow")}
	})
	return NEFilterBrowserFlowClass
}

type _NEFilterBrowserFlowClass struct {
	class objc.Class
}





// An interface definition for the [NEFilterBrowserFlow] class.
type INEFilterBrowserFlow interface {
	INEFilterFlow
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEFilterBrowserFlowClass) Alloc() NEFilterBrowserFlow {
	rv := objc.Send[NEFilterBrowserFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterBrowserFlowClass) New() NEFilterBrowserFlow {
	rv := objc.Send[NEFilterBrowserFlow](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterBrowserFlow) Init() NEFilterBrowserFlow {
	rv := objc.Send[NEFilterBrowserFlow](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterBrowserFlow) Autorelease() NEFilterBrowserFlow {
	rv := objc.Send[NEFilterBrowserFlow](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterBrowserFlow creates a new NEFilterBrowserFlow instance.
func NewNEFilterBrowserFlow() NEFilterBrowserFlow {
	return getNEFilterBrowserFlowClass().New()
}





// A flow of network data, originating from a WebKit-based browser, that the filter examines.


// A flow of network data, originating from a WebKit-based browser, that the filter examines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterBrowserFlow
type NEFilterBrowserFlow struct {
	NEFilterFlow
}

// NEFilterBrowserFlowFrom constructs a [NEFilterBrowserFlow] from an unsafe.Pointer.
//
// A flow of network data, originating from a WebKit-based browser, that the filter examines.
func NEFilterBrowserFlowFrom(ptr unsafe.Pointer) NEFilterBrowserFlow {
	return NEFilterBrowserFlow{
		NEFilterFlow: NEFilterFlowFrom(ptr),
	}
}






























