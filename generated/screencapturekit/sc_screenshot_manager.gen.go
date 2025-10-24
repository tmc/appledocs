// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SCScreenshotManager */


/* debug [class_header]: Header for SCScreenshotManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScreenshotManager */
// An interface definition for the [ScreenshotManager] class.
type IScreenshotManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ScreenshotManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScreenshotManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScreenshotManager */
// Alloc allocates a new instance without initialization.
func (sc _ScreenshotManagerClass) Alloc() ScreenshotManager {
	rv := objc.Send[ScreenshotManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScreenshotManager */
// An instance for the capture of single frames from a stream.


// An instance for the capture of single frames from a stream.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScreenshotManager */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScreenshotManager */

// Captures a single frame from a stream as an image, using a filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager/captureImage(contentFilter:configuration:completionHandler:)
func (sc _ScreenshotManagerClass) CaptureImageWithFilterConfigurationCompletionHandler(contentFilter ISCContentFilter, config ISCStreamConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("captureImageWithFilter:configuration:completionHandler:"), contentFilter, config, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CaptureImageWithFilterConfigurationCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager/captureImage(in:completionHandler:)
func (sc _ScreenshotManagerClass) CaptureImageInRectCompletionHandler(rect corefoundation.CGRect, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("captureImageInRect:completionHandler:"), rect, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CaptureImageInRectCompletionHandler) */


// Captures a single frame directly from a stream’s buffer, using a filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager/captureSampleBuffer(contentFilter:configuration:completionHandler:)
func (sc _ScreenshotManagerClass) CaptureSampleBufferWithFilterConfigurationCompletionHandler(contentFilter ISCContentFilter, config ISCStreamConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("captureSampleBufferWithFilter:configuration:completionHandler:"), contentFilter, config, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CaptureSampleBufferWithFilterConfigurationCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager/captureScreenshot(contentFilter:configuration:completionHandler:)
func (sc _ScreenshotManagerClass) CaptureScreenshotWithFilterConfigurationCompletionHandler(contentFilter ISCContentFilter, config ISCScreenshotConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("captureScreenshotWithFilter:configuration:completionHandler:"), contentFilter, config, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CaptureScreenshotWithFilterConfigurationCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotManager/captureScreenshot(rect:configuration:completionHandler:)
func (sc _ScreenshotManagerClass) CaptureScreenshotWithRectConfigurationCompletionHandler(rect corefoundation.CGRect, config ISCScreenshotConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("captureScreenshotWithRect:configuration:completionHandler:"), rect, config, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CaptureScreenshotWithRectConfigurationCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScreenshotManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScreenshotManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScreenshotManager */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SCScreenshotManager */


