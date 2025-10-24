// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CLSActivity */


/* debug [class_header]: Header for CLSActivity */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SActivity */
// An interface definition for the [SActivity] class.
type ISActivity interface {
	ISObject
	
/* debug [class_interface_properties]: Properties for SActivity */
	// properties:
	AdditionalActivityItems() []SActivityItem
	Duration() float64
	Started() bool
	PrimaryActivityItem() ICLSActivityItem
	SetPrimaryActivityItem(value ICLSActivityItem)
	Progress() float64
	SetProgress(value float64)
	IsStarted() bool
	SetIsStarted(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SActivity */
	// methods:
	AddAdditionalActivityItem(activityItem ICLSActivityItem)
	AddProgressRangeFromStartToEnd(start float64, end float64)
	RemoveAllActivityItems()
	Start()
	Stop()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SActivity */
// Alloc allocates a new instance without initialization.
func (sc _SActivityClass) Alloc() SActivity {
	rv := objc.Send[SActivity](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SActivity */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SActivity *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SActivity */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SActivity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SActivity */

// Adds an activity item to an activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/addAdditionalActivityItem(_:)
func (s_ SActivity) AddAdditionalActivityItem(activityItem ICLSActivityItem) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addAdditionalActivityItem:"), activityItem)
}/* debug [instance_methods/method]: AddAdditionalActivityItem */


// Adds a progress range to a given activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/addProgressRange(fromStart:toEnd:)
func (s_ SActivity) AddProgressRangeFromStartToEnd(start float64, end float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addProgressRangeFromStart:toEnd:"), start, end)
}/* debug [instance_methods/method]: AddProgressRangeFromStartToEnd */


// Deletes all activity items associated with the current activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/removeAllActivityItems()
func (s_ SActivity) RemoveAllActivityItems() {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeAllActivityItems"))
}/* debug [instance_methods/method]: RemoveAllActivityItems */


// Tells an activity to start recording duration and progress for a task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/start()
func (s_ SActivity) Start() {
	objc.Send[objc.ID](s_.ID, objc.Sel("start"))
}/* debug [instance_methods/method]: Start */


// Tells an activity to stop or pause recording duration and progress for a task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/stop()
func (s_ SActivity) Stop() {
	objc.Send[objc.ID](s_.ID, objc.Sel("stop"))
}/* debug [instance_methods/method]: Stop */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SActivity */

// The list of activity items associated with an activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/additionalActivityItems
func (s_ SActivity) AdditionalActivityItems() []SActivityItem {
	rv := objc.Send[[]SActivityItem](s_.ID, objc.Sel("additionalActivityItems"))
	return rv
}/* debug [instance_properties/getter]: additionalActivityItems */


// The cumulative time in seconds that an activity has been active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/duration
func (s_ SActivity) Duration() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// A Boolean that indicates whether an activity is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/isStarted
func (s_ SActivity) Started() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("started"))
	return rv
}/* debug [instance_properties/getter]: started */


// Adds an activity item to an activity and sets it as the primary activity item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/primaryActivityItem
func (s_ SActivity) PrimaryActivityItem() ICLSActivityItem {
	rv := objc.Send[SActivityItem](s_.ID, objc.Sel("primaryActivityItem"))
	return rv
}/* debug [instance_properties/getter]: primaryActivityItem */


// Adds an activity item to an activity and sets it as the primary activity item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/primaryActivityItem
func (s_ SActivity) SetPrimaryActivityItem(value ICLSActivityItem) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPrimaryActivityItem:"), value)
}/* debug [instance_properties/setter]: primaryActivityItem */


// A measure of progress through the task, given as a fraction in the range [0, 1].
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/progress
func (s_ SActivity) Progress() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("progress"))
	return rv
}/* debug [instance_properties/getter]: progress */


// A measure of progress through the task, given as a fraction in the range [0, 1].
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivity/progress
func (s_ SActivity) SetProgress(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setProgress:"), value)
}/* debug [instance_properties/setter]: progress */


// A Boolean that indicates whether an activity is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clsactivity/isstarted
func (s_ SActivity) IsStarted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isStarted"))
	return rv
}/* debug [instance_properties/getter]: isStarted */


// A Boolean that indicates whether an activity is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clsactivity/isstarted
func (s_ SActivity) SetIsStarted(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsStarted:"), value)
}/* debug [instance_properties/setter]: isStarted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CLSActivity */



