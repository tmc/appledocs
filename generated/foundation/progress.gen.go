// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Progress] class.
var progressClass = _ProgressClass{objc.GetClass("NSProgress")}

type _ProgressClass struct {
	class objc.Class
}

// An object that conveys ongoing progress to the user for a specified task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress

type Progress struct {
	objectivec.Object
}

// ProgressFrom constructs a [Progress] from an unsafe.Pointer.
//
// An object that conveys ongoing progress to the user for a specified task.
func ProgressFrom(ptr unsafe.Pointer) Progress {
	return Progress{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (pc _ProgressClass) Alloc() Progress {
	rv := objc.Send[Progress](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _ProgressClass) New() Progress {
	rv := objc.Send[Progress](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Progress) Init() Progress {
	rv := objc.Send[Progress](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Progress) Autorelease() Progress {
	rv := objc.Send[Progress](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProgress creates a new Progress instance.
func NewProgress() Progress {
	return progressClass.New()
}
// Creates a new progress instance. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/init(parent:userInfo:)
func NewProgressWithParentUserInfo(parentProgressOrNil unsafe.Pointer, userInfoOrNil unsafe.Pointer) Progress {
	instance := progressClass.Alloc()
	rv := objc.Send[Progress](instance.ID, objc.Sel("initWithParent:userInfo:"), parentProgressOrNil, userInfoOrNil)
	rv.Autorelease()
	return rv
}
// Creates and returns a progress instance. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/init(totalUnitCount:)
func NewProgressWithTotalUnitCount(unitCount unsafe.Pointer) Progress {
	rv := objc.Send[Progress](objc.ID(progressClass.class), objc.Sel("progressWithTotalUnitCount:"), unitCount)
	rv.Autorelease()
	return rv
}


// Creates and returns a progress instance. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/init(totalUnitCount:)
func (pc _ProgressClass) ProgressWithTotalUnitCount(unitCount unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("progressWithTotalUnitCount:"), unitCount)
	return rv
}
// Sets the progress object as the current object of the current thread, and assigns the amount of work for the next suboperation progress object to perform. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/becomeCurrent(withPendingUnitCount:)
func (p_ Progress) BecomeCurrentWithPendingUnitCount(unitCount unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("becomeCurrentWithPendingUnitCount:"), unitCount)
}
// Cancels progress tracking. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/cancel()
func (p_ Progress) Cancel() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancel"))
}

