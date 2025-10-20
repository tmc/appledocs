// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var BackgroundActivitySchedulerClass _BackgroundActivitySchedulerClass

func init() {
	BackgroundActivitySchedulerClass = _BackgroundActivitySchedulerClass{objc.GetClass("NSBackgroundActivityScheduler")}
}

type _BackgroundActivitySchedulerClass struct {
	class objc.Class
}

type BackgroundActivityScheduler struct {
	objc.ID
}

func BackgroundActivitySchedulerFrom(ptr unsafe.Pointer) BackgroundActivityScheduler {
	return BackgroundActivityScheduler{
		ID: objc.ID(ptr),
	}
}




