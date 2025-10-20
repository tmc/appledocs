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
	backgroundActivitySchedulerClass     _BackgroundActivitySchedulerClass
	backgroundActivitySchedulerClassOnce sync.Once
)

func getBackgroundActivitySchedulerClass() _BackgroundActivitySchedulerClass {
	backgroundActivitySchedulerClassOnce.Do(func() {
		backgroundActivitySchedulerClass = _BackgroundActivitySchedulerClass{objc.GetClass("NSBackgroundActivityScheduler")}
	})
	return backgroundActivitySchedulerClass
}

type _BackgroundActivitySchedulerClass struct {
	class objc.Class
}

// An interface definition for the [BackgroundActivityScheduler] class.
type IBackgroundActivityScheduler interface {
	objectivec.IObject
}

// A task scheduler suitable for low priority operations that can run in the background.
//
// Use an object to schedule an arbitrary maintenance or background task. It’s similar to an object, in that it lets you schedule a repeating or non-repeating task. However, gives the system flexibility to determine the most efficient time to execute based on energy usage, thermal conditions, and CPU use. For example, use an object to schedule: Automatic saves Backups Data maintenance Periodic content fetches Installation of updates Activities occurring in intervals of 10 minutes or more Any other deferrable task For information about performing non-deferrable tasks efficiently, see in .
//
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




