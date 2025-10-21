// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [BackgroundActivitySession] class.
var (
	BackgroundActivitySessionClass     _BackgroundActivitySessionClass
	BackgroundActivitySessionClassOnce sync.Once
)

func getBackgroundActivitySessionClass() _BackgroundActivitySessionClass {
	BackgroundActivitySessionClassOnce.Do(func() {
		BackgroundActivitySessionClass = _BackgroundActivitySessionClass{objc.GetClass("CLBackgroundActivitySession")}
	})
	return BackgroundActivitySessionClass
}

type _BackgroundActivitySessionClass struct {
	class objc.Class
}

// An interface definition for the [BackgroundActivitySession] class.
type IBackgroundActivitySession interface {
	objectivec.IObject
	Invalidate()
}

// An object that manages a visual indicator that keeps your app in use in the background, allowing it to receive updates or events.
//
// Use to start a background activity session that allows a when-in-use authorized app to receive location updates or monitoring events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySession-4nl4y
type BackgroundActivitySession struct {
	objectivec.Object
}

// BackgroundActivitySessionFrom constructs a [BackgroundActivitySession] from an unsafe.Pointer.
//
// An object that manages a visual indicator that keeps your app in use in the background, allowing it to receive updates or events.
func BackgroundActivitySessionFrom(ptr unsafe.Pointer) BackgroundActivitySession {
	return BackgroundActivitySession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BackgroundActivitySessionClass) Alloc() BackgroundActivitySession {
	rv := objc.Send[BackgroundActivitySession](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BackgroundActivitySessionClass) New() BackgroundActivitySession {
	rv := objc.Send[BackgroundActivitySession](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BackgroundActivitySession) Init() BackgroundActivitySession {
	rv := objc.Send[BackgroundActivitySession](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BackgroundActivitySession) Autorelease() BackgroundActivitySession {
	rv := objc.Send[BackgroundActivitySession](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBackgroundActivitySession creates a new BackgroundActivitySession instance.
func NewBackgroundActivitySession() BackgroundActivitySession {
	return getBackgroundActivitySessionClass().New()
}


// Creates a new background activity session.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySession-4nl4y/backgroundActivitySession
func (bc _BackgroundActivitySessionClass) BackgroundActivitySession() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("backgroundActivitySession"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySession-4nl4y/backgroundActivitySessionWithQueue:handler:
func (bc _BackgroundActivitySessionClass) BackgroundActivitySessionWithQueueHandler(queue unsafe.Pointer, handler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("backgroundActivitySessionWithQueue:handler:"), queue, handler)
	return rv
}

// Invalidates the background activity session.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySession-4nl4y/invalidate
func (b_ BackgroundActivitySession) Invalidate() {
	objc.Send[objc.ID](b_.ID, objc.Sel("invalidate"))
}



