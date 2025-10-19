// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SCScreenshotManager] class.
var (
	sCScreenshotManagerClass     _SCScreenshotManagerClass
	sCScreenshotManagerClassOnce sync.Once
)

func getSCScreenshotManagerClass() _SCScreenshotManagerClass {
	sCScreenshotManagerClassOnce.Do(func() {
		sCScreenshotManagerClass = _SCScreenshotManagerClass{objc.GetClass("SCScreenshotManager")}
	})
	return sCScreenshotManagerClass
}

type _SCScreenshotManagerClass struct {
	class objc.Class
}

// An interface definition for the [SCScreenshotManager] class.
type ISCScreenshotManager interface {
	objectivec.IObject
}

// An instance for the capture of single frames from a stream.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager
type SCScreenshotManager struct {
	objectivec.Object
}

// SCScreenshotManagerFrom constructs a [SCScreenshotManager] from an unsafe.Pointer.
//
// An instance for the capture of single frames from a stream.
func SCScreenshotManagerFrom(ptr unsafe.Pointer) SCScreenshotManager {
	return SCScreenshotManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SCScreenshotManagerClass) Alloc() SCScreenshotManager {
	rv := objc.Send[SCScreenshotManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SCScreenshotManagerClass) New() SCScreenshotManager {
	rv := objc.Send[SCScreenshotManager](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SCScreenshotManager) Init() SCScreenshotManager {
	rv := objc.Send[SCScreenshotManager](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SCScreenshotManager) Autorelease() SCScreenshotManager {
	rv := objc.Send[SCScreenshotManager](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCScreenshotManager creates a new SCScreenshotManager instance.
func NewSCScreenshotManager() SCScreenshotManager {
	return getSCScreenshotManagerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager/captureImage(in:completionHandler:)
func (sc _SCScreenshotManagerClass) CaptureImageInRectCompletionHandler(rect unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("captureImageInRect:completionHandler:"), rect, completionHandler)
}



