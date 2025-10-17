// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Progress] class.
var ProgressClass objc.Class

func init() {
	ProgressClass = objc.GetClass("NSProgress")
}

type Progress struct {
	objc.ID
}

func ProgressFrom(ptr unsafe.Pointer) Progress {
	return Progress{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc Progress) Alloc() Progress {
	ret := objc.ID(ProgressClass).Send(objc.RegisterName("alloc"))
	return Progress{ret}
}

// Init initializes the instance.
func (p_ Progress) Init() Progress {
	ret := p_.ID.Send(objc.RegisterName("init"))
	return Progress{ret}
}
// Creates a new progress instance. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Progress/init(parent:userInfo:)
func NewProgressWithParentUserInfo(parentProgressOrNil unsafe.Pointer, userInfoOrNil unsafe.Pointer) Progress {
	instance := Progress{}.Alloc()
	sel := objc.RegisterName("initWithParent:userInfo:")
	ret := instance.ID.Send(sel, parentProgressOrNil, userInfoOrNil)
	instance = Progress{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Creates and returns a progress instance. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Progress/init(totalUnitCount:)
func (pc Progress) ProgressWithTotalUnitCount(unitCount unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("progressWithTotalUnitCount:")
	ret := objc.ID(ProgressClass).Send(sel, unitCount)
	return unsafe.Pointer(ret)
}
// Sets the progress object as the current object of the current thread, and assigns the amount of work for the next suboperation progress object to perform. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Progress/becomeCurrent(withPendingUnitCount:)
func (p_ Progress) BecomeCurrentWithPendingUnitCount(unitCount unsafe.Pointer) {
	sel := objc.RegisterName("becomeCurrentWithPendingUnitCount:")
	p_.ID.Send(sel, unitCount)
}
// Cancels progress tracking. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Progress/cancel()
func (p_ Progress) Cancel() {
	sel := objc.RegisterName("cancel")
	p_.ID.Send(sel)
}

