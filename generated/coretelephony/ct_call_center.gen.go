// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CallCenter] class.
var (
	CallCenterClass     _CallCenterClass
	CallCenterClassOnce sync.Once
)

func getCallCenterClass() _CallCenterClass {
	CallCenterClassOnce.Do(func() {
		CallCenterClass = _CallCenterClass{objc.GetClass("CTCallCenter")}
	})
	return CallCenterClass
}

type _CallCenterClass struct {
	class objc.Class
}

// An interface definition for the [CallCenter] class.
type ICallCenter interface {
	objectivec.IObject
}

// An object that provides a list of current cellular calls, and provides the ability to respond to state changes for calls.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCallCenter
type CallCenter struct {
	objectivec.Object
}

// CallCenterFrom constructs a [CallCenter] from an unsafe.Pointer.
//
// An object that provides a list of current cellular calls, and provides the ability to respond to state changes for calls.
func CallCenterFrom(ptr unsafe.Pointer) CallCenter {
	return CallCenter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CallCenterClass) Alloc() CallCenter {
	rv := objc.Send[CallCenter](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CallCenterClass) New() CallCenter {
	rv := objc.Send[CallCenter](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CallCenter) Init() CallCenter {
	rv := objc.Send[CallCenter](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CallCenter) Autorelease() CallCenter {
	rv := objc.Send[CallCenter](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCallCenter creates a new CallCenter instance.
func NewCallCenter() CallCenter {
	return getCallCenterClass().New()
}


// An array representing the cellular calls in progress.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCallCenter/currentCalls
func (c_ CallCenter) CurrentCalls() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("currentCalls"))
	return rv
}



