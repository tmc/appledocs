// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ResponderState] class.
var (
	ResponderStateClass     _ResponderStateClass
	ResponderStateClassOnce sync.Once
)

func getResponderStateClass() _ResponderStateClass {
	ResponderStateClassOnce.Do(func() {
		ResponderStateClass = _ResponderStateClass{objc.GetClass("PKResponderState")}
	})
	return ResponderStateClass
}

type _ResponderStateClass struct {
	class objc.Class
}

// An interface definition for the [ResponderState] class.
type IResponderState interface {
	objectivec.IObject
	// properties:
	// methods:
}

// The state of PencilKit behavior related to a .
//
// Control the behavior of responders via the property.


// The state of PencilKit behavior related to a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKResponderState
type ResponderState struct {
	objectivec.Object
}

// ResponderStateFrom constructs a [ResponderState] from an unsafe.Pointer.
//
// The state of PencilKit behavior related to a .
func ResponderStateFrom(ptr unsafe.Pointer) ResponderState {
	return ResponderState{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _ResponderStateClass) Alloc() ResponderState {
	rv := objc.Send[ResponderState](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ResponderStateClass) New() ResponderState {
	rv := objc.Send[ResponderState](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ResponderState) Init() ResponderState {
	rv := objc.Send[ResponderState](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ResponderState) Autorelease() ResponderState {
	rv := objc.Send[ResponderState](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewResponderState creates a new ResponderState instance.
func NewResponderState() ResponderState {
	return getResponderStateClass().New()
}



