// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNScript] class.
var (
	CNScriptClass     _CNScriptClass
	CNScriptClassOnce sync.Once
)

func getCNScriptClass() _CNScriptClass {
	CNScriptClassOnce.Do(func() {
		CNScriptClass = _CNScriptClass{objc.GetClass("CNScript")}
	})
	return CNScriptClass
}

type _CNScriptClass struct {
	class objc.Class
}

// An interface definition for the [CNScript] class.
type ICNScript interface {
	objectivec.IObject
	DecisionBeforeTime(time unsafe.Pointer) unsafe.Pointer
	PrimaryDecisionAtTime(time unsafe.Pointer) unsafe.Pointer
	TimeRangeOfTransitionBeforeDecision(decision unsafe.Pointer) unsafe.Pointer
}

// A collection of focus decisions, focus transitions, detections, and detection tracks associated with a movie captured in Cinematic mode and methods to change them.
//
// The Cinematic script provides thread-safe access to information about the focus decisions made in the original recorded Cinematic movie. The script supports changing those decisions and obtaining updated information about where to focus each frame. You can snapshot changes to a script and later reload them.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn
type CNScript struct {
	objectivec.Object
}

// CNScriptFrom constructs a [CNScript] from an unsafe.Pointer.
//
// A collection of focus decisions, focus transitions, detections, and detection tracks associated with a movie captured in Cinematic mode and methods to change them.
func CNScriptFrom(ptr unsafe.Pointer) CNScript {
	return CNScript{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNScriptClass) Alloc() CNScript {
	rv := objc.Send[CNScript](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNScriptClass) New() CNScript {
	rv := objc.Send[CNScript](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNScript) Init() CNScript {
	rv := objc.Send[CNScript](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNScript) Autorelease() CNScript {
	rv := objc.Send[CNScript](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNScript creates a new CNScript instance.
func NewCNScript() CNScript {
	return getCNScriptClass().New()
}


// The decision that occurs before the given time.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/decisionBeforeTime:
func (c_ CNScript) DecisionBeforeTime(time unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("decisionBeforeTime:"), time)
	return rv
}

// The primary decision that’s in effect at the specified time, unless it’s outside the time range of the Cinematic script.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/primaryDecisionAtTime:
func (c_ CNScript) PrimaryDecisionAtTime(time unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("primaryDecisionAtTime:"), time)
	return rv
}

// The time range during which the focus transitioned toward the given decision.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/timeRangeOfTransitionBeforeDecision:
func (c_ CNScript) TimeRangeOfTransitionBeforeDecision(decision unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("timeRangeOfTransitionBeforeDecision:"), decision)
	return rv
}



