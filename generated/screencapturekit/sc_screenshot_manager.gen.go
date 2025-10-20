// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [ScreenshotManager] class.
var (
	ScreenshotManagerClass     _ScreenshotManagerClass
	ScreenshotManagerClassOnce sync.Once
)

func getScreenshotManagerClass() _ScreenshotManagerClass {
	ScreenshotManagerClassOnce.Do(func() {
		ScreenshotManagerClass = _ScreenshotManagerClass{objc.GetClass("SCScreenshotManager")}
	})
	return ScreenshotManagerClass
}

type _ScreenshotManagerClass struct {
	class objc.Class
}

// An interface definition for the [ScreenshotManager] class.
type IScreenshotManager interface {
	objectivec.IObject
}

// An instance for the capture of single frames from a stream.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager
type ScreenshotManager struct {
	objectivec.Object
}

// ScreenshotManagerFrom constructs a [ScreenshotManager] from an unsafe.Pointer.
//
// An instance for the capture of single frames from a stream.
func ScreenshotManagerFrom(ptr unsafe.Pointer) ScreenshotManager {
	return ScreenshotManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScreenshotManagerClass) Alloc() ScreenshotManager {
	rv := objc.Send[ScreenshotManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScreenshotManagerClass) New() ScreenshotManager {
	rv := objc.Send[ScreenshotManager](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScreenshotManager) Init() ScreenshotManager {
	rv := objc.Send[ScreenshotManager](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScreenshotManager) Autorelease() ScreenshotManager {
	rv := objc.Send[ScreenshotManager](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScreenshotManager creates a new ScreenshotManager instance.
func NewScreenshotManager() ScreenshotManager {
	return getScreenshotManagerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager/captureImage(in:completionHandler:)
func (sc _ScreenshotManagerClass) CaptureImageInRectCompletionHandler(rect coregraphics.CGRect, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("captureImageInRect:completionHandler:"), rect, completionHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager/captureScreenshot(contentFilter:configuration:completionHandler:)
func (sc _ScreenshotManagerClass) CaptureScreenshotWithFilterConfigurationCompletionHandler(contentFilter unsafe.Pointer, config unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("captureScreenshotWithFilter:configuration:completionHandler:"), contentFilter, config, completionHandler)
}



