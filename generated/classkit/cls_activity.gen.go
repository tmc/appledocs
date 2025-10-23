// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AdditionalActivityItems() []SActivityItem /* primitive/slice/pointer. */
	Duration() foundation.TimeInterval /* not a class type */
	Started() bool /* primitive/slice/pointer. */
	PrimaryActivityItem() ICLSActivityItem
	SetPrimaryActivityItem(value ICLSActivityItem)
	Progress() float64 /* primitive/slice/pointer. */
	SetProgress(value float64 /* primitive/slice/pointer. */)
	IsStarted() bool /* primitive/slice/pointer. */
	SetIsStarted(value bool /* primitive/slice/pointer. */)
	// methods:
	AddAdditionalActivityItem(activityItem ICLSActivityItem)
	AddProgressRangeFromStartToEnd(start float64 /* primitive/slice/pointer. */, end float64 /* primitive/slice/pointer. */)
	RemoveAllActivityItems()
	Start()
	Stop()
}

// A representation of user interaction with a context.
//
// An activity represents a student’s attempt to complete the task corresponding to a instance. For example, if a context represents a quiz, the associated activity represents the student’s attempt to take the quiz. As such, an activity is always associated with a context. In fact, you never initialize an activity in isolation or store a reference to it. Rather, you ask a context to create the activity and retrieve it from the context.


// A representation of user interaction with a context.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/addAdditionalActivityItem(_:)
func (s_ SActivity) AddAdditionalActivityItem(activityItem ICLSActivityItem) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addAdditionalActivityItem:"), activityItem)
}


// Adds a progress range to a given activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/addProgressRange(fromStart:toEnd:)
func (s_ SActivity) AddProgressRangeFromStartToEnd(start float64 /* primitive/slice/pointer. */, end float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addProgressRangeFromStart:toEnd:"), start, end)
}


// Deletes all activity items associated with the current activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/removeAllActivityItems()
func (s_ SActivity) RemoveAllActivityItems() {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeAllActivityItems"))
}


// Tells an activity to start recording duration and progress for a task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/start()
func (s_ SActivity) Start() {
	objc.Send[objc.ID](s_.ID, objc.Sel("start"))
}


// Tells an activity to stop or pause recording duration and progress for a task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/stop()
func (s_ SActivity) Stop() {
	objc.Send[objc.ID](s_.ID, objc.Sel("stop"))
}


// The list of activity items associated with an activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/additionalActivityItems
func (s_ SActivity) AdditionalActivityItems() []SActivityItem /* primitive/slice/pointer. */ {
	rv := objc.Send[[]SActivityItem](s_.ID, objc.Sel("additionalActivityItems"))
	return rv
}


// The cumulative time in seconds that an activity has been active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/duration
func (s_ SActivity) Duration() foundation.TimeInterval /* not a class type */ {
	rv := objc.Send[foundation.TimeInterval](s_.ID, objc.Sel("duration"))
	return rv
}


// A Boolean that indicates whether an activity is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/isStarted
func (s_ SActivity) Started() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("started"))
	return rv
}


// Adds an activity item to an activity and sets it as the primary activity item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/primaryActivityItem
func (s_ SActivity) PrimaryActivityItem() ICLSActivityItem {
	rv := objc.Send[SActivityItem](s_.ID, objc.Sel("primaryActivityItem"))
	return rv
}


// Adds an activity item to an activity and sets it as the primary activity item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/primaryActivityItem
func (s_ SActivity) SetPrimaryActivityItem(value ICLSActivityItem) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPrimaryActivityItem:"), value)
}


// A measure of progress through the task, given as a fraction in the range [0, 1].
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/progress
func (s_ SActivity) Progress() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](s_.ID, objc.Sel("progress"))
	return rv
}


// A measure of progress through the task, given as a fraction in the range [0, 1].
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/progress
func (s_ SActivity) SetProgress(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setProgress:"), value)
}


// A Boolean that indicates whether an activity is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clsactivity/isstarted
func (s_ SActivity) IsStarted() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("isStarted"))
	return rv
}


// A Boolean that indicates whether an activity is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clsactivity/isstarted
func (s_ SActivity) SetIsStarted(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsStarted:"), value)
}



