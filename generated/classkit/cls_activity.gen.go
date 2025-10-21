// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SActivity] class.
var (
	SActivityClass     _SActivityClass
	SActivityClassOnce sync.Once
)

func getSActivityClass() _SActivityClass {
	SActivityClassOnce.Do(func() {
		SActivityClass = _SActivityClass{objc.GetClass("CLSActivity")}
	})
	return SActivityClass
}

type _SActivityClass struct {
	class objc.Class
}

// An interface definition for the [SActivity] class.
type ISActivity interface {
	ISObject
	AddAdditionalActivityItem(activityItem unsafe.Pointer)
	AddProgressRangeFromStartToEnd(start unsafe.Pointer, end unsafe.Pointer)
	RemoveAllActivityItems()
	Start()
	Stop()
}

// A representation of user interaction with a context.
//
// An activity represents a student’s attempt to complete the task corresponding to a instance. For example, if a context represents a quiz, the associated activity represents the student’s attempt to take the quiz. As such, an activity is always associated with a context. In fact, you never initialize an activity in isolation or store a reference to it. Rather, you ask a context to create the activity and retrieve it from the context.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity
type SActivity struct {
	SObject
}

// SActivityFrom constructs a [SActivity] from an unsafe.Pointer.
//
// A representation of user interaction with a context.
func SActivityFrom(ptr unsafe.Pointer) SActivity {
	return SActivity{
		SObject: SObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SActivityClass) Alloc() SActivity {
	rv := objc.Send[SActivity](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SActivityClass) New() SActivity {
	rv := objc.Send[SActivity](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SActivity) Init() SActivity {
	rv := objc.Send[SActivity](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SActivity) Autorelease() SActivity {
	rv := objc.Send[SActivity](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSActivity creates a new SActivity instance.
func NewSActivity() SActivity {
	return getSActivityClass().New()
}


// Adds an activity item to an activity.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/addAdditionalActivityItem(_:)
func (s_ SActivity) AddAdditionalActivityItem(activityItem unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addAdditionalActivityItem:"), activityItem)
}

// Adds a progress range to a given activity.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/addProgressRange(fromStart:toEnd:)
func (s_ SActivity) AddProgressRangeFromStartToEnd(start unsafe.Pointer, end unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addProgressRangeFromStart:toEnd:"), start, end)
}

// Deletes all activity items associated with the current activity.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/removeAllActivityItems()
func (s_ SActivity) RemoveAllActivityItems() {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeAllActivityItems"))
}

// Tells an activity to start recording duration and progress for a task.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/start()
func (s_ SActivity) Start() {
	objc.Send[objc.ID](s_.ID, objc.Sel("start"))
}

// Tells an activity to stop or pause recording duration and progress for a task.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/stop()
func (s_ SActivity) Stop() {
	objc.Send[objc.ID](s_.ID, objc.Sel("stop"))
}

// The list of activity items associated with an activity.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/additionalActivityItems
func (s_ SActivity) AdditionalActivityItems() []SActivityItem {
	rv := objc.Send[[]SActivityItem](s_.ID, objc.Sel("additionalActivityItems"))
	return rv
}

// The cumulative time in seconds that an activity has been active.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/duration
func (s_ SActivity) Duration() TimeInterval {
	rv := objc.Send[TimeInterval](s_.ID, objc.Sel("duration"))
	return rv
}

// A Boolean that indicates whether an activity is running.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/isStarted
func (s_ SActivity) Started() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("started"))
	return rv
}

// Adds an activity item to an activity and sets it as the primary activity item.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/primaryActivityItem
func (s_ SActivity) PrimaryActivityItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("primaryActivityItem"))
	return rv
}


// SetPrimaryActivityItem sets the value of the primaryActivityItem property.
// Adds an activity item to an activity and sets it as the primary activity item.

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/primaryActivityItem
func (s_ SActivity) SetPrimaryActivityItem(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPrimaryActivityItem:"), value)
}
// A measure of progress through the task, given as a fraction in the range [0, 1].
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/progress
func (s_ SActivity) Progress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("progress"))
	return rv
}


// SetProgress sets the value of the progress property.
// A measure of progress through the task, given as a fraction in the range [0, 1].

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/progress
func (s_ SActivity) SetProgress(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setProgress:"), value)
}


