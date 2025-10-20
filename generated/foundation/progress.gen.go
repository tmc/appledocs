// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Progress] class.
var (
	progressClass     _ProgressClass
	progressClassOnce sync.Once
)

func getProgressClass() _ProgressClass {
	progressClassOnce.Do(func() {
		progressClass = _ProgressClass{objc.GetClass("NSProgress")}
	})
	return progressClass
}

type _ProgressClass struct {
	class objc.Class
}

// An interface definition for the [Progress] class.
type IProgress interface {
	objectivec.IObject
	BecomeCurrentWithPendingUnitCount(unitCount unsafe.Pointer)
	Cancel()
}

// An object that conveys ongoing progress to the user for a specified task.
//
// The class provides a self-contained mechanism for progress reporting. It makes it easy for code that performs work to report the progress of that work, and for user interface code to observe that progress for presentation to the user. Specifically, you can use a progress object to show the user a progress bar and explanatory text that update as you do work. It also allows the user to cancel or pause work.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getProgressClass().New()
}


// Creates a new progress instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/init(parent:userInfo:)
func NewProgressWithParentUserInfo(parentProgressOrNil unsafe.Pointer, userInfoOrNil unsafe.Pointer) Progress {
	instance := getProgressClass().Alloc()
	rv := objc.Send[Progress](instance.ID, objc.Sel("initWithParent:userInfo:"), parentProgressOrNil, userInfoOrNil)
	rv.Autorelease()
	return rv
}

// Creates and returns a progress instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/init(totalUnitCount:)
func NewProgressWithTotalUnitCount(unitCount unsafe.Pointer) Progress {
	rv := objc.Send[Progress](objc.ID(getProgressClass().class), objc.Sel("progressWithTotalUnitCount:"), unitCount)
	return rv
}


// Creates and returns a progress instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/init(totalUnitCount:)
func (pc _ProgressClass) ProgressWithTotalUnitCount(unitCount unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("progressWithTotalUnitCount:"), unitCount)
	return rv
}

// Sets the progress object as the current object of the current thread, and assigns the amount of work for the next suboperation progress object to perform.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/becomeCurrent(withPendingUnitCount:)
func (p_ Progress) BecomeCurrentWithPendingUnitCount(unitCount unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("becomeCurrentWithPendingUnitCount:"), unitCount)
}

// Cancels progress tracking.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/cancel()
func (p_ Progress) Cancel() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancel"))
}

// The number of completed units of work for the current job.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/completedUnitCount
func (p_ Progress) CompletedUnitCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("completedUnitCount"))
	return rv
}



// SetCompletedUnitCount sets the value of the completedUnitCount property.
// The number of completed units of work for the current job.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/completedUnitCount
func (p_ Progress) SetCompletedUnitCount(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCompletedUnitCount:"), value)
}
// The fraction of the overall work that the progress object completes, including work from its suboperations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/fractionCompleted
func (p_ Progress) FractionCompleted() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("fractionCompleted"))
	return rv
}


// A Boolean value that indicates the progress object is complete.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isFinished
func (p_ Progress) Finished() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("finished"))
	return rv
}


// The total number of tracked units of work for the current progress.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/totalUnitCount
func (p_ Progress) TotalUnitCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("totalUnitCount"))
	return rv
}



// SetTotalUnitCount sets the value of the totalUnitCount property.
// The total number of tracked units of work for the current progress.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/totalUnitCount
func (p_ Progress) SetTotalUnitCount(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTotalUnitCount:"), value)
}

