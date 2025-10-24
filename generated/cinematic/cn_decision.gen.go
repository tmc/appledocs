// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNDecision] class.
var (
	CNDecisionClass     _CNDecisionClass
	CNDecisionClassOnce sync.Once
)

func getCNDecisionClass() _CNDecisionClass {
	CNDecisionClassOnce.Do(func() {
		CNDecisionClass = _CNDecisionClass{objc.GetClass("CNDecision")}
	})
	return CNDecisionClass
}

type _CNDecisionClass struct {
	class objc.Class
}

// An interface definition for the [CNDecision] class.
type ICNDecision interface {
	objectivec.IObject
	// properties:
	GroupDecision() bool
	StrongDecision() bool
	// methods:
}

// An object that represents a decision to focus on a particular detection, or group of detections, at a particular time.


// An object that represents a decision to focus on a particular detection, or group of detections, at a particular time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDecision-c.class
type CNDecision struct {
	objectivec.Object
}

// CNDecisionFrom constructs a [CNDecision] from an unsafe.Pointer.
//
// An object that represents a decision to focus on a particular detection, or group of detections, at a particular time.
func CNDecisionFrom(ptr unsafe.Pointer) CNDecision {
	return CNDecision{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNDecisionClass) Alloc() CNDecision {
	rv := objc.Send[CNDecision](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNDecisionClass) New() CNDecision {
	rv := objc.Send[CNDecision](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNDecision) Init() CNDecision {
	rv := objc.Send[CNDecision](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNDecision) Autorelease() CNDecision {
	rv := objc.Send[CNDecision](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNDecision creates a new CNDecision instance.
func NewCNDecision() CNDecision {
	return getCNDecisionClass().New()
}



// A flag representing whether this is a group decision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDecision-c.class/groupDecision
func (c_ CNDecision) GroupDecision() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("groupDecision"))
	return rv
}


// A flag representing whether this is a strong decision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDecision-c.class/strongDecision
func (c_ CNDecision) StrongDecision() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("strongDecision"))
	return rv
}



