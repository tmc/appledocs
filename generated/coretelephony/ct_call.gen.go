// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Call] class.
var (
	CallClass     _CallClass
	CallClassOnce sync.Once
)

func getCallClass() _CallClass {
	CallClassOnce.Do(func() {
		CallClass = _CallClass{objc.GetClass("CTCall")}
	})
	return CallClass
}

type _CallClass struct {
	class objc.Class
}

// An interface definition for the [Call] class.
type ICall interface {
	objectivec.IObject
}

// An object used to identify a cellular call and determine its state.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCall
type Call struct {
	objectivec.Object
}

// CallFrom constructs a [Call] from an unsafe.Pointer.
//
// An object used to identify a cellular call and determine its state.
func CallFrom(ptr unsafe.Pointer) Call {
	return Call{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CallClass) Alloc() Call {
	rv := objc.Send[Call](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CallClass) New() Call {
	rv := objc.Send[Call](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Call) Init() Call {
	rv := objc.Send[Call](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Call) Autorelease() Call {
	rv := objc.Send[Call](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCall creates a new Call instance.
func NewCall() Call {
	return getCallClass().New()
}


// A unique identifier for the cellular call.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCall/callID
func (c_ Call) CallID() string {
	rv := objc.Send[string](c_.ID, objc.Sel("callID"))
	return rv
}

// The state of the cellular call.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCall/callState
func (c_ Call) CallState() string {
	rv := objc.Send[string](c_.ID, objc.Sel("callState"))
	return rv
}



