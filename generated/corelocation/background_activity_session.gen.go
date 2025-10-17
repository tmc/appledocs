// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BackgroundActivitySession] class.
var backgroundActivitySessionClass = _BackgroundActivitySessionClass{objc.GetClass("CLBackgroundActivitySession")}

type _BackgroundActivitySessionClass struct {
	class objc.Class
}

// An object that manages a visual indicator that keeps your app in use in the background, allowing it to receive updates or events. [Full Topic]
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

// Creates a new background activity session. [Full Topic]

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
// Invalidates the background activity session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySession-4nl4y/invalidate
func (b_ BackgroundActivitySession) Invalidate() {
	objc.Send[objc.ID](b_.ID, objc.Sel("invalidate"))
}


