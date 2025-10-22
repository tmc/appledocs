// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BackgroundActivityScheduler] class.
var (
	BackgroundActivitySchedulerClass     _BackgroundActivitySchedulerClass
	BackgroundActivitySchedulerClassOnce sync.Once
)

func getBackgroundActivitySchedulerClass() _BackgroundActivitySchedulerClass {
	BackgroundActivitySchedulerClassOnce.Do(func() {
		BackgroundActivitySchedulerClass = _BackgroundActivitySchedulerClass{objc.GetClass("NSBackgroundActivityScheduler")}
	})
	return BackgroundActivitySchedulerClass
}

type _BackgroundActivitySchedulerClass struct {
	class objc.Class
}

// An interface definition for the [BackgroundActivityScheduler] class.
type IBackgroundActivityScheduler interface {
	objectivec.IObject
	Invalidate()
	ScheduleWithBlock(block unsafe.Pointer)
	Identifier() string
	Interval() TimeInterval
	SetInterval(value ITimeInterval)
	QualityOfService() QualityOfService
	SetQualityOfService(value IQualityOfService)
	Repeats() bool
	SetRepeats(value bool)
	ShouldDefer() bool
	Tolerance() TimeInterval
	SetTolerance(value ITimeInterval)
}

// A task scheduler suitable for low priority operations that can run in the background.
//
// Use an object to schedule an arbitrary maintenance or background task. It’s similar to an object, in that it lets you schedule a repeating or non-repeating task. However, gives the system flexibility to determine the most efficient time to execute based on energy usage, thermal conditions, and CPU use. For example, use an object to schedule: Automatic saves Backups Data maintenance Periodic content fetches Installation of updates Activities occurring in intervals of 10 minutes or more Any other deferrable task For information about performing non-deferrable tasks efficiently, see in .


// A task scheduler suitable for low priority operations that can run in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler

type BackgroundActivityScheduler struct {
	objectivec.Object
}

// BackgroundActivitySchedulerFrom constructs a [BackgroundActivityScheduler] from an unsafe.Pointer.
//
// A task scheduler suitable for low priority operations that can run in the background.
func BackgroundActivitySchedulerFrom(ptr unsafe.Pointer) BackgroundActivityScheduler {
	return BackgroundActivityScheduler{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BackgroundActivitySchedulerClass) Alloc() BackgroundActivityScheduler {
	rv := objc.Send[BackgroundActivityScheduler](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BackgroundActivitySchedulerClass) New() BackgroundActivityScheduler {
	rv := objc.Send[BackgroundActivityScheduler](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BackgroundActivityScheduler) Init() BackgroundActivityScheduler {
	rv := objc.Send[BackgroundActivityScheduler](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BackgroundActivityScheduler) Autorelease() BackgroundActivityScheduler {
	rv := objc.Send[BackgroundActivityScheduler](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBackgroundActivityScheduler creates a new BackgroundActivityScheduler instance.
func NewBackgroundActivityScheduler() BackgroundActivityScheduler {
	return getBackgroundActivitySchedulerClass().New()
}




// Initializes a background activity scheduler object with a specified unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/init(identifier:)

func NewBackgroundActivitySchedulerWithIdentifier(identifier string) BackgroundActivityScheduler {
	instance := getBackgroundActivitySchedulerClass().Alloc()
	rv := objc.Send[BackgroundActivityScheduler](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(identifier))
	rv.Autorelease()
	return rv
}




// Prevents the background activity from being scheduled again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/invalidate()

func (b_ BackgroundActivityScheduler) Invalidate() {
	objc.Send[objc.ID](b_.ID, objc.Sel("invalidate"))
}



// Begins scheduling the background activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/schedule(_:)

func (b_ BackgroundActivityScheduler) ScheduleWithBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("scheduleWithBlock:"), block)
}


// A unique reverse DNS notation string, such as , that identifies the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/identifier

func (b_ BackgroundActivityScheduler) Identifier() string {
	rv := objc.Send[string](b_.ID, objc.Sel("identifier"))
	return rv
}


// An integer providing a suggested interval between scheduling and invoking the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/interval

func (b_ BackgroundActivityScheduler) Interval() TimeInterval {
	rv := objc.Send[TimeInterval](b_.ID, objc.Sel("interval"))
	return rv
}


// An integer providing a suggested interval between scheduling and invoking the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/interval

func (b_ BackgroundActivityScheduler) SetInterval(value ITimeInterval) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setInterval:"), value)
}


// A value of type , which controls how aggressively the system schedules the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/qualityOfService

func (b_ BackgroundActivityScheduler) QualityOfService() QualityOfService {
	rv := objc.Send[QualityOfService](b_.ID, objc.Sel("qualityOfService"))
	return rv
}


// A value of type , which controls how aggressively the system schedules the activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/qualityOfService

func (b_ BackgroundActivityScheduler) SetQualityOfService(value IQualityOfService) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setQualityOfService:"), value)
}


// A Boolean value indicating whether the activity should be rescheduled after it completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/repeats

func (b_ BackgroundActivityScheduler) Repeats() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("repeats"))
	return rv
}


// A Boolean value indicating whether the activity should be rescheduled after it completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/repeats

func (b_ BackgroundActivityScheduler) SetRepeats(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setRepeats:"), value)
}


// A Boolean value indicating whether your app should stop performing background activity and resume at a more optimal time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/shouldDefer

func (b_ BackgroundActivityScheduler) ShouldDefer() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("shouldDefer"))
	return rv
}


// A value of type , which specifies a range of time during which the background activity may occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/tolerance

func (b_ BackgroundActivityScheduler) Tolerance() TimeInterval {
	rv := objc.Send[TimeInterval](b_.ID, objc.Sel("tolerance"))
	return rv
}


// A value of type , which specifies a range of time during which the background activity may occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/tolerance

func (b_ BackgroundActivityScheduler) SetTolerance(value ITimeInterval) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTolerance:"), value)
}


