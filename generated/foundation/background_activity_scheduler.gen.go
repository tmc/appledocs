// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [BackgroundActivityScheduler] class.
var BackgroundActivitySchedulerClass objc.Class

func init() {
	BackgroundActivitySchedulerClass = objc.GetClass("NSBackgroundActivityScheduler")
}

type BackgroundActivityScheduler struct {
	objc.ID
}

func BackgroundActivitySchedulerFrom(ptr unsafe.Pointer) BackgroundActivityScheduler {
	return BackgroundActivityScheduler{
		ID: objc.ID(ptr),
	}
}



