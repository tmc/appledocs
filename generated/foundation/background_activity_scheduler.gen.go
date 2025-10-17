// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BackgroundActivityScheduler] class.
var backgroundActivitySchedulerClass = _BackgroundActivitySchedulerClass{objc.GetClass("NSBackgroundActivityScheduler")}

type _BackgroundActivitySchedulerClass struct {
	class objc.Class
}

// An interface definition for the [BackgroundActivityScheduler] class.
type IBackgroundActivityScheduler interface {
	objectivec.IObject
}

// A task scheduler suitable for low priority operations that can run in the background. [Full Topic]
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



